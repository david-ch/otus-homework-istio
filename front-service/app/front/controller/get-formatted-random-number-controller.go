package controller

import (
	"fmt"
	"github.com/david-ch/otus-architect/front-service/client"
	"github.com/david-ch/otus-architect/front-service/util"
	"log"
	"net/http"
)


func OnGetFormattedRandomNumber(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received OnGetFormattedRandomNumber")

	numberRs, err := client.GetRandomNumber()
	if err != nil {
		log.Printf("OnGetFormattedRandomNumber: error while requesting back service")
		util.Resp(w, http.StatusInternalServerError, util.FromError(err))
		return
	}

	num := numberRs["randomNumber"].(float64)
	rs := map[string]interface{} {
		"message": fmt.Sprintf("Your number is %v", num),
	}
	log.Printf("OnGetFormattedRandomNumber: number is %v", num)

	util.Resp(w, http.StatusOK, rs)
}
