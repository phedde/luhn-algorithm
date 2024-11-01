# Luhn Algorithm Package for Go

A Go package that implements the Luhn algorithm, also known as the "modulus 10" or "mod 10" algorithm. This package provides functionality to generate and validate numbers that follow the Luhn algorithm, commonly used for validating various identification numbers such as credit card numbers, Social Security Numbers, etc.

## Installation

To install the package, use `go get`:

```bash
go get github.com/phedde/luhn-algorithm
```

## Usage

Import the package in your Go code:

```go
import "github.com/phedde/luhn-algorithm"
```

### Available Functions

The package provides four main functions:

#### 1. Generate Full Number with Check Digit

```go
number, err := luhn.FullNumber(7992739871)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Full number with check digit: %d\n", number) // prints 79927398713
```

#### 2. Calculate Check Digit Only

```go
checkDigit, err := luhn.CheckDigit(7992739871)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Check digit: %d\n", checkDigit) // prints 3
```

#### 3. Validate a Number

```go
isValid := luhn.IsValid(79927398713)
fmt.Printf("Is valid: %v\n", isValid)
```

### Error Handling

All functions that can return errors should be properly handled. The package will return errors in the following cases:
- Input number is not positive
- Invalid digits in the input
- Parsing errors

## Example

Here's a complete example showing how to use the package:

```go
package main

import (
    "fmt"
    "log"
    "github.com/phedde/luhn-algorithm"
)

func main() {
    // Generate a full number with check digit
    number, err := luhn.FullNumber(7992739871)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Full number: %d\n", number) // prints 79927398713

    // Calculate just the check digit
    checkDigit, err := luhn.CheckDigit(7992739871)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Check digit: %d\n", checkDigit)

    // Validate a number
    isValid := luhn.IsValid(79927398713)
    fmt.Printf("Is valid: %v\n", isValid) // prints true
}
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.