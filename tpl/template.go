package tpl

import "fmt"

//Template defines a snippet template
type Template interface {
	TplString() string
}

//NewStringTemplate returns a new template
func NewStringTemplate(template []byte) StringTemplate {
	return StringTemplate{
		content: string(template),
	}
}

//StringTemplate is the standard template implementation
type StringTemplate struct {
	content string
}

//TplString returns the templates result string
func (s StringTemplate) TplString() string {
	return s.content + "\n"
}

//NewFormatTemplate returns a new formattable template
func NewFormatTemplate(template []byte) FormatTemplate {
	return FormatTemplate{
		content: string(template),
	}
}

//FormatTemplate is the standard template implementation
type FormatTemplate struct {
	content string
	params  []interface{}
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
