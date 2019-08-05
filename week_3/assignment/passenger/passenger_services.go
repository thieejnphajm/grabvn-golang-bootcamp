package passenger

import (
	"context"
	"github.com/jinzhu/gorm"
)

type PassengerFeedbackServer struct {
	FeedbackMap map[string]*PassengerFeedback
	DB          *gorm.DB
}

var Success = &ResponseCode{
	Code:    0,
	Message: "Success",
}

var NotFound = &ResponseCode{
	Code:    1,
	Message: "Not found",
}

var Error = &ResponseCode{
	Code:    2,
	Message: "Exists Feedback",
}

type FeedbackModel struct {
	gorm.Model
	BookingCode string
	PassengerID int32
	Feedback    string
}

func (s *PassengerFeedbackServer) AddPassengerFeedback(ctx context.Context, pf *PassengerFeedback) (*AddPassengerFeedbackResponse, error) {
	var responseCode = Success

	s.DB.Create(&FeedbackModel{BookingCode: pf.BookingCode, PassengerID: pf.PassengerID, Feedback: pf.Feedback})

	return &AddPassengerFeedbackResponse{
		ResponseCode: responseCode,
	}, nil
}

func (s *PassengerFeedbackServer) GetPassengerFeedbackByPassengerId(ctx context.Context, feedbackReq *GetPassengerFeedbackRequest) (*GetPassengerFeedbacksResponse, error) {
	var responseCode = Success
	var passengerFeedbacks []*PassengerFeedback

	var feedbacks []FeedbackModel

	s.DB.Where(&FeedbackModel{PassengerID: feedbackReq.PassengerID}).Find(&feedbacks)

	if len(feedbacks) == 0 {
		responseCode = NotFound
	}

	for _, v := range feedbacks {
		fb := PassengerFeedback{
			BookingCode: v.BookingCode,
			PassengerID: v.PassengerID,
			Feedback:    v.Feedback,
		}
		passengerFeedbacks = append(passengerFeedbacks, &fb)
	}

	return &GetPassengerFeedbacksResponse{
		ResponseCode:       responseCode,
		PassengerFeedbacks: passengerFeedbacks,
	}, nil

}

func (s *PassengerFeedbackServer) GetPassengerFeedbackByBookingCode(ctx context.Context, feedbackReq *GetPassengerFeedbackRequest) (*GetPassengerFeedbackResponse, error) {
	var feedback FeedbackModel

	s.DB.Where(&FeedbackModel{BookingCode: feedbackReq.BookingCode}).First(&feedback)

	return &GetPassengerFeedbackResponse{
		ResponseCode: Success,
		PassengerFeedback: &PassengerFeedback{
			BookingCode: feedback.BookingCode,
			PassengerID: feedback.PassengerID,
			Feedback:    feedback.Feedback,
		},
	}, nil
}

func (s *PassengerFeedbackServer) DeletePassengerFeedbackPassengerId(ctx context.Context, feedbackReq *DeletePassengerFeedbackRequest) (*DeletePassengerFeedbackResponse, error) {
	s.DB.Delete(&FeedbackModel{PassengerID: feedbackReq.PassengerID})

	return &DeletePassengerFeedbackResponse{
		ResponseCode: Success,
	}, nil
}
