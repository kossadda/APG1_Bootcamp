// Package wc provides the implementation for counting lines, words, or characters in files.
// It handles the reading of files and performs counting based on user-defined flags (-l, -w, -m).
// It also validates input and processes files concurrently using goroutines.
package wc

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

// Constants representing the flag masks for line, word, and character counts.
const (
	lMask   = 1 << iota             // Flag for counting lines
	mMask                           // Flag for counting characters (runes)
	wMask                           // Flag for counting words
	allMask = lMask | mMask | wMask // All flags combined
)

// WC type represents a bitmask to specify the counting type (lines, words, characters).
type WC uint8

// fileInfo method processes the file and counts based on the flag set in WC.
// It reads the file line by line and applies the appropriate counting logic.
func (wc WC) fileInfo(filename string) string {
	file, er := os.Open(filename)
	if er != nil {
		fmt.Fprintln(os.Stderr, er)
		return ""
	}
	defer file.Close()

	cnt := 0
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			fmt.Fprintln(os.Stderr, "Error while reading", err)
			return ""
		}

		switch wc & allMask {
		case lMask:
			if !errors.Is(err, io.EOF) {
				cnt++
			}
		case mMask:
			cnt += utf8.RuneCountInString(line)
		default:
			cnt += len(strings.Fields(line))
		}

		if errors.Is(err, io.EOF) {
			break
		}
	}

	return fmt.Sprint(cnt)
}

// validPaths validates the input file paths.
// It ensures that files exist and are not directories.
func validPaths(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("enter one or more text files")
	}

	for _, path := range args {
		if info, err := os.Stat(path); info != nil && info.IsDir() {
			return fmt.Errorf("'%s' is directory", path)
		} else if os.IsNotExist(err) {
			return fmt.Errorf("file '%s' does not exists", path)
		}
	}

	return nil
}

// New initializes a WC type based on the provided command-line arguments.
// It processes flags, validates file paths, and ensures only one flag is set.
func New(name string, args *[]string) (wc WC, err error) {
	m := make(map[string]*bool)
	fs := flag.NewFlagSet(name, flag.ContinueOnError)

	m["l"] = fs.Bool("l", false, "show permission of file")
	m["m"] = fs.Bool("m", false, "number of str")
	m["w"] = fs.Bool("w", false, "number of symbols")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] [FILE/DIR]...\n", os.Args[0])
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  -%s\t%s\n", f.Name, f.Usage)
		})
	}

	if err = fs.Parse(*args); err != nil {
		return 0, errors.New("try again")
	}
	*args = fs.Args()
	if err := validPaths(*args); err != nil {
		return 0, err
	}

	for fl, ok := range m {
		switch {
		case fl == "l" && *ok:
			wc |= lMask
		case fl == "m" && *ok:
			wc |= mMask
		case fl == "w" && *ok:
			wc |= wMask
		}
	}

	if wc&(wc-1) != 0 {
		return 0, errors.New("multiple or no flags are set")
	}

	return wc, nil
}

// Output starts concurrent processing for each file.
// It creates a goroutine for each file, processes it, and sends the results to a channel.
func Output(args []string, w WC) chan string {
	ch := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(args))

		for _, path := range args {
			go func(p string) {
				defer wg.Done()
				ch <- fmt.Sprintf("%s\t%s", w.fileInfo(p), p)
			}(path)
		}

		wg.Wait()
		close(ch)
	}()

	return ch
}
