package gokiwoom

import "C"
import (
	"syscall"
	"unsafe"
)

type OnEventConnect func(errCode int32)
type OnReceiveTrData func(scrNo string, rqName string, trCode string,
	recordName string, prevNext string, dataLength int32, errorCode string,
	message string, splmMsg string)

var (
	kw = syscall.NewLazyDLL("kw_.dll")

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
	kw_GetDataCountW              = kw.NewProc("kw_GetDataCountW")
	kw_GetOutputValueA            = kw.NewProc("kw_GetOutputValueA")
	kw_GetCommDataA               = kw.NewProc("kw_GetCommDataA")
	kw_GetCommDataW               = kw.NewProc("kw_GetCommDataW")
	kw_GetCommRealDataA           = kw.NewProc("kw_GetCommRealDataA")
	kw_GetChejanDataA             = kw.NewProc("kw_GetChejanDataA")
	kw_GetThemeGroupListA         = kw.NewProc("kw_GetThemeGroupListA")
	kw_GetThemeGroupCodeA         = kw.NewProc("kw_GetThemeGroupCodeA")
	kw_GetFutureListA             = kw.NewProc("kw_GetFutureListA")
	kw_GetFutureCodeByIndexA      = kw.NewProc("kw_GetFutureCodeByIndexA")
	kw_GetActPriceListA           = kw.NewProc("kw_GetActPriceListA")
	kw_GetMonthListA              = kw.NewProc("kw_GetMonthListA")
	kw_GetOptionCodeA             = kw.NewProc("kw_GetOptionCodeA")
	kw_GetOptionCodeByMonthA      = kw.NewProc("kw_GetOptionCodeByMonthA")
	kw_GetOptionCodeByActPriceA   = kw.NewProc("kw_GetOptionCodeByActPriceA")
	kw_GetSFutureListA            = kw.NewProc("kw_GetSFutureListA")
	kw_GetSFutureCodeByIndexA     = kw.NewProc("kw_GetSFutureCodeByIndexA")
	kw_GetSActPriceListA          = kw.NewProc("kw_GetSActPriceListA")
	kw_GetSMonthListA             = kw.NewProc("kw_GetSMonthListA")
	kw_GetSOptionCodeA            = kw.NewProc("kw_GetSOptionCodeA")
	kw_GetSOptionCodeByMonthA     = kw.NewProc("kw_GetSOptionCodeByMonthA")
	kw_GetSOptionCodeByActPriceA  = kw.NewProc("kw_GetSOptionCodeByActPriceA")
	kw_GetSFOBasisAssetListA      = kw.NewProc("kw_GetSFOBasisAssetListA")
	kw_GetOptionATMA              = kw.NewProc("kw_GetOptionATMA")
	kw_GetSOptionATMA             = kw.NewProc("kw_GetSOptionATMA")
	kw_GetBranchCodeNameA         = kw.NewProc("kw_GetBranchCodeNameA")
	kw_SendOrderCreditA           = kw.NewProc("kw_SendOrderCreditA")
	kw_KOA_FunctionsA             = kw.NewProc("kw_KOA_FunctionsA")
	kw_SetInfoDataA               = kw.NewProc("kw_SetInfoDataA")
	kw_SetInfoDataW               = kw.NewProc("kw_SetInfoDataW")
	kw_SetRealRegA                = kw.NewProc("kw_SetRealRegA")
	kw_GetConditionLoad           = kw.NewProc("kw_GetConditionLoad")
	kw_GetConditionNameListA      = kw.NewProc("kw_GetConditionNameListA")
	kw_SendConditionA             = kw.NewProc("kw_SendConditionA")
	kw_SendConditionStopA         = kw.NewProc("kw_SendConditionStopA")
	kw_GetCommDataExA             = kw.NewProc("kw_GetCommDataExA")
	kw_SetRealRemoveA             = kw.NewProc("kw_SetRealRemoveA")
	kw_GetMarketTypeA             = kw.NewProc("kw_GetMarketTypeA")
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
	/*
		typedef void (*kw_OnReceiveTrDataW)(ScrNo, PCWSTR sRQName,
			PCWSTR sTrCode, PCWSTR sRecordName, PCWSTR sPrevNext, long nDataLength,
			PCWSTR sErrorCode, PCWSTR sMessage, PCWSTR sSplmMsg);
		typedef void (*kw_OnReceiveTrDataA)(PCSTR sScrNo, PCSTR sRQName,
			PCSTR sTrCode, PCSTR sRecordName, PCSTR sPrevNext, long nDataLength,
			PCSTR sErrorCode, PCSTR sMessage, PCSTR sSplmMsg);

		typedef void (*kw_OnReceiveRealDataW)(PCWSTR sRealKey,
			PCWSTR sRealType, PCWSTR sRealData);
		typedef void (*kw_OnReceiveRealDataA)(PCSTR sRealKey,
			PCSTR sRealType, PCSTR sRealData);

		typedef void (*kw_OnReceiveMsgW)(PCWSTR sScrNo, PCWSTR sRQName,
			PCWSTR sTrCode, PCWSTR sMsg);
		typedef void (*kw_OnReceiveMsgA)(PCSTR sScrNo, PCSTR sRQName,
			PCSTR sTrCode, PCSTR sMsg);

		typedef void (*kw_OnReceiveChejanDataW)(PCWSTR sGubun, long nItemCnt,
			PCWSTR sFIdList);
		typedef void (*kw_OnReceiveChejanDataA)(PCSTR sGubun, long nItemCnt,
			PCSTR sFIdList);

		typedef void (*kw_OnReceiveRealConditionW)(PCWSTR sTrCode,
			PCWSTR strType, PCWSTR strConditionName, PCWSTR strConditionIndex);
		typedef void (*kw_OnReceiveRealConditionA)(PCSTR sTrCode,
			PCSTR strType, PCSTR strConditionName, PCSTR strConditionIndex);

		typedef void (*kw_OnReceiveTrConditionW)(PCWSTR sScrNo,
			PCWSTR strCodeList, PCWSTR strConditionName, int32 nIndex, int32 nNext);
		typedef void (*kw_OnReceiveTrConditionA)(PCSTR sScrNo,
			PCSTR strCodeList, PCSTR strConditionName, int32 nIndex, int32 nNext);

		typedef void (*kw_OnReceiveConditionVerW)(long lRet, PCWSTR sMsg);
		typedef void (*kw_OnReceiveConditionVerA)(long lRet, PCSTR sMsg);
	*/

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

func CommConnect() int32 {
	kw_SetCharsetUtf8.Call(uintptr(1))
	r, _, _ := kw_CommConnect.Call()
	return int32(r)
}
func CommRqData(rqName string, trCode string, prevNext int32, screenNo string) int32 {
	r, _, _ := kw_CommRqDataW.Call(wstr(rqName), wstr(trCode), uintptr(prevNext),
		wstr(screenNo))
	return int32(r)
}
func GetLoginInfo(tag string) string {
	r, _, _ := kw_GetLoginInfoA.Call(astr(tag))
	return kwastr2str(r)
}
func SendOrder(rqName string, screenNo string, accNo string, orderType int32,
	code string, qty int32, price int32, hogaGb string, orgOrderNo string) int32 {
	r, _, _ := kw_SendOrderW.Call(wstr(rqName), wstr(screenNo), wstr(accNo),
		uintptr(orderType), wstr(code), uintptr(qty), uintptr(price), wstr(hogaGb),
		wstr(orgOrderNo))
	return int32(r)
}
func SendOrderFO(rqName string, screenNo string, accNo string, code string,
	ordKind int32, slbyTp string, ordTp string, qty int32, price string,
	orgOrdNo string) int32 {
	r, _, _ := kw_SendOrderFOW.Call(wstr(rqName), wstr(screenNo), wstr(accNo),
		wstr(code), uintptr(ordKind), wstr(slbyTp), wstr(ordTp), uintptr(qty),
		wstr(price), wstr(orgOrdNo))
	return int32(r)
}

func SetInputValue(id string, value string) {
	kw_SetInputValueW.Call(wstr(id), wstr(value))
}

func DisconnectRealData(scnNo string) {
	kw_DisconnectRealDataW.Call(wstr(scnNo))
}

func GetRepeatCnt(trCode string, recordName string) int32 {
	r, _, _ := kw_GetRepeatCntW.Call(wstr(trCode), wstr(recordName))
	return int32(r)
}

func CommKwRqData(arrCode string, bNext int32, codeCount int32, typeFlag int32,
	rqName string, screenNo string) int32 {
	r, _, _ := kw_CommKwRqDataW.Call(wstr(arrCode), uintptr(bNext),
		uintptr(codeCount), uintptr(typeFlag), wstr(rqName), wstr(screenNo))
	return int32(r)
}

func GetAPIModulePath() string {
	r, _, _ := kw_GetAPIModulePathA.Call()
	return kwastr2str(r)
}

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

func GetCommData(trCode string, recordName string, index int32, itemName string) string {
	r, _, _ := kw_GetCommDataA.Call(astr(trCode), astr(recordName),
		uintptr(index), astr(itemName))
	return kwastr2str(r)
}

func GetCommRealData(trCode string, recordName string, index int32,
	itemName string) string {
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

func GetActPriceList()          {}
func GetMonthList()             {}
func GetOptionCode()            {}
func GetOptionCodeByMonth()     {}
func GetOptionCodeByActPrice()  {}
func GetSFutureList()           {}
func GetSFutureCodeByIndex()    {}
func GetSActPriceList()         {}
func GetSMonthList()            {}
func GetSOptionCode()           {}
func GetSOptionCodeByMonth()    {}
func GetSOptionCodeByActPrice() {}
func GetSFOBasisAssetList()     {}
func GetOptionATM()             {}
func GetSOptionATM()            {}
func GetBranchCodeName()        {}
func SendOrderCredit()          {}
func KOA_Functions()            {}
func SetInfoData()              {}
func SetRealReg()               {}
func GetConditionLoad()         {}
func GetConditionNameList()     {}
func SendCondition()            {}
func SendConditionStop()        {}
func GetCommDataEx()            {}
func SetRealRemove()            {}
func GetMarketType()            {}
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
func SetOnReceiveRealData()      {}
func SetOnReceiveMsg()           {}
func SetOnReceiveChejanData()    {}
func SetOnReceiveRealCondition() {}
func SetOnReceiveTrCondition()   {}
func SetOnReceiveConditionVer()  {}
func Wait() {
	syscall.Syscall(kw_Wait.Addr(), 0, 0, 0, 0)
}

func free(p uintptr) {
	kw_Free.Call(p)
}

func Disconnect()     {}
func SetCharsetUtf8() {}
