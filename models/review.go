package models

type LeaveReviewPayload struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}

type Review struct {
	ID      int64
	OrderID int64
	UserID  int64
	Rating  int64
	Comment string
}
