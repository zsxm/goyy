// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"net/http"
	"strconv"
)

type htmlServeMux struct{}

type directiveInfo struct {
	statement string // <!--#include file="/footer.html"-->content<!--#endinclude"-->
	directive string // include
	argKey    string // file
	argValue  string // /footer.html
	begin     int    // postion:<!--#include file="
	center    int    // postion:"-->
	end       int    // postion:<!--#endinclude"-->
}

type tagInfo struct {
	statement string // <link rel="shortcut icon" href="/favicon.ico" go:href="{{assets}}/favicon.ico">
	newstmt   string // <link rel="shortcut icon" href="/static/favicon.ico">
	attr      string // href
	tagBegin  int    // postion:<
	tagEnd    int    // postion:>
	srcBegin  int    // postion: href="
	srcEnd    int    // postion:"
	dstBegin  int    // postion: go:href="
	dstEnd    int    // postion:"
	dstVal    string // {{assets}}/favicon.ico
}

type tagTextInfo struct {
	statement string // <title go:title="/title.html">login</title>
	newstmt   string // <title>login-appendTitle</title>
	attr      string // title
	tagBegin  int    // postion:<
	tagEnd    int    // postion:</
	srcBegin  int    // postion:>
	srcEnd    int    // postion:</
	dstBegin  int    // postion: go:title="
	dstEnd    int    // postion:"
	dstVal    string // /title.html
}

var hsm = &htmlServeMux{}

func (me *htmlServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	if me.isHtml(r.URL.Path) {
		filename := Conf.Templates.Directory + r.URL.Path
		if files.IsExist(filename) {
			if me.isUseBrowserCache(w, r, filename) {
				return true
			}
			if c, err := files.Read(filename); err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(me.parseFile(w, r, c)))
			}
			return true
		}
	}
	return false
}

func (me *htmlServeMux) isHtml(path string) bool {
	if strings.HasSuffix(path, ".html") {
		return true
	}
	return false
}

func (me *htmlServeMux) isInclude(content string) bool {
	if strings.Index(content, directiveIncludeEnd) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) isSec(content string) bool {
	if strings.Index(content, directiveSecEnd) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) isUseBrowserCache(w http.ResponseWriter, r *http.Request, filename string) bool {
	if fileModTimeUnix, err := files.ModTime(filename); err == nil {
		var browserModTimeUnix int64
		// Browser save file last modified time
		browserModTime := r.Header.Get("If-Modified-Since")
		if strings.IsNotBlank(browserModTime) {
			if v, err := times.ParseGMT(browserModTime); err == nil {
				browserModTimeUnix = v
			}
		}
		if browserModTimeUnix < fileModTimeUnix {
			// Actual file last modified time
			fileModTime := times.Ugmt(fileModTimeUnix)
			// Tell the browser not to use cache
			w.Header().Set("last-modified", fileModTime)
			return false
		} else {
			var content string
			if c, err := files.Read(filename); err == nil {
				content = c
			}
			if me.isSec(content) {
				var lastLoginTimeUnix int64
				s := newSession4Redis(w, r)
				if v, err := s.Get(principalLoginTime); err == nil && strings.IsNotBlank(v) {
					if i, err := strconv.Atoi(v); err == nil {
						lastLoginTimeUnix = int64(i)
					}
				}
				if browserModTimeUnix < lastLoginTimeUnix {
					// Actual last login time
					lastLoginTime := times.Ugmt(lastLoginTimeUnix)
					// Tell the browser not to use cache
					w.Header().Set("last-modified", lastLoginTime)
					return false
				}
			}
			if me.isInclude(content) {
				var includeFileModTimeUnix int64
				directives := make([]directiveInfo, 0)
				directives = me.buildDirectiveInfo(content, directiveIncludeBegin, directiveIncludeEnd, directives)
				for _, v := range directives {
					if val, err := files.ModTime(v.argValue); err == nil {
						if includeFileModTimeUnix < val {
							includeFileModTimeUnix = val
						}
					}
				}
				if browserModTimeUnix < includeFileModTimeUnix {
					// The actual last modification time of the include file
					includeFileModTime := times.Ugmt(includeFileModTimeUnix)
					// Tell the browser not to use cache
					w.Header().Set("last-modified", includeFileModTime)
					return false
				}
			}
			// Tell the browser to use the cache
			w.WriteHeader(304)
			return true
		}
	}
	return false
}

