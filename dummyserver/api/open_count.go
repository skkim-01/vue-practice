package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/skkim-01/vue-practice/dummyserver/results"
)

func OpenCount(w http.ResponseWriter, r *http.Request) {
	tmStart := time.Now()

	result := results.NewResult()
	result.Success("open_count", 31)

	fmt.Fprintf(w, "%v", result.Build())

	tmElapsed := time.Since(tmStart)
	fmt.Println("#DBG\t[GET] /api/open/count\ttook:", tmElapsed)
}
