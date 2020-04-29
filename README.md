# Testground plan templates

Templates for `testground plan create`


## description:

This package contains template files as `go mod` compatible assets.

These are the tempalates used by the command `testground plan create`


## Adding new templates

* Create a new directory using the format <target>-templates
* generate `pkged.go` using the [pkger](https://github.com/markbates/pkger) with the -include flag
* double-check yourself, by running `go run main.go` from the root of this repository.

