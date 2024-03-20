package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
)

func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
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

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type Route struct {
	groupName  string
	method     string
	pattern    *regexp.Regexp
	rawPattern string
	handler    HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handle(groupName, method, pattern string, handler HandlerFunc) {
	regexPattern := regexp.MustCompile(`\{(\w+)\}`)
	regexStr := regexPattern.ReplaceAllString(pattern, `(?P<$1>[^/]+)`)
	compiledPattern, err := regexp.Compile("^" + regexStr + "$")
	if err != nil {
		log.Fatal("Could not compile regex pattern:", err)
	}

	r.routes = append(r.routes, Route{
		groupName:  groupName,
		method:     method,
		pattern:    compiledPattern,
		rawPattern: pattern,
		handler:    handler,
	})
}
