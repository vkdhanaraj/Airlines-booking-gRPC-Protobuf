package booking

import (
	"errors"
	"fmt"

	pb "Airlines-booking-gRPC-Protobuf/genfiles/booking"

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

//ListFlights method
func (s *Server) ListFlights(path *pb.JourneyPath, server pb.Flightinfo_ListFlightsServer) error {
	if path.GetSource() != "" && path.GetDestination() != "" {
		flights := [...]*pb.Flight{
			&pb.Flight{FlightName: "JET1"},
			&pb.Flight{FlightName: "JET2"},
			&pb.Flight{FlightName: "JET3"}}

		for _, j := range flights {
			if err := server.Send(j); err != nil {
				return err
			}
		}
		fmt.Println("Returned List of Flights")
	}
	return nil
}
