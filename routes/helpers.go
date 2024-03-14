package routes

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithText(w, http.StatusOK, "Pong")
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}

func RespondWithText(w http.ResponseWriter, statusCode int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(text))
}

// func prettyPrintKeys(data interface{}) {
// 	val := reflect.ValueOf(data)
// 	typ := val.Type()

// 	switch val.Kind() {
// 	case reflect.Struct:
// 		for i := 0; i < val.NumField(); i++ {
// 			field := typ.Field(i)
// 			fmt.Println(field.Name)
// 			// Recursive call to handle nested structs
// 			prettyPrintKeys(val.Field(i).Interface())
// 		}
// 	case reflect.Map:
// 		for _, key := range val.MapKeys() {
// 			fmt.Println(key.String())
// 			prettyPrintKeys(val.MapIndex(key).Interface())
// 		}
// 	default:
// 		// Not a struct or map, so no keys to print at this level.
// 	}
// }
