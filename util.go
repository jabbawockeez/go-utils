package util

import (
	"fmt"
	"unicode"
	//"errors"
	"strconv"
	"encoding/json"
    "net/url"
    "path"
	"github.com/axgle/mahonia"
	"strings"
)

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	src = strings.TrimSpace(src)
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

// Mute doc: http://blog.vladimirvivien.com/2014/03/hacking-go-filter-values-from-multi.html
func Mute(a ...interface{}) []interface{} {
	return a
}

func GenURL(args ...string) string {
    u, err := url.Parse(args[0])
    if err != nil {
        fmt.Println(err)
    }

    p := path.Join(args[1:]...)
    u.Path = path.Join(u.Path, p)

    return u.String()
}

func StructToMap(s interface{}) map[string]interface{} {
    var m map[string]interface{}
    j, _ := json.Marshal(s)
    json.Unmarshal(j, &m)

    return m
}

func EscapeRegex(src string) string {
    res := ""

    for _, i := range src {
        if unicode.IsLetter(i) || unicode.IsDigit(i) {
            res += string(i)
        } else {
            res += `\\` + string(i)
        }
    }
    //fmt.Println(res)
    return res
}

func Concat(args ...interface{}) string {
    var s string
    
    for _, i := range args {
        s += Str(i)
        // switch i.(type) {
        // case int:
        //     s += strconv.Itoa(i.(int))
        // case []byte:
        //     s += string(i.([]byte))
        // case string:
        //     s += i.(string)
        // }
    }
    return s
}

func Str(v interface{}) string {
    var s string

    switch v.(type) {
    case int:
         s = strconv.Itoa(v.(int))
    case []byte:
         s = string(v.([]byte))
    }

    return s
}
