package file

import "strings"

func FileName(path string) string {
	if i := strings.LastIndex(path, "/"); i >= 0 {
		return path[i:]
	} else {
		return path
	}
}
