package utility

import (
	"encoding/json"
	"net/http"
)

type jsonWriterFunc func(v interface{}, status int)

func NewJSONWriter(w http.ResponseWriter) jsonWriterFunc{
	return func(v interface{}, status int) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(v);
	}
}

func(f jsonWriterFunc) Write(v interface{}, status int){
	f(v, status);
}