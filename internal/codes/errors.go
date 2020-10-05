package codes

// RecordExists struct is an error that indicates, that record
// you tried to create, already exists
type RecordExists struct{}

func (err *RecordExists) Error() string {
	return "unable to insert value: already exists"
}

// SignalCanceled struct is an error that indicates, that context
// hadd error value equal to Canceled
type SignalCanceled struct{}

func (err *SignalCanceled) Error() string {
	return "recieved cancel signal"
}
