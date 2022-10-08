package fsutil

import (
	"os"
	"path"
)

func DirWalk(p string) ([]string, error) {
	res := make([]string, 0)
	l, err := os.ReadDir(p)

	if err != nil {
		return nil, err
	}

	for _, i := range l {
		if i.IsDir() {
			l2, err := DirWalk(path.Join(p, i.Name()))
			if err != nil {
				return nil, err
			}

			res = append(res, l2...)
		} else {
			res = append(res, path.Join(p, i.Name()))
		}
	}

	return res, nil
}
