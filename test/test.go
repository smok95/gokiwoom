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

			codelist := strings.Split(kw.GetCodeListByMarket("0"), ";")

			fmt.Println(len(codelist), codelist)
		} else {
			fmt.Printf("%d) 접속 실패\n", errCode)
			kw.Disconnect()
		}
	})

	kw.SetOnReceiveMsg(func(scrNo, rqName, trCode, msg string) {
		fmt.Println("OnReceiveMsg", scrNo, rqName, trCode, msg)
	})

	kw.SetOnReceiveTrData(func(scrNo, rqName, trCode, recordName, prevNext string,
		dataLength int32, errorCode, message, splmMsg string) {
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
