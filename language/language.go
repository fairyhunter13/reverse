// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package language

import (
	"html/template"

	"xorm.io/xorm/schemas"
)

// Language represents a languages supported when reverse codes
type Language struct {
	Name      string
	Template  string
	Types     map[string]string
	Funcs     template.FuncMap
	Formatter func(string) (string, error)
	Importter func([]*schemas.Table) []string
	ExtName   string
}

var (
	languages = make(map[string]*Language)
)

// RegisterLanguage registers a language
func RegisterLanguage(l *Language) {
	languages[l.Name] = l
}

// GetLanguage returns a language if exists
func GetLanguage(name string, tableName bool) *Language {
	language := languages[name]
	if tableName {
		language = languages[name]
		language.Template = defaultGolangTemplateTable
		return language
	}
	return language
}
