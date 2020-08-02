package rating

import (
	"errors"
	"fmt"
	//rating "Airlines-booking-gRPC-Protobuf/genfiles/rating"
	pb "Airlines-booking-gRPC-Protobuf/genfiles/rating"

	"golang.org/x/net/context"
)

//Server struct
type Server struct {
}

type RatingServiceServer interface {
	CreateFlight(context.Context, *CreateRateRequest)(*CreateRateResponse,error)
	RateFlights(RatingService_RateFlightsServer) error
}

type RatingStore interface {
	Add(flightID string,score int32)(*Rating,error)
}

type Rating struct {
	Count int32
	Sum float64
}

type InMemoryRatingStore struct {
	mutex sync.RWMutex
	ratings map[string]*Rating
}

type FlightServer struct {
	flightStore FlightStore
	ratingStore RatingStore
}
func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		ratings: make(map[string]*Rating),
	}
}

func (store *InMemoryRatingStore)Add(flightID string, score int32)(*Rating,error){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	ratings :=store.ratings[flightID]
	if ratings == nil{
		ratings = &Rating{
			Count:1,
			Sum:score,
		}
	}
	else{
		ratings.Count++
		ratings.Sum+=score
	}
	store.ratings[flightID] = ratings
	return ratings,nil
}

func NewFlightServer(flightStore FlightStore,ratingStore RatingStore) *FlightServer{
	return &FlightServer(flightStore,ratingStore)
}

func(s *FlightServer)RateFlights(stream pb.RatingService_RateFlightsServer) error {
	for {
		err:=contextError(stream.Context())
		if err!=nil{
			return err
		}
	
		req,err = stream.Recv()
		if err==io.EOF{
			log.Println("no more data")
			break
		}

		if err!=nil{
			return logError(status.Errorf(codes.Unknown,"cannot receive stream request %v",err))
		}

		flightID:=req.GetFlightId()
		score:=req.GetRating()

		log.Println("Received rating request:flight id %s ; rating %d :",flightID,score)

		found,err:=s.FlightStore.Find(flightID)
		if err!=nil{
			return logError(status.Errorf(codes.Internal,"cannot find flight %v",err))
		}

		if found==nil{
			return logError(status.Errorf(codes.NotFound,"cannot find flightid %s ",flightID))
		}

		rating,err : s.ratingStore.Add(flightID,score)
		if err!=nil{
			return logError(status.Errorf(codes.Internal,"cannot add rating to the store: %v ",err))
		}

		res := &pb.RatingResponse{
			id:=flightID,
			count:=rating.Count,
			avgscore:=rating.Sum/float64(rating.Count),
		}

		err = stream.Send(res)
		if err!=nil{
			return logError(status.Errorf(codes.Unknown,"cannot send stream response %v ",err))
		}

	}
	return nil
}
