package utils

import (
	"path"
	"regexp"
	"strings"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func ParsePath(p string) (pp []string) {
	a := regexp.MustCompile(`/+`)
	p = path.Clean("/" + p)
	pp = a.Split(p[1:]+"/", -1)
	return pp
}

