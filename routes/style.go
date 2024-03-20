package routes

import "fmt"

var methodColors = map[string]string{
	"GET":    "\033[32m", // Green
	"POST":   "\033[36m", // Blue
	"PUT":    "\033[34m", // Cyan
	"DELETE": "\033[31m", // Red
	"PATCH":  "\033[33m", // Yellow
}

const resetColor = "\033[0m"

func (r *Router) ListRoutes() {
	fmt.Println("\nCBAM API Routes\n---------------------------------------------")
	groupMap := make(map[string][]Route)
	for _, route := range r.routes {
		groupMap[route.groupName] = append(groupMap[route.groupName], route)
	}

	for groupName, routes := range groupMap {
		fmt.Printf("\n%s\n", groupName)
		for _, route := range routes {
			methodColor, ok := methodColors[route.method]
			if !ok {
				methodColor = ""
			}
			fmt.Printf("%s%-6s%s %s\n", methodColor, route.method, resetColor, route.rawPattern)
		}
	}
	fmt.Println()
}
