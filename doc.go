// Copyright 2020 The golang.design Initiative.
// All rights reserved. Use of this source code
// is governed by a MIT license that can be found
// in the LICENSE file.

// Command history renders README.md and README_cn.md into the
// golang.design/history website. Each document is parsed once and
// split at its "##" headings into a home page and one page per part,
// so heading ids stay what they were on the old single-page site.
package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	ghtml "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

//go:embed layout.html
var layoutHTML string

var outDir = flag.String("out", "public", "output directory")

// parts lists the pages after the home page, in order, by the English
// "##" heading each one shows. Every other section stays on the home
// page. The Chinese document is paired with the English one by section
// index.
var parts = []struct{ slug, heading string }{
	{"people", "Committers"},
	{"language", "Language Design"},
	{"toolchain", "Compiler Toolchain"},
	{"runtime", "Runtime Core"},
	{"stdlib", "Standard Library"},
	{"links", "Unclassified But Relevant Links"},
}

// A language is one translation of the document.
type language struct {
	File  string     // markdown source
	Tag   string     // html lang attribute
	Dir   string     // output directory, relative to the site root
	Index string     // search index, relative to the site root
	IDs   parser.IDs // heading ids; nil means goldmark's default
	UI    ui
}

// ui holds the interface strings of a language.
type ui struct {
	Name, Menu, Theme, Search, Count, More            string
	Documentary, RuntimeCompiler, RuntimeCompilerLong string
	RuntimeCompilerDesc, Contents, Entries, Also      string
	Previous, Next, Updated, Source, Views, Visitors  string
}

var en = &language{
	File: "README.md", Tag: "en", Dir: "", Index: "search-en.json",
	UI: ui{
		Name: "English", Menu: "Menu", Theme: "Theme",
		Search: "Search entries", Count: "%d results", More: "first %d of %n results",
		Documentary: "Documentary", RuntimeCompiler: "Runtime & compiler",
		RuntimeCompilerLong: "Runtime & compiler history",
		RuntimeCompilerDesc: "Go 1.0 to 1.28, reconstructed from commits: every topic, every release, and the paths not taken.",
		Contents:            "Contents", Entries: "entries", Also: "Also on this site",
		Previous: "Previous", Next: "Next", Updated: "Last updated", Source: "Source on GitHub",
		Views: "Views", Visitors: "Visitors",
	},
}

var cn = &language{
	File: "README_cn.md", Tag: "zh-CN", Dir: "cn/", Index: "search-cn.json", IDs: newIDs(),
	UI: ui{
		Name: "中文", Menu: "菜单", Theme: "主题",
		Search: "搜索条目", Count: "%d 条结果", More: "共 %n 条结果，显示前 %d 条",
		Documentary: "文档", RuntimeCompiler: "运行时与编译器",
		RuntimeCompilerLong: "运行时与编译器演进史（英文）",
		RuntimeCompilerDesc: "从提交历史重建 Go 1.0 至 1.28 的演进：每个主题、每个版本，以及未选择的路。",
		Contents:            "目录", Entries: "条", Also: "本站其他内容",
		Previous: "上一页", Next: "下一页", Updated: "最后更新", Source: "GitHub 源码",
		Views: "浏览", Visitors: "访客",
	},
}

var md = goldmark.New(
	goldmark.WithExtensions(extension.Table),
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	goldmark.WithRendererOptions(ghtml.WithUnsafe()),
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("history: ")

	e := parse(en)
	c := parse(cn)
	if len(e.sections) != len(c.sections) {
		log.Fatalf("%s has %d \"##\" sections and %s has %d; they are paired by index, so add the missing ones",
			en.File, len(e.sections), cn.File, len(c.sections))
	}
	var keys []string
	for _, s := range e.sections {
		keys = append(keys, s.title)
	}
	e.paginate(keys)
	c.paginate(keys)

	tmpl := template.Must(template.New("layout").Parse(layoutHTML))
	updated := time.Now().UTC().Format("2006-01-02")
	for _, d := range []*doc{e, c} {
		other := c
		if d == c {
			other = e
		}
		for i, p := range d.pages {
			var b bytes.Buffer
			if err := tmpl.Execute(&b, d.view(p, other.pages[i], other.lang, updated)); err != nil {
				log.Fatal(err)
			}
			write(p.url+"index.html", b.Bytes())
		}
		write(d.lang.Index, d.searchIndex())
	}
	write("cn.html", c.redirectPage())
	fmt.Println("golang-design/history: A Documentary of Go")
}

