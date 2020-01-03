package thtml

import (
	"html/template"
	"testing"
)

func TestHTMLParser_Div(t *testing.T) {
	desiredLine := template.HTML("<div class=\"test\" id=\"test\">This is a test div</div>")
	newHTMLParser := NewHTMLParser()
	newHTMLParser.Div("test", "test", "This is a test div").Close("div")
	if desiredLine == newHTMLParser.Body {
		t.Log("This worked: " + newHTMLParser.Body)
	} else {
		t.Error("Gone wrong, check why")
	}
}