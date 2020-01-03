package thtml

import "html/template"

type HTMLParser struct {
	Text template.HTML
}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{""}
}

func (p *HTMLParser) Body(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<body class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) Div(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<div class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) A(link, text, target string) *HTMLParser {
	p.Text += template.HTML("<a target=\"" + target + "\" href=\"" + link + "\">" + text)
	return p
}

func (p *HTMLParser) H2(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<h2 class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) H1(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<h1 class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) H3(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<h3 class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) H4(class, id, text string) *HTMLParser {
	p.Text += template.HTML("<h4 class=\"" + class + "\" id=\"" + id + "\">" + text)
	return p
}

func (p *HTMLParser) Strong(text string) *HTMLParser {
	p.Text += template.HTML("<strong>" + text)
	return p
}

func (p *HTMLParser) Center() *HTMLParser {
	p.Text += "<center>"
	return p
}
