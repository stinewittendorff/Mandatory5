package main

import (
	proto "Auction/proto"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Auction struct {
	state      proto.ResultResponse_STATE
	winner     int32
	winningBid int32
}

type peer struct {
	proto.UnimplementedBackupNodeServer
	id                int32
	clients           map[int32]proto.BackupNodeClient
	PrimaryServer     bool
	primaryServerPort int32
	Lamport           int32
	auction           Auction
	wg                sync.WaitGroup
	ctx               context.Context
}

func (p *peer) elect() {
	delete(p.clients, p.primaryServerPort)

	lowestID := p.id
	for id, client := range p.clients {
		reply, err := client.Ping(p.ctx, &proto.PingMessage{Id: p.id})
		if err != nil {
			fmt.Printf("Got no reply from %d\n", id)
		}
		fmt.Printf("Got reply from id %v: %v\n", id, reply.Id)
		if lowestID > reply.Id {
			lowestID = reply.Id
		}
	}

	log.Printf("New Primary lead is &d", lowestID)
	if p.primaryServerPort == p.id {
		p.PrimaryServer = true
	}
}

func (p *peer) incrementLamport(serverTimestamp int32) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	if p.Lamport < serverTimestamp {
		p.Lamport = serverTimestamp + 1
	} else {
		p.Lamport++
	}
}

func (p *peer) IsPrimary(ctx context.Context, in *proto.Empty) (*proto.ServerResponse, error) {
	return &proto.ServerResponse{IsPrimary: p.PrimaryServer}, nil
}

func (p *peer) Ping(ctx context.Context, in *proto.PingMessage) (*proto.PingMessage, error) {
	log.Printf("Received a Ping! \n")
	return &proto.PingMessage{
		Id: p.id,
	}, nil
}

func (p *peer) ResultBackup(ctx context.Context, in *proto.ResultRequest) (*proto.ResultResponse, error) {
	if p.PrimaryServer {
		var wg sync.WaitGroup
		for id, client := range p.clients {
			wg.Add(1)
			go func(clientId int32, requestClient proto.BackupNodeClient) {
				defer wg.Done()
				log.Printf("Sending request to %d", clientId)
				_, err := requestClient.ResultBackup(ctx, in)

				if err != nil {
					fmt.Printf("Backup server %d is not responding", clientId)
					delete(p.clients, clientId)
				}
			}(id, client)
		}
		wg.Wait()
		log.Printf("recieved acknowledgements from all backups. \n")
	}
	log.Printf("Server recieved result request.\nSent result reply with outcome %s, winner Id %d adn winneramount %d\n",
		p.auction.state.String(), p.auction.winner, p.auction.winningBid)
	return &proto.ResultResponse{
		Outcome:       p.auction.state,
		HighestBidder: p.auction.winner,
		HighestBid:    p.auction.winningBid,
	}, nil
}

func (p *peer) BidBackup(ctx context.Context, in *proto.BidRequest) (*proto.BidResponse, error) {
	p.incrementLamport(in.Lamport)

	log.Printf("Server has received a bid from client %d with amount %d at lamport time %d\n", in.BidderId, in.Amount, p.Lamport)

	var response *proto.BidResponse

	p.wg.Add(1)
	defer p.wg.Done()

	if in.Amount <= 0 {
		response = &proto.BidResponse{
			Lamport: p.Lamport,
			Result:  proto.BidResponse_EXCEPTION,
		}
	} else if p.auction.state == proto.ResultResponse_NOTSTARTED || p.auction.state == proto.ResultResponse_FINISHED {
		p.auction.state = proto.ResultResponse_ONGOING
		p.auction.winner = in.BidderId
		p.auction.winningBid = in.Amount
		go p.StartAuction()
		response = &proto.BidResponse{
			Lamport: p.Lamport,
			Result:  proto.BidResponse_SUCCESS,
		}
	} else if in.Amount < p.auction.winningBid {
		log.Printf("Bid %d from client %d is lower than the current highest bid %d\n", in.Amount, in.BidderId, p.auction.winningBid)
		response = &proto.BidResponse{
			Lamport: p.Lamport,
			Result:  proto.BidResponse_FAIL,
		}
	} else {
		log.Printf("New bid %d from client %d is the highest bid.\n", in.Amount, in.BidderId)
		p.auction.winner = in.BidderId
		p.auction.winningBid = in.Amount
		response = &proto.BidResponse{
			Lamport: p.Lamport,
			Result:  proto.BidResponse_SUCCESS,
		}
	}

	return response, nil
}

func (p *peer) StartAuction() {
	log.Printf("An auction has started\n")
	time.Sleep(1000 * time.Microsecond)
	p.wg.Wait()
	p.auction.state = proto.ResultResponse_FINISHED
	log.Printf("Auction has concluded!\n")
}

func (p *peer) ExtendAuction() {
	p.wg.Add(1)
	time.Sleep(5000 * time.Microsecond)
	p.wg.Done()
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 6000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f, err := os.OpenFile(fmt.Sprintf("logfile.%d", ownPort), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	p := &peer{
		id:                ownPort,
		clients:           make(map[int32]proto.BackupNodeClient),
		PrimaryServer:     false,
		primaryServerPort: 6000,
		Lamport:           1,
		auction: Auction{
			state:      proto.ResultResponse_NOTSTARTED,
			winner:     -1,
			winningBid: -1,
		},
		wg:  sync.WaitGroup{},
		ctx: ctx,
	}
	p.auction.winningBid = -1;

	if ownPort == p.primaryServerPort {
		p.PrimaryServer = true
	}

	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBackupNodeServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("Failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		port := int32(6000) + int32(i)
		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to connect to %d\n", port)
		conn, err = grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to %d", port)
		}
		log.Printf("Connected to server %d", port)
		client := proto.NewBackupNodeClient(conn)
		p.clients[port] = client
	}

	go func() {
		for !p.PrimaryServer {
			log.Print("Trying to become primary server\n")
			_, err := p.clients[p.primaryServerPort].Ping(ctx, &proto.PingMessage{Id: p.id})
			if err != nil {
				log.Printf("Primary server is down")
				p.elect()
			} else {
				log.Printf("Successfully pinged primary server. \n")
				time.Sleep(3000 * time.Millisecond)
			}
		}
	}()
	for {

	}
}
