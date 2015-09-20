// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/result"
)

type Service interface {
	NewEntity() entity.Interface
	NewEntities() entity.Interfaces
	NewEntityResult() *result.Entity
	NewEntitiesResult() *result.Entities
	NewPageResult() *result.Page
	Get(out entity.Interface) error
	SelectOne(out entity.Interface, query string, args ...interface{}) (err error)
	SelectList(out entity.Interfaces, query string, args ...interface{}) (err error)
	SelectPage(content entity.Interfaces, pageable domain.Pageable, dql string, args ...interface{}) (out domain.Page, err error)
	SelectInt(dql string, args ...interface{}) (int, error)
	SelectFloat(dql string, args ...interface{}) (float64, error)
	SelectStr(dql string, args ...interface{}) (string, error)
	SelectOneBySift(out entity.Interface, sifts ...domain.Sift) (err error)
	SelectListBySift(out entity.Interfaces, sifts ...domain.Sift) (err error)
	SelectPageBySift(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (out domain.Page, err error)
	Save(e entity.Interface) error
	Disable(e entity.Interface) (int64, error)
	Exec(dml string, args ...interface{}) (sql.Result, error)
}