package thtml

import "html/template"

// This comes in handy when you need to add a span or strong or even an a attribute.
func (p *HTMLParser) Next(text string) *HTMLParser {
	p.Text += template.HTML(text)
	return p
}

func (p *HTMLParser) Close(attribute string) *HTMLParser {
	p.Text += template.HTML("</" + attribute + ">")
	return p
}
