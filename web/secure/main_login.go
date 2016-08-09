// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/secure/internal"
	"gopkg.in/goyy/goyy.v0/web/session"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func Login(c xhttp.Context, loginName, passwd string) error {
	if strings.IsBlank(passwd) {
		return errors.NewNotBlank("passwd")
	}
	npasswd := EncryptPasswd(passwd)
	checked := internal.UserMgr.CheckPasswd(loginName, npasswd)
	if checked {
		u, err := internal.UserMgr.SelectUser(loginName)
		if err != nil {
			return err
		}
		ps, err := internal.PermissionMgr.SelectPermission(u.Id())
		if err != nil {
			return err
		}
		p := session.Principal{
			Id:          u.Id(),
			Name:        u.Name(),
			LoginName:   u.LoginName(),
			LoginTime:   times.NowUnixStr(),
			Permissions: ps,
		}
		setCookies(c, p)
		return c.Session().SetPrincipal(p)
	} else {
		return errors.New(i18N.Message("err.login"))
	}
	return nil
}

func setCookies(c xhttp.Context, p session.Principal) {
	ps := base64.StdEncoding.EncodeToString([]byte(p.Permissions))
	// GSESSIONUSER
	ucookie := &http.Cookie{
		Name:     "GSESSIONUSER",
		Value:    p.LoginName,
		Path:     xhttp.Conf.Session.Path,
		Domain:   xhttp.Conf.Session.Domain,
		Secure:   xhttp.Conf.Session.Secure,
		HttpOnly: xhttp.Conf.Session.HttpOnly,
	}
	http.SetCookie(c.ResponseWriter(), ucookie)
	// GSESSIONN
	pslen := len(ps)
	pstimes := pslen / 4000
	loop := int(pstimes)
	ncookie := &http.Cookie{
		Name:     "GSESSIONN",
		Value:    strconv.Itoa(loop),
		Path:     xhttp.Conf.Session.Path,
		Domain:   xhttp.Conf.Session.Domain,
		Secure:   xhttp.Conf.Session.Secure,
		HttpOnly: xhttp.Conf.Session.HttpOnly,
	}
	http.SetCookie(c.ResponseWriter(), ncookie)
	// GSESSION*
	for i := 0; i <= loop; i++ {
		logger.Println(i, loop, 4000*i, 4000*(i+1))
		psmax := 4000 * (i + 1)
		if psmax > pslen {
			psmax = pslen
		}
		pcookie := &http.Cookie{
			Name:     "GSESSION" + strconv.Itoa(i),
			Value:    ps[4000*i : psmax],
			Path:     xhttp.Conf.Session.Path,
			Domain:   xhttp.Conf.Session.Domain,
			Secure:   xhttp.Conf.Session.Secure,
			HttpOnly: xhttp.Conf.Session.HttpOnly,
		}
		http.SetCookie(c.ResponseWriter(), pcookie)
	}
}

func Logout(c xhttp.Context) error {
	return c.Session().ResetPrincipal()
}
