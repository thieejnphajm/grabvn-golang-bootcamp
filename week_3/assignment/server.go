package main

import (
	"./passenger"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net"
)

type FeedbackModel struct {
	gorm.Model
	BookingCode string
	PassengerID int32
	Feedback    string
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err2 := gorm.Open("sqlite3", "feedback.db")

	if err2 != nil {
		log.Fatalf("failed to listen: %v", err2)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&FeedbackModel{})

	s := grpc.NewServer()
	passenger.RegisterPassengerServiceServer(s, &passenger.PassengerFeedbackServer{
		DB: db,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
