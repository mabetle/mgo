package render

import (
	"log"
	"os"

	"github.com/mabetle/mgo/mcore/mprint"
)

type FileRender struct {
	fpath string
}

func NewFileRender(fpath string) Render {
	return &FileRender{fpath: fpath}
}

func (r FileRender) Rend(value interface{}) {
	out, err := os.Create(r.fpath)
	if err != nil {
		log.Printf("open file error: %v", err)
	}
	defer out.Close()
	mprint.Fprint(out, value)
}

func FileRend(fpath string, value interface{}) {
	NewFileRender(fpath).Rend(value)
}
