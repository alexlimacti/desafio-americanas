package routes

import (
	"net/http"
)

func GetError(err error, w http.ResponseWriter) {

	http.Error(w, err.Error(), http.StatusInternalServerError)

	//var response = ErrorResponse{
	//	ErrorMessage: err.Error(),
	//	StatusCode:   http.StatusInternalServerError,
	//}
	//
	//message, _ := json.Marshal(response)
	//
	//w.WriteHeader(response.StatusCode)
	//w.Write(message)
}
