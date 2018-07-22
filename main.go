package main

import (
	"io"
	"net/http"

	faktory "github.com/contribsys/faktory/client"
	worker "github.com/riquellopes/queue/worker"
)

// PostHandler -
func PostHandler(res http.ResponseWriter, req *http.Request) {

	client, _ := faktory.Open()

	defer client.Close()

	job := faktory.NewJob("name", "Henrique")
	_ = client.Push(job)

	io.WriteString(res, "Adding a new item in queue.")
}

func main() {

	go func() {
		worker.Run()
	}()

	http.HandleFunc("/", PostHandler)
	http.ListenAndServe(":5000", nil)
}
