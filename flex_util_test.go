package flex

import (
	"fmt"
)

func FmtString(object interface{}) string {
	return fmt.Sprintf("%v", object)
}

func FmtStringPlus(object interface{}) string {
	return fmt.Sprintf("%+v", object)
}

func FmtStringSharp(object interface{}) string {
	return fmt.Sprintf("%#v", object)
}
