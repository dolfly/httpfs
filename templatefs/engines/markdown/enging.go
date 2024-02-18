package markdown

import (
	"github.com/dolfly/httpfs/templatefs"
	"github.com/russross/blackfriday/v2"
)

type Engine struct {
	Extensions blackfriday.Extensions
	HtmlFlags  blackfriday.HTMLFlags
	CSS        string
}

func (e *Engine) Render(input []byte) []byte {
	return blackfriday.Run(input,
		blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: e.HtmlFlags})),
		blackfriday.WithExtensions(e.Extensions))
}

func (e *Engine) PageInfo(input []byte) *templatefs.Page {
	return &templatefs.Page{Title: getTitle(input), CSS: e.CSS}
}

func (e *Engine) Exts() []string {
	return []string{".md", ".markdown"}
}
