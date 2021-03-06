package rating

import (
	pb "Airlines-booking-gRPC-Protobuf/genfiles/rating"
	"fmt"
	"io"
	"log"
)

//Server struct
type Server struct {
}

//RateFlights method
func (s *Server) RateFlights(stream pb.RatingService_RateFlightsServer) error {

	req, err := stream.Recv()
	if err == io.EOF {
		fmt.Print("no more data")
		return nil
	}
	if err != nil {
		log.Fatalf("cannot receive stream request: %v", err)
		return nil
	}
	fmt.Println("request received from client", req)

	id := req.GetFlightId()
	rate := req.GetRating()
	res := []*pb.RatingResponse{&pb.RatingResponse{FlightId: id, Quality: "None"}}

	if rate <= 5 {
		res = []*pb.RatingResponse{&pb.RatingResponse{FlightId: id, Quality: "Bad"}}
	} else if rate >= 6 && rate < 8 {
		res = []*pb.RatingResponse{&pb.RatingResponse{FlightId: id, Quality: "Above Average"}}
	} else {
		res = []*pb.RatingResponse{&pb.RatingResponse{FlightId: id, Quality: "Good"}}
	}

	for _, info := range res {
		err = stream.Send(info)
		if err != nil {
			log.Fatalf("cannot send stream response: %v", err)
		}
	}
	return nil
}
