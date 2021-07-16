package gokiwoom

import "C"
import (
	"syscall"
	"unsafe"
)

type OnEventConnect func(errCode int)
type OnReceiveTrData func(scrNo string, rqName string, trCode string,
	recordName string, prevNext string, dataLength int, errorCode string,
	message string, splmMsg string)

var (
	kw = syscall.NewLazyDLL("kw_.dll")

	kw_CommConnect                = kw.NewProc("kw_CommConnect")
	kw_CommRqDataA                = kw.NewProc("kw_CommRqDataA")
	kw_CommRqDataW                = kw.NewProc("kw_CommRqDataW")
	kw_GetLoginInfoA              = kw.NewProc("kw_GetLoginInfoA")
	kw_SendOrderA                 = kw.NewProc("kw_SendOrderA")
	kw_SendOrderFOA               = kw.NewProc("kw_SendOrderFOA")
	kw_SetInputValueA             = kw.NewProc("kw_SetInputValueA")
	kw_SetInputValueW             = kw.NewProc("kw_SetInputValueW")
	kw_DisconnectRealDataA        = kw.NewProc("kw_DisconnectRealDataA")
	kw_GetRepeatCntA              = kw.NewProc("kw_GetRepeatCntA")
	kw_CommKwRqDataA              = kw.NewProc("kw_CommKwRqDataA")
	kw_GetAPIModulePathA          = kw.NewProc("kw_GetAPIModulePathA")
	kw_GetCodeListByMarketA       = kw.NewProc("kw_GetCodeListByMarketA")
	kw_GetConnectState            = kw.NewProc("kw_GetConnectState")
	kw_GetMasterCodeNameA         = kw.NewProc("kw_GetMasterCodeNameA")
	kw_GetMasterListedStockCntA   = kw.NewProc("kw_GetMasterListedStockCntA")
	kw_GetMasterConstructionA     = kw.NewProc("kw_GetMasterConstructionA")
	kw_GetMasterListedStockDateA  = kw.NewProc("kw_GetMasterListedStockDateA")
	kw_GetMasterLastPriceA        = kw.NewProc("kw_GetMasterLastPriceA")
	kw_GetMasterStockStateA       = kw.NewProc("kw_GetMasterStockStateA")
	kw_GetDataCountA              = kw.NewProc("kw_GetDataCountA")
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
		typedef void (*kw_OnReceiveTrDataW)(PCWSTR sScrNo, PCWSTR sRQName,
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
			PCWSTR strCodeList, PCWSTR strConditionName, int nIndex, int nNext);
		typedef void (*kw_OnReceiveTrConditionA)(PCSTR sScrNo,
			PCSTR strCodeList, PCSTR strConditionName, int nIndex, int nNext);

		typedef void (*kw_OnReceiveConditionVerW)(long lRet, PCWSTR sMsg);
		typedef void (*kw_OnReceiveConditionVerA)(long lRet, PCSTR sMsg);
	*/

)

// s2u string -> wchar_t*
func s2u(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// s2b string -> char*
func s2b(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}

func p2s(p unsafe.Pointer) string {
	return C.GoString((*C.char)(p))
}

func CommConnect() int {
	kw_SetCharsetUtf8.Call(uintptr(1))
	r1, _, _ := syscall.Syscall(kw_CommConnect.Addr(), 0, 0, 0, 0)
	return int(r1)
}
func CommRqData(rqName string, trCode string, prevNext int, screenNo string) int {
	r1, _, _ := syscall.Syscall6(kw_CommRqDataW.Addr(), 4, s2u(rqName), s2u(trCode),
		uintptr(prevNext), s2u(screenNo), 0, 0)
	return int(r1)
}
func GetLoginInfo() {}
func SendOrder()    {}
func SendOrderFO()  {}
func SetInputValue(id string, value string) {
	syscall.Syscall(kw_SetInputValueW.Addr(), 2, s2u(id), s2u(value), 0)
}
func DisconnectRealData()       {}
func GetRepeatCnt()             {}
func CommKwRqData()             {}
func GetAPIModulePath()         {}
func GetCodeListByMarket()      {}
func GetConnectState()          {}
func GetMasterCodeName()        {}
func GetMasterListedStockCnt()  {}
func GetMasterConstruction()    {}
func GetMasterListedStockDate() {}
func GetMasterLastPrice()       {}
func GetMasterStockState()      {}
func GetDataCount()             {}
func GetOutputValue()           {}
func GetCommData(trCode string, recordName string, index int, itemName string) string {
	r1, _, _ := kw_GetCommDataA.Call(s2b(trCode), s2b(recordName),
		uintptr(index), s2b(itemName))
	return p2s(unsafe.Pointer(r1))
}
func GetCommRealData()          {}
func GetChejanData()            {}
func GetThemeGroupList()        {}
func GetThemeGroupCode()        {}
func GetFutureList()            {}
func GetFutureCodeByIndex()     {}
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
	cb := syscall.NewCallbackCDecl(func(errCode int) (ret uintptr) {
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
		prevNext unsafe.Pointer, dataLength int, errorCode unsafe.Pointer,
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
func Free()           {}
func FreeString()     {}
func Disconnect()     {}
func SetCharsetUtf8() {}
