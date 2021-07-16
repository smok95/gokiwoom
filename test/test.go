package main

import (
	"fmt"

	kw "github.com/smok95/gokiwoom"
)

func main() {

	kw.SetOnEventConnect(func(errCode int) {
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
		recordName string, prevNext string, dataLength int, errorCode string,
		message string, splmMsg string) {
		fmt.Printf("OnReceiveTrData, scrNo=%s, rqName=%s, trCode=%s, recordName=%s\n",
			scrNo, rqName, trCode, recordName)
		temp := kw.GetCommData(trCode, recordName, 0, "종목명")
		fmt.Println(temp)
	})

	ret := kw.CommConnect()
	fmt.Println("CommConnect", ret)

	kw.Wait()
}
