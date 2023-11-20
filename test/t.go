package test

import (
	"fmt"

	"github.com/google/uuid"
)

func Hello() {
	fmt.Println("HELLO v2")
	fmt.Println(uuid.New())
}
