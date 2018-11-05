package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"sync"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "hoge")
	}
	flag.Parse()
}

func main() {
	xlsx2csv()
	os.Exit(0)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func xlsx2csv() {
	var err error
	xlsxFile, err := xlsx.OpenFile(flag.Arg(0))
	check(err)

	wg := sync.WaitGroup{}
	for _, sheet := range xlsxFile.Sheets {
		wg.Add(1)
		go func(sheet *xlsx.Sheet) {
			defer wg.Done()
			createCsv(sheet)
		}(sheet)
	}
	wg.Wait()
}

func createCsv(sheet *xlsx.Sheet) {
	file, err := os.Create(sheet.Name + ".csv")
	check(err)
	w := csv.NewWriter(file)
	defer w.Flush()
	for _, row := range sheet.Rows {
		out := make([]string, 0)
		for _, cell := range row.Cells {
			text := cell.String()
			out = append(out, text)
		}
		err := w.Write(out)
		check(err)
	}
}
