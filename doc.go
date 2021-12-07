// Copyright 2020 The golang.design Initiative.
// All rights reserved. Use of this source code
// is governed by a MIT license that can be found
// in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var (
	md   goldmark.Markdown
	pctx parser.Context
)

type data struct {
	Lang       string
	Selected   string
	Navigation template.HTML
	Content    template.HTML
}

func init() {
	md = goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
		// html.WithHardWraps(), // do this someday.
		),
	)
	pctx = parser.NewContext(parser.WithIDs(newIDs()))
}

func main() {
	index := convertMD("README.md")
	parseTemplate("en", "index", index)

	cn := convertMD("README_cn.md", parser.WithContext(pctx))
	parseTemplate("cn", "cn", cn)

	fmt.Println("golang-design/history: A Documentary of Go")
}

// convertMD convert md to html
func convertMD(filename string, opts ...parser.ParseOption) bytes.Buffer {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read: cannot read README.md, err: %v", err)
	}

	var b bytes.Buffer
	err = md.Convert(d, &b, opts...)
	if err != nil {
		log.Fatalf("Convert: cannot convert README from markdown to html, err: %v", err)
	}
	return b
}

// parseTemplate Add md data to the template
func parseTemplate(lang, target string, b bytes.Buffer) {
	f, err := os.Create("public/" + target + ".html")
	if err != nil {
		log.Fatalf("Create: cannot create index.html, err: %v", err)
	}

	// Find TOC, and remove all hard coded all back to top. Then use
	// back to top button for navigation.
	//
	// This is urgly, I know. At least this works so far with least efforts.
	// A better and generic way obviously is to do DOM tree analysis.
	dom := b.String()
	tocStart := strings.Index(dom, "<p><strong>Table of Contents</strong></p>")
	tocEnd := strings.Index(dom, `<hr>`)
	toc := dom[tocStart:tocEnd]
	dom = dom[:tocStart] + `<div class="doc-nav-mobile">` +
		dom[tocStart:tocEnd] + `</div>` + dom[tocEnd:]
	dom = strings.ReplaceAll(dom, `<p><a href="#top">Back To Top</a></p>`, "")

	tmpl := template.Must(template.New(target).Parse(indexTemplate))

	var selected string
	if "index" != target {
		selected = "selected"
	}
	tmpl.Execute(f, data{
		Lang:       lang,
		Selected:   selected,
		Navigation: template.HTML(toc),
		Content:    template.HTML(dom),
	})
}

const indexTemplate = `
<!DOCTYPE html>
<html lang="{{.Lang}}">
<head>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-80889616-4"></script>
    <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());
    gtag('config', 'UA-80889616-4');
    </script>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go: A Documentary</title>
    <link rel="stylesheet" href="bootstrap.css">
    <link rel="stylesheet" href="style.css">
</head>
<body>
<div class="header row">
	<div class="col-3 dark-switch">
		<label>
   			<select Onchange="window.open(this.options[this.selectedIndex].value,target='_self')">
	 		<option value="index.html">English</option>
	 		<option {{.Selected}} value="cn.html">简体中文</option>
      		</select>
		</label>
    </div>
    <div class="col-6"></div>
    <div class="col-3 dark-switch">
    <div class="custom-control custom-switch">
    <input type="checkbox" class="custom-control-input" id="darkSwitch" />
        <label class="custom-control-label" for="darkSwitch">Dark Mode</label>
    </div>
    </div>
</div>
<div id="btt"><a href="#top">⬆</a></div>
<div id="container" class="row">
<div class="doc-content col-lg-9">{{.Content}}</div>
<nav class="doc-nav col-lg-3">{{.Navigation}}</nav>
</div>
<script src="dark.js"></script>
<script async src="//changkun.de/urlstat/client.js"></script>
</body>
</html>
`
