package goecmserver

import "net/http"

type GinFS struct {
	FS http.FileSystem
}

func (g *GinFS) Open(name string) (http.File, error) {
	return g.FS.Open(name)
}

func (g *GinFS) Exists(prefix string, filepath string) bool {
	if _, err := g.FS.Open(filepath); err != nil {
		return false
	}
	return true
}
