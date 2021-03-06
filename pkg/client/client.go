package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbBooking "Airlines-booking-gRPC-Protobuf/genfiles/booking"
	pbRating "Airlines-booking-gRPC-Protobuf/genfiles/rating"
	pbUserInfo "Airlines-booking-gRPC-Protobuf/genfiles/userInfo"
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
	fmt.Println("Enter \n 1: Book Ticket \n 2: List flights between airports \n 3: Rate your flight \n 4: UserInfo")
	fmt.Scanln(&ch)

	switch ch {
	case "1": //Book Ticket
		bookTicket(conn)
	case "2": //List Flights
		listFlights(conn)
	case "3": //Rate flights
		rateFlights(conn)
	case "4": //UserInfo
		getUserInfo(conn)
	default:
		fmt.Println("Invalid Input")
	}

}

func bookTicket(conn *grpc.ClientConn) {
	var name string
	fmt.Println("Enter Your Name: ")
	fmt.Scanln(&name)

	c := pbBooking.NewTicketServiceClient(conn)
	res, err := c.BookTicket(context.Background(), &pbBooking.Passenger{Name: name})
	if err != nil {
		log.Fatalf("Error when calling BookTicket: %s", err)
	}
	log.Printf("Your seat number is : %d", res.SeatNumber)
}

func listFlights(conn *grpc.ClientConn) {
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

func rateFlights(conn *grpc.ClientConn) {
	var id string
	var rate int32
	var n int

	fmt.Println("No. of flights to rate: ")
	fmt.Scanln(&n)

	c := pbRating.NewRatingServiceClient(conn)

	for i := 1; i <= n; i++ {
		fmt.Println("Enter your fligtid: ")
		fmt.Scanln(&id)
		fmt.Println("Enter rating: ")
		fmt.Scanln(&rate)

		data := []*pbRating.RatingRequest{&pbRating.RatingRequest{FlightId: id, Rating: rate}}

		stream, err := c.RateFlights(context.Background())
		if err != nil {
			log.Fatalf("Error when calling RateFlights: %s", err)
		}

		for _, info := range data {
			stream.Send(info)
		}

		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		fmt.Println("Response recieved: \n" + res.Quality) //check
	}

}

func getUserInfo(conn *grpc.ClientConn) {
	var name string
	var dob string
	var country string
	var email string

	//Get data from user
	fmt.Println("Enter Name:")
	fmt.Scanln(&name)
	fmt.Println("Enter dob:")
	fmt.Scanln(&dob)
	fmt.Println("Enter country:")
	fmt.Scanln(&country)
	fmt.Println("Enter email:")
	fmt.Scanln(&email)

	data := [...]*pbUserInfo.Map{&pbUserInfo.Map{Key: "Name", Value: name},
		&pbUserInfo.Map{Key: "DOB", Value: dob},
		&pbUserInfo.Map{Key: "Country", Value: country},
		&pbUserInfo.Map{Key: "Email", Value: email}}

	c := pbUserInfo.NewFormClient(conn)
	stream, err := c.UploadFormData(context.Background())
	if err != nil {
		log.Fatalf("Error when calling UploadFormData: %s", err)
	}

	for _, info := range data {
		stream.Send(info)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	fmt.Println("Response recieved: \n" + res.Message)

}
