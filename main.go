package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/gocolly/colly"
	"github.com/rivo/tview"
)

func main() {

	var exchanges = []string{
		"koineks",
		"btcturk",
		"ovis",
		"paribu",
		"vebitcoin",
		"koinim",
		"bithesap",
		"sistemkoin",
	}

	app := tview.NewApplication()
	cols := tview.NewTable().SetSeparator(tview.GraphicsVertBar)
	cols.SetBorder(true).SetTitle("Turkish Cryptocurrency Market Capitalizations").SetTitleColor(tcell.ColorWhite)

	loadData(cols, exchanges)

	err := app.SetRoot(cols, true).Run()

	if err != nil {
		panic(err)
	}

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func loadData(t *tview.Table, exchanges []string) {

	t.SetCell(0, 0, tview.NewTableCell("exchange").SetTextColor(tcell.ColorYellow))
	t.SetCell(0, 1, tview.NewTableCell("capital").SetTextColor(tcell.ColorYellow))

	list := dataList(exchanges)

	for i := 0; i < len(list); i++ {
		t.SetCell(i+1, 0, tview.NewTableCell(list.ToMarket(i)).SetTextColor(tcell.ColorDarkCyan))
		t.SetCell(i+1, 1, tview.NewTableCell(list.ToCap(i)).SetAlign(tview.AlignRight))
	}

	t.SetCell(len(list)+1, 0, tview.NewTableCell(""))
	t.SetCell(len(list)+1, 1, tview.NewTableCell(""))

	t.SetCell(len(list)+2, 0, tview.NewTableCell("total ").SetAlign(tview.AlignRight).SetTextColor(tcell.ColorDarkRed))
	t.SetCell(len(list)+2, 1, tview.NewTableCell(list.Total()))
}

func dataList(exchanges []string) (list capdata) {

	c := colly.NewCollector()
	c.DisableCookies()

	for i := 0; i < len(exchanges); i++ {

		d := data{}
		name := exchanges[i]
		cap, capstr, err := getCapital(c, name)

		d.market = name
		d.sizeStr = capstr

		if err != nil {
			d.size = -1
		} else {
			d.size = cap
		}

		list = append(list, d)
	}

	sort.Sort(list)

	return
}

func getCapital(c *colly.Collector, exchange string) (cap float64, capStr string, err error) {

	var currencyValue = ""

	c.OnHTML("span[data-currency-value]", func(e *colly.HTMLElement) {
		capStr = "   " + e.Text
		currencyValue = strings.TrimPrefix(e.Text, "$")
		currencyValue = strings.Replace(currencyValue, ",", "", -1)
	})

	err = c.Visit("https://coinmarketcap.com/exchanges/" + exchange + "/")

	if err != nil {
		return cap, capStr, err
	}

	cap, err = strconv.ParseFloat(currencyValue, 32)

	return cap, capStr, err
}
