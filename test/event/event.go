package main

import (
	"fmt"
	"strings"

	kw "github.com/smok95/gokiwoom"
)

func main() {
	kw.Initialize(1)

	// 연결상태 이벤트
	kw.SetOnEventConnect(func(errCode int32) {
		if errCode == 0 {
			fmt.Println("연결 성공")

			// 삼성전자 시세 조회
			fmt.Println("삼성전자 시세 조회")
			kw.SetInputValue("종목코드", "005930")
			kw.CommRqData("rqName1", "opt10001", 0, "scrNo1")

		} else {
			fmt.Println("연결 종료")
		}
	})

	// TR응답 이벤트
	kw.SetOnReceiveTrData(func(scrNo, rqName, trCode, recordName, prevNext string,
		dataLength int32, errorCode, message, splmMsg string) {
		fmt.Printf("TrData) scrNo=%s, rqName=%s\n", scrNo, rqName)
		name := kw.GetCommData(trCode, recordName, 0, "종목명")
		price := kw.GetCommData(trCode, recordName, 0, "현재가")
		fmt.Printf("%s 현재가: %s\n", strings.TrimSpace(name), strings.TrimSpace(price))
	})

	// 실시간 이벤트
	kw.SetOnReceiveRealData(func(realKey, realType, realData string) {
		fmt.Printf("realData) %s\n", realData)
	})

	// 접속
	eCode := kw.CommConnect()
	fmt.Printf("CommConnect() => %d : '%s'\n", eCode, kw.OpErrText(int(eCode)))

	kw.Wait()

	kw.Uninitialize()
}
