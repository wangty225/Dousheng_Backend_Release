package load

import (
	"fmt"
	"testing"
)

func TestInitMysqlConfig(t *testing.T) {
	fmt.Printf("test mysql dsn: %v \n", InitMysqlConfig().GetDSN())
}
