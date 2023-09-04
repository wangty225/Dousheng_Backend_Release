package utils

import (
	"fmt"
	"testing"
)

func TestCryptPassword(t *testing.T) {
	//salt := GenerateRandomSalt()
	password := "qwerasdf"
	salt := "Y1wlIS1aWOIl4mQPUBzcAA=="
	passwdCrypt := GenerateSaltedMD5(password, salt)
	fmt.Printf("%v\n", passwdCrypt)
	fmt.Printf("%v\n", len(passwdCrypt))

	fmt.Printf("%v\n", ContainsInvalidCharacters("qwerasdf~"))
}
