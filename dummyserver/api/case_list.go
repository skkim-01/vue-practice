package api

import (
	"fmt"
	"net/http"
	"time"

	JsonMapper "github.com/skkim-01/json-mapper/src"
	"github.com/skkim-01/vue-practice/dummyserver/results"
)

func CaseList(w http.ResponseWriter, r *http.Request) {
	tmStart := time.Now()
	result := results.NewResult()
	var tmElapsed time.Duration

	SetCors(&w)
	w.Header().Add("content-type", "application/json")

	jsonResponse, _ := JsonMapper.NewBytes([]byte("{}"))
	jsonResponse.Insert("", "account_id", "account 1")
	jsonResponse.Insert("", "category_code", "account 1")
	jsonResponse.Insert("", "display_id", "account 1")
	jsonResponse.Insert("", "language", "account 1")
	jsonResponse.Insert("", "service_code", "account 1")
	jsonResponse.Insert("", "status", "account 1")
	jsonResponse.Insert("", "subject", "account 1")
	jsonResponse.Insert("", "severity_code", "account 1")
	jsonResponse.Insert("", "onwer_email", "account 1")
	jsonResponse.Insert("", "use_yn", "account 1")
	jsonResponse.Insert("", "case_create_time", "account 1")
	jsonResponse.Insert("", "create_time", "account 1")
	jsonResponse.Insert("", "update_time", "account 1")

	result.Success("items", jsonResponse.Find(""))

	fmt.Fprintf(w, "%v", result.Build())

	tmElapsed = time.Since(tmStart)
	fmt.Println("#DBG\t[GET] /v1/support-case\tresult:", result.Build(), "\ttook:", tmElapsed)
}
