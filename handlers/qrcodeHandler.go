package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func QRCodeHandler(c *gin.Context) {
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	base64Image := base64.StdEncoding.EncodeToString(png)

	c.JSON(http.StatusOK, gin.H{
		"image": base64Image,
	})
}
