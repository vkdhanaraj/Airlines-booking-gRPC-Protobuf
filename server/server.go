package main

import (
	"fmt"
	"log"
	"net"

	"airlines-booking/server/booking"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Airlines Booking Server")

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := booking.Server{}
	grpcServer := grpc.NewServer()
	booking.RegisterTicketServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
