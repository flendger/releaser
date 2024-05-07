package version

import "time"

type Version struct {
	ID            uint64
	ApplicationID uint64
	Number        string
	ReleaseDate   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
