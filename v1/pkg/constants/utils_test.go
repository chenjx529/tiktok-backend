package constants

import (
	"fmt"
	"testing"
)

func TestGetOutBoundIP(t *testing.T) {
	ip, _ := GetOutBoundIP()
	fmt.Println(ip)
}