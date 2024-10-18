package main

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/api"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/ssr"
)

func main() {
	go ssr.ListenAndServe()
	go api.ListenAndServe()
}
