package handlers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ShortenHandler struct {
	Urls map[string]string
}

func (s *ShortenHandler) AddNewLink(c *gin.Context) {
	validator := validator.New()
	var url RequestURLShorten
	ok := c.BindJSON(&url)
	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := validator.Struct(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if url.CustomAlias != "" {
		if _, ok := s.Urls[url.CustomAlias]; ok {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Alias already exists",
			})
			return
		}
		s.Urls[url.CustomAlias] = url.URL
		c.JSON(http.StatusOK, gin.H{
			"shortened_url": url.CustomAlias,
		})
		return
	}
	url.CustomAlias = randomString(5)
	s.Urls[url.CustomAlias] = url.URL

	fmt.Println(s.Urls)
	c.JSON(http.StatusOK, gin.H{
		"shortened_url": url.CustomAlias,
	})
}

func (s *ShortenHandler) Redirect(c *gin.Context) {
	alias := c.Param("alias")
	if val, ok := s.Urls[alias]; ok {
		c.Redirect(http.StatusMovedPermanently, val)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Alias not found",
	})
}

type RequestURLShorten struct {
	URL         string `json:"url" validate:"required,url"`
	CustomAlias string `json:"customAlias"`
}

func randomString(length int) string {
	ranStr := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ranStr[i] = byte(65 + rand.Intn(25))
	}
	return string(ranStr)
}
