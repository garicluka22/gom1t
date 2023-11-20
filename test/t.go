package test

import (
	"fmt"

	"github.com/google/uuid"
)

func Hello() {
	fmt.Println("HELLO")
	fmt.Println(uuid.New())
}