// A doc is a parsed document.
type doc struct {
	lang     *language
	src      []byte
	title    string     // text of the "#" heading
	titleID  string     // its id, kept for links to the old single page
	desc     string     // first long paragraph, for the description meta tag
	preamble []ast.Node // between the "#" heading and the first section
	sections []*section
	pages    []*page
	pageOf   map[string]*page // heading id to the page that shows it
}

// A section is a "##" heading and everything up to the next one.
type section struct {
	heading *ast.Heading
	title   string
	nodes   []ast.Node
}

// A page is one output page.
type page struct {
	slug     string // "" for the home page
	title    string
	sections []*section
	url      string // relative to the site root; "" for the English home
	entries  int    // list items
}

func parse(l *language) *doc {
	src, err := os.ReadFile(l.File)
	if err != nil {
		log.Fatal(err)
	}
	ctx := parser.NewContext()
	if l.IDs != nil {
		ctx = parser.NewContext(parser.WithIDs(l.IDs))
	}
	root := md.Parser().Parse(text.NewReader(src), parser.WithContext(ctx))

	d := &doc{lang: l, src: src, pageOf: map[string]*page{}}
	inContents := false // the hand-written table of contents
	for n := root.FirstChild(); n != nil; n = n.NextSibling() {
		switch {
		case isHeading(n, 1):
			d.title = string(n.Text(src))
			d.titleID = id(n.(*ast.Heading))
			continue
		case isHeading(n, 2):
			d.sections = append(d.sections, &section{heading: n.(*ast.Heading), title: string(n.Text(src))})
			continue
		case (n.Kind() == ast.KindParagraph || n.Kind() == ast.KindHTMLBlock) && bytes.Contains(lines(n, src), []byte(`id="top"`)):
			inContents = true
			continue
		case inContents:
			inContents = n.Kind() != ast.KindThematicBreak
			continue
		case isBackToTop(n):
			continue
		}
		if len(d.sections) == 0 {
			d.preamble = append(d.preamble, n)
			if t := strings.Join(strings.Fields(string(n.Text(src))), " "); d.desc == "" && n.Kind() == ast.KindParagraph && len(t) > 60 {
				d.desc = t
			}
			continue
		}
		s := d.sections[len(d.sections)-1]
		s.nodes = append(s.nodes, n)
	}
	return d
}

// paginate assigns the sections to pages. keys are the English
// section titles, by index.
func (d *doc) paginate(keys []string) {
	home := &page{title: d.title, url: d.lang.Dir}
	d.pages = []*page{home}
	bySection := map[string]*page{}
	for _, p := range parts {
		pg := &page{slug: p.slug, url: d.lang.Dir + p.slug + "/"}
		bySection[p.heading] = pg
		d.pages = append(d.pages, pg)
	}
	for i, s := range d.sections {
		p := bySection[keys[i]]
		if p == nil {
			p = home
		} else {
			if len(p.sections) > 0 {
				log.Fatalf("%s: two sections titled %q", d.lang.File, keys[i])
			}
			p.title = s.title
		}
		p.sections = append(p.sections, s)
	}
	for _, p := range d.pages {
		if len(p.sections) == 0 && p != home {
			log.Fatalf("%s: no section for page %q", d.lang.File, p.slug)
		}
		for _, s := range p.sections {
			for _, h := range s.headings() {
				d.pageOf[id(h)] = p
			}
		}
	}
	d.pageOf[d.titleID] = home
	for _, p := range d.pages {
		d.rewrite(p)
	}
}

