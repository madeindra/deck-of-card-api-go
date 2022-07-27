package response

import (
	"encoding/json"
	"net/http"
)

type ResultData struct {
	Code int
	Data interface{}
}

type FailedResult struct {
	Message string
}

func JSON(w http.ResponseWriter, result ResultData) {
	response, err := json.Marshal(result.Data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("{message: \"An Error Occured\"}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Code)
	w.Write(response)
	return
}
