# diceroll

![CI](https://github.com/spenserblack/diceroll/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/spenserblack/diceroll)](https://goreportcard.com/report/github.com/spenserblack/diceroll)
[![Go Reference](https://pkg.go.dev/badge/github.com/spenserblack/diceroll.svg)](https://pkg.go.dev/github.com/spenserblack/diceroll)

Roll some dice

## Usage

### Library

```go
combo, err := dice.ParseCombo("3d4 + 1d6 + 2")

// handle error

fmt.Println(combo.Roll())
```

### cmd

```bash
# install
go get -u github.com/spenserblack/diceroll/cmd/...

# roll 3d4 + 1d6 + 2
diceroll 3d4 + 1d6 + 2
```
