package users

import "github.com/kkakoz/gquery"

var Table = "users"

var Name = gquery.C("users.name")

var Age = gquery.C("users.age")

var ClassId = gquery.C("users.class_id")

func NameEQ(name string) *gquery.Expr {
	return Name.EQ(name)
}

func AgeEQ(age int) *gquery.Expr {
	return Age.EQ(age)
}

type userAssign struct {
	val map[string]any
}

func (u *userAssign) Assign() map[string]any {
	return u.val
}

func SetName(name string) *userAssign {
	return &userAssign{val: map[string]any{
		"name": name,
	}}
}
