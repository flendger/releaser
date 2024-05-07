package application

import "time"

type Application struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
