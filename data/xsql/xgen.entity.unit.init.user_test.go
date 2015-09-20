// generated by xgen -- DO NOT EDIT
package xsql_test

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"time"
)

var (
	USER          = schema.TABLE("users")
	USER_ID       = USER.PRIMARY("id")
	USER_CODE     = USER.COLUMN("code")
	USER_NAME     = USER.COLUMN("name")
	USER_PASSWORD = USER.COLUMN("password")
	USER_MEMO     = USER.COLUMN("memo")
	USER_GENRE    = USER.COLUMN("genre")
	USER_STATUS   = USER.COLUMN("status")
	USER_ROLES    = USER.COLUMN("roles")
	USER_POSTS    = USER.COLUMN("posts")
	USER_ORG      = USER.COLUMN("org")
	USER_AREA     = USER.COLUMN("area")
	USER_CREATER  = USER.CREATER("creater")
	USER_CREATED  = USER.CREATED("created")
	USER_MODIFIER = USER.MODIFIER("modifier")
	USER_MODIFIED = USER.MODIFIED("modified")
	USER_VERSION  = USER.VERSION("version")
	USER_DELETION = USER.DELETION("deletion")
)

func NewUser() *User {
	e := &User{}
	e.init()
	return e
}

func (me *User) Id() string {
	return me.id.Value()
}

func (me *User) SetId(v string) {
	me.id.SetValue(v)
}

func (me *User) Code() string {
	return me.code.Value()
}

func (me *User) SetCode(v string) {
	me.code.SetValue(v)
}

func (me *User) Name() string {
	return me.name.Value()
}

func (me *User) SetName(v string) {
	me.name.SetValue(v)
}

func (me *User) Password() string {
	return me.password.Value()
}

func (me *User) SetPassword(v string) {
	me.password.SetValue(v)
}

func (me *User) Memo() string {
	return me.memo.Value()
}

func (me *User) SetMemo(v string) {
	me.memo.SetValue(v)
}

func (me *User) Genre() string {
	return me.genre.Value()
}

func (me *User) SetGenre(v string) {
	me.genre.SetValue(v)
}

func (me *User) Status() string {
	return me.status.Value()
}

func (me *User) SetStatus(v string) {
	me.status.SetValue(v)
}

func (me *User) Roles() string {
	return me.roles.Value()
}

func (me *User) SetRoles(v string) {
	me.roles.SetValue(v)
}

func (me *User) Posts() string {
	return me.posts.Value()
}

func (me *User) SetPosts(v string) {
	me.posts.SetValue(v)
}

func (me *User) Org() string {
	return me.org.Value()
}

func (me *User) SetOrg(v string) {
	me.org.SetValue(v)
}

func (me *User) Area() string {
	return me.area.Value()
}

func (me *User) SetArea(v string) {
	me.area.SetValue(v)
}

func (me *User) Creater() string {
	return me.creater.Value()
}

func (me *User) SetCreater(v string) {
	me.creater.SetValue(v)
}

func (me *User) Created() time.Time {
	return me.created.Value()
}

func (me *User) SetCreated(v time.Time) {
	me.created.SetValue(v)
}

func (me *User) Modifier() string {
	return me.modifier.Value()
}

func (me *User) SetModifier(v string) {
	me.modifier.SetValue(v)
}

func (me *User) Modified() time.Time {
	return me.modified.Value()
}

func (me *User) SetModified(v time.Time) {
	me.modified.SetValue(v)
}

func (me *User) Version() int {
	return me.version.Value()
}

func (me *User) SetVersion(v int) {
	me.version.SetValue(v)
}

func (me *User) Deletion() int {
	return me.deletion.Value()
}

func (me *User) SetDeletion(v int) {
	me.deletion.SetValue(v)
}

func (me *User) init() {
	me.table = USER

	me.id.SetColumn(USER_ID)
	me.code.SetColumn(USER_CODE)
	me.name.SetColumn(USER_NAME)
	me.password.SetColumn(USER_PASSWORD)
	me.memo.SetColumn(USER_MEMO)
	me.genre.SetColumn(USER_GENRE)
	me.status.SetColumn(USER_STATUS)
	me.roles.SetColumn(USER_ROLES)
	me.posts.SetColumn(USER_POSTS)
	me.org.SetColumn(USER_ORG)
	me.area.SetColumn(USER_AREA)
	me.creater.SetColumn(USER_CREATER)
	me.created.SetColumn(USER_CREATED)
	me.modifier.SetColumn(USER_MODIFIER)
	me.modified.SetColumn(USER_MODIFIED)
	me.version.SetColumn(USER_VERSION)
	me.deletion.SetColumn(USER_DELETION)


	me.id.SetField(entity.DefaultField())
	me.code.SetField(entity.DefaultField())
	me.name.SetField(entity.DefaultField())
	me.password.SetField(entity.DefaultField())
	me.memo.SetField(entity.DefaultField())
	me.genre.SetField(entity.DefaultField())
	me.status.SetField(entity.DefaultField())
	me.roles.SetField(entity.DefaultField())
	me.posts.SetField(entity.DefaultField())
	me.org.SetField(entity.DefaultField())
	me.area.SetField(entity.DefaultField())
	me.creater.SetField(entity.DefaultField())
	me.created.SetField(entity.DefaultField())
	me.modifier.SetField(entity.DefaultField())
	me.modified.SetField(entity.DefaultField())
	me.version.SetField(entity.DefaultField())
	me.deletion.SetField(entity.DefaultField())
}

