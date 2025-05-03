package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b) // Just buffer it; do not write to the real writer yet
}

func ResponseHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := new(bytes.Buffer)
		writer := &responseWriter{body: buf, ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		statusCode := c.Writer.Status()
		contentType := c.Writer.Header().Get("Content-Type")

		// Handle errors
		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			defErr := core.ToDefaultError(lastErr, c.GetString("RequestID"))

			// Reset and write error response
			c.Writer = writer.ResponseWriter
			c.JSON(defErr.StatusCode(), defErr)
			return
		}

		// Handle successful JSON response
		if contentType == "application/json" && statusCode < 300 {
			var originalData interface{}
			if err := json.Unmarshal(writer.body.Bytes(), &originalData); err != nil {
				c.Writer = writer.ResponseWriter
				c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
				return
			}

			// Reset and write wrapped success response
			c.Writer = writer.ResponseWriter
			c.JSON(statusCode, gin.H{
				"success": true,
				"data":    originalData,
				"message": http.StatusText(statusCode),
			})
			return
		}

		// For non-JSON or other responses: write raw content
		c.Writer = writer.ResponseWriter
		c.Writer.WriteHeaderNow()
		_, _ = c.Writer.Write(writer.body.Bytes())
	}
}
