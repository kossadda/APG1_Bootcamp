// Package find provides functionality for scanning directories based on specific filters
// such as file type, extension, directory, and symbolic links. The 'Scan' function walks
// through a directory and applies these filters to the files and directories it encounters.
package find

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kossadda/APG1_Bootcamp/pkg/find/param"
	"github.com/kossadda/APG1_Bootcamp/pkg/message"
)

// Scan walks through a directory (and its subdirectories) and filters items based on the provided
// parameters. It returns a slice of strings containing the paths of the matching files, directories,
// or symbolic links.
func Scan(prm *param.Param) ([]string, error) {
	var sys []string

	err := filepath.WalkDir(prm.Path, func(path string, d fs.DirEntry, err error) error {
		if os.IsNotExist(err) {
			return message.NotExists(prm.Path)
		} else if err != nil {
			return nil
		}

		res := item(*prm, path, d)
		if res != "" {
			if filepath.IsAbs(path) {
				sys = append(sys, res)
			} else {
				if len(prm.Path) > 1 && prm.Path[0:2] == "./" {
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

// item applies the necessary filters to each item in the directory
// and returns the path of the item if it matches the filter criteria.
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

// fileFilter checks if the item is a regular file and matches the extension filter.
func fileFilter(p param.Param, path string, d fs.DirEntry) bool {
	if d.Type().IsRegular() {
		if p.IsSetExt() {
			return filepath.Ext(path) == p.Ext
		}
		return true
	}
	return false
}

// folderFilter checks if the item is a directory.
func folderFilter(d fs.DirEntry) bool {
	return d.IsDir()
}

// symlinkFilter checks if the item is a symbolic link.
func symlinkFilter(d fs.DirEntry) bool {
	return d.Type()&os.ModeSymlink != 0
}
