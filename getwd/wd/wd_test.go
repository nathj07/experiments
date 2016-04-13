package wd

import (
	"fmt"
	"os"
	"testing"
)

func TestWD(t *testing.T) {
	d, _ := os.Getwd()
	fmt.Println(d)
}
