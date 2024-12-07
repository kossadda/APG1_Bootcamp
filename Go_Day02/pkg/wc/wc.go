package wc

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sync"
	"unicode/utf8"
)

const (
	lMask = 1 << iota
	mMask
	wMask
	allMask = lMask | mMask | wMask
)

type WC uint8

func (wc WC) FileInfo(filename string) string {
	if mask := wc & allMask; mask == lMask || mask == mMask {
		return wc.fileReader(filename)
	} else {
		return wc.permission(filename)
	}
}

func (wc WC) fileReader(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	defer file.Close()

	cnt := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if wc&allMask == lMask {
			cnt++
		} else {
			cnt += utf8.RuneCountInString(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	return fmt.Sprint(cnt)
}

func (wc WC) permission(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	return info.Mode().String()
}

func New(name string, args *[]string) (wc WC, err error) {
	m := make(map[string]*bool)
	fs := flag.NewFlagSet(name, flag.ContinueOnError)

	m["l"] = fs.Bool("l", false, "show permission of file")
	m["m"] = fs.Bool("m", false, "number of str")
	m["w"] = fs.Bool("w", false, "number of symbols")

	if err = fs.Parse(*args); err != nil {
		return 0, errors.New("try again")
	}
	*args = fs.Args()
	if path, ok := validPaths(*args); !ok {
		return 0, fmt.Errorf("invalid filepath %s", path)
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

func validPaths(args []string) (string, bool) {
	for _, path := range args {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return path, false
		}
	}

	return "", true
}

func Output(args []string, w WC) chan string {
	ch := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(args))

		for _, path := range args {
			go func() {
				defer wg.Done()
				ch <- fmt.Sprintf("%s\t%s", w.FileInfo(path), path)
			}()
		}

		wg.Wait()
		close(ch)
	}()

	return ch
}
