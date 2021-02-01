# diceroll

![CI](https://github.com/spenserblack/diceroll/workflows/CI/badge.svg)

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
