package main

import (
	"math"
	"net/http"
	"os"
	"os/signal"

	"github.com/dshipenok/coverapi/service"
)

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		return math.Inf(1)
	}
	return a / b
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/multiply", service.NewHandlerBinaryOp(multiply))
	mux.Handle("/divide", service.NewHandlerBinaryOp(divide))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() { http.ListenAndServe(":8080", mux) }()
	<-stop
}
