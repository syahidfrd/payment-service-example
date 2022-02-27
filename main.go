package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"payment/model"

	"google.golang.org/grpc"
)

type PaymentService struct {
	model.UnimplementedPaymentServiceServer
}

func (PaymentService) CreatePayment(ctx context.Context, req *model.CreatePaymentRequest) (*model.CreatePaymentResponse, error) {
	paymentLink := fmt.Sprintf("https://test.payment.com/code=%s", req.OrderId)
	return &model.CreatePaymentResponse{
		PaymentLink: paymentLink,
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	var paymentSvc PaymentService
	model.RegisterPaymentServiceServer(grpcServer, paymentSvc)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("gRPC server starting on port :3000")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}
