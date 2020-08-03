package main

import (
	"fmt"
	"log"
	"net"

	pbBooking "Airlines-booking-gRPC-Protobuf/genfiles/booking"
	pbUserinfo "Airlines-booking-gRPC-Protobuf/genfiles/userInfo"
	pbRating "Airlines-booking-gRPC-Protobuf/genfiles/rating"

	booking "Airlines-booking-gRPC-Protobuf/pkg/server/booking"
	userInfo "Airlines-booking-gRPC-Protobuf/pkg/server/userInfo"
	rating "Airlines-booking-gRPC-Protobuf/pkg/server/rating"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Airlines Booking Server")

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := booking.Server{}
	grpcServer := grpc.NewServer()
	pbBooking.RegisterTicketServiceServer(grpcServer, &s)
	pbBooking.RegisterFlightinfoServer(grpcServer, &s)

	serverUserInfo := userInfo.Server{}
	pbUserinfo.RegisterFormServer(grpcServer, &serverUserInfo)

	ratingServer := rating.Server{}
	pbRating.RegisterRatingServiceServer(grpcServer, &ratingServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
