package gokiwoom

import "C"
import (
	"syscall"
	"unsafe"
)

type OnEventConnect func(errCode int32)
type OnReceiveTrData func(scrNo, rqName, trCode, recordName,
	prevNext string, dataLength int32, errorCode, message, splmMsg string)

// 리턴코드
const (
	OP_ERR_NONE               = 0    // 정상처리
	OP_ERR_FAIL               = -10  // 실패
	OP_ERR_COND_NOTFOUND      = -11  // 조건번호 없음
	OP_ERR_COND_MISMATCH      = -12  // 조건번호와 조건식 틀림
	OP_ERR_COND_OVERFLOW      = -13  // 조건검색 조회요청 초과
	OP_ERR_LOGIN              = -100 // 사용자정보 교환실패
	OP_ERR_CONNECT            = -101 // 서버접속 실패
	OP_ERR_VERSION            = -102 // 버전처리 실패
	OP_ERR_FIREWALL           = -103 // 개인방화벽 실패
	OP_ERR_MEMORY             = -104 // 메모리보호 실패
	OP_ERR_INPUT              = -105 // 함수입력값 오류
	OP_ERR_SOCKET_CLOSED      = -106 // 통신 연결종료
	OP_ERR_SISE_OVERFLOW      = -200 // 시세조회 과부하
	OP_ERR_RQ_STRUCT_FAIL     = -201 // 전문작성 초기화 실패
	OP_ERR_RQ_STRING_FAIL     = -202 // 전문작성 입력값 오류
	OP_ERR_NO_DATA            = -203 // 데이터 없음
	OP_ERR_OVER_MAX_DATA      = -204 // 조회 가능한 종목수 초과
	OP_ERR_DATA_RCV_FAIL      = -205 // 데이터수신 실패
	OP_ERR_OVER_MAX_FID       = -206 // 조회 가능한 FID수초과
	OP_ERR_REAL_CANCEL        = -207 // 실시간 해제 오류
	OP_ERR_ORD_WRONG_INPUT    = -300 // 입력값 오류
	OP_ERR_ORD_WRONG_ACCTNO   = -301 // 계좌 비밀번호 없음
	OP_ERR_OTHER_ACC_USE      = -302 // 타인계좌사용 오류
	OP_ERR_MIS_2BILL_EXC      = -303 // 주문가격이 20억원을 초과
	OP_ERR_MIS_5BILL_EXC      = -304 // 주문가격이 50억원을 초과
	OP_ERR_MIS_1PER_EXC       = -305 // 주문수량이 총발행주수의 1%초과오류
	OP_ERR_MIS_3PER_EXC       = -306 // 주문수량이 총발행주수의 3%초과오류
	OP_ERR_SEND_FAIL          = -307 // 주문전송 실패
	OP_ERR_ORD_OVERFLOW       = -308 // 주문전송 과부하
	OP_ERR_ORD_OVERFLOW2      = -311 // 주문전송 과부하
	OP_ERR_MIS_300CNT_EXC     = -309 // 주문수량 300계약 초과
	OP_ERR_MIS_500CNT_EXC     = -310 // 주문수량 500계약 초과
	OP_ERR_ORD_WRONG_ACCTINFO = -340 // 계좌정보없음
	OP_ERR_ORD_SYMCODE_EMPTY  = -500 // 종목코드없음
)

