package api

import (
	"fmt"
	"net/http"
	"time"

	HttpsUtil "github.com/skkim-01/vue-practice/dummyserver/httpsutil"
	"github.com/skkim-01/vue-practice/dummyserver/results"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmStart := time.Now()
	result := results.NewResult()
	var tmElapsed time.Duration

	requestBody, err := HttpsUtil.RequestBodyParser(r)
	if err != nil {
		result.Error(404, "Login Failed")
		goto EXIT_BLOCK
	}
	if fmt.Sprintf("%v", requestBody.Find("user")) == "" {
		result.Error(404, "Login Failed")
		goto EXIT_BLOCK
	}
	result.Success("session", true)

EXIT_BLOCK:
	SetCors(&w)
	fmt.Fprintf(w, "%v", result.Build())

	tmElapsed = time.Since(tmStart)
	fmt.Println("#DBG\t[GET] /v1/support-case\tresult:", result.Build(), "\ttook:", tmElapsed)
}
