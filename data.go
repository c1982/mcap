package main

import (
	"fmt"
	"strings"
)

type data struct {
	market  string
	size    float64
	sizeStr string
}

type capdata []data

func (c capdata) Len() int              { return len(c) }
func (c capdata) Swap(i, j int)         { c[i], c[j] = c[j], c[i] }
func (c capdata) Less(i, j int) bool    { return c[i].size > c[j].size }
func (c capdata) ToCap(i int) string    { return c[i].sizeStr }
func (c capdata) ToMarket(i int) string { return fmt.Sprintf("%s   ", c[i].market) }

func (c capdata) Total() string {

	var total float64

	for _, v := range c {
		total += v.size
	}

	return " $" + seperateFloat(total)
}

func seperateFloat(f float64) (currency string) {

	var fstr = fmt.Sprintf("%0.f", f)
	var offset = 3

	for i := len(fstr); i > 0; i -= 3 {
		if i < 3 {
			offset = i
		}

		sliceText := fstr[i-offset : i]
		currency = strings.TrimSuffix(sliceText+","+currency, ",")
	}

	return
}
