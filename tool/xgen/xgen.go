// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
)

func main() {
	epath := flag.String("entity", "", "entity file path")
	clipath := flag.String("clipath", "", "import path for client project")
	apipath := flag.String("apipath", "", "import path for api project")
	tstpath := flag.String("tstpath", "", "import path for test project")
	hasScaffold := flag.Bool("scaffold", false, "is generated service and controller")
	hasService := flag.Bool("service", false, "is generated service")
	hasController := flag.Bool("controller", false, "is generated controller")
	hasDto := flag.Bool("dto", false, "is generated dto")
	hasApi := flag.Bool("api", false, "is generated api")
	hasSql := flag.Bool("sql", false, "is generated sql")
	hasLog := flag.Bool("log", false, "is generated log")
	hasUtil := flag.Bool("util", false, "is generated util")
	hasConst := flag.Bool("const", false, "is generated const")
	hasHtml := flag.Bool("html", false, "is generated html")
	hasJs := flag.Bool("js", false, "is generated js")
	flag.Parse()
	f := factory{
		Clipath:          *clipath,
		Apipath:          *apipath,
		Tstpath:          *tstpath,
		HasGenService:    *hasService,
		HasGenController: *hasController,
		HasGenDto:        *hasDto,
		HasGenApi:        *hasApi,
		HasGenSql:        *hasSql,
		HasGenLog:        *hasLog,
		HasGenUtil:       *hasUtil,
		HasGenConst:      *hasConst,
		HasGenHtml:       *hasHtml,
		HasGenJs:         *hasJs,
	}
	if *hasScaffold {
		f.HasGenService = true
		f.HasGenController = true
		f.HasGenDto = false
		f.HasGenApi = true
		f.HasGenSql = true
		f.HasGenLog = true
		f.HasGenUtil = true
		f.HasGenConst = true
		f.HasGenHtml = true
		f.HasGenJs = true
	}
	if err := f.Init(*epath); err != nil {
		log.Printf("%v", err)
		return
	}
	if err := f.Write(); err != nil {
		log.Printf("%v", err)
		return
	}
}