var opText = map[int]string{
	OP_ERR_NONE:               "정상처리",
	OP_ERR_FAIL:               "실패",
	OP_ERR_COND_NOTFOUND:      "조건번호 없음",
	OP_ERR_COND_MISMATCH:      "조건번호와 조건식 틀림",
	OP_ERR_COND_OVERFLOW:      "조건검색 조회요청 초과",
	OP_ERR_LOGIN:              "사용자정보 교환실패",
	OP_ERR_CONNECT:            "서버접속 실패",
	OP_ERR_VERSION:            "버전처리 실패",
	OP_ERR_FIREWALL:           "개인방화벽 실패",
	OP_ERR_MEMORY:             "메모리보호 실패",
	OP_ERR_INPUT:              "함수입력값 오류",
	OP_ERR_SOCKET_CLOSED:      "통신 연결종료",
	OP_ERR_SISE_OVERFLOW:      "시세조회 과부하",
	OP_ERR_RQ_STRUCT_FAIL:     "전문작성 초기화 실패",
	OP_ERR_RQ_STRING_FAIL:     "전문작성 입력값 오류",
	OP_ERR_NO_DATA:            "데이터 없음",
	OP_ERR_OVER_MAX_DATA:      "조회 가능한 종목수 초과",
	OP_ERR_DATA_RCV_FAIL:      "데이터수신 실패",
	OP_ERR_OVER_MAX_FID:       "조회 가능한 FID수초과",
	OP_ERR_REAL_CANCEL:        "실시간 해제 오류",
	OP_ERR_ORD_WRONG_INPUT:    "입력값 오류",
	OP_ERR_ORD_WRONG_ACCTNO:   "계좌 비밀번호 없음",
	OP_ERR_OTHER_ACC_USE:      "타인계좌사용 오류",
	OP_ERR_MIS_2BILL_EXC:      "주문가격이 20억원을 초과",
	OP_ERR_MIS_5BILL_EXC:      "주문가격이 50억원을 초과",
	OP_ERR_MIS_1PER_EXC:       "주문수량이 총발행주수의 1%초과오류",
	OP_ERR_MIS_3PER_EXC:       "주문수량이 총발행주수의 3%초과오류",
	OP_ERR_SEND_FAIL:          "주문전송 실패",
	OP_ERR_ORD_OVERFLOW:       "주문전송 과부하",
	OP_ERR_ORD_OVERFLOW2:      "주문전송 과부하",
	OP_ERR_MIS_300CNT_EXC:     "주문수량 300계약 초과",
	OP_ERR_MIS_500CNT_EXC:     "주문수량 500계약 초과",
	OP_ERR_ORD_WRONG_ACCTINFO: "계좌정보없음",
	OP_ERR_ORD_SYMCODE_EMPTY:  "종목코드없음",
}

func OpErrText(code int) string {
	return opText[code]
}

