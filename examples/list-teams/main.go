package main

import (
	"context"
	"log"

	connect "github.com/bufbuild/connect-go"
	"github.com/gopad/gopad-go/gopad"
	teamsv1 "github.com/gopad/gopad-go/gopad/teams/v1"
)

func main() {
	client := gopad.New(
		gopad.WithBaseURL("http://localhost:8080"),
	)

	resp, err := client.Teams.List(
		context.Background(),
		&connect.Request[teamsv1.ListRequest]{},
	)

	if err != nil {
		log.Fatalf("Failed to get teams: %s", err)
	}

	for _, team := range resp.Msg.Teams {
		log.Println(team.Name)
	}
}
