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
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// ResponseHandlerMiddleware wraps the response with a standard format
func ResponseHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the response writer to capture output
		buf := new(bytes.Buffer)
		writer := &responseWriter{ResponseWriter: c.Writer, body: buf}
		c.Writer = writer

		// Process the request
		c.Next()

		// If there is an error in context, handle it
		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			defErr := core.ToDefaultError(lastErr, c.GetString("RequestID"))
			c.JSON(defErr.StatusCode(), gin.H{
				"success": false,
				"error":   defErr,
			})
			return
		}

		// Handle success responses for JSON
		statusCode := c.Writer.Status()
		contentType := c.Writer.Header().Get("Content-Type")

		// Only wrap JSON responses
		if contentType == "application/json" && statusCode < 300 {
			var originalData interface{}
			if err := json.Unmarshal(writer.body.Bytes(), &originalData); err != nil {
				defErr := core.ErrInternalServerError.WithDebugf("Failed to parse JSON: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   defErr,
				})
				return
			}

			// Success response, wrap data in standard format
			c.JSON(statusCode, gin.H{
				"success": true,
				"data":    originalData,
				"message": http.StatusText(statusCode),
				"request": c.GetString("RequestID"),
			})
		}
	}
}
