package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvCreator interface {
	NewCsv(f io.Writer)
}

func setupCsv(c CsvCreator) {
	if csvExists() {
		return
	}

	f, err := os.Create(csvPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}

	defer f.Close()

	c.NewCsv(f)
}

func csvExists() bool {
	if _, err := os.Stat(csvPath); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			fmt.Fprintf(os.Stderr, "error: %s", err)
		}
	}

	return true
}

func appendToCsv(f io.Writer, s []string) error {
	w := csv.NewWriter(f)
	w.Write(s)
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
