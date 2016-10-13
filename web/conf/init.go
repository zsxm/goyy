// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/templates"
)

func Init(envName, defaultProfile string, activesProfile ...string) {
	initProfile(defaultProfile, activesProfile...)
	initLog()
	initApi(envName)
	initAsset(envName)
	initExport(envName)
	initSession(envName)
	initTemplate(envName)
}

func initProfile(defaults string, actives ...string) {
	profile.SetDefault(defaults)
	profile.SetActives(actives...)
}

func initLog() {
	if profile.Accepts(profile.DEV) {
		log.SetDefaultOutput(log.Oconsole)
	} else {
		log.SetDefaultOutput(log.Odailyfile)
	}
}

func initApi(envName string) {
	if v, err := env.Api(envName); err == nil {
		Conf.Api.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}

func initAsset(envName string) {
	if v, err := env.Asset(envName); err == nil {
		Conf.Asset.Enable = v.Enable
		Conf.Asset.Dir = v.Dir
		Conf.Asset.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Static(envName); err == nil {
		Conf.Static.Enable = v.Enable
		Conf.Static.Dir = v.Dir
		Conf.Static.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Developer(envName); err == nil {
		Conf.Developer.Enable = v.Enable
		Conf.Developer.Dir = v.Dir
		Conf.Developer.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Operation(envName); err == nil {
		Conf.Operation.Enable = v.Enable
		Conf.Operation.Dir = v.Dir
		Conf.Operation.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Upload(envName); err == nil {
		Conf.Upload.Enable = v.Enable
		Conf.Upload.Dir = v.Dir
		Conf.Upload.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}

func initExport(envName string) {
	if v, err := env.Export(envName); err == nil {
		Conf.Export.Dir = v.Dir
	} else {
		log.Println(err.Error())
	}
}

func initSession(envName string) {
	if v, err := env.Session(envName); err == nil {
		Conf.Session.Addr = v.Addr
		Conf.Session.Password = v.Password
	} else {
		log.Println(err.Error())
	}
}

func initTemplate(envName string) {
	if v, err := env.Html(envName); err == nil {
		Conf.Html.Enable = v.Enable
		Conf.Html.Reloaded = v.Reloaded
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Template(envName); err == nil {
		Conf.Template.Enable = v.Enable
		Conf.Template.Reloaded = v.Reloaded
		if v.Enable {
			templates.GetApis = func() string { return Conf.Api.URL }
			templates.GetAssets = func() string { return Conf.Asset.URL }
			templates.GetDevelopers = func() string { return Conf.Developer.URL }
			templates.GetOperations = func() string { return Conf.Operation.URL }
			templates.GetStatics = func() string { return Conf.Static.URL }
			templates.GetUploads = func() string { return Conf.Upload.URL }
		}
	} else {
		log.Println(err.Error())
	}

}