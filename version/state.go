package version

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
