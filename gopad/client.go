package gopad

import (
	"net/http"

	"github.com/gopad/gopad-go/gopad/members/v1/membersv1connect"
	"github.com/gopad/gopad-go/gopad/teams/v1/teamsv1connect"
	"github.com/gopad/gopad-go/gopad/users/v1/usersv1connect"
)

// DefaultHTTPClient defines a default HTTP client.
func DefaultHTTPClient() *http.Client {
	return &http.Client{}
}

// New initializes a new client.
func New(options ...ClientOption) *Client {
	client := &Client{
		httpClient: DefaultHTTPClient(),
	}

	for _, option := range options {
		option(client)
	}

	client.Users = usersv1connect.NewUsersServiceClient(
		client.httpClient,
		client.baseURL,
	)

	client.Teams = teamsv1connect.NewTeamsServiceClient(
		client.httpClient,
		client.baseURL,
	)

	client.Members = membersv1connect.NewMembersServiceClient(
		client.httpClient,
		client.baseURL,
	)

	return client
}

// Client defines the client.
type Client struct {
	httpClient *http.Client
	baseURL    string

	Users   usersv1connect.UsersServiceClient
	Teams   teamsv1connect.TeamsServiceClient
	Members membersv1connect.MembersServiceClient
}

// ClientOption is used to configure a Client.
type ClientOption func(*Client)

// WithHTTPClient configures a Client to use the specified HTTP client.
func WithHTTPClient(value *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = value
	}
}

// WithBaseURL configures a Client to use the specified base URL.
func WithBaseURL(value string) ClientOption {
	return func(client *Client) {
		client.baseURL = value
	}
}
