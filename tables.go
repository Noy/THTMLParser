package thtml

func (p *HTMLParser) Table(class, id string) *HTMLParser {
	p.generateTag("table", class, id, "")
	return p
}

func (p *HTMLParser) THead(class, id string) *HTMLParser {
	p.generateTag("thead", class, id, "")
	return p
}

func (p *HTMLParser) Tr(class, id string) *HTMLParser {
	p.generateTag("tr", class, id, "")
	return p
}

func (p *HTMLParser) Td(class, id, text string) *HTMLParser {
	p.generateTag("td", class, id, text)
	return p
}

func (p *HTMLParser) Th(class, id, text string) *HTMLParser {
	p.generateTag("th", class, id, text)
	return p
}

func (p *HTMLParser) TBody() *HTMLParser {
	p.Text += "<tbody>"
	return p
}
