package find

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func Start(prm *param.Param) error {
	info, err := os.Stat(prm.Path)
	if err != nil {
		return fmt.Errorf("directory %s does not exist", prm.Path)
	} else if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", prm.Path)
	}

	if prm.IsSetD() {
		fmt.Println(prm.Path)
	}
	return filepath.WalkDir(prm.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err = d.Info()
		if err != nil {
			return err
		}

		if prm.IsNoFlags() {
			printSymlink(info, path)
			printDir(info, path)
			printFile(info, path)
		} else {
			if prm.IsSetSl() {
				printSymlink(info, path)
			}
			if prm.IsSetD() {
				printDir(info, path)
			}
			if prm.IsSetExt() {
				printFileWithExt(info, path, prm.Ext)
			} else {
				printFile(info, path)
			}
		}

		return nil
	})
}

func printFileWithExt(info fs.FileInfo, path, ext string) {
	if info.Mode()&os.ModeDir == 0 && info.Mode()&os.ModeSymlink == 0 && filepath.Ext(path) == ext {
		if !filepath.IsAbs(path) {
			fmt.Print("./")
		}
		fmt.Println(path)
	}
}

func printFile(info fs.FileInfo, path string) {
	if info.Mode()&os.ModeDir == 0 && info.Mode()&os.ModeSymlink == 0 {
		if !filepath.IsAbs(path) {
			fmt.Print("./")
		}
		fmt.Println(path)
	}
}

func printDir(info fs.FileInfo, path string) {
	if info.Mode()&os.ModeDir != 0 {
		if !filepath.IsAbs(path) {
			fmt.Print("./")
		}
		fmt.Println(path)
	}
}

func printSymlink(info fs.FileInfo, path string) {
	if info.Mode()&os.ModeSymlink != 0 {
		if !filepath.IsAbs(path) {
			fmt.Print("./")
		}

		realPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			fmt.Println(path, "-> [broken]")
		} else {
			fmt.Println(path, "->", realPath)
		}
	}
}
