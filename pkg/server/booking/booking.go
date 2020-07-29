package booking

import (
	"errors"
	"fmt"

	booking "Airlines-booking-gRPC-Protobuf/genfiles"
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

//FlightDetails method
func (s *Server) FlightDetails(ctx context.Context, company *pb.FlightFilter) (*pb.Flight, error) {
	if company.AirlineCompany != "" {
		fmt.Println("Flight details returned for " + company.AirlineCompany)
		return &pb.Flight{FlightName: "JETX123"}, nil
	}
	return &pb.Flight{}, errors.New("Flight info failed")
}

//ListFlights method
func (s *Server) ListFlights(path *pb.JourneyPath, server pb.Flightinfo_ListFlightsServer) error {
	if path.GetSource() != "" && path.GetDestination() != "" {
		names := [...]*booking.Flight{
			&pb.Flight{FlightName: "JET1"},
			&pb.Flight{FlightName: "JET2"},
			&pb.Flight{FlightName: "JET3"}}

		for _, i := range names {
			j := i
			if err := server.Send(j); err != nil {
				return err
			}
		}
		fmt.Println("Returned List of Flights")
	}
	return errors.New("Flight info failed")
}
