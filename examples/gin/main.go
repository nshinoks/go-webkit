package main

import (
	"fmt"
	"net/http"
	"time"

	ginadapter "github.com/nshinoks/go-webkit/adapters/gin"

	"github.com/gin-gonic/gin"
	"github.com/nshinoks/go-webkit/errors"
	"github.com/nshinoks/go-webkit/middleware"
	"github.com/nshinoks/go-webkit/respond"
)

func main() {
	r := gin.New()

	r.Use(ginadapter.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.AccessLog(func(req *http.Request, status int, dur time.Duration) {
			id, _ := middleware.RequestIDFrom(req.Context())
			fmt.Printf("status=%d dur=%s method=%s path=%s request_id=%s\n",
				status, dur.String(), req.Method, req.URL.Path, id)
		}),
	))

	r.GET("/health", func(c *gin.Context) {
		respond.JSON(c.Writer, 200, map[string]any{"ok": true})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		if c.Param("id") == "0" {
			respond.Error(c.Writer, errors.BadRequest("id must not be 0"))
			return
		}
		// respond.JSON(c.Writer, 200, map[string]any{"id": c.Param("id")})
		user := gin.H{"id": c.Param("id")}
		respond.OK(c.Writer, c.Request, user)
	})

	_ = r.Run(":8080")
}
