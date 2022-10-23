package main

import (
	"context"
	"log"

	connect "github.com/bufbuild/connect-go"
	"github.com/gopad/gopad-go/gopad"
	usersv1 "github.com/gopad/gopad-go/gopad/users/v1"
)

func main() {
	client := gopad.New(
		gopad.WithBaseURL("http://localhost:8080"),
	)

	resp, err := client.Users.List(
		context.Background(),
		&connect.Request[usersv1.ListRequest]{},
	)

	if err != nil {
		log.Fatalf("Failed to get users: %s", err)
	}

	for _, user := range resp.Msg.Users {
		log.Println(user.Username)
	}
}
