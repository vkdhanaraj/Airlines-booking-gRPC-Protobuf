package userinfo

import (
	pb "Airlines-booking-gRPC-Protobuf/genfiles/userInfo"
	"fmt"
	"io"
)

//Server Struct
type Server struct {
}

//UploadFormData method
func (s *Server) UploadFormData(stream pb.Form_UploadFormDataServer) error {
	var info pb.Message
	var flag bool = true

	fmt.Println("Recieving streamed data from client")
	for {
		pair, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if pair.Key == "" || pair.Value == "" {
			flag = false
		}
		fmt.Println(pair)
	}
	if flag {
		info.Message = "Data valid"
	} else {
		info.Message = "Data invalid"
	}
	return stream.SendAndClose(&info)

}
