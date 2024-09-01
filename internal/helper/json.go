package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	IfPanicError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	IfPanicError(err)
}
