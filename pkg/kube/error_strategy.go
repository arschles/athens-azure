package kube

// ErrorStrategy is the strategy that a profile takes when installing,
// uninstalling, and updating resources
type ErrorStrategy string

const (
	ErrorStrategyStop     ErrorStrategy = "stop"
	ErrorStrategyRollback ErrorStrategy = "rollback"
	ErrorStrategyContinue ErrorStrategy = "continue"
)
