package main

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/cmd/cliapp"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/ssr"
)

func main() {
	go cliapp.RunCli()
	go ssr.ListenAndServe()
}
