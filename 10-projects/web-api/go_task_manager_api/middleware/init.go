package middleware

import (
	"net/http"
)

// WrapWithMiddlewares applies all global middlewares to the router
func WrapWithMiddlewares(next http.Handler) http.Handler {
	// You can chain more middlewares here like logging, CORS, etc.
	return CORSMiddleware(LoggingMiddleware(next))
	// return LoggingMiddleware(next) // example: wrap with logging
}
