package resource

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.ResourceSetup = setup
	internal.ResourceCleanup = cleanup
}

////////////////////////////////////////////////////////////////////////////////

func setup() error {
	internal.Log.Println("Loading resources...")
	already := map[string]bool{}

	for _, s := range sources {
		switch s := s.(type) {

		case pack:
			for _, f := range s.Reader.File {
				if already[f.Name] {
					continue
				}
				already[f.Name] = true
				n, k, t, e := splitpath(f.Name)
				h := handlers[k]
				if h != nil {
					r, err := f.Open()
					if err != nil {
						return internal.Wrap("resource walk (in packed source)", err)
					}
					err = h(n, t, e, r)
					if err != nil {
						return internal.Wrap("resource walk (in packed source)", err)
					}
					err = r.Close()
					if err != nil {
						return internal.Wrap("resource walk (in packed source)", err)
					}
				}
			}

		case path:
			err := filepath.Walk(string(s), func(p string, inf os.FileInfo, err error) error {
				if inf.IsDir() {
					return nil
				}
				pp := strings.TrimPrefix(p, string(s)) //TODO: something more robust
				if already[pp] {
					return nil
				}
				already[pp] = true
				n, k, t, e := splitpath(pp)
				h := handlers[k]
				if h != nil {
					r, err := os.Open(p)
					if err != nil {
						return internal.Wrap("resource walk (in packed source)", err)
					}
					err = h(n, t, e, r)
					if err != nil {
						return internal.Wrap("resource walk (in packed source)", err)
					}
				}
				return nil
			})
			if err != nil {
				return internal.Wrap("resource walk (in packed source)", err)
			}

		default:
			panic("should not be here") //TODO:
		}
	}
	return nil
}

func splitpath(p string) (name string, kind string, tags []string, ext string) {
	tags = strings.Split(p, ".")
	switch len(tags) {
	case 0, 1:
		return p, "", []string{}, ""
	case 2:
		return tags[0], "", []string{}, tags[1]
	case 3:
		return tags[0], tags[1], []string{}, tags[2]
	default:
		return tags[0], tags[1], tags[2 : len(tags)-1], tags[len(tags)-1]
	}
}

////////////////////////////////////////////////////////////////////////////////

func cleanup() error {
	return nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