func (me *htmlServeMux) parseFile(w http.ResponseWriter, r *http.Request, content string) string {
	content = me.parseIncludeFile(content)
	content = me.parseIfFile(content)
	content = me.parseTagAttrFile(content, tagAttrHref)
	content = me.parseTagAttrFile(content, tagAttrSrc)
	content = me.parseTagAttrFile(content, tagAttrAction)
	content = me.parseTagTextFile(content, tagTextTitle)
	if me.isSec(content) {
		content = me.parseSecUserFile(w, r, content)
		content = me.parseSecLoginFile(w, r, content)
		content = me.parseSecHasRoleFile(w, r, content)
		content = me.parseSecHasAnyRoleFile(w, r, content)
	}
	return content
}

func (me *htmlServeMux) parseIncludeFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveIncludeBegin, directiveIncludeEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		v, err := files.Read(directives[i].argValue)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		content = strings.Replace(content, directives[i].statement, v, -1)
	}
	return content
}

func (me *htmlServeMux) parseIfFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveIfBegin, directiveIfEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if directives[i].argValue == "false" {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
	}
	return content
}

func (me *htmlServeMux) parseTagAttrFile(content, attr string) string {
	tags := make([]tagInfo, 0)
	tags = me.buildTagInfo(content, attr, tags)
	for i := len(tags) - 1; i >= 0; i-- {
		content = strings.Replace(content, tags[i].statement, tags[i].newstmt, -1)
	}
	return content
}

func (me *htmlServeMux) parseTagTextFile(content, attr string) string {
	tags := make([]tagTextInfo, 0)
	tags = me.buildTagTextInfo(content, attr, tags)
	for i := len(tags) - 1; i >= 0; i-- {
		content = strings.Replace(content, tags[i].statement, tags[i].newstmt, -1)
	}
	return content
}

func (me *htmlServeMux) parseSecLoginFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecLoginBegin, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		isLogin := false
		s := newSession4Redis(w, r)
		if v, err := s.Get(principalId); err == nil && strings.IsNotBlank(v) {
			isLogin = true
		}
		if directives[i].argValue == "true" && !isLogin {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
		if directives[i].argValue == "false" && isLogin {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
	}
	return content
}

func (me *htmlServeMux) parseSecUserFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecUserBegin, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if directives[i].argValue == "name" {
			s := newSession4Redis(w, r)
			v, err := s.Get(principalLoginName)
			if err != nil {
				logger.Error(err.Error())
				break
			}
			content = strings.Replace(content, directives[i].statement, v, -1)
			break
		}
	}
	return content
}

func (me *htmlServeMux) parseSecHasRoleFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecHasRoleBegin, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if strings.IsNotBlank(directives[i].argValue) {
			isLogin := false
			s := newSession4Redis(w, r)
			if v, err := s.Get(principalId); err == nil && strings.IsNotBlank(v) {
				isLogin = true
			}
			isContains := false
			if v, err := s.Get(principalPermissions); err == nil {
				if strings.Contains(v, directives[i].argValue) {
					isContains = true
				}
			}
			if !isLogin || !isContains {
				content = strings.Replace(content, directives[i].statement, "", -1)
			}
			break
		}
	}
	return content
}

func (me *htmlServeMux) parseSecHasAnyRoleFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecHasAnyRoleBegin, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if strings.IsNotBlank(directives[i].argValue) {
			isLogin := false
			s := newSession4Redis(w, r)
			if v, err := s.Get(principalId); err == nil && strings.IsNotBlank(v) {
				isLogin = true
			}
			isContains := false
			if v, err := s.Get(principalPermissions); err == nil {
				avs := strings.Split(directives[i].argValue, ",")
				if strings.ContainsSliceAny(v, avs) {
					isContains = true
				}
			}
			if !isLogin || !isContains {
				content = strings.Replace(content, directives[i].statement, "", -1)
			}
			break
		}
	}
	return content
}

func (me *htmlServeMux) buildDirectiveInfo(content, directiveBegin, directiveEnd string, directives []directiveInfo) []directiveInfo {
	pos := 0
	for {
		begin := strings.IndexStart(content, directiveBegin, pos)
		if begin == -1 {
			if pos == 0 {
				return directives
			}
			break
		}
		center := strings.IndexStart(content, directiveArgEnd, begin)
		if center == -1 {
			if pos == 0 {
				return directives
			}
			break
		}
		end := strings.IndexStart(content, directiveEnd, center)
		if end == -1 {
			if pos == 0 {
				return directives
			}
			break
		}
		pos = end
		argValue := strings.Slice(content, begin+len(directiveBegin), center)
		if directiveBegin == directiveIncludeBegin {
			argValue = Conf.Templates.Directory + argValue
			if !files.IsExist(argValue) {
				continue
			}
		}
		statement := strings.Slice(content, begin, end+len(directiveEnd))
		directive := "include"
		argKey := "file"
		switch directiveBegin {
		case directiveIfEnd:
			directive = "if"
			argKey = "expr"
		case directiveSecUserBegin:
			directive = "sec"
			argKey = "user"
		case directiveSecHasRoleBegin:
			directive = "sec"
			argKey = "hasRole"
		case directiveSecHasAnyRoleBegin:
			directive = "sec"
			argKey = "hasAnyRole"
		}
		ii := directiveInfo{
			statement: statement,
			directive: directive,
			argKey:    argKey,
			argValue:  argValue,
			begin:     begin,
			center:    center,
			end:       end,
		}
		directives = append(directives, ii)
	}
	return directives
}

