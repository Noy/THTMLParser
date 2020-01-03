package thtml

import "html/template"

// This comes in handy when you need to add a span or strong or even an a attribute.
func (p *HTMLParser) Next(text string) *HTMLParser {
	p.Body += template.HTML(text)
	return p
}

func (p *HTMLParser) Close(attribute string) *HTMLParser {
	p.Body += template.HTML("</"+attribute+">")
	return p
}