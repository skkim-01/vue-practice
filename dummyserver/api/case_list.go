package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/skkim-01/vue-practice/dummyserver/results"

	"github.com/tyler-smith/go-bip39"
)

func CaseList(w http.ResponseWriter, r *http.Request) {
	tmStart := time.Now()
	result := results.NewResult()
	var tmElapsed time.Duration

	SetCors(&w)
	w.Header().Add("content-type", "application/json")

	// slResult := make([]map[string]interface{}, 3)
	// slResult[0] = _createDummyData()
	// slResult[1] = _createDummyData()
	// slResult[2] = _createDummyData()

	slResult := _createDummys(45)
	result.Success("total", 45, "items", slResult)

	fmt.Fprintf(w, "%v", result.Build())

	tmElapsed = time.Since(tmStart)
	fmt.Println(fmt.Sprintf("#DBG\t[%v] %v%v\ttook: %v", r.Method, r.Host, r.URL, tmElapsed))
	// fmt.Println(result.Build())
}

func _createDummys(count int) []map[string]interface{} {
	retv := make([]map[string]interface{}, count)

	for i := 0; i < count; i++ {
		dummy := make(map[string]interface{}, 0)
		dummy["case_id"] = fmt.Sprintf("caseid_%03d", i)
		dummy["account_id"] = fmt.Sprintf("account %v", (i+1)%5)
		dummy["category_code"] = fmt.Sprintf("category %v", (i+1)%7)
		dummy["display_id"] = fmt.Sprintf("display id %v", (i+1)%5)
		dummy["language"] = "en"
		dummy["service_code"] = fmt.Sprintf("svc code %v", (i+1)%13)
		dummy["status"] = "OPENED" // TODO
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		dummy["subject"] = mnemonic
		dummy["severity_code"] = "CRITICAL" // TODO
		dummy["onwer_email"] = "test@account.com"
		dummy["use_yn"] = "y"
		dummy["case_create_time"] = fmt.Sprintf("2023-06-07 %02d:%02d:%02d", (i+1)%3, (i+1)%58, (i+1)%58)
		dummy["create_time"] = fmt.Sprintf("2023-06-07 %02d:%02d:%02d", (i+2)%7, (i+2)%23, (i+2)%29)
		dummy["update_time"] = fmt.Sprintf("2023-06-07 %02d:%02d:%02d", (i+3)%11, (i+3)%37, (i+3)%18)
		retv[i] = dummy
	}
	return retv
}

func _createDummyData() map[string]interface{} {
	dummy := make(map[string]interface{}, 0)
	dummy["account_id"] = "account 1"
	dummy["category_code"] = "cat_10103"
	dummy["display_id"] = "disp_1013"
	dummy["language"] = "en"
	dummy["service_code"] = "svc_1013"
	dummy["status"] = "OPENED"
	dummy["subject"] = "test subject"
	dummy["severity_code"] = "CRITICAL"
	dummy["onwer_email"] = "test@account.com"
	dummy["use_yn"] = "y"
	dummy["case_create_time"] = "2023-06-07 12:01:02"
	dummy["create_time"] = "2023-06-07 12:01:02"
	dummy["update_time"] = "2023-06-07 12:01:02"

	return dummy
}
