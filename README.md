# Gopad: SDK for Go

[![General Workflow](https://github.com/gopad/gopad-go/actions/workflows/general.yml/badge.svg)](https://github.com/gopad/gopad-go/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/6aafa031df1746baa55287204ccea99f)](https://www.codacy.com/gh/gopad/gopad-go/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gopad/gopad-go&amp;utm_campaign=Badge_Grade) [![Go Reference](https://pkg.go.dev/badge/github.com/gopad/gopad-go.svg)](https://pkg.go.dev/github.com/gopad/gopad-go)

This repository provides a client SDK for Go.

## Installation

TBD

## Tests

TBD

## Examples

### List users

[embedmd]:# (examples/list-users/main.go go)
```go
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
```

### List teams

[embedmd]:# (examples/list-teams/main.go go)
```go
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
```

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.18, at least that's the version we are using.

```console
git clone https://github.com/gopad/gopad-go.git
cd gopad-go

make clean generate test
```

## Security

If you find a security issue please contact
[gopad@webhippie.de](mailto:gopad@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[golang]: http://golang.org/doc/install.html
