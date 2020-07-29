package main

import (
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

	c := pbBooking.NewTicketServiceClient(conn)

	response, err := c.BookTicket(context.Background(), &pbBooking.Passenger{Name: "Kidiyoor"})
	if err != nil {
		log.Fatalf("Error when calling BookTicket: %s", err)
	}
	log.Printf("Your seat number is : %d", response.SeatNumber)

	fc := pbBooking.NewFlightinfoClient(conn)

	res,er := fc.FlightDetails(context.Background(),&pbBooking.FlightFilter{AirlineCompany:"JET Airways"})
	if er!=nil {
		log.Fatalf("Error when calling FlightDetails: %s",er)
	}
	log.Printf("JetAirways Flights available: %s",res.FlightName)

}
