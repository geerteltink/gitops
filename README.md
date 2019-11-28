# GitOps

## Usage

```bash
# Check git status of current and all sub directories
$ gitops status

# Set upstream remote
$ gitops upstream git@github.com:<original_organization>/<project>.git
```

## Develop

Run:

```bash
$ go run main.go
```

Compile and run:

```bash
$ go build .
$ .\gitops.exe
```

Update dependencies:

```bash
$ go mod tidy
$ go mod tidy
```

Test:

```bash
$ go test ./...
$ go test -coverage -coverprofile=coverage.out ./...
$ go tool cover -html=coverage.out -o coverage.html
```

## Resources

- https://golang.org/src/os/exec/example_test.go
- https://github.com/spf13/cobra
- https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html
- https://marcofranssen.nl/start-on-your-first-golang-project/
- https://simplyitinc.blogspot.com/2016/10/testing-code-using-execcommand-in-go.html

## Contributing

***BEFORE you start work on a feature or fix***, please read & follow the
[contributing guidelines](https://github.com/xtreamwayz/.github/blob/master/CONTRIBUTING.md#contributing)
to help avoid any wasted or duplicate effort.

## Copyright and license

Code released under the [MIT License](https://github.com/xtreamwayz/.github/blob/master/LICENSE.md).
Documentation distributed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
