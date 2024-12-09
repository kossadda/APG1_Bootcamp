// Package message provides error-handling utilities and structured error messages
// for use across different command-line tools. It includes custom error types and
// usage messages for better user feedback.
package message

import (
	"flag"
	"fmt"
	"os"
)

// resp is a custom error type used to return structured error messages.
type resp struct {
	err string
}

// Error implements the error interface for the resp type.
func (e resp) Error() string {
	return e.err
}

// FindUsage generates and returns a usage error message for tools similar to `find`.
// Accepts a flag.FlagSet to include specific option descriptions in the message.
func FindUsage(fs *flag.FlagSet) error {
	return resp{err: fsusage("[OPTION]... [FILE/DIR]", fs)}
}

// WCUsage generates and returns a usage error message for tools similar to `wc`.
// Accepts a flag.FlagSet to include specific option descriptions in the message.
func WCUsage(fs *flag.FlagSet) error {
	return resp{err: fsusage("[OPTION] [FILE/DIR]...", fs)}
}

func ArchiveUsage(fs *flag.FlagSet) error {
	return resp{err: fsusage("[OPTION] [FILE/DIR]...", fs)}
}

// NotExists returns an error indicating that the specified file or directory does not exist.
// Takes the filename or directory name as a parameter.
func NotExists(filename string) error {
	return resp{err: fmt.Sprintf("‘%s’: No such file or directory", filename)}
}

// InvalidExtUse returns an error indicating that the `-ext` flag was used without the `-f` flag.
func InvalidExtUse() error {
	return resp{err: "flag -ext provided but -f is not used"}
}

// InvalidArgument returns an error indicating that the path argument was not provided at the end of the command.
func InvalidArgument() error {
	return resp{err: "provide path argument at the end"}
}

// InvalidFlag returns an error indicating that more than one exclusive flag was provided.
func InvalidFlag() error {
	return resp{err: "provide only one flag"}
}

// EmptyExt returns an error indicating that the `-ext` flag was provided but no extension was specified.
func EmptyExt() error {
	return resp{err: "flag -ext provided but extension is empty"}
}

// EmptyCommand returns an error indicating that no command name was provided in the input.
func EmptyCommand() error {
	return resp{err: fmt.Sprintf("input command name for %s", os.Args[0])}
}

// IsDirectory returns an error indicating that the specified path is a directory, not a file.
// Takes the path as input.
func IsDirectory(path string) error {
	return resp{err: fmt.Sprintf("'%s' is directory", path)}
}

func IsNotDirectory(path string) error {
	return resp{err: fmt.Sprintf("'%s' is not directory", path)}
}

func EmptyDirectory() error {
	return resp{err: "-a provided by empty directory"}
}

func FailedArchive(err error) error {
	return resp{err: fmt.Sprintf("failed to write header to archive: %s", err)}
}

// fsusage generates a formatted usage string based on the given usage format and FlagSet.
// This is an internal utility function used by FindUsage and WCUsage to build error messages.
// Usage example (not directly called):
//
//	str := fsusage("[OPTION] [FILE/DIR]...", fs)
func fsusage(usage string, fs *flag.FlagSet) (str string) {
	str += fmt.Sprintf("Usage: %s %s\n", os.Args[0], usage)
	fs.VisitAll(func(f *flag.Flag) {
		str += fmt.Sprintf("  -%s\t%s\n", f.Name, f.Usage)
	})
	return str
}
