package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	// Blank-import the function package so the init() runs
	_ "github.com/mark.francis/evidence"
)

func main() {
	if err := funcframework.Start("8090"); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
