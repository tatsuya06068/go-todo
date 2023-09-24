package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	auth := controller.NewAuthController(dbDriver.NewSqlHandler())

	// signup
	mux.HandleFunc("/auth/signup/", auth.Signup)
	// signin
	mux.HandleFunc("/auth/signin/", auth.Signin)

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", mux))
}