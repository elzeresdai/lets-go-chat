package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"runtime/debug"
)

func PanicRecovery(next echo.HandlerFunc) echo.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})
}