var refLabel = regexp.MustCompile(`^[a-z][a-z0-9.\-]*/\S+$`)

// rewrite points links to headings on other pages at those pages,
// labels list entries like issue/123, and, on part pages, raises
// headings one level since the "##" heading becomes the page title.
func (d *doc) rewrite(p *page) {
	for _, s := range p.sections {
		for _, n := range s.nodes {
			ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
				if !entering {
					return ast.WalkContinue, nil
				}
				switch n := n.(type) {
				case *ast.Heading:
					if p.slug != "" {
						n.Level--
					}
				case *ast.ListItem:
					p.entries++
				case *ast.Link:
					dest := string(n.Destination)
					if strings.HasPrefix(dest, "#") {
						if q := d.pageOf[dest[1:]]; q == nil {
							log.Printf("%s: link to missing heading %s", d.lang.File, dest)
						} else if q != p {
							n.Destination = []byte(rel(p, q) + dest)
						}
					}
					if isRef(n, d.src) {
						n.SetAttributeString("class", []byte("ref"))
					}
				}
				return ast.WalkContinue, nil
			})
		}
	}
}

// navItem is an entry of the sidebar or the pager.
type navItem struct {
	Title, URL string
	Current    bool
	Subs       []navItem
}

// view is the data for layout.html.
type view struct {
	Lang, Other               *language
	Title, Brand, Description string
	Base, HomeURL, OtherURL   string
	Nav                       []navItem
	Body                      template.HTML
	Prev, Next                *navItem
	Updated                   string
	Script                    template.JS
}

func (d *doc) view(p, alt *page, other *language, updated string) view {
	v := view{
		Lang: d.lang, Other: other, Brand: d.title, Description: d.desc,
		Base: strings.Repeat("../", depth(p)), HomeURL: rel(p, d.pages[0]), OtherURL: rel(p, alt),
		Body: template.HTML(d.body(p)), Updated: updated,
	}
	v.Title = p.title + " · " + d.title
	if p.slug == "" {
		v.Title = d.title
	}
	for i, q := range d.pages {
		it := navItem{Title: q.title, URL: rel(p, q), Current: q == p}
		if q != p {
			v.Nav = append(v.Nav, it)
			continue
		}
		it.Subs = d.outline(p)
		v.Nav = append(v.Nav, it)
		if i > 0 {
			v.Prev = &navItem{Title: d.pages[i-1].title, URL: rel(p, d.pages[i-1])}
		}
		if i+1 < len(d.pages) {
			v.Next = &navItem{Title: d.pages[i+1].title, URL: rel(p, d.pages[i+1])}
		}
	}
	if p.slug == "" && d.lang.Dir == "" {
		v.Script = template.JS(d.redirectScript(p))
	}
	return v
}

// outline lists the headings of p for the sidebar: the sections of the
// home page, or two heading levels of a part page.
func (d *doc) outline(p *page) []navItem {
	var out []navItem
	if p.slug == "" {
		for _, s := range p.sections {
			out = append(out, navItem{Title: s.title, URL: "#" + id(s.heading)})
		}
		return out
	}
	for _, h := range p.sections[0].headings()[1:] {
		it := navItem{Title: string(h.Text(d.src)), URL: "#" + id(h)}
		switch {
		case h.Level == 2:
			out = append(out, it)
		case h.Level == 3 && len(out) > 0:
			out[len(out)-1].Subs = append(out[len(out)-1].Subs, it)
		}
	}
	return out
}

var headingTag = regexp.MustCompile(`(?s)<h([1-4]) id="([^"]+)">(.*?)</h[1-4]>`)

