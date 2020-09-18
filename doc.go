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

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type data struct {
	Navigation string
	Content    template.HTML
}

var md goldmark.Markdown

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
}

func main() {
	d, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatalf("Read: cannot read README.md, err: %v", err)
	}

	var b bytes.Buffer
	err = md.Convert(d, &b)
	if err != nil {
		log.Fatalf("Convert: cannot convert README from markdown to html, err: %v", err)
	}

	a := data{
		Navigation: "",
		Content:    template.HTML(b.Bytes()),
	}

	f, err := os.Create("public/index.html")
	if err != nil {
		log.Fatalf("Create: cannot create index.html, err: %v", err)
	}

	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	tmpl.Execute(f, a)
	fmt.Println("golang-design/history: A Documentary of Go")
}

const indexTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go: A Documentary</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    <link rel="stylesheet" href="style.css">
</head>
<body>
<div class="header row">
    <div class="col-9"></div>
    <div class="col-3 dark-switch">
    <div class="custom-control custom-switch">
    <input type="checkbox" class="custom-control-input" id="darkSwitch" />
        <label class="custom-control-label" for="darkSwitch">Dark Mode</label>
    </div>
    </div>
</div>
<div id="container">
<nav class="doc-nav">{{.Navigation}}</nav>
<div class="doc-content">{{.Content}}</div>
</div>
<script src="dark.js"></script>
</body>
</html>
`
