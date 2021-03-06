// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Uint16 struct {
	base
	value uint16
}

func (me *Uint16) Value() uint16 {
	return me.value
}

func (me *Uint16) ValuePtr() *uint16 {
	return &me.value
}

func (me *Uint16) SetValue(v uint16) {
	me.value = v
	me.field.SetModified(true)
}

func (me *Uint16) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = uint16(val)
	return nil
}

func (me *Uint16) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Uint16) String() string {
	return strconv.Itoa(int(me.value))
}

func (me *Uint16) Name() string {
	return "uint16"
}