func (d *doc) body(p *page) string {
	var b bytes.Buffer
	if p.slug == "" {
		fmt.Fprintf(&b, "<h1 id=\"%s\">%s</h1>\n<div class=\"lede\">\n", html.EscapeString(d.titleID), html.EscapeString(d.title))
		d.render(&b, d.preamble...)
		b.WriteString("</div>\n")
		d.contents(&b, p)
		for _, s := range p.sections {
			d.render(&b, s.heading)
			d.render(&b, s.nodes...)
		}
	} else {
		s := p.sections[0]
		fmt.Fprintf(&b, "<h1 id=\"%s\">%s</h1>\n", html.EscapeString(id(s.heading)), html.EscapeString(s.title))
		d.render(&b, s.nodes...)
	}
	return headingTag.ReplaceAllString(b.String(),
		`<h$1 id="$2">$3<a class="anchor" href="#$2" aria-hidden="true" tabindex="-1">#</a></h$1>`)
}

// contents writes the home page's cards, one per part.
func (d *doc) contents(b *bytes.Buffer, home *page) {
	fmt.Fprintf(b, "<div class=\"cards\" role=\"list\" aria-label=\"%s\">\n", d.lang.UI.Contents)
	for _, p := range d.pages[1:] {
		var subs []string
		for _, h := range p.sections[0].headings()[1:] {
			if h.Level == 2 {
				subs = append(subs, string(h.Text(d.src)))
			}
		}
		if len(subs) > 6 {
			subs = append(subs[:6], "…")
		}
		fmt.Fprintf(b, "<a class=\"card\" role=\"listitem\" href=\"%s\"><b>%s</b><span>%s</span><small>%d %s</small></a>\n",
			rel(home, p), html.EscapeString(p.title), html.EscapeString(strings.Join(subs, " · ")), p.entries, d.lang.UI.Entries)
	}
	fmt.Fprintf(b, "<a class=\"card feature\" role=\"listitem\" href=\"%sruntime-compiler/\"><b>%s</b><span>%s</span><small>Parts I–III</small></a>\n</div>\n",
		strings.Repeat("../", depth(home)), d.lang.UI.RuntimeCompilerLong, d.lang.UI.RuntimeCompilerDesc)
}

func (d *doc) render(w *bytes.Buffer, nodes ...ast.Node) {
	for _, n := range nodes {
		if err := md.Renderer().Render(w, d.src, n); err != nil {
			log.Fatal(err)
		}
	}
}

// redirectScript sends links to the old single page, such as
// /history/#generics, to the page that now holds the heading.
func (d *doc) redirectScript(home *page) string {
	m := map[string]string{}
	for h, p := range d.pageOf {
		if p != home {
			m[h] = rel(home, p)
		}
	}
	return fmt.Sprintf(`(function(m){var h=location.hash.slice(1);if(h&&!document.getElementById(h)&&m[h]){window.moved=true;location.replace(m[h]+'#'+h)}})(%s);`, jsonString(m))
}

// redirectPage replaces the old cn.html with a redirect to the page
// that now holds the heading in its URL fragment.
func (d *doc) redirectPage() []byte {
	m := map[string]string{}
	for h, p := range d.pageOf {
		m[h] = p.url
	}
	return []byte(fmt.Sprintf(`<!DOCTYPE html>
<html lang="%s">
<head>
<meta charset="utf-8">
<title>%s</title>
<meta name="robots" content="noindex">
<link rel="canonical" href="%s">
<script>(function(m){var h=location.hash.slice(1);location.replace((m[h]||%q)+(h?'#'+h:''))})(%s);</script>
</head>
<body><p><a href="%s">%s</a></p></body>
</html>
`, d.lang.Tag, html.EscapeString(d.title), d.lang.Dir, d.lang.Dir, jsonString(m), d.lang.Dir, html.EscapeString(d.title)))
}

var tags = regexp.MustCompile(`<[^>]*>`)

