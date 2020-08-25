package transfer

import (
	"github.com/Sheerley/pluggabl/pkg/plog"
	"golang.org/x/net/context"
)

// Server xd
type Server struct {
}

// SendPackage se
func (s *Server) SendPackage(ctx context.Context, in *Message) (*Response, error) {
	plog.Messagef("Receive message body from client: %s\n", in.Body)
	return &Response{Code: StatusCode_Ok}, nil
}
