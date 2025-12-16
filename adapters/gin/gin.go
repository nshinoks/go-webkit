package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware func(http.Handler) http.Handler

func Use(mw ...Middleware) gin.HandlerFunc {
	// Convert net/http middleware chain to gin
	return func(c *gin.Context) {
		final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})

		h := http.Handler(final)
		for i := len(mw) - 1; i >= 0; i-- {
			h = mw[i](h)
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}
