package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"Airlines-booking-gRPC-Protobuf/client/booking"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := booking.NewTicketServiceClient(conn)

	response, err := c.BookTicket(context.Background(), &booking.Passenger{Name: "Kidiyoor"})
	if err != nil {
		log.Fatalf("Error when calling BookTicket: %s", err)
	}
	log.Printf("Response from server: %d", response.SeatNumber)

}
