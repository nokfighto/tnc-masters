package response

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type XResponseTimer struct {
	gin.ResponseWriter
	start time.Time
}

func (w *XResponseTimer) WriteHeader(statusCode int) {
	// calculate time duration
	duration := time.Since(w.start)

	w.Header().Set("X-Response-Time", fmt.Sprintf("%v", duration))
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *XResponseTimer) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

func NewXResponseTimer(c *gin.Context) {
	blw := &XResponseTimer{
		ResponseWriter: c.Writer,
		start:          time.Now(),
	}
	c.Writer = blw
	c.Next()
}
