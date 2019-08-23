package format

import (
	"io"

	"github.com/k0kubun/pp"
	"github.com/morikuni/failure"
)

type ppFormatter struct {
	w io.Writer
}

func (f *ppFormatter) Format(v interface{}) error {
	if _, err := pp.Fprintln(f.w, v); err != nil {
		return failure.Wrap(err)
	}
	return nil
}

func NewPP(w io.Writer) Formatter {
	return &ppFormatter{
		w: w,
	}
}
