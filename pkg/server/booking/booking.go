package booking

import (
	"errors"
	"fmt"

	pb "Airlines-booking-gRPC-Protobuf/genfiles"

	"golang.org/x/net/context"
)

//Server struct
type Server struct {
}

//BookTicket method
func (s *Server) BookTicket(ctx context.Context, passenger *pb.Passenger) (*pb.Booking, error) {
	if passenger.Name != "" {
		fmt.Println("Ticket booked for " + passenger.Name)
		return &pb.Booking{SeatNumber: 1, PassengerName: passenger.Name}, nil
	}
	return &pb.Booking{}, errors.New("Booking failed")
}
