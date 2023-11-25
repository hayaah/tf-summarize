package parser

import (
	"strings"

	"github.com/hayaah/tf-summarize/reader"
	"github.com/hayaah/tf-summarize/terraformstate"
)

type Parser interface {
	Parse() (terraformstate.TerraformState, error)
}

func CreateParser(data []byte, fileName string) (Parser, error) {
	if fileName != reader.StdinFileName && !strings.HasSuffix(fileName, ".json") {
		return NewBinaryParser(fileName), nil
	}
	return NewJSONParser(data), nil
}
