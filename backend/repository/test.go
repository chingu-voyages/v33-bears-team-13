package main

// This package tests the package `record` by writing to and reading from the record.

import (
	"fmt"
	"context"
	"time"
	"log"
	// "os"
	"summaries_record/record"
)

func main() {
	// os.Setenv("MONGODB_URI", )

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := record.Open(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = r.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// Add summary x 2

	err = r.AddSummary(ctx, "Current temperature in București...")
	if err != nil {
	    log.Fatal(err)
	}
	err = r.AddSummary(ctx, "Current temperature in București...")
	if err != nil {
	    log.Fatal(err)
	}

	// Read record

	if r_, err := r.Read(ctx); err != nil {
			log.Fatal(err)
	} else {
		fmt.Println(r_)
	}
}
