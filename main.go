package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/wynt/chia-pool-web/service"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*.html")

	// get details of a single farmer by id
	r.GET("/farmers/:id", func(c *gin.Context) {
		farmerId := c.Param("id")
		farmerDetail, err := service.GetFarmerById(farmerId)

		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"Message": fmt.Sprintf("Error retrieving farmer %s: %+v\n", farmerId, err.Error()),
			})
		}

		c.HTML(http.StatusOK, "farmerDetail.html", gin.H{
			"Farmer": farmerDetail,
		})
	})

	// update farmer view (only name is currently updateable)
	// TODO should be protected as an "admin" function
	r.GET("/farmers/:id/edit", func(c *gin.Context) {
		farmerId := c.Param("id")

		farmer, err := service.GetFarmerById(farmerId)

		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"Message": fmt.Sprintf("Error getting farmer %s for update: %+v\n", farmerId, err.Error()),
			})
		}

		c.HTML(http.StatusOK, "farmerEdit.html", gin.H{
			"Farmer": farmer,
		})

	})

	r.POST("/farmers/:id/edit", func(c *gin.Context) {
		id := c.Param("id")
		service.UpdateFarmer(id, c.PostForm("name"))
		c.Redirect(http.StatusFound, "/farmers/"+id)
	})

	// get summary of all farmers
	// TODO should be protected as an "admin" function
	r.GET("/farmers", func(c *gin.Context) {
		// list all farmers
		farmers, err := service.GetAllFarmers()

		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"Message": fmt.Sprintf("Error retrieving farmers: %+v\n", err.Error()),
			})
		}

		fmt.Println(farmers)
		c.HTML(http.StatusOK, "farmers.html", gin.H{
			"Farmers": farmers,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