var (
	kw                     = syscall.NewLazyDLL("kw_.dll")
	kw_Initialize          = kw.NewProc("kw_Initialize")
	kw_Uninitizlize        = kw.NewProc("kw_Uninitialize")
	kw_CommConnect         = kw.NewProc("kw_CommConnect")
	kw_CommRqDataA         = kw.NewProc("kw_CommRqDataA")
	kw_CommRqDataW         = kw.NewProc("kw_CommRqDataW")
	kw_GetLoginInfoA       = kw.NewProc("kw_GetLoginInfoA")
	kw_SendOrderA          = kw.NewProc("kw_SendOrderA")
	kw_SendOrderW          = kw.NewProc("kw_SendOrderW")
	kw_SendOrderFOA        = kw.NewProc("kw_SendOrderFOA")
	kw_SendOrderFOW        = kw.NewProc("kw_SendOrderFOW")
	kw_SetInputValueA      = kw.NewProc("kw_SetInputValueA")
	kw_SetInputValueW      = kw.NewProc("kw_SetInputValueW")
	kw_DisconnectRealDataA = kw.NewProc("kw_DisconnectRealDataA")
	kw_DisconnectRealDataW = kw.NewProc("kw_DisconnectRealDataW")
	kw_GetRepeatCntA       = kw.NewProc("kw_GetRepeatCntA")
	kw_GetRepeatCntW       = kw.NewProc("kw_GetRepeatCntW")
	//kw_CommKwRqDataA              = kw.NewProc("kw_CommKwRqDataA")
	kw_CommKwRqDataW             = kw.NewProc("kw_CommKwRqDataW")
	kw_GetAPIModulePathA         = kw.NewProc("kw_GetAPIModulePathA")
	kw_GetCodeListByMarketA      = kw.NewProc("kw_GetCodeListByMarketA")
	kw_GetConnectState           = kw.NewProc("kw_GetConnectState")
	kw_GetMasterCodeNameA        = kw.NewProc("kw_GetMasterCodeNameA")
	kw_GetMasterListedStockCntA  = kw.NewProc("kw_GetMasterListedStockCntA")
	kw_GetMasterConstructionA    = kw.NewProc("kw_GetMasterConstructionA")
	kw_GetMasterListedStockDateA = kw.NewProc("kw_GetMasterListedStockDateA")
	kw_GetMasterLastPriceA       = kw.NewProc("kw_GetMasterLastPriceA")
	kw_GetMasterStockStateA      = kw.NewProc("kw_GetMasterStockStateA")
	//kw_GetDataCountA              = kw.NewProc("kw_GetDataCountA")
	kw_GetDataCountW             = kw.NewProc("kw_GetDataCountW")
	kw_GetOutputValueA           = kw.NewProc("kw_GetOutputValueA")
	kw_GetCommDataA              = kw.NewProc("kw_GetCommDataA")
	kw_GetCommDataW              = kw.NewProc("kw_GetCommDataW")
	kw_GetCommRealDataA          = kw.NewProc("kw_GetCommRealDataA")
	kw_GetChejanDataA            = kw.NewProc("kw_GetChejanDataA")
	kw_GetThemeGroupListA        = kw.NewProc("kw_GetThemeGroupListA")
	kw_GetThemeGroupCodeA        = kw.NewProc("kw_GetThemeGroupCodeA")
	kw_GetFutureListA            = kw.NewProc("kw_GetFutureListA")
	kw_GetFutureCodeByIndexA     = kw.NewProc("kw_GetFutureCodeByIndexA")
	kw_GetActPriceListA          = kw.NewProc("kw_GetActPriceListA")
	kw_GetMonthListA             = kw.NewProc("kw_GetMonthListA")
	kw_GetOptionCodeA            = kw.NewProc("kw_GetOptionCodeA")
	kw_GetOptionCodeByMonthA     = kw.NewProc("kw_GetOptionCodeByMonthA")
	kw_GetOptionCodeByActPriceA  = kw.NewProc("kw_GetOptionCodeByActPriceA")
	kw_GetSFutureListA           = kw.NewProc("kw_GetSFutureListA")
	kw_GetSFutureCodeByIndexA    = kw.NewProc("kw_GetSFutureCodeByIndexA")
	kw_GetSActPriceListA         = kw.NewProc("kw_GetSActPriceListA")
	kw_GetSMonthListA            = kw.NewProc("kw_GetSMonthListA")
	kw_GetSOptionCodeA           = kw.NewProc("kw_GetSOptionCodeA")
	kw_GetSOptionCodeByMonthA    = kw.NewProc("kw_GetSOptionCodeByMonthA")
	kw_GetSOptionCodeByActPriceA = kw.NewProc("kw_GetSOptionCodeByActPriceA")
	kw_GetSFOBasisAssetListA     = kw.NewProc("kw_GetSFOBasisAssetListA")
	kw_GetOptionATMA             = kw.NewProc("kw_GetOptionATMA")
	kw_GetSOptionATMA            = kw.NewProc("kw_GetSOptionATMA")
	kw_GetBranchCodeNameA        = kw.NewProc("kw_GetBranchCodeNameA")
	//kw_SendOrderCreditA           = kw.NewProc("kw_SendOrderCreditA")
	kw_SendOrderCreditW = kw.NewProc("kw_SendOrderCreditW")
	kw_KOA_FunctionsA   = kw.NewProc("kw_KOA_FunctionsA")
	//kw_SetInfoDataA               = kw.NewProc("kw_SetInfoDataA")
	kw_SetInfoDataW = kw.NewProc("kw_SetInfoDataW")
	//kw_SetRealRegA                = kw.NewProc("kw_SetRealRegA")
	kw_SetRealRegW           = kw.NewProc("kw_SetRealRegW")
	kw_GetConditionLoad      = kw.NewProc("kw_GetConditionLoad")
	kw_GetConditionNameListA = kw.NewProc("kw_GetConditionNameListA")
	//kw_SendConditionA             = kw.NewProc("kw_SendConditionA")
	kw_SendConditionW = kw.NewProc("kw_SendConditionW")
	//kw_SendConditionStopA         = kw.NewProc("kw_SendConditionStopA")
	kw_SendConditionStopW = kw.NewProc("kw_SendConditionStopW")
	kw_GetCommDataExA     = kw.NewProc("kw_GetCommDataExA")
	//kw_SetRealRemoveA             = kw.NewProc("kw_SetRealRemoveA")
	kw_SetRealRemoveW = kw.NewProc("kw_SetRealRemoveW")
	//kw_GetMarketTypeA             = kw.NewProc("kw_GetMarketTypeA")
	kw_GetMarketTypeW             = kw.NewProc("kw_GetMarketTypeW")
	kw_SetOnEventConnect          = kw.NewProc("kw_SetOnEventConnect")
	kw_SetOnReceiveTrDataA        = kw.NewProc("kw_SetOnReceiveTrDataA")
	kw_SetOnReceiveRealDataA      = kw.NewProc("kw_SetOnReceiveRealDataA")
	kw_SetOnReceiveMsgA           = kw.NewProc("kw_SetOnReceiveMsgA")
	kw_SetOnReceiveChejanDataA    = kw.NewProc("kw_SetOnReceiveChejanDataA")
	kw_SetOnReceiveRealConditionA = kw.NewProc("kw_SetOnReceiveRealConditionA")
	kw_SetOnReceiveTrConditionA   = kw.NewProc("kw_SetOnReceiveTrConditionA")
	kw_SetOnReceiveConditionVerA  = kw.NewProc("kw_SetOnReceiveConditionVerA")
	kw_Wait                       = kw.NewProc("kw_Wait")
	kw_Free                       = kw.NewProc("kw_Free")
	kw_FreeStringA                = kw.NewProc("kw_FreeStringA")
	kw_Disconnect                 = kw.NewProc("kw_Disconnect")
	kw_SetCharsetUtf8             = kw.NewProc("kw_SetCharsetUtf8")
	kw_Sleep                      = kw.NewProc("kw_Sleep")
)

