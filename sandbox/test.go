package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	l "sandbox/logger"
)

type LoginInput struct {
	Username string `json:"username"`
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	logger := l.GetReqLogger(req.URL.Path)
	logger.Info("received login request")

	raw, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("an error occured while reading request body: %s", err))
		res.Write([]byte("an error occured"))
		return
	}

	// logger.Info(fmt.Sprintf("RAW %s", raw))

	var input LoginInput
	err = json.Unmarshal(raw, &input)
	if err != nil {
		logger.Error(fmt.Sprintf("an error occured while unmarshaling request body: %s", err))
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("an error occured"))
		return
	}

	// Add some req validation

	logger.Info(fmt.Sprintf("user %s logged in", input.Username))
}

func testHandler() {
	// http.HandleFunc("/login", loginHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /login", loginHandler)
	http.ListenAndServe(":8080", nil)
}
