package service

import "time"

type Service struct {
	ID            uint64
	Name          string
	ApplicationID uint64
	Git           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
