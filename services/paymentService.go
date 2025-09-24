package services

import (
	"context"
	"errors"
	"time"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var ErrPaymentMethodNotFound = errors.New("payment method not found or you do not have permission to delete it")
var ErrProcessorFailed = errors.New("failed to validate token with payment processor")

func AddPaymentMethod(ctx context.Context, userID int64, payload *models.AddPaymentMethodPayload) error {

	cardDetails := struct {
		Brand    string
		LastFour string
		ExpMonth int
		ExpYear  int
	}{
		Brand:    "Visa",
		LastFour: "4242",
		ExpMonth: int(time.Now().Month()),
		ExpYear:  time.Now().Year() + 2,
	}

	newMethod := &models.PaymentMethod{
		UserID:         userID,
		ProcessorToken: payload.ProcessorToken,
		CardBrand:      cardDetails.Brand,
		LastFour:       cardDetails.LastFour,
		ExpiryMonth:    cardDetails.ExpMonth,
		ExpiryYear:     cardDetails.ExpYear,
	}

	return repository.CreatePaymentMethod(ctx, newMethod)
}

func GetMyPaymentMethods(ctx context.Context, userID int64) ([]models.PaymentMethod, error) {
	return repository.GetPaymentMethodsByUserID(ctx, userID)
}

func DeletePaymentMethod(ctx context.Context, userID, methodID int64) error {
	rowsAffected, err := repository.DeletePaymentMethod(ctx, userID, methodID)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrPaymentMethodNotFound
	}
	return nil
}
