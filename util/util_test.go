package util

import (
	"fmt"
	"testing"
)

func TestU(t *testing.T) {
	filec, err := WriteFile("./test.txt", "test")
	if err != nil {
		fmt.Println(err)
		return
	}

	filec.Close()
	Interval(func() bool {
		if FileValid(filec) {
			return true
		}
		return false
	}, 10)

	// Delay(func(){
		
	// },10)
}
