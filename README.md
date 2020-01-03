# tSQL
Basic HTML parser written in go using string concat.

### Example

```go
package main

import (
	"github.com/Noy/thtml"
    "log"
)

func main() {
    parser := thtml.NewHTMLParser()
    //<div class="row"><div class="col-lg-12 col-sm-12 col-xs-12">..
    //..<div class="test id="nice-div"><center>
    //..<h2 class="second-title" id="second-title-id">..
    //..🎉 Congratulations! 🎉</h2></center></div></div></div>
     html := parser.Div("row", "", "").
    	  Div("col-lg-12 col-sm-12 col-xs-12", "", "").
    	  Div("test", "nice-div", "").
    	  Center().
          H2("second-title", "second-title-id", "🎉 Congratulations! 🎉").
    	  Close("h2").Close("center").Close("div").
          Close("div").Close("div").Body
    log.Println(html)	
}

```

#### Still needs a lot of work, adding other HTML tags soon...


