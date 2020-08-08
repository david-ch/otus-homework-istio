package controller

import (
	"github.com/david-ch/otus-architect/back-service/util"
	"log"
	"math/rand"
	"net/http"
)

func OnGetRandomNumber(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received OnGetRandomNumber")
	rs := map[string]interface{} {
		"randomNumber": rand.Intn(100),
	}

	log.Printf("OnGetRandomNumber: response is %+v", rs)
	util.Resp(w, http.StatusOK, rs)
}
