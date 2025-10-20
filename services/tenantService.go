package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
	"github.com/go-redis/redis/v8"
)

var ErrTenantNotFound = errors.New("tenant not found")

func GetTenantConfig(ctx context.Context, id string) (*models.TenantConfig, error) {
	cacheKey := fmt.Sprintf("tenant_config:%s", id)

	val, err := db.Rdb.Get(db.Ctx, cacheKey).Result()
	if err == nil {
		var config models.TenantConfig
		if json.Unmarshal([]byte(val), &config) == nil {
			return &config, nil
		}
	} else if err != redis.Nil {
		fmt.Printf("Redis error on get: %v\n", err)
	}

	tenant, err := repository.GetTenantByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	configJSON, err := json.Marshal(tenant.Config)
	if err == nil {
		db.Rdb.Set(db.Ctx, cacheKey, configJSON, 10*time.Minute)
	}

	return &tenant.Config, nil
}
