package go_custom_utils

import "github.com/gin-gonic/gin"

func SetRunningMode(isDebug bool) {
	if isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func GetRunningMode(isDebug bool) string {
	if isDebug {
		return gin.DebugMode
	} else {
		return gin.ReleaseMode
	}
}
