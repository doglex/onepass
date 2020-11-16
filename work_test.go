package main

import (
	"fmt"
	"testing"
)

func TestGenPassword(t *testing.T) {
	fmt.Println(GenPassword(8))
}
