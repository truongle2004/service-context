package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// responseBodyWriter captures response body
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)    // Capture
	return len(b), nil // Don't write yet
}

// ResponseFormatterMiddleware wraps JSON responses into standard format
func ResponseFormatterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Replace response writer
		writer := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		// Restore real writer
		c.Writer = writer.ResponseWriter

		status := c.Writer.Status()
		contentType := c.Writer.Header().Get("Content-Type")

		// Use captured response body if JSON
		if contentType == "application/json" {
			var parsed any
			_ = json.Unmarshal(writer.body.Bytes(), &parsed)

			if len(c.Errors) > 0 || status >= 400 {
				// Error case
				c.JSON(status, gin.H{
					"success": false,
					"error":   parsed,
					"message": http.StatusText(status),
				})
			} else {
				// Success case
				c.JSON(status, gin.H{
					"success": true,
					"data":    parsed,
					"message": http.StatusText(status),
				})
			}
			return
		}

		// Non-JSON: write as is
		c.Writer.WriteHeaderNow()
		_, _ = c.Writer.Write(writer.body.Bytes())
	}
}
