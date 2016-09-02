// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"errors"

	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

var DB xsql.DB

type Manager struct {
	db       xsql.DB
	Entity   func() entity.Interface
	Entities func() entity.Interfaces
	Pre      func()
}

func (me *Manager) NewEntity() entity.Interface {
	if me.Entity != nil {
		return me.Entity()
	}
	return nil
}

func (me *Manager) NewEntities() entity.Interfaces {
	if me.Entities != nil {
		return me.Entities()
	}
	return nil
}

func (me *Manager) NewEntityResult() *result.Entity {
	return &result.Entity{Data: me.NewEntity()}
}

func (me *Manager) NewEntitiesResult() *result.Entities {
	return &result.Entities{Data: me.NewEntities()}
}

func (me *Manager) NewPageResult() *result.Page {
	out := domain.NewPageDefault(me.NewEntities())
	return &result.Page{Data: out}
}

func (me *Manager) Get(out entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Get(out)
}

func (me *Manager) SelectOne(out entity.Interface, query string, args ...interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(query, args...).Row(out)
}

func (me *Manager) SelectList(out entity.Interfaces, query string, args ...interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(query, args...).Rows(out)
}

func (me *Manager) SelectPage(content entity.Interfaces, pageable domain.Pageable, dql string, args ...interface{}) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Page(content, pageable)
}

func (me *Manager) SelectCount(dql string, args ...interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	countSql := sqls.ParseCountSql(dql)
	return me.SelectInt(countSql, args...)
}

func (me *Manager) SelectInt(dql string, args ...interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Int()
}

func (me *Manager) SelectFloat(dql string, args ...interface{}) (float64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Float()
}

func (me *Manager) SelectStr(dql string, args ...interface{}) (string, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Str()
}

func (me *Manager) SelectOneByNamed(out entity.Interface, dql string, args map[string]interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return err
	}
	err = query.Row(out)
	return err
}

func (me *Manager) SelectListByNamed(out entity.Interfaces, dql string, args map[string]interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return err
	}
	err = query.Rows(out)
	return err
}

func (me *Manager) SelectPageByNamed(content entity.Interfaces, pageable domain.Pageable, dql string, args map[string]interface{}) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	return query.Page(content, pageable)
}

func (me *Manager) SelectCountByNamed(dql string, args map[string]interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	countSql := sqls.ParseCountSql(dql)
	return me.SelectIntByNamed(countSql, args)
}

func (me *Manager) SelectIntByNamed(dql string, args map[string]interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	return query.Int()
}

func (me *Manager) SelectFloatByNamed(dql string, args map[string]interface{}) (float64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return 0.0, err
	}
	return query.Float()
}

func (me *Manager) SelectStrByNamed(dql string, args map[string]interface{}) (string, error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	return query.Str()
}

func (me *Manager) SelectOneBySift(out entity.Interface, sifts ...domain.Sift) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Row(out)
}

func (me *Manager) SelectListBySift(out entity.Interfaces, sifts ...domain.Sift) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Rows(out)
}

func (me *Manager) SelectPageBySift(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Page(content, pageable)
}

func (me *Manager) SelectCountBySift(sifts ...domain.Sift) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Count(me.NewEntity())
}

func (me *Manager) Save(c xhttp.Context, e entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	if err := e.Validate(); err != nil {
		return err
	}
	tx, err := me.DB().Begin()
	if err != nil {
		return err
	}
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		if c != nil && c.Session().IsLogin() {
			if p, err := c.Session().Principal(); err == nil {
				e.SetString(creater, p.Id)
				e.SetString(modifier, p.Id)
			}
		}
		e.SetString(created, times.NowUnixStr())
		e.SetString(modified, times.NowUnixStr())
		_, err = me.DB().Insert(e)
	} else {
		if c != nil && c.Session().IsLogin() {
			if p, err := c.Session().Principal(); err == nil {
				e.SetString(modifier, p.Id)
			}
		}
		e.SetString(modified, times.NowUnixStr())
		_, err = me.DB().Update(e)
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (me *Manager) Disable(c xhttp.Context, e entity.Interface) (int64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		return 0, errors.New("Gets the primary key value failed")
	}
	tx, err := me.DB().Begin()
	if err != nil {
		return 0, err
	}
	if c != nil && c.Session().IsLogin() {
		if p, err := c.Session().Principal(); err == nil {
			e.SetString(modifier, p.Id)
			e.SetString(modified, times.NowUnixStr())
		}
	}
	r, err := me.DB().Disable(e)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return r, nil
}

func (me *Manager) Exec(dml string, args ...interface{}) (sql.Result, error) {
	if me.Pre != nil {
		me.Pre()
	}
	tx, err := me.DB().Begin()
	if err != nil {
		return nil, err
	}
	r, err := me.DB().Exec(dml, args...)
	if err != nil {
		tx.Rollback()
		return r, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return r, err
	}
	return r, nil
}

func (me *Manager) SetDB(val xsql.DB) {
	me.db = val
}

func (me *Manager) DB() xsql.DB {
	if me.db == nil {
		me.SetDB(DB)
	}
	return me.db
}
