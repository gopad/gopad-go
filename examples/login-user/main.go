package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/gopad/gopad-go/v1/gopad"
	"github.com/gopad/gopad-go/v1/gopad/auth"
	"github.com/gopad/gopad-go/v1/models"
)

func main() {
	username := "your-username"
	password := strfmt.Password("your-password")

	resp, err := gopad.Default.Auth.LoginUser(
		auth.NewLoginUserParams().WithAuthLogin(
			&models.AuthLogin{
				Username: &username,
				Password: &password,
			},
		),
	)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
