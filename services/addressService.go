package services

import (
	"context"
	"errors"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var ErrAddressNotFound = errors.New("address not found or you do not have permission to delete it")

func AddAddress(ctx context.Context, userID int64, payload *models.AddAddressPayload) error {
	return repository.CreateAddress(ctx, userID, payload)
}

func GetMyAddresses(ctx context.Context, userID int64) ([]models.Address, error) {
	return repository.GetAddressesByUserID(ctx, userID)
}

func DeleteAddress(ctx context.Context, userID int64, addressID int64) error {
	rowsAffected, err := repository.DeleteAddress(ctx, userID, addressID)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrAddressNotFound
	}
	return nil
}
