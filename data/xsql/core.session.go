// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dml"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/dql"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"time"
)

type session struct {
	db      *sql.DB
	dialect dialect.Interface
	dml     dml.Interface
	dql     dql.Interface
}

// New Query
func (me *session) Query(dql string, args ...interface{}) Query {
	return &query{db: me.db, session: me, dql: dql, args: args}
}

// Select one SQL
func (me *session) Get(out entity.Interface) error {
	dql, arg := me.dql.SelectOne(out)
	err := me.Query(dql, arg).Row(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Select one SQL
func (me *session) SelectOne(out entity.Interface, sifts ...domain.Sift) error {
	dql, args, err := me.dql.SelectListBySift(out, sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	err = me.Query(dql, args...).Row(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Select list SQL
func (me *session) SelectList(out entity.Interfaces, sifts ...domain.Sift) error {
	dql, args, err := me.dql.SelectListBySift(out.New(), sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	err = me.Query(dql, args...).Rows(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Select page SQL
func (me *session) SelectPage(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (domain.Page, error) {
	dql, args, err := me.dql.SelectListBySift(content.New(), sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}
	page, err := me.Query(dql, args...).Page(content, pageable)
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}
	return page, nil
}

// Insert SQL
func (me *session) Insert(e entity.Interface) (int64, error) {
	pk := e.Table().Primary().Name()
	if t, ok := e.Column(pk); ok {
		e.SetString(t.Name(), newId())
	}
	dml, args := me.dml.Insert(e)
	res, err := me.Exec(dml, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Update SQL
func (me *session) Update(e entity.Interface) (int64, error) {
	dml, args := me.dml.Update(e)
	res, err := me.Exec(dml, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Delete SQL
func (me *session) Delete(e entity.Interface) (int64, error) {
	dml, arg := me.dml.Delete(e)
	res, err := me.Exec(dml, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Logical Delete SQL
func (me *session) Disable(e entity.Interface) (int64, error) {
	dml, arg := me.dml.Disable(e)
	res, err := me.Exec(dml, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (me *session) Exec(dml string, args ...interface{}) (sql.Result, error) {
	if isDebug() {
		now := time.Now()
		defer debugLog(now, dml, args...)
	}
	stmt, err := me.db.Prepare(dml)
	if err != nil {
		panic(err)
	}
	return stmt.Exec(args...)
}

// Begin Transaction
func (me *session) Begin() (Tx, error) {
	return me.db.Begin()
}

// Close Session
func (me *session) Close() error {
	return nil
}

// Get Database Type
func (me *session) DBType() string {
	return me.dialect.Type()
}

// newId returns the id string.
func newId() string {
	id := uuid.NewV1()
	return strings.Replace(id.String(), "-", "", -1)
}