package entities

// BuildableStatus defines values supported by the statuses constraint.
type BuildableStatus string

var (
	// BuildableStatusPreparing - builbable is being prepared.
	BuildableStatusPreparing = "preparing"
	// BuildableStatusBuilding - building is in progress.
	BuildableStatusBuilding = "building"
	// BuildableStatusPassed - all blocking builds of builtable have passed.
	BuildableStatusPassed = "passed"
	// BuildableStatusFailed - some builds of buildable have failed.
	BuildableStatusFailed = "failed"
)
