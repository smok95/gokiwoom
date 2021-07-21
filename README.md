# gokiwoom
Go wrapper for 키움증권 OpenAPI+

## 예제
```go
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
```
  
    
      

## 주의사항
32bit만 지원하며, kw_.dll이 프로세스 실행경로와 같은 위치에 존재하여 합니다.

