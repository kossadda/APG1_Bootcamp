package response

import (
	"flag"
	"fmt"
	"os"
)

type resp struct {
	err string
}

func FindUsage(fs *flag.FlagSet) error {
	err := resp{
		err: fsusage("[OPTION]... [FILE/DIR]", fs),
	}

	return err
}

func WCUsage(fs *flag.FlagSet) error {
	err := resp{
		err: fsusage("[OPTION] [FILE/DIR]...", fs),
	}

	return err
}

func fsusage(usage string, fs *flag.FlagSet) (str string) {
	str += fmt.Sprintf("Usage: %s %s\n", os.Args[0], usage)
	fs.VisitAll(func(f *flag.Flag) {
		str += fmt.Sprintf("  -%s\t%s\n", f.Name, f.Usage)
	})
	return str
}

func (e resp) Error() string {
	return e.err
}

func NotExists(filename string) error {
	err := resp{
		err: fmt.Sprintf("‘%s’: No such file or directory", filename),
	}

	return err
}

func InvalidExtUse() error {
	err := resp{
		err: "flag -ext provided but -f is not used",
	}

	return err
}

func InvalidArgument() error {
	err := resp{
		err: "provide path argument at the end",
	}

	return err
}

func InvalidFlag() error {
	err := resp{
		err: " provide only one flag",
	}

	return err
}

func EmptyExt() error {
	err := resp{
		err: "flag -ext provided but extension is empty",
	}

	return err
}

func EmptyCommand() error {
	err := resp{
		err: fmt.Sprintf("input command name for %s", os.Args[0]),
	}

	return err
}

func IsDirectory(path string) error {
	err := resp{
		err: fmt.Sprintf("'%s' is directory", os.Args[0]),
	}

	return err
}
