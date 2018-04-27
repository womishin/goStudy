package iniconfig

import (
	"strings"
	"fmt"
	"reflect"
	"errors"
	"strconv"
)

func Marshal(i interface{}) (data []byte, err error) {
	return
}

func UnMarshal(data []byte, i interface{}) (err error) {
	var lastFieldName string
	lineArr := strings.Split(string(data), "\n")
	t := reflect.TypeOf(i)
	//判断是否是指针
	if t.Kind() != reflect.Ptr {
		err := errors.New("please pass address")
		return err
	}
	typeStruct := t.Elem()
	//判断是否是结构体
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return err
	}
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		//判断是否为空或者是否为备注
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			lastFieldName, err = parsesection(line, typeStruct, index)
			if err != nil {
				err = fmt.Errorf("%v lineNo:%d", err, index+1)
				return
			}
			continue
		}

		err = parseItem(lastFieldName, line, i)
		if err != nil {
			err = fmt.Errorf("%v lineNo:%d", err, index+1)
			return
		}
	}
	return
}

func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}

	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])

	if len(key) == 0 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}

	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFieldName)

	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFieldName)
		return
	}

	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}

	if len(keyFieldName) == 0 {
		return
	}

	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}

	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)

	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		intVal, errRet := strconv.ParseUint(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			return
		}

		fieldValue.SetFloat(floatVal)

	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
	}

	return
}

func parsesection(line string, typeStruct reflect.Type, index int) (fieldName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error, invalid section:%s, lineNo:%d", line, index+1)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s, lineNo:%d", line, index+1)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		space := strings.TrimSpace(line[1 : len(line)-1])
		if len(space) == 0 {
			err = fmt.Errorf("syntax error, invalid section:%s, lineNo:%d", line, index+1)
			return
		}

		for i := 0; i < typeStruct.NumField(); i++ {
			field := typeStruct.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == space {
				fieldName = field.Name
				break
			}
		}
	}
	return
}