// wstr string -> wchar_t*
func wstr(s string) uintptr {
	u, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(u))
}

// astr string -> char*
func astr(s string) uintptr {
	b, _ := syscall.BytePtrFromString(s)
	return uintptr(unsafe.Pointer(b))
}

// p2s char*(utf8) -> string
func p2s(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	return C.GoString((*C.char)(p))
}

// kwastr2str utf8 char* -> string (with free)
func kwastr2str(psz uintptr) string {
	if psz == 0 {
		return ""
	}
	defer free(psz)
	return C.GoString((*C.char)(unsafe.Pointer(psz)))
}

func Initialize(option int32) int32 {
	r, _, _ := kw_Initialize.Call(uintptr(option))
	return int32(r)
}

func Uninitialize() {
	kw_Uninitizlize.Call()
}

// CommConnect 로그인 윈도우를 실행한다.
//
// 반환값 : 0-성공, 음수값은 실패
func CommConnect() int32 {
	kw_SetCharsetUtf8.Call(uintptr(1))
	r, _, _ := kw_CommConnect.Call()
	return int32(r)
}

// CommRqData Tran을 서버로 송신한다.
func CommRqData(rqName, trCode string, prevNext int32, screenNo string) int32 {
	r, _, _ := kw_CommRqDataW.Call(wstr(rqName), wstr(trCode), uintptr(prevNext),
		wstr(screenNo))
	return int32(r)
}
func GetLoginInfo(tag string) string {
	r, _, _ := kw_GetLoginInfoA.Call(astr(tag))
	return kwastr2str(r)
}
func SendOrder(rqName, screenNo, accNo string, orderType int32,
	code string, qty, price int32, hogaGb, orgOrderNo string) int32 {
	r, _, _ := kw_SendOrderW.Call(wstr(rqName), wstr(screenNo), wstr(accNo),
		uintptr(orderType), wstr(code), uintptr(qty), uintptr(price), wstr(hogaGb),
		wstr(orgOrderNo))
	return int32(r)
}
func SendOrderFO(rqName, screenNo, accNo, code string, ordKind int32,
	slbyTp, ordTp string, qty int32, price, orgOrdNo string) int32 {
	r, _, _ := kw_SendOrderFOW.Call(wstr(rqName), wstr(screenNo), wstr(accNo),
		wstr(code), uintptr(ordKind), wstr(slbyTp), wstr(ordTp), uintptr(qty),
		wstr(price), wstr(orgOrdNo))
	return int32(r)
}

