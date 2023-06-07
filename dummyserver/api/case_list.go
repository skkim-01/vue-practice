package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/skkim-01/vue-practice/dummyserver/results"
)

func CaseList(w http.ResponseWriter, r *http.Request) {
	tmStart := time.Now()
	result := results.NewResult()
	var tmElapsed time.Duration

	SetCors(&w)
	w.Header().Add("content-type", "application/json")

	slResult := make([]map[string]interface{}, 3)
	slResult[0] = _createDummyData()
	slResult[1] = _createDummyData()
	slResult[2] = _createDummyData()

	result.Success("items", slResult)

	fmt.Fprintf(w, "%v", result.Build())

	tmElapsed = time.Since(tmStart)
	fmt.Println("#DBG\t[GET] /v1/support-case\tresult:", result.Build(), "\ttook:", tmElapsed)
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
