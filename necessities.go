package thtml

import "html/template"

func (p *HTMLParser) Body(class, id, text string) *HTMLParser {
	p.generateTag("body", class, id, text)
	return p
}

func (p *HTMLParser) Div(class, id, text string) *HTMLParser {
	p.generateTag("div", class, id, text)
	return p
}

func (p *HTMLParser) A(link, text, target string) *HTMLParser {
	p.Text += template.HTML("<a target=\"" + target + "\" href=\"" + link + "\">" + text)
	return p
}

func (p *HTMLParser) H1(class, id, text string) *HTMLParser {
	p.generateTag("h1", class, id, text)
	return p
}

func (p *HTMLParser) H2(class, id, text string) *HTMLParser {
	p.generateTag("h2", class, id, text)
	return p
}

func (p *HTMLParser) H3(class, id, text string) *HTMLParser {
	p.generateTag("h3", class, id, text)
	return p
}

func (p *HTMLParser) H4(class, id, text string) *HTMLParser {
	p.generateTag("h4", class, id, text)
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