func SetInputValue(id, value string) {
	kw_SetInputValueW.Call(wstr(id), wstr(value))
}

func DisconnectRealData(scnNo string) {
	kw_DisconnectRealDataW.Call(wstr(scnNo))
}

func GetRepeatCnt(trCode, recordName string) int32 {
	r, _, _ := kw_GetRepeatCntW.Call(wstr(trCode), wstr(recordName))
	return int32(r)
}

func CommKwRqData(arrCode string, bNext, codeCount, typeFlag int32,
	rqName, screenNo string) int32 {
	r, _, _ := kw_CommKwRqDataW.Call(wstr(arrCode), uintptr(bNext),
		uintptr(codeCount), uintptr(typeFlag), wstr(rqName), wstr(screenNo))
	return int32(r)
}

func GetAPIModulePath() string {
	r, _, _ := kw_GetAPIModulePathA.Call()
	return kwastr2str(r)
}

// GetCodeListByMarket 시장구분에 따른 종목코드를 반환한다.
//
// market
//	0:장내, 3:ELW, 4:뮤추얼펀드, 5:신주인수권, 6:리츠,
//	8:ETF, 9:하이일드펀드, 10:코스닥, 30:K-OTC, 50:코넥스(KONEX)
//
// 종목코드 리스트, 종목간 구분은 ";"이다.
// https://download.kiwoom.com/web/openapi/kiwoom_openapi_plus_devguide_ver_1.5.pdf
func GetCodeListByMarket(market string) string {
	r, _, _ := kw_GetCodeListByMarketA.Call(astr(market))
	return kwastr2str(r)
}

func GetConnectState() int32 {
	r, _, _ := kw_GetConnectState.Call()
	return int32(r)
}

func GetMasterCodeName(trCode string) string {
	r, _, _ := kw_GetMasterCodeNameA.Call(astr(trCode))
	return kwastr2str(r)
}

func GetMasterListedStockCnt(trCode string) string {
	r, _, _ := kw_GetMasterListedStockCntA.Call(astr(trCode))
	return kwastr2str(r)
}

func GetMasterConstruction(trCode string) string {
	r, _, _ := kw_GetMasterConstructionA.Call(astr(trCode))
	return kwastr2str(r)
}

func GetMasterListedStockDate(trCode string) string {
	r, _, _ := kw_GetMasterListedStockDateA.Call(astr(trCode))
	return kwastr2str(r)
}

func GetMasterLastPrice(trCode string) string {
	r, _, _ := kw_GetMasterLastPriceA.Call(astr(trCode))
	return kwastr2str(r)
}

func GetMasterStockState(trCode string) string {
	r, _, _ := kw_GetMasterStockStateA.Call(astr(trCode))
	return kwastr2str(r)
}

func GetDataCount(recordName string) int32 {
	r, _, _ := kw_GetDataCountW.Call(wstr(recordName))
	return int32(r)
}

func GetOutputValue(recordName string, repeatIdx int32) string {
	r, _, _ := kw_GetOutputValueA.Call(astr(recordName), uintptr(repeatIdx))
	return kwastr2str(r)
}

func GetCommData(trCode, recordName string, index int32, itemName string) string {
	r, _, _ := kw_GetCommDataA.Call(astr(trCode), astr(recordName),
		uintptr(index), astr(itemName))
	return kwastr2str(r)
}

func GetCommRealData(trCode, recordName string, index int32, itemName string) string {
	r, _, _ := kw_GetCommRealDataA.Call(astr(trCode), astr(recordName),
		uintptr(index), astr(itemName))
	return kwastr2str(r)
}

