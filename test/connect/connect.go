package main

import (
	"fmt"

	kw "github.com/smok95/gokiwoom"
)

func main() {

	// 연결상태 이벤트
	kw.SetOnEventConnect(func(errCode int32) {
		if errCode == 0 {
			fmt.Println("접속 성공")
		} else {
			fmt.Println("접속 실패")
		}

		// 종료
		kw.Disconnect()
	})

	// 접속
	kw.CommConnect()

	kw.Wait()
}
