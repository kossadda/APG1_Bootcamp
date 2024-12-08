package wc

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
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

func Output(args []string, w WC) chan string {
	ch := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(args))

		for _, path := range args {
			go func() {
				defer wg.Done()
				ch <- fmt.Sprintf("%s\t%s", w.fileInfo(path), path)
			}()
		}

		wg.Wait()
		close(ch)
	}()

	return ch
}

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
		line, ok := reader.ReadString('\n')
		err := func() string {
			if ok != nil {
				return ok.Error()
			}
			return ""
		}

		if ok != nil && ok.Error() != "EOF" {
			fmt.Fprintln(os.Stderr, "Error while reading", ok)
			return ""
		}

		switch wc & allMask {
		case lMask:
			if err() == "" {
				cnt++
			}
		case mMask:
			cnt += utf8.RuneCountInString(line)
		default:
			cnt += len(strings.Fields(line))
		}

		if err() == "EOF" {
			break
		}
	}

	return fmt.Sprint(cnt)
}

func validPaths(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("enter one or more text files")
	}

	for _, path := range args {
		if info, err := os.Stat(path); info != nil && info.IsDir() {
			return fmt.Errorf("'%s' is directory", path)
		} else if os.IsNotExist(err) {
			return fmt.Errorf("file '%s' is not exists", path)
		}
	}

	return nil
}
