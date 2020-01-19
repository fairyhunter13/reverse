// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io/ioutil"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var result = fmt.Sprintf(`package models

type A struct {
	Id int %sxorm:"integer"%s
}

type B struct {
	Id int %sxorm:"INTEGER"%s
}
`, "`", "`", "`", "`")

func TestReverse(t *testing.T) {
	err := reverse("../example/goxorm.yml")
	assert.NoError(t, err)

	bs, err := ioutil.ReadFile("../models/models.go")
	assert.NoError(t, err)
	assert.EqualValues(t, result, string(bs))
}
