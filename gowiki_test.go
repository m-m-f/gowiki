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
	"encoding/json"
	//	"os"
	//	"strings"
	"testing"
)

func TestParseArticle(t *testing.T) {
	mw := "* ''[[The Album (ABBA album)|''The Album'']]'' (1977)"
	t.Log(mw)
	a, err := ParseArticle("Test", mw, &DummyPageGetter{})
	if err != nil {
		t.Error("Error:", err)
	}
	b, err := json.MarshalIndent(a.Tokens, "", "\t")
	if err != nil {
		t.Error("Error:", err)
	}
	t.Log("Tokens\n")
	t.Log(string(b))
}

func TestWikiCanonicalFormNamespaceEsc(t *testing.T) {
	wl := StandardNamespaces.WikiCanonicalFormNamespaceEsc("WiKIpEdia:pagename#section", "", true)
	if wl.Namespace != "Wikipedia" || wl.PageName != "Pagename" || wl.Anchor != "section" {
		t.Error("Error: wikilink not parsed correctly", wl)
	}
}