func (me *htmlServeMux) buildTagInfo(content, attr string, tags []tagInfo) []tagInfo {
	pos := 0
	srcBeginPre := " " + attr + tagAttrPost
	dstBeginPre := tagAttrPre + attr + tagAttrPost
	for {
		dstBegin := strings.IndexStart(content, dstBeginPre, pos)
		if dstBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		dstEnd := strings.IndexStart(content, tagAttrEnd, dstBegin+len(dstBeginPre))
		if dstEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagBegin := strings.IndexForward(content, tagBeginPre, dstBegin)
		if tagBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagEnd := strings.IndexStart(content, tagEndPre, dstEnd)
		if tagEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		pos = tagEnd
		statement := strings.Slice(content, tagBegin, tagEnd+len(tagEndPre))
		newstmt := statement
		dstVal := strings.Slice(content, dstBegin+len(dstBeginPre), dstEnd)
		srcBegin := strings.Index(newstmt, srcBeginPre)
		var srcEnd int
		if srcBegin > 0 {
			srcEnd = strings.IndexStart(newstmt, tagAttrEnd, srcBegin+len(srcBeginPre))
			newstmt = strings.Overlay(newstmt, "", srcBegin, srcEnd+len(tagAttrEnd))
		}
		newstmt = strings.Replace(newstmt, tagStaticAssets, Conf.Static.Assets, -1)
		newstmt = strings.Replace(newstmt, tagStaticConsumers, Conf.Static.Consumers, -1)
		newstmt = strings.Replace(newstmt, tagStaticOperations, Conf.Static.Operations, -1)
		newstmt = strings.Replace(newstmt, dstBeginPre, srcBeginPre, -1)

		ti := tagInfo{
			statement: statement,
			newstmt:   newstmt,
			attr:      attr,
			tagBegin:  tagBegin,
			tagEnd:    tagEnd,
			srcBegin:  srcBegin,
			srcEnd:    srcEnd,
			dstBegin:  dstBegin,
			dstEnd:    dstEnd,
			dstVal:    dstVal,
		}
		tags = append(tags, ti)
	}
	return tags
}

func (me *htmlServeMux) buildTagTextInfo(content, attr string, tags []tagTextInfo) []tagTextInfo {
	pos := 0
	dstBeginPre := tagAttrPre + attr + tagAttrPost
	for {
		dstBegin := strings.IndexStart(content, dstBeginPre, pos)
		if dstBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		dstEnd := strings.IndexStart(content, tagAttrEnd, dstBegin+len(dstBeginPre))
		if dstEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagBegin := strings.IndexForward(content, tagBeginPre, dstBegin)
		if tagBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagEnd := strings.IndexStart(content, tagTextEndPre, dstEnd)
		if tagEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		srcBegin := strings.IndexStart(content, tagEndPre, dstEnd)
		if srcBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		srcEnd := tagEnd
		pos = tagEnd
		statement := strings.Slice(content, tagBegin, tagEnd+len(tagTextEndPre))
		newstmt := statement
		dstVal := strings.Slice(content, dstBegin+len(dstBeginPre), dstEnd)
		if strings.IsNotBlank(dstVal) {
			filename := Conf.Templates.Directory + dstVal
			if files.IsExist(filename) {
				if c, err := files.Read(filename); err == nil {
					title := strings.Slice(content, srcBegin+len(tagBeginPre), srcEnd)
					newstmt = strings.Replace(newstmt, title, title+c, 1)
					dst := dstBeginPre + dstVal + tagAttrEnd
					newstmt = strings.Replace(newstmt, dst, "", 1)
				}
			}
		}

		tti := tagTextInfo{
			statement: statement,
			newstmt:   newstmt,
			attr:      attr,
			tagBegin:  tagBegin,
			tagEnd:    tagEnd,
			srcBegin:  srcBegin,
			srcEnd:    srcEnd,
			dstBegin:  dstBegin,
			dstEnd:    dstEnd,
			dstVal:    dstVal,
		}
		tags = append(tags, tti)
	}
	return tags
}