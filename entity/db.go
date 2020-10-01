package entity

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Database struct {
	Postgres *gorm.DB
	Redis    *redis.Client
}
