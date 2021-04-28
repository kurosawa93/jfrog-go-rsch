package rsch

import (
	"context"
	"fmt"

	otel "bitbucket.org/tunaiku/amargo-otel"
)

func Diagnostic() {
	fmt.Printf("Diagnosting OK!\n")

	id := otel.GetTraceID(context.Background())
	fmt.Printf("Trace ID: %v\n", id)
}

func Hello(name string) {
	fmt.Printf("Hello, Sir " + name)
}

func Salute(name string) {
	fmt.Printf("SALUTE TO " + name + "!!!!")
}
