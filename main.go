package main

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {
	fmt.Print(uuid.Must(uuid.NewV4()))
}
