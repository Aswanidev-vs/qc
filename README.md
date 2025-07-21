# quickcalc

A simple Go package providing basic arithmetic operations and utility functions like prime checking and factorial calculation.

## Features

- Addition, Subtraction, Multiplication, Division
- Prime number check
- Factorial calculation

## Installation

Clone the repository and use Go modules to import the package.

```sh
git clone https://github.com/Aswanidev-vs/qc.git
go mod tidy
```

## Usage

Import the package in your Go code:

```go
go get github.com/Aswanidev-vs/qc/quickcalc
```

## Example

See `main.go` for a sample implementation:

```go
package main

import (
	"fmt"

	q "github.com/Aswanidev-vs/qc/quickcalc"
)

func main() {

	fmt.Println("testing the my math package")
	var a, b, c, d, e, f, n int
	a, b = read(a, b)
	sum := q.Add(a, b)
	fmt.Println("sum is :", sum)
	c, d = read(c, d)
	sub := q.Sub(c, d)
	fmt.Println("sub is :", sub)
	e, f = read(e, f)
	multi := q.Multi(e, f)
	fmt.Println("mutli is :", multi)
	e, f = read(e, f)
	g := float64(e)
	h := float64(f)
	div := q.Div(g, h)
	if g/h == 0 {
		fmt.Println(e, " cannot be divided by Zero!")
	} else {
		fmt.Println("div is :", div)
	}

	fmt.Println("enter number to perform prime:")
	fmt.Scan(&n)
	prime := q.Prime(n)
	if prime != 0 {
		fmt.Println("This is a prime number:", prime)
	}
	var fact int
	fmt.Println("Enter number to perform factorial:")
	fmt.Scan(&fact)

	factor := q.Fact(fact)
	if factor != 0 {
		fmt.Println("Factorial is:", factor)
	}

}
func read(p, q int) (int, int) {
	fmt.Print("enter 2 number:")
	fmt.Scan(&p)
	fmt.Scan(&q)
	return p, q
}

```


 
