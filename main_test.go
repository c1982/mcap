package main

import (
	"fmt"
	"testing"
)

func TestFormatFloatLikeBoss(t *testing.T) {

	var f = 73793993.0
	fstr := seperateFloat(f)

	if fstr != "73,793,993" {
		t.Error("not expected value:", fstr)
	}

	fmt.Println(fstr)
}
