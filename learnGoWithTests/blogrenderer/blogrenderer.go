package blogrenderer

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostViewModel struct {
	Title, Description string
	Body               template.HTML
	Tags               []string
}

func (p Post) String() string {
	return fmt.Sprintf("Title: %s; Description: %s; Tags: %v; Body: %s", p.Title, p.Description, p.Tags, p.Body)
}

func (p Post) SanitizeTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostViewModel(p))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func newPostViewModel(p Post) PostViewModel {
	return PostViewModel{
		Title:       p.Title,
		Description: p.Description,
		Body:        template.HTML(mdToHTML(p.Body)),
		Tags:        p.Tags}
}

func mdToHTML(markdownString string) string {
	p := parser.New()
	doc := p.Parse([]byte(markdownString))
	renderer := html.NewRenderer(html.RendererOptions{})
	return string(markdown.Render(doc, renderer))
}
