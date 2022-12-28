package main

import (
	"github.com/lucasvmiguel/stock-api/cmd/api/starter"
)

func main() {
	s := starter.New()
	s.Start()
}