func (me User) New() entity.Interface {
	return NewUser()
}

func (me *User) Get(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.Value()
	case USER_CODE.Name():
		return me.code.Value()
	case USER_NAME.Name():
		return me.name.Value()
	case USER_PASSWORD.Name():
		return me.password.Value()
	case USER_MEMO.Name():
		return me.memo.Value()
	case USER_GENRE.Name():
		return me.genre.Value()
	case USER_STATUS.Name():
		return me.status.Value()
	case USER_ROLES.Name():
		return me.roles.Value()
	case USER_POSTS.Name():
		return me.posts.Value()
	case USER_ORG.Name():
		return me.org.Value()
	case USER_AREA.Name():
		return me.area.Value()
	case USER_CREATER.Name():
		return me.creater.Value()
	case USER_CREATED.Name():
		return me.created.Value()
	case USER_MODIFIER.Name():
		return me.modifier.Value()
	case USER_MODIFIED.Name():
		return me.modified.Value()
	case USER_VERSION.Name():
		return me.version.Value()
	case USER_DELETION.Name():
		return me.deletion.Value()
	default:
		return nil
	}
}

func (me *User) GetPtr(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.ValuePtr()
	case USER_CODE.Name():
		return me.code.ValuePtr()
	case USER_NAME.Name():
		return me.name.ValuePtr()
	case USER_PASSWORD.Name():
		return me.password.ValuePtr()
	case USER_MEMO.Name():
		return me.memo.ValuePtr()
	case USER_GENRE.Name():
		return me.genre.ValuePtr()
	case USER_STATUS.Name():
		return me.status.ValuePtr()
	case USER_ROLES.Name():
		return me.roles.ValuePtr()
	case USER_POSTS.Name():
		return me.posts.ValuePtr()
	case USER_ORG.Name():
		return me.org.ValuePtr()
	case USER_AREA.Name():
		return me.area.ValuePtr()
	case USER_CREATER.Name():
		return me.creater.ValuePtr()
	case USER_CREATED.Name():
		return me.created.ValuePtr()
	case USER_MODIFIER.Name():
		return me.modifier.ValuePtr()
	case USER_MODIFIED.Name():
		return me.modified.ValuePtr()
	case USER_VERSION.Name():
		return me.version.ValuePtr()
	case USER_DELETION.Name():
		return me.deletion.ValuePtr()
	default:
		return nil
	}
}

func (me *User) Table() schema.Table {
	return me.table
}

func (me *User) Type(name string) (entity.Type, bool) {
	switch name {
	case USER_ID.Name():
		return &me.id, true
	case USER_CODE.Name():
		return &me.code, true
	case USER_NAME.Name():
		return &me.name, true
	case USER_PASSWORD.Name():
		return &me.password, true
	case USER_MEMO.Name():
		return &me.memo, true
	case USER_GENRE.Name():
		return &me.genre, true
	case USER_STATUS.Name():
		return &me.status, true
	case USER_ROLES.Name():
		return &me.roles, true
	case USER_POSTS.Name():
		return &me.posts, true
	case USER_ORG.Name():
		return &me.org, true
	case USER_AREA.Name():
		return &me.area, true
	case USER_CREATER.Name():
		return &me.creater, true
	case USER_CREATED.Name():
		return &me.created, true
	case USER_MODIFIER.Name():
		return &me.modifier, true
	case USER_MODIFIED.Name():
		return &me.modified, true
	case USER_VERSION.Name():
		return &me.version, true
	case USER_DELETION.Name():
		return &me.deletion, true
	}
	return nil, false
}

func (me *User) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return USER_ID, true
	case "code":
		return USER_CODE, true
	case "name":
		return USER_NAME, true
	case "password":
		return USER_PASSWORD, true
	case "memo":
		return USER_MEMO, true
	case "genre":
		return USER_GENRE, true
	case "status":
		return USER_STATUS, true
	case "roles":
		return USER_ROLES, true
	case "posts":
		return USER_POSTS, true
	case "org":
		return USER_ORG, true
	case "area":
		return USER_AREA, true
	case "creater":
		return USER_CREATER, true
	case "created":
		return USER_CREATED, true
	case "modifier":
		return USER_MODIFIER, true
	case "modified":
		return USER_MODIFIED, true
	case "version":
		return USER_VERSION, true
	case "deletion":
		return USER_DELETION, true
	}
	return nil, false
}

