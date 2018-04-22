package main

import (
	"fmt"
	"strings"
)

type data struct {
	ID        int
	CreatedAt int64
	Market    string
	Size      float64
	SizeStr   string
}

type capdata []data

func (c capdata) Len() int              { return len(c) }
func (c capdata) Swap(i, j int)         { c[i], c[j] = c[j], c[i] }
func (c capdata) Less(i, j int) bool    { return c[i].Size > c[j].Size }
func (c capdata) ToCap(i int) string    { return c[i].SizeStr }
func (c capdata) ToMarket(i int) string { return fmt.Sprintf("%s   ", c[i].Market) }

func (c capdata) Total() string {

	var total float64

	for _, v := range c {
		total += v.Size
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
