package main

import (
	"fmt"
	"strings"

	kw "github.com/smok95/gokiwoom"
)

func main() {

	kw.SetOnEventConnect(func(errCode int32) {
		if errCode == 0 {
			fmt.Println("접속 성공")

			kw.SetInputValue("종목코드", "000660")
			ret := kw.CommRqData("jktest", "opt10001", 0, "jktestscr")
			fmt.Printf("CommRqData=%d\n", ret)
		} else {
			fmt.Println("접속 실패")
		}
	})

	kw.SetOnReceiveTrData(func(scrNo string, rqName string, trCode string,
		recordName string, prevNext string, dataLength int32, errorCode string,
		message string, splmMsg string) {
		fmt.Printf("OnReceiveTrData, scrNo=%s, rqName=%s, trCode=%s, recordName=%s\n",
			scrNo, rqName, trCode, recordName)
		name := kw.GetCommData(trCode, recordName, 0, "종목명")
		fmt.Println(name)

		code := kw.GetCommData(trCode, recordName, 0, "종목코드")
		code = strings.TrimSpace(code)
		name2 := kw.GetMasterCodeName(code)
		fmt.Printf("GetMasterCodeName('%s') => '%s'\n", code, name2)
	})

	ret := kw.CommConnect()
	fmt.Println("CommConnect", ret)

	kw.Wait()
}
