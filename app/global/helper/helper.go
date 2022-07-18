package helper

import (
	"encoding/json"
	"goformatv2/app/global"
	"goformatv2/app/global/errorcode"
	"goformatv2/app/global/structs"
	"time"
)

// Success 回傳成功API
func Success(result interface{}) *structs.APIResult {
	res := &structs.APIResult{
		Result: result,
		Status: structs.RespStatus{
			ErrorCode:   1,
			ErrorMsg:    "SUCCESS",
			Datetime:    time.Now().String(),
			LogIDentity: "",
		},
	}

	return res
}

// Fail 回傳失敗API
func Fail(err errorcode.Error) *structs.APIResult {
	res := &structs.APIResult{}
	status := structs.RespStatus{}

	status.ErrorCode = err.GetErrorCode()
	status.ErrorMsg = err.GetErrorText()
	status.LogIDentity = err.GetLogID()

	res.Result = map[string]string{}
	res.Status = status

	return res
}

// ParseTime 轉換時間格式(string ---> time.Time)
func ParseTime(myTime string) (t time.Time, apiErr errorcode.Error) {
	var err error

	if myTime == "0000-00-00 00:00:00" {
		return
	}

	// 指定時區
	local, err := time.LoadLocation("Local")
	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, errorcode.Code.GetTimeZoneError, err.Error())
		return
	}

	t, err = time.ParseInLocation("2006-01-02 15:04:05", myTime, local)
	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, errorcode.Code.ParseTimeError, err.Error())
		return
	}

	return
}

// StructToMap struct型態 轉 map型態 (For DB 使用)
func StructToMap(myStruct interface{}) (myMap map[string]interface{}, apiErr errorcode.Error) {

	// 轉形成map，才可以處理空字串,0,false值
	myMap = make(map[string]interface{})

	// 資料轉型
	byteData, err := json.Marshal(myStruct)
	if err != nil {
		apiErr = ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, err.Error())
		return
	}

	if err := json.Unmarshal(byteData, &myMap); err != nil {
		apiErr = ErrorHandle(global.WarnLog, errorcode.Code.JSONUnMarshalError, err.Error())
		return
	}

	return
}
