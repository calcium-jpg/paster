package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type paste struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Paste string `json:"paste"`
}

var pastes = []paste{}

func getPastes(c *gin.Context) {
	c.JSON(http.StatusOK, pastes)
}

func getPasteByID(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}

	for _, p := range pastes {
		if p.ID == uint32(uid) {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "paste not found"})
}

func postPastes(c *gin.Context) {
	var newPaste paste

	if !(len(pastes) == 0) {
		newPaste.ID = pastes[len(pastes)-1].ID + 1
	} else {
		newPaste.ID = 0
	}

	if err := c.BindJSON(&newPaste); err != nil {
		return
	}

	pastes = append(pastes, newPaste)

	c.JSON(http.StatusCreated, newPaste)
}

func putPasteByID(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}

	var newPaste paste

	newPaste.ID = uint32(uid)

	if err := c.BindJSON(&newPaste); err != nil {
		return
	}

	for i, p := range pastes {
		if p.ID == uint32(uid) {
			pastes[i] = newPaste
			c.JSON(http.StatusOK, pastes[i])
			return
		}
	}
}

func deletePasteByID(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}

	for i, p := range pastes {
		if p.ID == uint32(uid) {
			pastes = append(pastes[:i], pastes[i+1:]...)
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "paste not found"})
}

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/pastes", getPastes)
	router.GET("/pastes/:id", getPasteByID)
	router.POST("/pastes", postPastes)
	router.PUT("/pastes/:id", putPasteByID)
	router.DELETE("/pastes/:id", deletePasteByID)

	router.Run("localhost:8080")
}
