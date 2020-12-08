package codes

import "errors"

// ErrRecordExists struct is an error that indicates, that record
// you tried to create, already exists
var ErrRecordExists = errors.New("unable to insert value: already exists")

// ErrSignalCanceled struct is an error that indicates, that context
// had error value equal to Canceled
var ErrSignalCanceled = errors.New("recieved cancel signal")

// ErrNotAuthenticated struct is an error that indicates,
// that user cannot be authenticated
var ErrNotAuthenticated = errors.New("username and password mismatch")

// ErrNoFreeNode struct is an error that indicates,
// that there is no node with available thread
var ErrNoFreeNode = errors.New("no node with free thread avaliable")

// ErrNoActiveNode struct is an error that indicates,
// that there is active node
var ErrNoActiveNode = errors.New("no active node avaliable")
