package main

import (
	"context"
	"grpc/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	log.Println("Server is hit.")
	amount:=req.GetAmount()
	from:=req.GetFrom()
	to:=req.GetTo()

	log.Println("Amount is ",amount)
	log.Println("From address is ",from)
	log.Println("To address is ",to)


	return &invoicer.CreateResponse{
		Pdf:  []byte("This is the PDF."),
		Docx: []byte("This is the Document."),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()

	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
