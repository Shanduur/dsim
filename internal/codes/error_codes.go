package codes

// Errors:
// - NoError indicates no error
// - UnknownErro indicates unknown / internal error
// - ConfError indicates configuration error
// - LogError indicates error with logs
// - ComputeError indicates error with computational goroutine
// - DbError indicates error with database connection OR query
// - ManagerConnectionError indicates error while connectiong to server
// - WorkerConnectionError indicates error while connectiong to worker
// - ClientConnectionError indicates error while exchanging data with client
const (
	NoError                = iota
	UnknownError           = iota
	ConfError              = iota
	LogError               = iota
	ComputeError           = iota
	DbError                = iota
	ManagerConnectionError = iota
	WorkerConnectionError  = iota
	ClientConnectionError  = iota
)
