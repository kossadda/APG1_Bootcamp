package find

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func Scan(prm *param.Param) ([]string, error) {
	var sys []string

	err := filepath.WalkDir(prm.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil || path == prm.Path {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "%s: ‘%s’: No such file or directory\n", os.Args[0], prm.Path)
			}
			return nil
		}

		res := item(*prm, path, d)
		if res != "" {
			if filepath.IsAbs(path) {
				sys = append(sys, res) // Добавляем результат в слайс sys
			} else {
				if prm.Path[0:2] == "./" {
					sys = append(sys, "./"+res)
				} else {
					sys = append(sys, res)
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return sys, nil
}

func item(p param.Param, path string, d fs.DirEntry) string {
	if p.IsSetF() && fileFilter(p, path, d) {
		return path
	}
	if p.IsSetD() && folderFilter(d) {
		return path
	}
	if p.IsSetSl() && symlinkFilter(d) {
		realPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return path + " -> [broken]"
		} else {
			return path + " -> " + realPath
		}
	}

	return ""
}

func fileFilter(p param.Param, path string, d fs.DirEntry) bool {
	if d.Type().IsRegular() {
		if p.IsSetExt() {
			return filepath.Ext(path) == p.Ext
		}
		return true
	}
	return false
}

func folderFilter(d fs.DirEntry) bool {
	return d.IsDir()
}

func symlinkFilter(d fs.DirEntry) bool {
	return d.Type()&os.ModeSymlink != 0
}
