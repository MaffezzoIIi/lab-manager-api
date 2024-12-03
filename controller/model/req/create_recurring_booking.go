package req

type CreateRecurringBookingRequest struct {
	LabID      string `json:"lab_id"`     // Lab ID
	UserID     string `json:"user_id"`    // User ID
	StartTime  string `json:"start_time"` // Start time in ISO 8601 format
	EndTime    string `json:"end_time"`   // End time in ISO 8601 format
	Period     string `json:"period"`     // Booking period
	DaysOfWeek []int  `json:"daysOfWeek"` // Days of the week (0 = Sunday, 6 = Saturday)
	Count      int    `json:"count"`      // Number of bookings to create
}
