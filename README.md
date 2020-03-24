> generate a uuid from the command-line

![go](https://github.com/wayneashleyberry/uuid/workflows/go/badge.svg)
![security](https://github.com/wayneashleyberry/uuid/workflows/security/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/uuid)](https://goreportcard.com/report/github.com/wayneashleyberry/uuid)

### Installation

```sh
go get github.com/wayneashleyberry/uuid
```

### Usage

```sh
Usage:
  uuid [flags]
  uuid [command]

Available Commands:
  help        Help about any command
  new         Generate a new UUID
  test        Test if a uuid is valid
  version     Print version information

Flags:
  -h, --help   help for uuid

Use "uuid [command] --help" for more information about a command.
```

### Example

```sh
uuid n | pbcopy
```
