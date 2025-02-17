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

	resp, err := client.ListGroupsWithResponse(
		context.Background(),
		&gopad.ListGroupsParams{},
	)

	if err != nil {
		log.Fatalf("Failed to get teams: %s", err)
	}

	for _, t := range resp.JSON200.Groups {
		log.Println(gopad.FromPtr(t.Name))
	}
}
