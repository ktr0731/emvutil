package format

import (
	"encoding/json"
	"io"

	"github.com/morikuni/failure"
)

type jsonFormatter struct {
	enc *json.Encoder
}

func (p *jsonFormatter) Format(v interface{}) error {
	if err := p.enc.Encode(v); err != nil {
		return failure.Wrap(err)
	}
	return nil
}

func NewJSON(w io.Writer) Formatter {
	return &jsonFormatter{
		enc: json.NewEncoder(w),
	}
}
