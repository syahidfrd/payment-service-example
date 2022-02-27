package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "payment/proto-store-example"

	"google.golang.org/grpc"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
}

func (PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	paymentLink := fmt.Sprintf("https://test.payment.com/code=%s", req.OrderId)
	return &pb.CreatePaymentResponse{
		PaymentLink: paymentLink,
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	var paymentSvc PaymentService
	pb.RegisterPaymentServiceServer(grpcServer, paymentSvc)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("gRPC server starting on port :3000")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
