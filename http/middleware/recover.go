package middleware

import (
	"github.com/justinas/alice"
	"log"
	"net/http"
	"runtime"
)

func Recoverer() alice.Constructor {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Panic: %v\n", err)
					stack := make([]byte, 4<<10) // 4 KB stack buffer
					length := runtime.Stack(stack, true)
					log.Printf("%s", stack[:length])
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}
