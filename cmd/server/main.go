package main

import (
	"fmt"
	"net/http"
)

const PORT = 9090

func main() {
	fmt.Println("Starting server on port:", PORT)
	err := http.ListenAndServe(":"+fmt.Sprint(PORT), http.FileServer(http.Dir("../../assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
