package pb

// Unimplemented interface is empty interface for any type of request or response
type Unimplemented interface{}

// FUnimplemented is func that can be put in place of unimplemented function that accepts
// request and returns response and error
func FUnimplemented(req Unimplemented) (res Unimplemented, err error) {
	return
}
