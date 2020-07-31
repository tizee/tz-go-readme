package parsers

import (
	"fmt"
	"reflect"
	"strconv"
	"tz-go-readme/mdblock"
)


type DefaultParser struct {
}

func (p DefaultParser) Parse(url string) ([]byte,error)  {
    return nil,nil
}

func init()  {
    var parser DefaultParser 
    mdblock.Register("default",parser)
}

// helper
func Display(name string, x interface {})  {
	fmt.Printf("Display %s (%T):\n",name,x)	
	display(name,reflect.ValueOf(x))
}

func display(name string, x reflect.Value)  {
	switch x.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid", name)
	case reflect.Slice, reflect.Array:
		for i := 0; i < x.Len(); i++ {
			display(fmt.Sprintf("%s[%d]",name,i),x.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s",name,x.Type().Field(i).Name)
			display(fieldPath,x.Field(i))
		}
	case reflect.Map:
		for _, val := range x.MapKeys() {
			// need format key value
			display(fmt.Sprintf("%s[%s]",name,formatAtom(val)),x.MapIndex(val))
		}
	case reflect.Ptr:
		if x.IsNil(){
			fmt.Printf("%s = nil\n",name)
		} else{
			display(fmt.Sprintf("(*%s)",name),x.Elem())
		}
	case reflect.Interface:
		if x.IsNil(){
			fmt.Printf("%s = nil\n",name)
		} else{
			fmt.Printf("%s.type = %s\n",name,x.Elem().Type())
			display(name+ ".value", x.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", name, formatAtom(x))
	}
}

// forma
func formatAtom(v reflect.Value) string  {
	switch v.Kind(){
	case reflect.Invalid:
		return "invalid"
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(v.Int(),10)
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return strconv.FormatUint(v.Uint(),10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan,reflect.Func,reflect.Ptr,reflect.Slice,reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64 (v.Pointer()),16)
	default:
		return v.Type().String() + " value"
	}
	
}