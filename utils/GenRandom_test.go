package utils

import (
	"fmt"
	"testing"
)

func TestGenerateRandomID(t *testing.T) {
	x := GenerateRandomID()
	fmt.Printf("%v\n", x)

	y := GenerateRandomSalt()
	fmt.Printf("%v\n", y)

}
