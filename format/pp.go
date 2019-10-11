package format

import (
	"io"

	"github.com/k0kubun/pp"
	"github.com/morikuni/failure"
)

type ppFormatter struct {
	w     io.Writer
	color bool
}

func (f *ppFormatter) Format(v interface{}) error {
	if !f.color {
		pp.ColoringEnabled = false
		defer func() {
			pp.ColoringEnabled = true
		}()
	}
	if _, err := pp.Fprintln(f.w, v); err != nil {
		return failure.Wrap(err)
	}
	return nil
}

func NewPP(w io.Writer, color bool) Formatter {
	return &ppFormatter{
		w:     w,
		color: color,
	}
}
