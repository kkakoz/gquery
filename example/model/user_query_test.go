package model

import (
	"fmt"
	. "github.com/kkakoz/gquery"
	"github.com/kkakoz/gquery/example/model/class"
	"github.com/kkakoz/gquery/example/model/users"
	"reflect"
	"testing"
)

func TestUserQuery(t *testing.T) {

	userList, err := Query[User]().Where(C("name").EQ("zhangsan")).ListRes().Result()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(userList)

	Query[User]().Where(EQ("name", "zhangsan"), EQ("age", 12))

	Query[User]().Where()

	userList, err = Query[User]().Where(users.NameEQ("name"), users.AgeEQ(12), EQ(users.Name, "zhangsan")).
		Join(class.Table, On(EQ(users.ClassId, class.Id))).
		Join(class.Table, On(EQ(users.ClassId, class.Id))).Select(users.Name).List()
	if err != nil {
		t.Fatal(err)
	}

	err = Query[User]().Where(users.NameEQ("zhangsan"), users.AgeEQ(12)).Update(users.SetName("zhangsan")).Err()

	fmt.Println(userList)
}

func TestC(t *testing.T) {
	fmt.Println(reflect.TypeOf(users.Name).Name())
}
