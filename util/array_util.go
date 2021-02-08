package util

import (
	"reflect"
)

func init() {

}

//Includes 判断元素是否在列表里
func Includes(array interface{}, item interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(array)
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(item, val.Index(i).Interface()) {
				return true
			}
		}
		break

	}
	return false
}

//MakeRange 生成区间
//interval 间隔值
//intervalTotal 间隔的数量
func MakeRange(interval int, intervalTotal int) []int {
	list := []int{}
	index := 0

	for {
		index ++
		list = append(list, interval*index)
		if index >= intervalTotal {
			break
		}
	}

	return list
}

// func List2Map(list []interface{},data map[string]interface{}){
// 	tmp := make(map[string]interface{})

// 	for index,val := range list {

// 	}
// }