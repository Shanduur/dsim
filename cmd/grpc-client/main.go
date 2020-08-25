package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Sheerley/pluggabl/pkg/plog"
	"github.com/Sheerley/pluggabl/pkg/transfer"
)

type lockable struct {
	id    int
	taken bool
}

func firesub(lc *lockable) {
	defer func() { (*lc).taken = false }()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		plog.Fatalf(1, "did not connect: %s", err)
	}
	defer conn.Close()

	c := transfer.NewTransferServiceClient(conn)

	s := fmt.Sprintf("Hello From Client #%v", (*lc).id)

	response, err := c.SendPackage(context.Background(), &transfer.Message{Body: s})
	if err != nil {
		plog.Fatalf(1, "Error when calling SendPackage: %s\n", err)
	}
	plog.Messagef("Response for client %v from server: %s\n", (*lc).id, response.Code)

	time.Sleep(105 * time.Millisecond)
}

func main() {

	var process [4]lockable

	for i := 0; i < 4; i++ {
		process[i].id = i
		process[i].taken = false
	}

	for {
		for i := 0; i < 4; i++ {
			if !(process[i].taken) {
				process[i].taken = true
				go firesub(&process[i])
				time.Sleep(100 * time.Millisecond)
				plog.Messagef("fired %v\n", process[i].id)
			}
		}
	}

}
