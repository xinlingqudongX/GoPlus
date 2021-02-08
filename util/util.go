package util

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

)

func init() {

}

//GetType 获取类结构的名称
func GetType(class interface{}) string {
	if t := reflect.TypeOf(class); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

//GetField 获取类的私有属性
func GetField(class interface{}, name string) interface{} {
	a := reflect.ValueOf(class).Elem().FieldByName(name)
	return a.Interface()
}

//SetField 设置类实例的属性值
func SetField(classInst interface{}, name string, val interface{}) bool {

	classVal := reflect.ValueOf(classInst)
	if classVal.Kind() != reflect.Ptr {
		return false
	}

	classVal.Elem()
	classElem := classVal.Elem()
	if classElem.Kind() != reflect.Struct {
		classElem = classElem.Elem()
	}

	tmp := reflect.New(classElem.Type()).Elem()
	tmp.Set(classElem)
	field := tmp.FieldByName(name)
	if !field.IsValid() || !field.CanSet() {
		return false
	}
	field.Set(reflect.ValueOf(val))
	classVal.Elem().Set(tmp)

	return true
}

// 获取类的方法
// func GetMethod(class interface{}, name string) func() {
// 	a := reflect.ValueOf(class).Elem().MethodByName(name);
// 	return a.Call;
// }

//		传递一个结构对象 没鸟用 根本不支持
// func NewPacket(pa interface{}) *interface{}{
// 	packet := &(pa);
// 	SetField(packet,"Pt",GetType(packet));
// 	packet.CSPacket = &CSPacket{};
// 	packet.SCPacket = &SCPacket{};
// 	return packet
// }

//CopyClass 参数类型和返回类型相同，地址同样返回地址
//接受参数类型是结构体或者地址都行
func CopyClass(class interface{}) interface{} {
	//	第一版复制无效
	var val reflect.Value

	// fmt.Println(reflect.ValueOf(class).Kind())

	// if reflect.ValueOf(class).Kind() != reflect.Ptr {
	// 	val = reflect.ValueOf(class)
	// } else {
	// 	val = reflect.ValueOf(&class).Elem()
	// }

	// inst := val.Interface()

	//	第二版也无效，信息丢失，类型不对
	// newInst := reflect.New(reflect.TypeOf(class))
	// inst := newInst.Elem().Interface()

	//	第三版
	if reflect.ValueOf(class).Kind() != reflect.Ptr {
		val = reflect.ValueOf(class)
	} else {
		val = reflect.ValueOf(&class).Elem()
	}

	newInst := reflect.New(val.Type()).Interface()

	return newInst
}

//Instance 判断类型是否相同
func Instance(a interface{}, b interface{}) bool {
	return GetType(a) == GetType(b)
}

//SetPacket 设置Pt属性
func SetPacket(inst *interface{}) *interface{} {
	SetField(inst, "Pt", GetType(inst))
	return inst
}

//Struct2Map 结构体转字典
func Struct2Map(structClass interface{}, tag string) map[string]interface{} {
	data := make(map[string]interface{})

	rType := reflect.TypeOf(structClass)
	rVal := reflect.ValueOf(structClass)

	//	类型判断
	switch rType.Kind() {
	case reflect.Ptr:
		rType = rType.Elem()
		rVal = rVal.Elem()
		break
	case reflect.Map:
		return structClass.(map[string]interface{})
	default:
		break
	}

	for i := 0; i < rVal.NumField(); i++ {
		keyField := rType.Field(i)
		valField := rVal.Field(i)
		if len(tag) > 0 {
			if _, ok := keyField.Tag.Lookup(tag); ok {
				switch valField.Kind() {
				case reflect.Int:
					data[Lowerfirst(keyField.Name)] = valField.Int()
					break
				default:
					data[Lowerfirst(keyField.Name)] = valField.String()
					break
				}
			} else {
				continue
			}
		} else {

			switch valField.Kind() {
			case reflect.Int:
				data[Lowerfirst(keyField.Name)] = valField.Int()
				break
			default:
				data[Lowerfirst(keyField.Name)] = valField.String()
				break
			}
		}

	}

	return data
}

//Struct2Array 多个结构对象转列表
func Struct2Array(structClassArray ...interface{}) []interface{} {

	array := []interface{}{}
	for _, structClass := range structClassArray {

		rVal := ReflectStruct(structClass)

		length := rVal.NumField()

		for i := 0; i < length; i++ {
			// keyField := rType.Field(i)
			valField := rVal.Field(i)
			if valField.CanAddr() {
				array = append(array, valField.Addr().Interface())
			}
		}
	}

	return array
}

//ReflectStruct 获取反射的值
func ReflectStruct(any interface{}) reflect.Value {

	var rVal reflect.Value

	rVal, ok := any.(reflect.Value)
	if !ok {
		rVal = reflect.ValueOf(any)
	}

	switch rVal.Kind() {
	case reflect.Interface:
		rVal = rVal.Elem()
		break
	case reflect.Ptr:
		rVal = rVal.Elem()
		break
	case reflect.Struct:
		return rVal
	default:
		return rVal
	}

	if rVal.Kind() != reflect.Struct {
		// fmt.Println(rVal.Kind())
		return ReflectStruct(rVal)
	} else {
		// fmt.Println(rVal.CanAddr())
		return rVal
	}
}

//Valid 判断类是否有效
//空的字符串或者空白填充的字符串也算是无效
func Valid(any interface{}) bool {
	if any == nil {
		return false
	}

	switch any.(type) {
	case int, int16, int8, int64:
		val := any.(int)
		return val != 0
	case string:
		val := any.(string)
		val = strings.TrimSpace(val)
		return len(val) > 0
	case map[string]interface{}:
		val := any.(map[string]interface{})
		return len(val) > 0
	case map[int]interface{}:
		val := any.(map[int]interface{})
		return len(val) > 0
	case []string:
		val := any.([]string)
		return len(val) > 0
	case []int:
		val := any.([]int)
		return len(val) > 0
	default:

		println(any)
		break
	}

	return false
}

//Now 当前时间 毫秒
func Now() int {
	return int(time.Now().Local().UnixNano() / 1e6)
}

//StartOf 获取时间当天的开始时间
func StartOf(t time.Time) int {
	// cTime := time.Now();
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return int(zeroTime.UnixNano() / 1e6)
}

//EndOf 获取时间当天的结束时间
func EndOf(cTime int) int {
	// cTime := time.Now();
	t := time.Unix(int64(cTime/1000), 0)
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	zeroTime = zeroTime.Add(23 * time.Hour)
	zeroTime = zeroTime.Add(59 * time.Minute)
	zeroTime = zeroTime.Add(59 * time.Second)
	zeroTime = zeroTime.Add(999 * time.Millisecond)

	return int(zeroTime.UnixNano() / 1e6)
}

//TimeRange 获取时间的周期
func TimeRange(start int, end int) []int {

	timeArray := make([]int, 0)

	endTime := time.Unix(int64(end/1000), 0)
	startTime := time.Unix(int64(start/1000), 0)

	day := int(endTime.Sub(startTime).Hours() / 24)
	now := time.Now()
	now.Add(2 * time.Hour)

	for i := 0; i < day; i++ {

		add, _ := time.ParseDuration(fmt.Sprintf("%vh", i*24))
		d := startTime.Add(add)
		timeArray = append(timeArray, int(d.Local().UnixNano()/1e6))
	}
	timeArray = append(timeArray, StartOf(endTime))

	return timeArray
}

//Debug 打印行号文件信息
func Debug(any ...interface{}) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	pathArray := strings.Split(pwd, `\`)
	pathName := pathArray[len(pathArray)-1]
	//	堆栈层数
	const MaxStep int = 10
	for step := 1; step <= MaxStep; step++ {
		funcPtr, file, line, ok := runtime.Caller(step)
		if !ok {
			continue
		}

		funcName := runtime.FuncForPC(funcPtr).Name()
		strArray := strings.Split(funcName, "/")
		for _, name := range strArray {
			if name != pathName {
				continue
			}
			// fmt.Println(`堆栈层数`,step)
			// fmt.Print(`Debug:`)
			fmt.Println(fmt.Sprintf(`Debug:%v:%v:%v`, file, funcName, line))
			// fmt.Println(any...)
		}

	}

	fmt.Println(any...)
}
