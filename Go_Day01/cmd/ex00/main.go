// Package main is the entry point for the application.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/readdb"
)

func main() {
	var reader readdb.DBReader
	filename := flag.String("f", "", "Filename to read ('xml' or 'json')")
	flag.Parse()

	file, err := readdb.DefineFile(&reader, *filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	if err = reader.DBRead(file); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(reader)
}
