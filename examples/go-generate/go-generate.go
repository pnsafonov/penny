package gogenerate

import "github.com/pnsafonov/penny/generic"

//go:generate penny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int ValueType=string,int"

type KeyType generic.Type
type ValueType generic.Type

func GetDefault() KeyType {
	return KeyType(nil)
}

type KeyTypeValueTypeMap map[KeyType]ValueType

func NewKeyTypeValueTypeMap() map[KeyType]ValueType {
	return make(map[KeyType]ValueType)
}
