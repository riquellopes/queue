package worker

import (
	"fmt"

	worker "github.com/contribsys/faktory_worker_go"
)

// GetName -
func GetName(ctx worker.Context, args ...interface{}) error {
	fmt.Println("Working on job", ctx.Jid(), args)
	return nil
}

// Run -
func Run() {
	mgr := worker.NewManager()

	// Register function
	mgr.Register("name", GetName)

	mgr.Run()
}
