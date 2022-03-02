package go_grammer

import "github.com/gin-gonic/gin"

type ErrorHandler func(c *gin.Context, err error)
