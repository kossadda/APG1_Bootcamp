package archiver

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kossadda/APG1_Bootcamp/pkg/message"
)

type Archiver string

func (a Archiver) RotateFiles(files []string) {
	erch := make(chan error)
	for _, path := range files {
		go func(file string) {
			erch <- a.RotateFile(path)
		}(path)
	}

	for range files {
		if err := <-erch; err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		}
	}
}

func (a Archiver) RotateFile(file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return message.FailedArchive(err)
	}

	if info.IsDir() {
		return message.IsDirectory(file)
	}

	filename := strings.TrimSuffix(info.Name(), filepath.Ext(file))
	archName := fmt.Sprintf("%s_%d.tar.gz", filename, info.ModTime().Unix())
	archAbs := filepath.Join(string(a), archName)

	archiveFile, err := os.Create(archAbs)
	if err != nil {
		return message.FailedArchive(err)
	}
	defer archiveFile.Close()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	srcFile, err := os.Open(file)
	if err != nil {
		return message.FailedArchive(err)
	}
	defer srcFile.Close()

	header := &tar.Header{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    int64(info.Mode()),
		ModTime: info.ModTime(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		return message.FailedArchive(err)
	}

	if _, err := io.Copy(tarWriter, srcFile); err != nil {
		return message.FailedArchive(err)
	}

	return nil
}

func New(args *[]string) (*Archiver, error) {
	fs := flag.NewFlagSet("archiver", flag.ContinueOnError)

	path := fs.String("a", "", "Create archive in specified folder")

	fs.Usage = func() {
		fmt.Fprint(os.Stderr, message.ArchiveUsage(fs))
	}

	err := fs.Parse(*args)
	*args = fs.Args()
	if *path == "" {
		*path = "."
	}

	if err != nil {
		return nil, err
	}

	if info, err := os.Stat(*path); os.IsNotExist(err) {
		return nil, message.NotExists(*path)
	} else if !info.IsDir() {
		return nil, message.IsNotDirectory(*path)
	}

	if len(*args) == 0 {
		return nil, message.InvalidArgument()
	}

	return (*Archiver)(path), nil
}
