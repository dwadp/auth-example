package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/dwadp/auth-example/auth"
	"github.com/go-redis/redis/v8"
)

type authRedisRepository struct {
	rdb *redis.Client
	ctx context.Context
}

func NewAuthRedisRepository(rdb *redis.Client, ctx context.Context) auth.Repository {
	return &authRedisRepository{
		rdb: rdb,
		ctx: ctx,
	}
}

func (a *authRedisRepository) Set(key string, value interface{}, expired time.Duration) error {
	if err := a.rdb.Set(a.ctx, key, value, expired).Err(); err != nil {
		return err
	}

	return nil
}

func (a *authRedisRepository) SetJSON(key string, value interface{}, expired time.Duration) error {
	jsonStr, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return a.Set(key, string(jsonStr), expired)
}

func (a *authRedisRepository) Get(key string, expectedType auth.Type) (interface{}, error) {
	value, err := a.rdb.Get(a.ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	if expectedType != "" {
		switch expectedType {
		case auth.Float:
			result, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return nil, err
			}

			return result, nil
		case auth.Integer:
			result, err := strconv.ParseInt(value, 0, 0)

			if err != nil {
				return nil, err
			}

			return result, nil
		case auth.Uint:
			result, err := strconv.ParseUint(value, 0, 0)

			if err != nil {
				return nil, err
			}

			return result, nil
		case auth.Boolean:
			result, err := strconv.ParseBool(value)

			if err != nil {
				return nil, err
			}

			return result, nil
		case auth.JSON:
			result := map[string]interface{}{}

			if err := json.Unmarshal([]byte(value), &result); err != nil {
				return nil, err
			}

			return result, nil
		}
	}

	return strings.Trim(value, "\""), nil
}

func (a *authRedisRepository) Delete(key string) error {
	if err := a.rdb.Do(a.ctx, "DEL", key).Err(); err != nil {
		return err
	}

	return nil
}
