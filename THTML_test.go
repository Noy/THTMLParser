package thtml

import (
	"html/template"
	"testing"
)

func TestHTMLParser_Div(t *testing.T) {
	desiredLine := template.HTML("<div class=\"test\" id=\"test\">This is a test div</div>")
	newHTMLParser := NewHTMLParser()
	newHTMLParser.Div("test", "test", "This is a test div").Close("div")
	if desiredLine == newHTMLParser.Text {
		t.Log("This worked: " + newHTMLParser.Text)
	} else {
		t.Error("Gone wrong, check why")
	}
}

func TestHTMLParser_Table(t *testing.T) {
	table := template.HTML("<table><thead><tr><th>Header 1</th><th>Header 2</th></tr></thead><tbody><tr class=\"row1\">" +
		"<td>Test 1</td><td>Test 2</td></tr></tbody></table>")
	tHTMLTable := NewHTMLParser().Table("", "").THead().Tr("", "").Th("", "", "Header 1").
		Close("th").Th("", "", "Header 2").Close("tr").Close("thead").TBody().Tr("row1", "").
		Td("", "", "Test 1").Close("td").Td("", "", "Test 2").Close("td").Close("tr").
		Close("tbody").Close("table")
	if table == tHTMLTable.Text {
		t.Log("This worked: " + tHTMLTable.Text)
	} else {
		t.Error("Gone wrong, check why." + tHTMLTable.Text)
	}
}
