package bstr

import (
	"fmt"
	"reflect"
	"strings"
)

// DD 用于打印的函数
func DD(values ...interface{}) {
	fmt.Println("")
	for _, value := range values {
		dump("", reflect.ValueOf(value), 0)
	}
	fmt.Println("")
}

func dump(name interface{}, value reflect.Value, deepin int) {
	if value.IsValid() && value.CanInterface() {
		t := value.Type()
		switch t.Kind() {
		case reflect.Array, reflect.Slice:
			printType(value.Interface(), deepin, "=>{")
			deepin++
			for i := 0; i < value.Len(); i++ {
				dump("", value.Index(i), deepin)
			}
			deepin--
			printOrigin("}", deepin)

		case reflect.Map:
			keys := value.MapKeys()
			printType(value.Interface(), deepin, "=>{")
			deepin++
			for _, key := range keys {
				dump(key.Interface(),
					reflect.ValueOf(value.MapIndex(key).Interface()), deepin)
			}
			deepin--
			printOrigin("}", deepin)

		case reflect.Ptr:
			printType(value.Interface(), deepin, "=>(")
			deepin++
			dump("", value.Elem(), deepin)
			deepin--
			printOrigin(")", deepin)
		case reflect.Struct:
			printType(value.Interface(), deepin, "=>{")
			deepin++
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				dump(field.Name, value.FieldByIndex([]int{i}), deepin)
			}
			deepin--
			printOrigin("}", deepin)
		case reflect.Interface:
			printValue(name, value.Interface(), deepin)
			deepin++
			dump(name, reflect.ValueOf(value.Interface()), deepin)
			deepin--
			printOrigin("}", deepin)
		default:
			printValue(name, value.Interface(), deepin)
		}
	} else {
		printValue(name, "", deepin)
	}
}

func printType(value interface{}, deepin int, separator string) {
	printDeepin(deepin)
	fmt.Printf("(%T) %s\n", value, separator)
}

func printValue(name interface{}, value interface{}, deepin int) {
	printDeepin(deepin)
	if name != "" {
		fmt.Printf("(%T)%[1]v: ", name)
		fmt.Printf("(%T)%#v\n", value, value)
	} else {
		fmt.Printf("(%T)%#v\n", value, value)
	}
}

func printDeepin(deepin int) {
	s := strings.Repeat(" ", deepin*4)
	fmt.Print(s)
}

func printOrigin(value interface{}, deepin int) {
	printDeepin(deepin)
	fmt.Println(value)
}
