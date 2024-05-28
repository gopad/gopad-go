package main

import (
	"context"
	"log"

	"github.com/gopad/gopad-go/gopad"
)

func main() {
	client, err := gopad.NewClientWithResponses(
		"http://localhost:8080/api/v1",
	)

	if err != nil {
		log.Fatalf("Failed to initialize client: %s", err)
	}

	resp, err := client.ListTeamsWithResponse(
		context.Background(),
		&gopad.ListTeamsParams{},
	)

	if err != nil {
		log.Fatalf("Failed to get teams: %s", err)
	}

	for _, t := range gopad.FromPtr(resp.JSON200.Teams) {
		log.Println(gopad.FromPtr(t.Name))
	}
}
