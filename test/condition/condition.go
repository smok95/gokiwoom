package main

import (
	"fmt"

	kw "github.com/smok95/gokiwoom"
)

func OnReceiveTrCondition(scrNo, codeList, conditionName string,
	index, next int32) {
	fmt.Printf("OnReceiveTrCondition() => scrNo:'%s'\ncodeList:'%s'\n", scrNo, codeList)
}

func main() {
	kw.Initialize(1)

	kw.SetOnReceiveTrCondition(OnReceiveTrCondition)

	kw.SetOnReceiveConditionVer(func(ret int32, msg string) {
		fmt.Printf("OnReceiveConditionVer() => ret:%d, msg:'%s'\n", ret, msg)
		kw.SendCondition("test1", "조건식1", 5, 0)
	})

	// 연결상태 이벤트
	kw.SetOnEventConnect(func(errCode int32) {
		if errCode == 0 {
			fmt.Println("연결 성공")

			fmt.Println("조건검색 로드")
			kw.GetConditionLoad()
		} else {
			fmt.Println("연결 종료")
		}
	})

	// 접속
	eCode := kw.CommConnect()
	fmt.Printf("CommConnect() => %d : '%s'\n", eCode, kw.OpErrText(int(eCode)))

	kw.Wait()

	kw.Uninitialize()
}
