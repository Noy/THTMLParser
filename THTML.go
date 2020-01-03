package thtml

import "html/template"

type HTMLParser struct {
	Body template.HTML
}

func NewHTMLParser() *HTMLParser{
	return &HTMLParser{""}
}

func (p *HTMLParser) Div(class, id, text string) *HTMLParser {
	p.Body += template.HTML("<div class=\""+class+"\" id=\""+id+"\">" + text)
	return p
}

func (p *HTMLParser) A(link, text, target string) *HTMLParser {
	p.Body += template.HTML("<a target=\""+target+"\" href=\""+link+"\">"+text)
	return p
}

func (p *HTMLParser) H2(class, id, text string) *HTMLParser {
	p.Body += template.HTML("<h2 class=\""+class+"\" id=\""+id+"\">" + text)
	return p
}

func (p *HTMLParser) H1(class, id, text string) *HTMLParser {
	p.Body += template.HTML("<h1 class=\""+class+"\" id=\""+id+"\">" + text)
	return p
}

func (p *HTMLParser) H3(class, id, text string) *HTMLParser {
	p.Body += template.HTML("<h3 class=\""+class+"\" id=\""+id+"\">" + text)
	return p
}

func (p *HTMLParser) H4(class, id, text string) *HTMLParser {
	p.Body += template.HTML("<h4 class=\""+class+"\" id=\""+id+"\">" + text)
	return p
}

func (p *HTMLParser) Strong(text string) *HTMLParser {
	p.Body += template.HTML("<strong>"+ text)
	return p
}

func (p *HTMLParser) Center() *HTMLParser {
	p.Body += "<center>"
	return p
}