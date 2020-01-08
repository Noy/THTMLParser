package thtml

import "html/template"

type HTMLParser struct {
	Text template.HTML
}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{""}
}

func (p *HTMLParser) generateTag(tag, class, id, text string) template.HTML {
	if class == "" && id == "" {
		p.Text += template.HTML("<" + tag + ">" + text)
		return p.Text
	} else if id == "" && class != "" {
		p.Text += template.HTML("<" + tag + " class=\"" + class + "\">" + text)
		return p.Text
	} else if class == "" && id != "" {
		p.Text += template.HTML("<" + tag + " id=\"" + id + "\">" + text)
		return p.Text
	} else {
		p.Text += template.HTML("<" + tag + " class=\"" + class + "\" id=\"" + id + "\">" + text)
		return p.Text
	}
}
