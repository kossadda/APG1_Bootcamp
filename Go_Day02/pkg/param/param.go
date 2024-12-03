package param

import (
	"flag"
	"fmt"
	"os"
)

const (
	slMask = 1 << iota
	dMask
	fMask
	extMask
)

type Param struct {
	Path  string
	Ext   string
	flags int8
}

func (p *Param) IsSetSl() bool {
	return p.flags&slMask != 0
}

func (p *Param) IsSetD() bool {
	return p.flags&dMask != 0
}

func (p *Param) IsSetF() bool {
	return p.flags&fMask != 0
}

func (p *Param) IsSetExt() bool {
	return p.flags&extMask != 0
}

func New(setName string, args []string) (*Param, error) {
	fs := flag.NewFlagSet(setName, flag.ContinueOnError)

	sl := fs.Bool("sl", false, "Set search pattern: symbolic link")
	f := fs.Bool("f", false, "Set search pattern: files")
	d := fs.Bool("d", false, "Set search pattern: directories")
	ext := fs.String("ext", "", "Set search pattern: file extensions (use with -f)")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  -%s\t%s\n", f.Name, f.Usage)
		})
	}

	err := fs.Parse(args)
	if err != nil {
		return &Param{}, err
	}

	return parseFlags(fs, *sl, *d, *f, *ext)
}

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
			return &Param{}, fmt.Errorf("flag -ext provided but -f is not used")
		}

		p.flags |= extMask
		p.Ext = "." + ext
	}

	if len(fs.Args()) != 1 {
		return &Param{}, fmt.Errorf("provide one path argument at the end")
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
		return &Param{}, fmt.Errorf("flag -ext provided but extension is empty")
	}

	p.Path = fs.Args()[0]
	if p.flags == 0 {
		p.flags = dMask | fMask | slMask
	}

	return &p, nil
}
