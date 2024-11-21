package main

import (
	proto "Auction/proto"

	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Bidder struct {
	id         int
	serverPort int
	Lamport    int32
	ctx        context.Context
}

// Method to increment our lamport timestamp
func (b *Bidder) incrementLamport(serverTimestamp int32) {
	var mu sync.Mutex
	defer mu.Unlock()
	if b.Lamport < serverTimestamp {
		b.Lamport = serverTimestamp + 1
	} else {
		b.Lamport++
	}
}

func (b *Bidder) connect() (proto.AuctionClient, error) {
	//Creates a connection to our GRPC server
	// the second half after the comma is used to establish an unsafe connection
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", b.serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	log.Print("Connected to server %d", b.serverPort)
	return proto.NewAuctionClient(conn), nil
}

func main() {
	//Start by setting up a port of the bidder and the server it connects to
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	arg2, _ := strconv.ParseInt(os.Args[0], 10, 32)

	//adjusting the port numbers
	bidderPort := int(arg1) + 5000
	serverPort := int(arg2) + 5000

	// Create a context, so we can cancel processes
	ctx, cancel := context.WithCancel(context.Background())
	// makes sure when the program is finished that our resources is released
	defer cancel()

	// Creates a new Bidder
	b := &Bidder{
		id:         bidderPort,
		serverPort: serverPort,
		Lamport:    0,
		ctx:        ctx,
	}

	// Opens a log file for the bidder
	f, err := os.OpenFile(fmt.Sprintf("bidder-%d.log", bidderPort), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mulWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mulWriter)

	// creates a connection to our server
	serverConn, _ := b.connect()

	//Reads input from the user
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		splitted := strings.Fields(input)
		//If the users input start starts with "bid" with an int after it, it will send the bid given to the server.
		if splitted[0] == "bid" {
			b.incrementLamport(-10)
			amount, err := strconv.Atoi(splitted[1])
			// If the amount is not an int, it will log an error
			if err != nil {
				log.Fatal(err)
			}
			// sends the bid to server
			log.Printf("Bidder %d bids %d, at timestamp %d\n", b.id, amount, b.Lamport)
			ack, err := serverConn.Bid(b.ctx, &proto.BidRequest{
				Lamport:  b.Lamport,
				BidderId: int32(b.id),
				Amount:   int32(amount),
			})
			//If the bid fails on the Serverside, it will log an error
			if err != nil {
				log.Fatal(err)
			}

			b.incrementLamport(ack.Lamport)
			log.Printf("Bidder %d received %s at timestamp %d\n", b.id, ack.Result, b.Lamport)
		} else if splitted[0] == "result" {
			//If the user instead inputs "result", it will get the current state of the auction, and winner and highest bid if auction is completed.
			outcome, err := serverConn.Result(b.ctx, &proto.ResultRequest{})
			//If the result fails on the Serverside, it will log an error
			if err != nil {
				log.Fatalf("Could not get result")
			}
			log.Printf("The auction is currently %s and the highest bid is %d by bidder %d\n", outcome.Outcome, outcome.HighestBid, outcome.HighestBidder)
		}
	}

}
