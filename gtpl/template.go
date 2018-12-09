package gtpl

import (
	"fmt"
	"strings"
)

//Template defines a snippet template
type Template interface {
	ID() string
	TplString() string
}

//NewStringTemplate returns a new template
func NewStringTemplate(id string, assetFunc func(string) []byte) StringTemplate {
	return StringTemplate{
		id:      id,
		content: string(assetFunc(id)),
	}
}

//NewStringTemplateSlice returns a slice of templates generated from all assets
//matching a given prefix (for example "tips/" to get all tips).
func NewStringTemplateSlice(prefix string, assetFunc func(string) []byte) []StringTemplate {
	var tpls = []StringTemplate{}
	for _, n := range AssetNames() {
		if !strings.HasPrefix(n, prefix) {
			continue
		}
		tpls = append(tpls, NewStringTemplate(n, assetFunc))
	}
	return tpls
}

//StringTemplate is the standard template implementation
type StringTemplate struct {
	id      string
	content string
}

//ID returns the templates identifier
func (s StringTemplate) ID() string {
	return s.id
}

//TplString returns the templates result string
func (s StringTemplate) TplString() string {
	return s.content + "\n"
}

//NewFormatTemplate returns a new formattable template
func NewFormatTemplate(id string, assetFunc func(string) []byte) FormatTemplate {
	return FormatTemplate{
		id:      id,
		content: string(assetFunc(id)),
	}
}

//FormatTemplate is the standard template implementation
type FormatTemplate struct {
	id      string
	content string
	params  []interface{}
}

//ID returns the templates identifier
func (s FormatTemplate) ID() string {
	return s.id
}

//Format adds formatting parameters to the template
func (s FormatTemplate) Format(params ...interface{}) FormatTemplate {
	s.params = params
	return s
}

//TplString returns the templates result string
func (s FormatTemplate) TplString() string {
	return fmt.Sprintf(s.content, s.params...) + "\n"
}
