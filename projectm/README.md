# Write packages, no programs

> All design decisions start and end with the package. The purpose of a
> package is to provide a solution to a specific problem domain. To be purposeful,
> packages must provide, not contain. The more focused each package’s purpose is,
> the more clear it should be what the package provides.
> —Bill Kennedy, “Design Philosophy On Packaging”

[Reference](https://bitfieldconsulting.com/golang/packages)

This repo contains the following components:

- One module with 3 packages (including `main`).
- One package contains one file where the `SayHello` func lives, the other one
contains one file where `SayBye` lives in.
- Another `main` package living in `cmd` directory which allows us to generate a
  binary using importing the `SayHello` func.

## How-to run

```sh
go run cmd/main.go
```

## How-to build the binary

```sh
go build cmd/main.go
```
