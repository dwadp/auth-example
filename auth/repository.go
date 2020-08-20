package auth

import "time"

type Repository interface {
	Set(key string, value interface{}, expired time.Duration) error
	SetJSON(key string, value interface{}, expired time.Duration) error
	Get(key string, expectedType Type) (interface{}, error)
	Delete(key string) error
}