func (me *User) Columns(filter int) []schema.Column {
	switch filter {
	case entity.ColAll:
		return []schema.Column{
			USER_ID,
			USER_CODE,
			USER_NAME,
			USER_PASSWORD,
			USER_MEMO,
			USER_GENRE,
			USER_STATUS,
			USER_ROLES,
			USER_POSTS,
			USER_ORG,
			USER_AREA,
			USER_CREATER,
			USER_CREATED,
			USER_MODIFIER,
			USER_MODIFIED,
			USER_VERSION,
			USER_DELETION,
		}
	case entity.ColUpdateable:
		cols := make([]schema.Column, 0)
		if me.id.HasUpdate() {
			cols = append(cols, USER_ID)
		}
		if me.code.HasUpdate() {
			cols = append(cols, USER_CODE)
		}
		if me.name.HasUpdate() {
			cols = append(cols, USER_NAME)
		}
		if me.password.HasUpdate() {
			cols = append(cols, USER_PASSWORD)
		}
		if me.memo.HasUpdate() {
			cols = append(cols, USER_MEMO)
		}
		if me.genre.HasUpdate() {
			cols = append(cols, USER_GENRE)
		}
		if me.status.HasUpdate() {
			cols = append(cols, USER_STATUS)
		}
		if me.roles.HasUpdate() {
			cols = append(cols, USER_ROLES)
		}
		if me.posts.HasUpdate() {
			cols = append(cols, USER_POSTS)
		}
		if me.org.HasUpdate() {
			cols = append(cols, USER_ORG)
		}
		if me.area.HasUpdate() {
			cols = append(cols, USER_AREA)
		}
		if me.creater.HasUpdate() {
			cols = append(cols, USER_CREATER)
		}
		if me.created.HasUpdate() {
			cols = append(cols, USER_CREATED)
		}
		if me.modifier.HasUpdate() {
			cols = append(cols, USER_MODIFIER)
		}
		if me.modified.HasUpdate() {
			cols = append(cols, USER_MODIFIED)
		}
		if me.version.HasUpdate() {
			cols = append(cols, USER_VERSION)
		}
		if me.deletion.HasUpdate() {
			cols = append(cols, USER_DELETION)
		}
		return cols
	case entity.ColInsertable:
		cols := make([]schema.Column, 0)
		if me.id.HasInsert() {
			cols = append(cols, USER_ID)
		}
		if me.code.HasInsert() {
			cols = append(cols, USER_CODE)
		}
		if me.name.HasInsert() {
			cols = append(cols, USER_NAME)
		}
		if me.password.HasInsert() {
			cols = append(cols, USER_PASSWORD)
		}
		if me.memo.HasInsert() {
			cols = append(cols, USER_MEMO)
		}
		if me.genre.HasInsert() {
			cols = append(cols, USER_GENRE)
		}
		if me.status.HasInsert() {
			cols = append(cols, USER_STATUS)
		}
		if me.roles.HasInsert() {
			cols = append(cols, USER_ROLES)
		}
		if me.posts.HasInsert() {
			cols = append(cols, USER_POSTS)
		}
		if me.org.HasInsert() {
			cols = append(cols, USER_ORG)
		}
		if me.area.HasInsert() {
			cols = append(cols, USER_AREA)
		}
		if me.creater.HasInsert() {
			cols = append(cols, USER_CREATER)
		}
		if me.created.HasInsert() {
			cols = append(cols, USER_CREATED)
		}
		if me.modifier.HasInsert() {
			cols = append(cols, USER_MODIFIER)
		}
		if me.modified.HasInsert() {
			cols = append(cols, USER_MODIFIED)
		}
		if me.version.HasInsert() {
			cols = append(cols, USER_VERSION)
		}
		if me.deletion.HasInsert() {
			cols = append(cols, USER_DELETION)
		}
		return cols
	}
	return nil
}

func (me *User) Names() []string {
	return []string{
		"id",
		"code",
		"name",
		"password",
		"memo",
		"genre",
		"status",
		"roles",
		"posts",
		"org",
		"area",
		"creater",
		"created",
		"modifier",
		"modified",
		"version",
		"deletion",
	}
}

func (me *User) Value() *User {
	return me
}

func (me *User) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	case "code":
		return me.code.SetString(value)
	case "name":
		return me.name.SetString(value)
	case "password":
		return me.password.SetString(value)
	case "memo":
		return me.memo.SetString(value)
	case "genre":
		return me.genre.SetString(value)
	case "status":
		return me.status.SetString(value)
	case "roles":
		return me.roles.SetString(value)
	case "posts":
		return me.posts.SetString(value)
	case "org":
		return me.org.SetString(value)
	case "area":
		return me.area.SetString(value)
	case "creater":
		return me.creater.SetString(value)
	case "created":
		return me.created.SetString(value)
	case "modifier":
		return me.modifier.SetString(value)
	case "modified":
		return me.modified.SetString(value)
	case "version":
		return me.version.SetString(value)
	case "deletion":
		return me.deletion.SetString(value)
	}
	return nil
}