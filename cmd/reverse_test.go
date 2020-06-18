// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
)

var (
	result = fmt.Sprintf(`package models

type A struct {
	Id int %sxorm:"integer"%s
}

func (m *A) TableName() string {
	return "a"
}

type B struct {
	Id int %sxorm:"INTEGER"%s
}

func (m *B) TableName() string {
	return "b"
}
`, "`", "`", "`", "`")
)

func TestReverse(t *testing.T) {
	err := reverse("../example/goxorm.yml")
	assert.NoError(t, err)

	bs, err := ioutil.ReadFile("../models/models.go")
	assert.NoError(t, err)
	assert.EqualValues(t, result, string(bs))
}

func TestReverse2(t *testing.T) {
	type Outfw struct {
		Id       int    `xorm:"not null pk autoincr"`
		Sql      string `xorm:"default '' TEXT"`
		Template string `xorm:"default '' TEXT"`
		Filename string `xorm:"VARCHAR(50)"`
	}

	dir, err := ioutil.TempDir(os.TempDir(), "reverse")
	assert.NoError(t, err)

	e, err := xorm.NewEngine("sqlite3", filepath.Join(dir, "db.db"))
	assert.NoError(t, err)

	assert.NoError(t, e.Sync2(new(Outfw)))

	err = reverseFromReader(strings.NewReader(`
kind: reverse
name: mydb
source:
  database: sqlite3
  conn_str: '../testdata/test.db'
targets:
- type: codes
  include_tables:
  - a
  - b
  exclude_tables:
  - c
  language: golang
  output_dir: ../models
`))
	assert.NoError(t, err)
}
