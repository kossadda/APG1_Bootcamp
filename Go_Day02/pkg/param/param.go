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

func New() (Param, error) {
	sl := flag.Bool("sl", false, "Set search pattern: symbolic link")
	f := flag.Bool("f", false, "Set search pattern: files")
	d := flag.Bool("d", false, "Set search pattern: directories")
	ext := flag.String("ext", "", "Set search pattern: file extensions (use with -f)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  -%s\t%s\n", f.Name, f.Usage)
		})
	}

	flag.Parse()

	return parseFlags(*sl, *d, *f, *ext)
}

func parseFlags(sl, d, f bool, ext string) (Param, error) {
	var p Param

	args := flag.Args()
	if len(args) != 1 {
		return p, fmt.Errorf("provide one path argument at the end")
	}

	p.Path = args[0]

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
			return p, fmt.Errorf("flag -ext provided but -f is not used")
		}

		p.flags |= extMask
		p.Ext = ext
	} else if func() bool {
		for _, arg := range os.Args {
			if arg == "-ext" {
				return true
			}
		}

		return false
	}() {
		return p, fmt.Errorf("flag -ext provided but extension is empty")
	}

	return p, nil
}
