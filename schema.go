package gquery

import "reflect"

var typeMap = map[string]schema{}

type schema struct {
	modelName   string
	packageName string
	fields      map[string]field
}

type field struct {
	column string
	kind   reflect.Kind
}

func parseModel(val any) schema {
	typeOf := reflect.TypeOf(val)
	key := typeOf.PkgPath() + typeOf.Name()
	res, ok := typeMap[key]
	if ok {
		return res
	}
	return schema{}
	//if typeOf.Kind() == reflect.Struct {
	//	for i := 0; i < typeOf.NumField(); i++ {
	//		field := typeOf.Field(i)
	//	}
	//}
}
