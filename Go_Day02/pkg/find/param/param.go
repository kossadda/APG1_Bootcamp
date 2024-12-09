// Package param provides functionality for parsing command-line flags.
// It supports flags for filtering by file type (file, directory, symbolic link)
// and file extensions.
package param

import (
	"flag"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/response"
)

const (
	slMask  = 1 << iota // Bitmask for symbolic links filter
	dMask               // Bitmask for directories filter
	fMask               // Bitmask for files filter
	extMask             // Bitmask for file extension filter
)

// Param holds the parsed parameters, including the path to scan, file extension, and selected filters.
type Param struct {
	Path  string // Path to the directory to scan
	Ext   string // File extension to filter by (if -f flag is used)
	flags int8   // Bitfield holding which flags are set
}

// IsSetSl checks if the symbolic link filter is set.
func (p *Param) IsSetSl() bool {
	return p.flags&slMask != 0
}

// IsSetD checks if the directory filter is set.
func (p *Param) IsSetD() bool {
	return p.flags&dMask != 0
}

// IsSetF checks if the file filter is set.
func (p *Param) IsSetF() bool {
	return p.flags&fMask != 0
}

// IsSetExt checks if the file extension filter is set.
func (p *Param) IsSetExt() bool {
	return p.flags&extMask != 0
}

// New parses the provided command-line arguments and returns a Param struct
// containing the parsed filters and the directory path. It returns an error
// if the flags are invalid or incomplete.
func New(setName string, args []string) (*Param, error) {
	fs := flag.NewFlagSet(setName, flag.ContinueOnError)

	sl := fs.Bool("sl", false, "Set search pattern: symbolic link")
	f := fs.Bool("f", false, "Set search pattern: files")
	d := fs.Bool("d", false, "Set search pattern: directories")
	ext := fs.String("ext", "", "Set search pattern: file extensions (use with -f)")

	fs.Usage = func() {
		fmt.Fprint(os.Stderr, response.FindUsage(fs))
	}

	err := fs.Parse(args)
	if err != nil {
		return &Param{}, err
	}

	return parseFlags(fs, *sl, *d, *f, *ext)
}

// parseFlags processes the flags and sets the appropriate flags in the Param struct.
func parseFlags(fs *flag.FlagSet, sl, d, f bool, ext string) (*Param, error) {
	var p Param

	if sl {
		p.flags |= slMask
	}
	if f {
		p.flags |= fMask
	}
	if d {
		p.flags |= dMask
	}

	if ext != "" {
		if !f {
			return &Param{}, response.InvalidExtUse()
		}

		p.flags |= extMask
		p.Ext = "." + ext
	}

	if len(fs.Args()) != 1 {
		return &Param{}, response.InvalidArgument()
	}

	if func() bool {
		extUsed := false
		fs.Visit(func(f *flag.Flag) {
			if f.Name == "ext" {
				extUsed = true
			}
		})
		return extUsed
	}() && ext == "" {
		return &Param{}, response.EmptyExt()
	}

	p.Path = fs.Args()[0]
	if p.flags == 0 {
		p.flags = dMask | fMask | slMask
	}

	return &p, nil
}
