package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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
	ch := strings.ToLower(os.Args[1])

	switch ch {
	case "bookticket": //Book Ticket
		bookTicket(conn)
	case "listflights": //List Flights
		listFlights(conn)
	case "rateflights": //Rate flights
		rateFlights(conn)
	case "uploadform": //UserInfo
		uploadForm(conn)
	default:
		fmt.Println("Invalid API")
	}

}

func bookTicket(conn *grpc.ClientConn) {
	name := os.Args[2]

	c := pbBooking.NewTicketServiceClient(conn)
	res, err := c.BookTicket(context.Background(), &pbBooking.Passenger{Name: name})
	if err != nil {
		log.Fatalf("Error when calling BookTicket: %s", err)
	}
	log.Printf("Your seat number is : %d", res.SeatNumber)
}

func listFlights(conn *grpc.ClientConn) {
	src := os.Args[2]
	dest := os.Args[3]
	c := pbBooking.NewFlightinfoClient(conn)

	path := &pbBooking.JourneyPath{Source: src, Destination: dest}

	stream, err := c.ListFlights(context.Background(), path)
	if err != nil {
		log.Fatalf("Error when calling ListFlights: %s", err)
	}
	fmt.Println("Available Flights: ")
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
	c := pbRating.NewRatingServiceClient(conn)
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("Error Parsing n")
	}

	for i := 1; i <= n; i++ {
		id := os.Args[2*i+1]
		rate, err := strconv.ParseInt(os.Args[2*(i+1)], 10, 32)
		data := &pbRating.RatingRequest{FlightId: id, Rating: int32(rate)}

		stream, err := c.RateFlights(context.Background())
		if err != nil {
			log.Fatalf("Error when calling RateFlights: %s", err)
		}

		stream.Send(data)

		res, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
			break
		}
		fmt.Println("Response recieved: \n" + res.Quality)
	}
}

func uploadForm(conn *grpc.ClientConn) {
	data := [...]*pbUserInfo.Map{
		&pbUserInfo.Map{Key: "Name", Value: os.Args[2]},
		&pbUserInfo.Map{Key: "DOB", Value: os.Args[3]},
		&pbUserInfo.Map{Key: "Country", Value: os.Args[4]},
		&pbUserInfo.Map{Key: "Email", Value: os.Args[5]}}

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

/* Send all req continuously - recieve all responses continuously
func rateFlights(conn *grpc.ClientConn) {
	var id string
	var rate int32
	var n int

	fmt.Println("No. of flights to rate: ")
	fmt.Scanln(&n)

	c := pbRating.NewRatingServiceClient(conn)
	stream, err := c.RateFlights(context.Background())
	if err != nil {
		log.Fatalf("Error when calling RateFlights: %s", err)
	}

	var data = make([]*pbRating.RatingRequest, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter your fligtid: ")
		fmt.Scanln(&id)
		fmt.Println("Enter rating: ")
		fmt.Scanln(&rate)
		data[i] = &pbRating.RatingRequest{FlightId: id, Rating: rate}
	}
	for _, info := range data {
		stream.Send(info)
	}
	stream.Send(&pbRating.RatingRequest{}) // Final request

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
			break
		}
		fmt.Println("Response recieved: \n" + res.Quality)
	}
}
*/
