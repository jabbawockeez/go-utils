package util

import (
	"fmt"
	"log"
	"unicode"
	//"errors"
	"strconv"
	"encoding/json"
    "net/url"
    "path"
	"github.com/axgle/mahonia"
	"strings"
    
    "github.com/astaxie/beego/logs"
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

func GenURL(args ...interface{}) string {
    var args_str = make([]string, len(args))
    for i, v := range args {
        args_str[i] = Str(v)
    }

    u, err := url.Parse(args_str[0])
    if err != nil {
        fmt.Println(err)
    }

    p := path.Join(args_str[1:]...)
    u.Path = path.Join(u.Path, p)

    result, err := url.PathUnescape(u.String())
    if err != nil {
        panic(err)
    }

    return result
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

func Join(sep string, args ...interface{}) string {
    /*
        return a string that join 
        every argument passed in, 
        no matter what type they were.
    */

    s := []string{}
    
    for _, i := range args {
        s = append(s, Str(i))
    }

    return strings.Join(s, sep)
}

func Str(val interface{}) string {
    var s string

    switch val.(type) {
    case int:
        s = strconv.Itoa(val.(int))
    case []byte:
        s = string(val.([]byte))
    default:
        s = fmt.Sprintf("%v", val)
    }

    return s
}

// print
func P(args ...interface{}) {
    //log.Println(args...)
    logs.Info(args)
}

// printf
func Pf(fmt string, args ...interface{}) {
    log.Printf(fmt, args...)
}

