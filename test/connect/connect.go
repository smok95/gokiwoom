package main

import (
	"fmt"

	kw "github.com/smok95/gokiwoom"
)

func main() {
	kw.Initialize(1)

	// 연결상태 이벤트
	kw.SetOnEventConnect(func(errCode int32) {
		if errCode == 0 {
			fmt.Println("연결 성공")
		} else {
			fmt.Println("연결 종료")
		}

		// 종료
		kw.Disconnect()
	})

	// 접속
	eCode := kw.CommConnect()
	fmt.Printf("CommConnect() => %d : '%s'\n", eCode, kw.OpErrText(int(eCode)))

	kw.Wait()

	kw.Uninitialize()
}
