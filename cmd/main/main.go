package main

import (
	"fmt"
	"net/http"

	"git.ytrack.learn.ynov.com/BMETEHRI/Forum/pkg/routes"
)

func main() {
	routes.ForumRoutes()

	port := "5555"
	fmt.Println("\nlien du Forum --> http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
