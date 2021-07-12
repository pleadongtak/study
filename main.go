package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func test(s interface{}) (res string, err error) {

	if s != s.(string) {
		err1 := errors.New("error1 occurred")
		return res, err1
	}
	r := errors.New("error2 occurred")
	return res, r
}

func searchBracket(s string) (res bool, err error) { //마지막 글자확인

	fmt.Println(lastChar(s), "마지막글자")
	if lastChar(s) != "]" || lastChar(s) != `"` || lastChar(s) != " " {
		return true, nil
	}

	var ErrInvalidParam = errors.New(" invalid parameter")
	return false, ErrInvalidParam

}

func formatString(configTypeCheck interface{}) string {
	// 인터페이스에 저장된 타입에 따라 case 실행
	switch configTypeCheck.(type) {
	case int: // configTypeCheck int형이라면
		intType := configTypeCheck.(int) // int형으로 값을 가져옴

		return strconv.Itoa(intType) // strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	case float32: // configTypeCheck float32형이라면
		float32Type := configTypeCheck.(float32) // float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(float32Type), 'f', -1, 32)
		// strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case float64: // configTypeCheck float64형이라면
		float64Type := configTypeCheck.(float64) // float64형으로 값을 가져옴
		return strconv.FormatFloat(float64Type, 'f', -1, 64)
		// strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case string: // configTypeCheck가  string이라면
		stringType := configTypeCheck.(string) // string으로 값을 가져옴
		return stringType                      // string이므로 그대로 리턴

	default:
		return "Error"
	}
}

//문자 앞뒤 글자 자르기
//https://stackoverflow.com/questions/57004213/how-in-golang-to-remove-the-last-letter-from-the-string
func removeLastRune(s string) string {
	r := []rune(s)

	return string(r[:len(r)-1])
}

func removeFirstRune(s string) string {
	r := []rune(s)
	return string(r[1:len(r)])
}

func lastChar(s string) string {
	r := []rune(s)

	return string(r[len(r)-1 : len(r)])
}
func FilterString(s string) string {

	var removeLastChar = removeLastRune(s)
	if len(removeLastChar) == 0 {
		return ""
	} else {
		var removeFirstChar = removeFirstRune(removeLastChar)

		return removeFirstChar
	}
}

func errorCheck(s interface{}) (res interface{}, err error) {
	parserError := errors.New("error occurred")

	return s, parserError

}
func main() {

	fmt.Println("---------------------------")
	//parsing 하기위한 문자열
	var unixFrame string = `[server]
			"key1" = "value1"
			"key2" = "value2"
			"key3" = "value3"
		`

	configTypeChecking := formatString(unixFrame)
	//fmt.Println(test(unixFrame))

	fmt.Println("---------------------------")
	m := make(map[string]interface{})
	keyFlag := false
	keyTrack := "" // 키,벨류 키벨류로 들어가는것을 임시로 저장
	//titleCount := 0

	// 1. 개행 및 띄어쓰기 기준으로 분리
	temp := strings.Fields(configTypeChecking)

	// 개행과 띄어쓰기가 제거된 temp라는 변수안에서 반복문을 돈다.

	for i := 0; i < len(temp); i++ {
		var tmp = temp[i] // 각 인덱스마다의 string값을 tmp라는 변수에 저장
		var firstChar = tmp[0:1]
		rr, err := searchBracket(tmp)
		if err != nil && rr == true {
			if firstChar == "[" {

				var titleString = FilterString(tmp)

				m[titleString] = "" //null 이면 title
				keyFlag = true
				//fmt.Println(m[titleString])

			}

		} else {
			var resultString = FilterString(tmp)
			if keyFlag {

				if len(resultString) > 0 {
					m[resultString] = ""
					keyFlag = false
					keyTrack = resultString
				}

			} else {
				if len(resultString) > 0 {
					m[keyTrack] = resultString
					keyFlag = true
					keyTrack = ""
				}
			}
		}
	}
	for k, v := range m {
		if v == "" {
			fmt.Println(k)
		} else {
			fmt.Println(k, v)
		}
	}

	fmt.Println("\n---------------------------")

}
