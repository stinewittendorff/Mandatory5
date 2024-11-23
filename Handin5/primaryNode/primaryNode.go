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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Setting up our node/servers abilities to match the requirements
type primaryNode struct {
	proto.UnimplementedAuctionServer
	ID           int32
	backupBidder BackupNodeClient
	backupServer map[int32]proto.BackupNodeClient
	ctx          context.Context
}

type BackupNodeClient struct {
	ID       int32
	serverID int32
	ctx      context.Context
}

// Function used to launch the primary server
func launch(p *primaryNode) {
	grpcServer := grpc.NewServer()
	//Creates a listener for our primary server
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", p.ID))
	//Checks if it is possible to start the server on the given port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Print("Listening on port %d", p.ID)

	//Registers the server
	proto.RegisterAuctionServer(grpcServer, p)

	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Failed to serve: %v", serveError)
	}
}

func (p *primaryNode) NewPrimaryServer() {
	delete(p.backupServer, p.backupBidder.serverID)

	for id, client := range p.backupServer {
		reply, err := client.IsPrimary(p.ctx, &proto.Empty{})
		if err != nil {
			log.Printf("Could not ping %d\n", id)
		} else {
			if reply.IsPrimary {
				log.Printf("The new primary server is %d\n", id)
				p.backupBidder.serverID = id
				break
			}
		}
	}
}

func (p *primaryNode) Bid(ctx context.Context, in *proto.BidRequest) (*proto.BidResponse, error) {
	var ack *proto.BidResponse
	ack, err := p.backupServer[p.backupBidder.serverID].BidBackup(ctx, in)
	if err != nil {
		log.Printf("The primary server is not repsonding. \n")
		p.NewPrimaryServer()
		ack, _ = p.backupServer[p.backupBidder.serverID].BidBackup(ctx, in)
	}

	return ack, nil
}
func (p *primaryNode) Result(ctx context.Context, in *proto.ResultRequest) (*proto.ResultResponse, error) {
	var outcome *proto.ResultResponse
	outcome, err := p.backupServer[p.backupBidder.serverID].ResultBackup(ctx, in)

	if err != nil {
		log.Printf("Primary server is not responding. \n")
		p.NewPrimaryServer()
		outcome, _ = p.backupServer[p.backupBidder.serverID].ResultBackup(ctx, in)
	}

	return outcome, nil
}

func (p *primaryNode) connect(connectionPort int32) (proto.BackupNodeClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", connectionPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %d", connectionPort)
	}
	log.Printf("Connected to server %d", connectionPort)
	return proto.NewBackupNodeClient(conn), nil
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	arg2, _ := strconv.ParseInt(os.Args[2], 10, 32)
	ownPort := int32(arg1) + 5000
	clientPort := int32(arg2) + 6000
	serverPort := int32(6000)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b := &BackupNodeClient{
		ID:       clientPort,
		serverID: serverPort,
		ctx:      ctx,
	}

	p := &primaryNode{
		ID:           ownPort,
		backupBidder: *b,
		backupServer: make(map[int32]proto.BackupNodeClient),
		ctx:          ctx,
	}

	for i := 0; i < 3; i++ {
		port := int32(6000) + int32(i)
		serverConn, err := p.connect(port)
		if err != nil {
			log.Fatalf("Failed to connect to %d", port)
		}
		p.backupServer[port] = serverConn
	}

	f, err := os.OpenFile(fmt.Sprintf("primary-%d.log", ownPort), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mulWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mulWriter)

	log.Printf("Starting primary server on port %d", ownPort)

	launch(p)
}
