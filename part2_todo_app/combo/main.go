package main

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/cmd/cliapp"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/inMemory"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/sql"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/api"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/ssr"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var repoType string
	var sharedStore repo.Repo

	flag.StringVar(&repoType, "r", "memory", "type of repository")
	flag.Parse()

	switch repoType {
	case "memory":
		sharedStore = new(inMemory.InMemory)
	case "sql":
		dbStore := new(sql.PostgresStore)
		err := dbStore.InitDB()
		if err != nil {
			panic(err)
		}

		sharedStore = dbStore
	default:
		sharedStore = new(inMemory.InMemory)
	}

	go cliapp.RunCli(sharedStore)
	go ssr.ListenAndServe(sharedStore)
	go api.ListenAndServe(sharedStore)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	doneChan := make(chan bool, 1)

	go func() {
		<-signalChan
		fmt.Println("\nReceived an interrupt, performing cleanup...")
		// Perform any cleanup here
		doneChan <- true
	}()

	fmt.Println("Press Ctrl+C to exit")
	<-doneChan
	fmt.Println("Cleanup complete, exiting.")
}
