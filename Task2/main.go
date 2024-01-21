package main

import (
	"fmt"
	"net/http"
	"task2/controller"
	"task2/router"
)

var version = "1.0.0"

func main() {
	versionController := controller.NewVersionController(version)

	r := router.NewRouter()
	r.RegisterController("/", controller.NewDefaultController())
	r.RegisterController("/GetVersion", versionController)
	r.RegisterController("/PutVersion", versionController)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	http.ListenAndServe(port, r)
}
