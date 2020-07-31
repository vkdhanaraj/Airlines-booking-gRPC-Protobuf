package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbBooking "Airlines-booking-gRPC-Protobuf/genfiles"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	//Services Menu (Take input from user)
	var ch string
	fmt.Println("Menu")
	fmt.Println("Enter \n 1: Book Ticket \n 2: List flights between airports")
	fmt.Scanln(&ch)

	switch ch {
	case "1": //Book Ticket
		var name string
		fmt.Println("Enter Your Name: ")
		fmt.Scanln(&name)

		c := pbBooking.NewTicketServiceClient(conn)
		response, err := c.BookTicket(context.Background(), &pbBooking.Passenger{Name: name})
		if err != nil {
			log.Fatalf("Error when calling BookTicket: %s", err)
		}
		log.Printf("Your seat number is : %d", response.SeatNumber)

	case "2": //List Flights
		var src string
		var dest string
		fmt.Println("Enter Source")
		fmt.Scanln(&src)
		fmt.Println("Enter Destination")
		fmt.Scanln(&dest)

		c := pbBooking.NewFlightinfoClient(conn)

		path := &pbBooking.JourneyPath{Source: src, Destination: dest}

		stream, err := c.ListFlights(context.Background(), path)
		if err != nil {
			log.Fatalf("Error when calling ListFlights: %s", err)
		}
		log.Printf("Available Flights: ")
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal("cannot receive response: ", err)
			}
			fmt.Println(res.FlightName)
		}

	}

}
