package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"grpc/invoicer"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := invoicer.NewInvoicerClient(conn)

	// Contact the server and print out its response.
	response, err := c.Create(context.Background(), &invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:   100, // Example amount value
			Currency: "USD", // Example currency
		},
		From: "Sender",
		To:   "Receiver",
	})
	if err != nil {
		log.Fatalf("could not create invoice: %v", err)
	}
	fmt.Printf("PDF: %s\n", response.Pdf)
	fmt.Printf("DOCX: %s\n", response.Docx)
}