func GetChejanData(nFid int32) string {
	r, _, _ := kw_GetChejanDataA.Call(uintptr(nFid))
	return kwastr2str(r)
}

func GetThemeGroupList(nType int32) string {
	r, _, _ := kw_GetThemeGroupCodeA.Call(uintptr(nType))
	return kwastr2str(r)
}

func GetThemeGroupCode(themeCode string) string {
	r, _, _ := kw_GetThemeGroupCodeA.Call(astr(themeCode))
	return kwastr2str(r)
}

func GetFutureList() string {
	r, _, _ := kw_GetFutureListA.Call()
	return kwastr2str(r)
}
func GetFutureCodeByIndex(index int32) string {
	r, _, _ := kw_GetFutureCodeByIndexA.Call(uintptr(index))
	return kwastr2str(r)
}

func GetActPriceList() string {
	r, _, _ := kw_GetActPriceListA.Call()
	return kwastr2str(r)
}

func GetMonthList() string {
	r, _, _ := kw_GetMonthListA.Call()
	return kwastr2str(r)
}

func GetOptionCode(actPrice string, cp int32, month string) string {
	r, _, _ := kw_GetOptionCodeA.Call(astr(actPrice), uintptr(cp),
		astr(month))
	return kwastr2str(r)
}

func GetOptionCodeByMonth(trCode string, cp int32, month string) string {
	r, _, _ := kw_GetOptionCodeByMonthA.Call(astr(trCode), uintptr(cp),
		astr(month))
	return kwastr2str(r)
}

func GetOptionCodeByActPrice(trCode string, cp int32, tick int32) string {
	r, _, _ := kw_GetOptionCodeByActPriceA.Call(astr(trCode), uintptr(cp),
		uintptr(tick))
	return kwastr2str(r)
}

func GetSFutureList(baseAssetCode string) string {
	r, _, _ := kw_GetSFutureListA.Call(astr(baseAssetCode))
	return kwastr2str(r)
}

func GetSFutureCodeByIndex(baseAssetCode string, index int32) string {
	r, _, _ := kw_GetSFutureCodeByIndexA.Call(astr(baseAssetCode),
		uintptr(index))
	return kwastr2str(r)
}

func GetSActPriceList(baseAssetGb string) string {
	r, _, _ := kw_GetSActPriceListA.Call(astr(baseAssetGb))
	return kwastr2str(r)
}

func GetSMonthList(baseAssetGb string) string {
	r, _, _ := kw_GetSMonthListA.Call(astr(baseAssetGb))
	return kwastr2str(r)
}

func GetSOptionCode(baseAssetGb, actPrice string, cp int32, month string) string {
	r, _, _ := kw_GetSOptionCodeA.Call(astr(baseAssetGb), astr(actPrice),
		uintptr(cp), astr(month))
	return kwastr2str(r)
}

func GetSOptionCodeByMonth(baseAssetGb, trCode string, cp int32, month string) string {
	r, _, _ := kw_GetSOptionCodeByMonthA.Call(astr(baseAssetGb), astr(trCode),
		uintptr(cp), astr(month))
	return kwastr2str(r)
}

func GetSOptionCodeByActPrice(baseAssetGb, trCode string, cp int32, tick int32) string {
	r, _, _ := kw_GetSOptionCodeByActPriceA.Call(astr(baseAssetGb), astr(trCode),
		uintptr(cp), uintptr(tick))
	return kwastr2str(r)
}

func GetSFOBasisAssetList() string {
	r, _, _ := kw_GetSFOBasisAssetListA.Call()
	return kwastr2str(r)
}

func GetOptionATM() string {
	r, _, _ := kw_GetOptionATMA.Call()
	return kwastr2str(r)
}

func GetSOptionATM(baseAssetGb string) string {
	r, _, _ := kw_GetSOptionATMA.Call(astr(baseAssetGb))
	return kwastr2str(r)
}

func GetBranchCodeName() string {
	r, _, _ := kw_GetBranchCodeNameA.Call()
	return kwastr2str(r)
}