// searchIndex lists every list item of the document with the heading
// it sits under, for site.js.
func (d *doc) searchIndex() []byte {
	var idx struct {
		Pages    [][2]string `json:"pages"`    // url, title
		Sections [][]any     `json:"sections"` // page, heading id, heading
		Items    [][]any     `json:"items"`    // label, text, link, section
	}
	for pi, p := range d.pages {
		idx.Pages = append(idx.Pages, [2]string{p.url, p.title})
		idx.Sections = append(idx.Sections, []any{pi, "", p.title})
		for _, s := range p.sections {
			if p.slug == "" {
				idx.Sections = append(idx.Sections, []any{pi, id(s.heading), s.title})
			}
			for _, n := range s.nodes {
				ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
					if !entering {
						return ast.WalkContinue, nil
					}
					switch n := n.(type) {
					case *ast.Heading:
						idx.Sections = append(idx.Sections, []any{pi, id(n), string(n.Text(d.src))})
					case *ast.ListItem:
						if label, text, link := d.entry(n); text != "" || label != "" {
							idx.Items = append(idx.Items, []any{label, text, link, len(idx.Sections) - 1})
						}
					}
					return ast.WalkContinue, nil
				})
			}
		}
	}
	return []byte(jsonString(idx) + "\n")
}

// entry returns the label, plain text and first link of a list item,
// leaving out its nested lists.
func (d *doc) entry(li *ast.ListItem) (label, text, link string) {
	first := li.FirstChild()
	if first == nil {
		return "", "", ""
	}
	var b bytes.Buffer
	d.render(&b, first)
	s := b.String()
	if i := strings.Index(s, "</summary>"); i >= 0 {
		s = s[:i]
	}
	text = strings.Join(strings.Fields(html.UnescapeString(tags.ReplaceAllString(s, " "))), " ")
	if l, ok := first.FirstChild().(*ast.Link); ok {
		link = string(l.Destination)
		if isRef(l, d.src) {
			label = string(l.Text(d.src))
			text = strings.TrimSpace(strings.TrimPrefix(text, label))
		}
	}
	if r := []rune(text); len(r) > 200 {
		text = string(r[:200]) + "…"
	}
	return label, text, link
}

// headings returns the section's headings, its own first.
func (s *section) headings() []*ast.Heading {
	hs := []*ast.Heading{s.heading}
	for _, n := range s.nodes {
		ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			if h, ok := n.(*ast.Heading); ok && entering {
				hs = append(hs, h)
			}
			return ast.WalkContinue, nil
		})
	}
	return hs
}

func isHeading(n ast.Node, level int) bool {
	h, ok := n.(*ast.Heading)
	return ok && h.Level == level
}

// isBackToTop reports whether n is a "[Back To Top](#top)" paragraph.
func isBackToTop(n ast.Node) bool {
	if n.Kind() != ast.KindParagraph || n.ChildCount() != 1 {
		return false
	}
	l, ok := n.FirstChild().(*ast.Link)
	return ok && string(l.Destination) == "#top"
}

// isRef reports whether l is the label that starts a list entry, such
// as [issue/123](...) or [design/go13compiler](...).
func isRef(l *ast.Link, src []byte) bool {
	p := l.Parent()
	if p == nil || l.PreviousSibling() != nil || p.Parent() == nil || p.Parent().Kind() != ast.KindListItem {
		return false
	}
	return refLabel.Match(l.Text(src))
}

func lines(n ast.Node, src []byte) []byte {
	var b []byte
	for i := 0; i < n.Lines().Len(); i++ {
		seg := n.Lines().At(i)
		b = append(b, seg.Value(src)...)
	}
	return b
}

func id(h *ast.Heading) string {
	v, _ := h.AttributeString("id")
	b, _ := v.([]byte)
	return string(b)
}

// depth is the number of directories between p and the site root.
func depth(p *page) int { return strings.Count(p.url, "/") }

// rel returns the URL of page q relative to page p.
func rel(p, q *page) string {
	u := strings.Repeat("../", depth(p)) + q.url
	if u == "" {
		return "./"
	}
	return u
}

func jsonString(v any) string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(b.String())
}

func write(name string, data []byte) {
	name = filepath.Join(*outDir, name)
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(name, data, 0o644); err != nil {
		log.Fatal(err)
	}
}
