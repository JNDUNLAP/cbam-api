package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (r *Router) ListRoutes() {
	fmt.Println("\nCBAM API Routes\n---------------------------------------------")
	groupMap := make(map[string][]Route)
	for _, route := range r.routes {
		groupMap[route.groupName] = append(groupMap[route.groupName], route)
	}

	for groupName, routes := range groupMap {
		fmt.Printf("\n%s\n", groupName)
		for _, route := range routes {

			fmt.Printf("%s %s\n", route.method, route.rawPattern)
		}
	}
	fmt.Println()
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func colorizeStatusCode(statusCode int) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return fmt.Sprintf("%s%d%s", colorGreen, statusCode, colorReset)
	case statusCode >= 300 && statusCode < 400:
		return fmt.Sprintf("%s%d%s", colorBlue, statusCode, colorReset)
	case statusCode >= 400 && statusCode < 500:
		return fmt.Sprintf("%s%d%s", colorYellow, statusCode, colorReset)
	case statusCode >= 500:
		return fmt.Sprintf("%s%d%s", colorRed, statusCode, colorReset)
	default:
		return fmt.Sprintf("%d", statusCode)
	}
}

func LogRequest(handler HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		start := time.Now()

		wrappedWriter := NewResponseWriter(w)

		handler(wrappedWriter, r, params)

		duration := time.Since(start)

		log.Printf("[%s]: [%s], Duration: [%v]  -  %s", r.Method, colorizeStatusCode(wrappedWriter.statusCode), duration, r.URL.Path)
	}
}
