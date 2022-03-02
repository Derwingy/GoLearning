package go_grammer

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type rate struct {
	OnError ErrorHandler
}

func DefaultLimitReachedHandler(c *gin.Context) {
	c.String(http.StatusTooManyRequests, "Limit exceeded")
}

func NameStyle(boilingF float64) float64 {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("first is %g F, second is %g C", f, c)
	return c
}

type NameProperties struct {
	Name string
	mu   sync.RWMutex
}

type LoadName map[string]*NameProperties

func statusName() *NameProperties {
	// np_instance = NameProperties("dy")
	return &NameProperties{
		Name: "Young",
	}
}
