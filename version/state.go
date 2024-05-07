package version

import "time"

type Status int

const (
	New Status = iota
	BuildInProgress
	Built
	BuildFail
	ReleaseInstalling
	ReleaseInstalled
	ReleaseFail
	MasterInstalling
	MasterInstalled
	MasterFail
	MasterTagged
)

type State struct {
	ID                  uint64
	VersionID           uint64
	Status              Status
	CreatedAt           time.Time
	UpdatedAt           time.Time
	BuildInProgressAt   time.Time
	BuiltAt             time.Time
	BuildFailAt         time.Time
	ReleaseInstallingAt time.Time
	ReleaseInstalledAt  time.Time
	ReleaseFailAt       time.Time
	MasterInstallingAt  time.Time
	MasterFailAt        time.Time
	MasterTaggedAt      time.Time
}
