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

// Router and route definitions
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
	wrappedHandler := LogRequest(handler)
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
		handler:    wrappedHandler,
	})
}

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK}
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if req.Method == route.method && route.pattern.MatchString(req.URL.Path) {
			matches := route.pattern.FindStringSubmatch(req.URL.Path)
			params := make(map[string]string)
			for i, name := range route.pattern.SubexpNames() {
				if i > 0 && i <= len(matches) {
					params[name] = matches[i]
				}
			}
			route.handler(w, req, params)
			return
		}
	}
	http.NotFound(w, req)
}
