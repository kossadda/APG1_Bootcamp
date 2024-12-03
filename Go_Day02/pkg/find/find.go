package find

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func Scan(prm *param.Param) error {
	info, err := os.Stat(prm.Path)
	if err != nil {
		return fmt.Errorf("directory %s does not exist", prm.Path)
	} else if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", prm.Path)
	}

	return filepath.WalkDir(prm.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil && path == prm.Path {
			return nil
		}

		res := item(*prm, path, d)
		if res != "" {
			if filepath.IsAbs(path) {
				fmt.Println(res)
			} else {
				fmt.Println("./" + res)
			}
		}

		return nil
	})
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
