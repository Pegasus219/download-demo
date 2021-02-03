package services

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
)

func readFromCSV(filePath string) (data [][]string, err error) {
	cntb, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	r := csv.NewReader(strings.NewReader(string(cntb)))
	return r.ReadAll()
}

func writeIntoCSV(filePath string, data [][]string) (err error) {
	f, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err = w.WriteAll(data); err == nil {
		w.Flush()
	}
	return
}

func buildBufferCSV(data [][]string) (buffer bytes.Buffer, err error) {
	w := csv.NewWriter(&buffer)
	if err = w.WriteAll(data); err == nil {
		w.Flush()
	}
	return
}
