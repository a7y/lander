package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/structs"
)

type Subscriber struct {
	Email string
	Host  string
	When  string
}

func (s Subscriber) NewCsv(f io.Writer) {
	err := appendToCsv(f, s.fields())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}
}

func (s Subscriber) Save(csvPath string) error {
	f, err := os.OpenFile(csvPath, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	err = appendToCsv(f, s.values())
	if err != nil {
		return err
	}

	return nil
}

func (s Subscriber) fields() []string {
	return structs.Names(s)
}

func (s Subscriber) values() []string {
	values := structs.Values(s)

	strings := make([]string, len(values))
	for i, v := range values {
		strings[i] = v.(string)
	}

	return strings
}
