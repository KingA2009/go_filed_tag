package main

import (
	"fmt"
	"reflect"
)
const tagName = "validate"

type User struct {
  Name  string `validate:"true"`
  Password string `validate:"true"`
}

func main() {
  user := User{
    Name:  "John Doe",
	Password: "password",
  }
  t := reflect.TypeOf(user)
  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    tag := field.Tag.Get(tagName)

    fmt.Printf("%d. %v %v,'%v'\n", i+1, field.Name, field.Type.Name(), tag)
  }
}