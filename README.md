# Gopad: SDK for Go

[![General Workflow](https://github.com/gopad/gopad-go/actions/workflows/general.yml/badge.svg)](https://github.com/gopad/gopad-go/actions/workflows/general.yml) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/6aafa031df1746baa55287204ccea99f)](https://www.codacy.com/gh/gopad/gopad-go/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gopad/gopad-go&amp;utm_campaign=Badge_Grade) [![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org) [![Go Reference](https://pkg.go.dev/badge/github.com/gopad/gopad-go.svg)](https://pkg.go.dev/github.com/gopad/gopad-go)

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

func main() {

}
```

### List teams

[embedmd]:# (examples/list-teams/main.go go)
```go
package main

func main() {

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

If you find a security issue please contact gopad@webhippie.de first.

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