func SendOrderCredit(rqName, screenNo, accNo string, orderType int32,
	code string, qty, price int32, hogaGb, creditGb, loanDate,
	orgOrderNo string) int32 {
	r, _, _ := kw_SendOrderCreditW.Call(wstr(rqName), wstr(screenNo),
		wstr(accNo), uintptr(orderType), wstr(code), uintptr(qty), uintptr(price),
		wstr(hogaGb), wstr(creditGb), wstr(loanDate), wstr(orgOrderNo))
	return int32(r)
}

func KOA_Functions(functionName, param string) string {
	r, _, _ := kw_KOA_FunctionsA.Call(astr(functionName), astr(param))
	return kwastr2str(r)
}

func SetInfoData(infoData string) int32 {
	r, _, _ := kw_SetInfoDataW.Call(wstr(infoData))
	return int32(r)
}

func SetRealReg(screenNo, codeList, fidList, optType string) int32 {
	r, _, _ := kw_SetRealRegW.Call(wstr(screenNo), wstr(codeList),
		wstr(fidList), wstr(optType))
	return int32(r)
}

func GetConditionLoad() int32 {
	r, _, _ := kw_GetConditionLoad.Call()
	return int32(r)
}

func GetConditionNameList() string {
	r, _, _ := kw_GetConditionNameListA.Call()
	return kwastr2str(r)
}

func SendCondition(scrNo, conditionName string, index, search int32) int32 {
	r, _, _ := kw_SendConditionW.Call(wstr(scrNo), wstr(conditionName),
		uintptr(index), uintptr(search))
	return int32(r)
}

func SendConditionStop(scrNo, conditionName string, index int32) {
	kw_SendConditionStopW.Call(wstr(scrNo), wstr(conditionName),
		uintptr(index))
}

// VARIANT 어떤식으로 처리하지? func GetCommDataEx() {}

func SetRealRemove(scrNo, delCode string) {
	kw_SetRealRemoveW.Call(wstr(scrNo), wstr(delCode))
}

func GetMarketType(trCode string) int32 {
	r, _, _ := kw_GetMarketTypeW.Call(wstr(trCode))
	return int32(r)
}

func SetOnEventConnect(callback OnEventConnect) {
	cb := syscall.NewCallbackCDecl(func(errCode int32) (ret uintptr) {
		if callback != nil {
			callback(errCode)
		}
		return
	})

	syscall.Syscall(kw_SetOnEventConnect.Addr(), 1, cb, 0, 0)
}
func SetOnReceiveTrData(callback OnReceiveTrData) {
	cb := syscall.NewCallbackCDecl(func(scrNo unsafe.Pointer,
		rqName unsafe.Pointer, trCode unsafe.Pointer, recordName unsafe.Pointer,
		prevNext unsafe.Pointer, dataLength int32, errorCode unsafe.Pointer,
		message unsafe.Pointer, splmMsg unsafe.Pointer) (ret uintptr) {
		if callback != nil {
			callback(p2s(scrNo), p2s(rqName), p2s(trCode), p2s(recordName),
				p2s(prevNext), dataLength, p2s(errorCode), p2s(message),
				p2s(splmMsg))
		}
		return
	})

	syscall.Syscall(kw_SetOnReceiveTrDataA.Addr(), 1, cb, 0, 0)
}

type OnReceiveRealData func(realKey, realType, realData string)

func SetOnReceiveRealData(callback OnReceiveRealData) {
	cb := syscall.NewCallbackCDecl(func(realKey unsafe.Pointer,
		realType unsafe.Pointer, realData unsafe.Pointer) (ret uintptr) {
		if callback != nil {
			callback(p2s(realKey), p2s(realType), p2s(realData))
		}
		return
	})

	kw_SetOnReceiveRealDataA.Call(cb)
}

type OnReceiveMsg func(scrNo, rqName, trCode, msg string)

