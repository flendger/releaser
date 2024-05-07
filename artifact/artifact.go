package artifact

import "time"

type Artifact struct {
	ID        uint64
	VersionID uint64
	ServiceID uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
