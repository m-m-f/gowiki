/*
Copyright (C) IBM Corporation 2015, Michele Franceschini <franceschini@us.ibm.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gowiki

import (
	"bytes"
	"unicode/utf8"
)

func (a *Article) appendText(t string) {
	a.nchar += utf8.RuneCountInString(t)
	a.text.WriteString(t)
}

func (a *Article) genTextInternal(root *ParseNode, indent int) {
	lastwasspace := false
	for _, n := range root.Nodes {
		var linkStart int
		var fl FullWikiLink
		isLink := false
		tappend := ""
		switch n.NType {
		case "break":
			a.appendText("\n")
		case "space":
			if !lastwasspace {
				a.appendText(" ")
			}
		case "text":
			a.appendText(n.Contents)
		case "image":
			a.appendText("\n")
			tappend = "\n"
		case "link":
			isLink = true
			linkStart = len(a.text.Bytes())
			fl = FullWikiLink{Link: n.Link, Start: a.nchar}
		case "html":
			switch n.NSubType {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				a.appendText("\n")
				tappend = "\n"
				if len(a.AbstractText) == 0 {
					a.AbstractText = a.text.String()
				}
			case "br":
				a.appendText("\n")
			case "ref":
				a.appendText(" ")
			}
		}
		if len(n.Nodes) > 0 {
			a.genTextInternal(n, 0)
		}
		if isLink {
			ttmp := a.text.Bytes()
			fl.End = a.nchar
			fl.Text = string(ttmp[linkStart:])
			a.TextLinks = append(a.TextLinks, fl)
		}
		lastwasspace = false
		if n.NType == "space" {
			lastwasspace = true
		}
		//		a.Text += tappend
		a.appendText(tappend)
	}

	return
}

func (a *Article) genText() error {
	a.text = bytes.NewBuffer(make([]byte, 1024*1024, 1024*1024))
	a.text.Truncate(0)
	a.nchar = 0
	a.AbstractText = ""
	a.genTextInternal(a.Root, 0)
	a.Text = string(a.text.Bytes())
	if len(a.AbstractText) == 0 {
		a.AbstractText = a.Text
	}
	a.gt = true
	return nil
}

func (a *Article) GenText() error {
	return a.genText()
}
