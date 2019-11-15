package panic

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		// fmt.Println("Finally!")
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("start")
	panic(errors.New("something wrong!"))
	// os.Exit(-1)
}
