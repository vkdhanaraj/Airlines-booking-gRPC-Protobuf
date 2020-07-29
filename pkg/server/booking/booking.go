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

//FlightDetails method
func (s *Server) FlightDetails(ctx context.Context, company *pb.FlightFilter) (*pb.Flight,error) {
	if company.AirlineCompany !=""{
		fmt.Println("Flight details returned for " + company.AirlineCompany)
		return &pb.Flight{FlightName:"JETX123"},nil
	}
	return &pb.Flight{},errors.New("Flight info failed")
}

/*
func (s *Server) FlightDetails(ctx context.Context,company *pb.FlightFilter) (stream pb.Flight,error) {
	if company.AirlineCompany !=""{
		names :=string {"JET1","JET2","JET3"}


		for i,j :=range names{
			if err:=stream.Send(j);
			err!=nil{return err}
		}
		fmt.Println("Flight details returned for " + company.AirlineCompany)
	}
	return &pb.Flight{},errors.New("Flight info failed")
}
*/


