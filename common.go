package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ServeResults serve the results at the end
// of most handlers
func ServeResult(c *gin.Context, scriptName string, result gin.H) {
	if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json; charset=utf-8") {
		c.JSON(http.StatusOK, result)
	} else if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
		c.JSON(http.StatusOK, result)
	}else {
		c.HTML(http.StatusOK, "htmlwrapper.tmpl",
			gin.H{"scriptname": scriptName, "body": result})
	}
}
