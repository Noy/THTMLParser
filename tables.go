package thtml

import "html/template"

func (p *HTMLParser) Table(class, id string) *HTMLParser {
	p.Text += template.HTML("<table class=\"" + class + "\" id=\"" + id + "\">")
	return p
}

func (p *HTMLParser) THead() *HTMLParser {
	p.Text += "<thead>"
	return p
}

func (p *HTMLParser) Tr(class, id string) *HTMLParser {
	p.Text += template.HTML("<tr class=\"" + class + "\" id=\"" + id + "\">")
	return p
}

func (p *HTMLParser) Td(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<td class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) Th(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<td class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) TBody() *HTMLParser {
	p.Text += "<tbody>"
	return p
}
