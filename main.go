package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	fmt.Print(uuid.NewV4())
}
