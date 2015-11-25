package base

import (
	"fmt"
	"net/http"

	"github.com/ttacon/chalk"
)

func LogError(e interface{}) {
	message := chalk.Red.Color(fmt.Sprintf("[ERROR] %s", e))
	fmt.Println(message)
}

func LogStatusError(r *http.Request, err *statusError) {
	message := fmt.Sprintf("<%d> %s (%s)", err.code, err.message, r.URL)
	LogError(message)
}
