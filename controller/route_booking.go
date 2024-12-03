package controller

import (
	"lab-manager-api/controller/model/req"
	"lab-manager-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param booking body req.CreateBooking true "Booking object that needs to be created"
// @Success 201 {object} res.BookingResponse
// @Failure 400 {object} rest_err.RestErr
// @Router /api/v1/bookings/create [post]
func CreateBooking(c *gin.Context) {
	var createBooking req.CreateBooking
	if err := c.ShouldBindJSON(&createBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	labID, err := primitive.ObjectIDFromHex(createBooking.LabID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid LabID"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(createBooking.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UserID"})
		return
	}

	booking, err := models.NewBooking(labID, userID, createBooking.StartTime, createBooking.EndTime, createBooking.Period, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating booking"})
		return
	}

	if booking, err = models.SaveBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving booking"})
		return
	}

	c.JSON(http.StatusCreated, MapToBookingResponse(booking))
}

// CancelBooking godoc
// @Summary Cancel a booking
// @Description Cancel a booking
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param id path string true "Booking ID"
// @Success 200 {object} nil
// @Failure 400 {object} rest_err.RestErr
// @Router /api/v1/bookings/{id} [delete]
func CancelBooking(c *gin.Context) {
	bookingIDHex := c.Param("id")
	bookingID, err := primitive.ObjectIDFromHex(bookingIDHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	err = models.DeleteBooking(bookingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking cancelled successfully"})
}

// CreateRecurringBooking godoc
// @Summary Create recurring bookings
// @Description Create recurring bookings
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param booking body req.CreateRecurringBookingRequest true "Recurring booking object that needs to be created"
// @Success 201 {object} res.CreateRecurringBookingResponse
// @Failure 400 {object} rest_err.RestErr
// @Router /api/v1/bookings/recurring [post]
func CreateRecurringBooking(c *gin.Context) {
	var createBooking struct {
		LabID      string         `json:"lab_id"`
		UserID     string         `json:"user_id"`
		StartTime  time.Time      `json:"start_time"`
		EndTime    time.Time      `json:"end_time"`
		Period     string         `json:"period"`
		DaysOfWeek []time.Weekday `json:"days_of_week"`
		Count      int            `json:"count"`
	}

	if err := c.ShouldBindJSON(&createBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	labID, err := primitive.ObjectIDFromHex(createBooking.LabID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid LabID"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(createBooking.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UserID"})
		return
	}

	bookings, err := models.CreateRecurringBookings(labID, userID, createBooking.StartTime, createBooking.EndTime, createBooking.Period, createBooking.DaysOfWeek, createBooking.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating recurring bookings"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"bookings": bookings})
}

// CreateSoftwareRequest godoc
// @Summary Create a new software request
// @Description Create a new software request
// @Tags software_requests
// @Accept  json
// @Produce  json
// @Param request body req.CreateSoftwareRequest true "Software request object that needs to be created"
// @Success 201 {object} models.SoftwareRequest
// @Failure 400 {object} rest_err.RestErr
// @Router /api/v1/software_requests/create [post]
func CreateSoftwareRequest(c *gin.Context) {
	var createRequest struct {
		LabID    string `json:"lab_id"`
		UserID   string `json:"user_id"`
		Software string `json:"software"`
	}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	labID, err := primitive.ObjectIDFromHex(createRequest.LabID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid LabID"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(createRequest.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UserID"})
		return
	}

	request, err := models.NewSoftwareRequest(labID, userID, createRequest.Software)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating software request"})
		return
	}

	if request, err = models.SaveSoftwareRequest(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving software request"})
		return
	}

	notifyAdmin(request)

	c.JSON(http.StatusCreated, request)
}

func notifyAdmin(request models.SoftwareRequest) {
	// Implementar lógica de notificação aqui (ex: enviar email, mensagem, etc.)
}

func MapToBookingResponse(booking models.Booking) gin.H {
	return gin.H{
		"id":         booking.ID,
		"lab_id":     booking.LabID,
		"user_id":    booking.UserID,
		"start_time": booking.StartTime,
		"end_time":   booking.EndTime,
		"period":     booking.Period,
	}
}
