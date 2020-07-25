package booking

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"
)

//Server struct
type Server struct {
}

//BookTicket method
func (s *Server) BookTicket(ctx context.Context, passenger *Passenger) (*Booking, error) {
	if passenger.Name != "" {
		fmt.Println("Ticket booked for " + passenger.Name)
		return &Booking{SeatNumber: 1, PassengerName: passenger.Name}, nil
	}
	return &Booking{}, errors.New("Booking failed")
}
