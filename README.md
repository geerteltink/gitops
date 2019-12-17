# GitOps

A collection of git tools:

- Show status from projects in sub directories.
- Easy merging hotfix and feature pull requests.

## Install

Install GitOps from source.

```bash
git clone git@github.com:xtreamwayz/gitops.git
cd gitops
go install
```

## Usage

```bash
gitops help
gitops <command> help
gitops <command> --verbose

# Check git status of current and all sub directories
gitops status

# Set upstream remote
gitops upstream git@github.com:<original_organization>/<project>.git

# Create hotfix from pull request
gitops hotfix --pr 123  [--branch master]

# Create feature from pull request
gitops feature --pr 123 [--branch develop]

# Merge hotfix / feature
gitops merge [--branch master,develop]
```

## Develop

Run:

```bash
go run main.go
```

Compile and run:

```bash
go build .
.\gitops.exe
```

Update dependencies:

```bash
go get -u ./...
go mod tidy
```

Test:

```bash
go test ./...
go test -coverage -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## Contributing

***BEFORE you start work on a feature or fix***, please read & follow the
[contributing guidelines](https://github.com/xtreamwayz/.github/blob/master/CONTRIBUTING.md#contributing)
to help avoid any wasted or duplicate effort.

## Copyright and license

Code released under the [MIT License](https://github.com/xtreamwayz/.github/blob/master/LICENSE.md).
Documentation distributed under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
