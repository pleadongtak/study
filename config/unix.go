package config

import (
	"errors"
	"reflect"
	"strconv"
)

//인터페이스는 사용할 메소드들로 묶임
type unixConfigHandler interface {
	GetConfigType(key string) interface{}
	GenerateReturnString(k string) string
	LoadConfig(val string) (bool, error) //읽었으면  오류가있으면 error 처리
}

//구조체를 생성하고
type UnixConfig struct {
	//필요한 정보
	Filename string //path + file name   path 파일이 저장될 경로....
}

//구조체를 리시브로 받는 함수를 생성
func (u UnixConfig) LoadConfig(val string) (bool, error) { // 스트링을 파라미터로 받고 , 그스트링에 길이로 판단

	if len(val) <= 0 {
		return false, errors.New("")
	} else {
		return true, nil
	}
}

//구조체를 리시브로 받는 함수를 생성
func (u UnixConfig) GetConfigType(val string) interface{} {
	if i, err := strconv.ParseInt(val, 10, 64); err == nil {
		return i
	} else if b, err := strconv.ParseBool(val); err == nil {
		return b
	} else {
		return val
	}
}

//구조체를 리시브로 받는 함수를 생성
func (u UnixConfig) GenerateReturnString(k string) string {
	tempReturnResult := ""
	retType := u.GetConfigType(k)
	if reflect.TypeOf(retType).Kind() == reflect.Int64 {
		tempReturnResult += "(INT)"
		tempReturnResult += k
	} else if reflect.TypeOf(retType).Kind() == reflect.Bool {
		tempReturnResult += "(BOOLEAN)"
		tempReturnResult += k
	} else {
		tempReturnResult += "(STRING)"
		tempReturnResult += k
	}
	return tempReturnResult
}

//인터페이스를 사용하기위한 함수.
func NewConfig(filename string) unixConfigHandler {
	//객체를생성할떄 파일을 읽어야한다면 성공실패로해서 객체리턴해도되고 자유롭게 ,
	return &UnixConfig{Filename: filename}

}
