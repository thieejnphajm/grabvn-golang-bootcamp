package main

import (
	pb "./passenger"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8080"
)

var passengerServiceClient pb.PassengerServiceClient

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	defer conn.Close()

	passengerServiceClient = pb.NewPassengerServiceClient(conn)

	router := gin.Default()

	router.POST("/feedback", addPassengerFeedback)
	router.DELETE("/feedback", deletePassengerFeedback)
	router.GET("/feedback_by_booking_code", getFeedbackByBookingCode)
	router.GET("/feedback_by_passenger", getFeedbackByPassengerID)
	router.Run(":8088")

}

func getFeedbackByPassengerID(c *gin.Context) {
	var params struct {
		PassengerID int32
	}

	err := c.BindJSON(&params)

	if err != nil {
		c.String(400, "invalid params")
		return
	}

	request := &pb.GetPassengerFeedbackRequest{
		BookingCode: params.BookingCode,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := passengerServiceClient.GetPassengerFeedbackByPassengerId(ctx, request)

	if err != nil {
		log.Fatalf("Couldn't found feedback: %v", err)
	}

	c.JSON(200, response)
}

func getFeedbackByBookingCode(c *gin.Context) {
	var params struct {
		BookingCode string
	}

	err := c.BindJSON(&params)

	if err != nil {
		c.String(400, "invalid params")
		return
	}

	request := &pb.GetPassengerFeedbackRequest{
		BookingCode: params.BookingCode,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := passengerServiceClient.GetPassengerFeedbackByBookingCode(ctx, request)

	if err != nil {
		log.Fatalf("Couldn't found feedback: %v", err)
	}

	c.JSON(200, response)
}

func deletePassengerFeedback(c *gin.Context) {
	var params struct {
		PassengerID int32
	}

	err := c.BindJSON(&params)

	if err != nil {
		c.String(400, "invalid params")
		return
	}

	request := &pb.DeletePassengerFeedbackRequest{
		PassengerID: params.PassengerID,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := passengerServiceClient.DeletePassengerFeedbackPassengerId(ctx, request)

	if err != nil {
		log.Fatalf("Couldn't delete feedback: %v", err)
	}

	c.JSON(200, response)
}

func addPassengerFeedback(c *gin.Context) {

	var params struct {
		PassengerID int32
		BookingCode string
		Feedback    string
	}

	err := c.BindJSON(&params)

	if err != nil {
		c.String(400, "invalid params")
		return
	}

	var feedback *pb.PassengerFeedback

	feedback = &pb.PassengerFeedback{
		PassengerID: params.PassengerID,
		BookingCode: params.BookingCode,
		Feedback:    params.Feedback,
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := passengerServiceClient.AddPassengerFeedback(ctx, feedback)

	if err != nil {
		log.Fatalf("Couldn't add feedback: %v", err)
	}

	c.JSON(200, response)
}
