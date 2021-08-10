package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Args = []string{"cmd", "--number-of-groups=100", "--batch-size=20", "--interval=3", "--output-directory=C:/Users/jtolmay/Desktop"}
	generateEvents()
}
