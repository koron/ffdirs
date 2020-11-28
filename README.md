# koron/ffdirs

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron/ffdirs)](https://pkg.go.dev/github.com/koron/ffdirs)
[![GoDoc](https://godoc.org/github.com/koron/ffdirs?status.svg)](https://godoc.org/github.com/koron/ffdirs)
[![Actions/Go](https://github.com/koron/ffdirs/workflows/Go/badge.svg)](https://github.com/koron/ffdirs/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/ffdirs)](https://goreportcard.com/report/github.com/koron/ffdirs)

**ffdirs** finds files from PATH directory list.

## Getting started

Install or update with...

```console
$ go get -u github.com/koron/ffdirs
```

## Usage

### Format

```
ffdirs [OPTIONS] {file names...}
```

### Options

*   `-srcenv` string
    name of environment variables of directory list (default "PATH")
*   `-verbose`
    verbose message

### Example

```console
$ ffdirs -srcenv INCLUDE stdio.h
```

Find `stdio.h` file from directory list in INCLUDE environment variable.
