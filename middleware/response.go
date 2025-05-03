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

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// ResponseHandlerMiddleware wraps the response body with a standard format
func ResponseHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture the response body
		buf := new(bytes.Buffer)
		writer := &responseWriter{body: buf, ResponseWriter: c.Writer}
		c.Writer = writer

		// Process the request
		c.Next()

		statusCode := c.Writer.Status()
		contentType := c.Writer.Header().Get("Content-Type")

		// If there's an error set in context, process it
		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			defErr := core.ToDefaultError(lastErr, c.GetString("RequestID"))
			c.JSON(defErr.StatusCode(), defErr)
			return
		}

		// Only wrap JSON responses (skip if already structured)
		if contentType == "application/json" && statusCode < 300 {
			var originalData interface{}
			if err := json.Unmarshal(writer.body.Bytes(), &originalData); err != nil {
				defErr := core.ErrInternalServerError.WithDebugf("Failed to parse JSON: %v", err)
				c.JSON(http.StatusInternalServerError, defErr)
				return
			}

			// Overwrite the response with a wrapped format
			c.JSON(statusCode, gin.H{
				"success": true,
				"data":    originalData,
				"message": http.StatusText(statusCode),
			})
		}
	}
}
