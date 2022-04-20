package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type organization struct {
	ID uint64 `json:"id"`
}

var organizations = []organization{
	{ID: 1},
	{ID: 2},
	{ID: 3},
}

func getOrganizationByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range organizations {
		if strconv.FormatUint(a.ID, 10) == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "organization not found"})
}

func postOrganizations(c *gin.Context) {
	var newOrg organization
	if err := c.BindJSON(&newOrg); err != nil {
		return
	}
	organizations = append(organizations, newOrg)
	c.IndentedJSON(http.StatusCreated, newOrg)
}

func main() {
	router := gin.Default()
	router.GET("/organizations/:id", getOrganizationByID)
	router.POST("/organizations", postOrganizations)

	router.Run("localhost:8040")
}
