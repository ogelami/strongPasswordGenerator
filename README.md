# Strong Password Generator

A secure command-line password generator written in Go that uses cryptographic random number generation.

## Features

- Generate passwords with customizable character sets
- Support for hex and base64 encoded strings
- Cryptographically secure using Go's `crypto/rand`
- Simple and easy-to-use CLI interface

## Installation

```bash
go install github.com/ogelami/strongPasswordGenerator
make build
make install
```

## Usage

```bash
spwgen
```

## Security

This project uses cryptographic random number generation to ensure the security of the generated passwords.

## Building from Source

1. Clone the repository
2. Run `make build` to compile
3. The binary will be created as `spwgen`
