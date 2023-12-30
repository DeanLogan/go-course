package database

import (
	"time"
	
	"github.com/google/uuid"
)


type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}