func SetOnReceiveMsg(callback OnReceiveMsg) {
	cb := syscall.NewCallbackCDecl(func(r1, r2, r3, r4 unsafe.Pointer) (ret uintptr) {
		if callback != nil {
			callback(p2s(r1), p2s(r2), p2s(r3), p2s(r4))
		}
		return
	})
	kw_SetOnReceiveMsgA.Call(cb)
}

type OnReceiveChejanData func(gubun string, itemCnt int32, fidList string)

func SetOnReceiveChejanData(callback OnReceiveChejanData) {
	cb := syscall.NewCallbackCDecl(func(r1 unsafe.Pointer, r2 int32,
		r3 unsafe.Pointer) (ret uintptr) {
		if callback != nil {
			callback(p2s(r1), r2, p2s(r3))
		}
		return
	})

	kw_SetOnReceiveChejanDataA.Call(cb)
}

type OnReceiveRealCondition func(code, sType, conditionName, conditionIndex string)

func SetOnReceiveRealCondition(callback OnReceiveRealCondition) {
	cb := syscall.NewCallbackCDecl(func(r1, r2, r3,
		r4 unsafe.Pointer) (ret uintptr) {
		if callback != nil {
			callback(p2s(r1), p2s(r2), p2s(r3), p2s(r4))
		}
		return
	})

	kw_SetOnReceiveRealConditionA.Call(cb)
}

type OnReceiveTrCondition func(scrNo, codeList, conditionName string,
	index, next int32)

func SetOnReceiveTrCondition(callback OnReceiveTrCondition) {
	cb := syscall.NewCallbackCDecl(func(r1, r2, r3 unsafe.Pointer,
		r4, r5 int32) (ret uintptr) {
		if callback != nil {
			callback(p2s(r1), p2s(r2), p2s(r3), r4, r5)
		}
		return
	})

	kw_SetOnReceiveTrConditionA.Call(cb)
}

type OnReceiveConditionVer func(ret int32, msg string)

func SetOnReceiveConditionVer(callback OnReceiveConditionVer) {
	cb := syscall.NewCallbackCDecl(func(r1 int32, r2 unsafe.Pointer) (ret uintptr) {
		if callback != nil {
			callback(r1, p2s(r2))
		}
		return
	})

	kw_SetOnReceiveConditionVerA.Call(cb)
}

func Wait() {
	kw_Wait.Call()
}

func free(p uintptr) {
	kw_Free.Call(p)
}

func Disconnect() {
	kw_Disconnect.Call()
}

func Sleep(millisecond int32) {
	kw_Sleep.Call(uintptr(millisecond))
}

/* cgo 방식

// 라이브러리 경로 및 라이브러리 설정
//#cgo LDFLAGS: -L./ -lkw_
//#include <stdlib.h>
//#include "kw_.h"
//
//void onEventConnect(int nErrCode);
//
//void onReceiveRealData(char* sRealKey, char* sRealType, char* sRealData);
import "C"
import "unsafe"

func setInputValue(id string, value string) {
	cId := C.CString(id)
	cVal := C.CString(value)
	C.kw_SetInputValueA(cId, cVal)
	C.free(unsafe.Pointer(cId))
	C.free(unsafe.Pointer(cVal))
}

//export onEventConnect
func onEventConnect(nErrCode C.int) {
	if kw.onConnectResult != nil {
		kw.onConnectResult(kw, int(nErrCode))
		cstr := C.CString("005930")
		ret := C.kw_GetMasterCodeNameA(cstr)
		C.free(unsafe.Pointer(cstr))
	}
}

//export onReceiveTrData
func onReceiveTrData(sScrNo *C.char, sRQName *C.char, sTrCode *C.char,
	sRecordName *C.char, sPrevNext *C.char, nDataLength C.long,
	sErrorCode *C.char, sMessage *C.char, sSplmMsg *C.char) {
	fmt.Println("onReceiveTrData")
}

//export onReceiveRealData
func onReceiveRealData(sRealKey *C.char, sRealType *C.char, sRealData *C.char) {
	fmt.Printf("key=%s, type=%s, %s\n", C.GoString(sRealKey), C.GoString(sRealType), C.GoString(sRealData))
}

*/
