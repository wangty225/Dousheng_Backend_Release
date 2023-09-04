package load

import (
	"fmt"
	"testing"
)

func TestInitRedisConfig(t *testing.T) {
	r := InitRedisConfig()
	fmt.Printf("%+v\n", r)
}
