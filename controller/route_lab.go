package controller

import (
	"fmt"
	"lab-manager-api/controller/model/req"
	"lab-manager-api/models"

	"github.com/gin-gonic/gin"
)

func CreateLab(c *gin.Context) {
	var labReq req.CreateLabRequest

	if err := c.ShouldBindJSON(&labReq); err != nil {
		c.JSON(400, gin.H{"error": "invalid json body"})
		return
	}

	lab, err := models.NewLab(labReq.Name, labReq.Local, labReq.Acessible,
		labReq.PcNumbers, models.LabStatus(labReq.Status), labReq.Softwares)

	if err != nil {
		c.JSON(500, gin.H{"error": "error creating lab"})
		return
	}

	lab, err = models.Save(lab)
	if err != nil {
		c.JSON(500, gin.H{"error": "error saving lab"})
		return
	}

	c.JSON(201, MapLabToResponse(lab))
}

func GetLab(c *gin.Context) {
	lab, err := models.FindLab(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "lab not found"})
		return
	}

	c.JSON(200, MapLabToResponse(lab))
}

func GetLabs(c *gin.Context) {
	labs, err := models.FindAllLabs()
	if err != nil {
		c.JSON(500, gin.H{"error": "error finding labs"})
		return
	}

	var labsResponse []interface{}
	for _, lab := range labs {
		labsResponse = append(labsResponse, MapLabToResponse(lab))
	}

	c.JSON(200, labsResponse)
}

func UpdateLab(c *gin.Context) {
	var labReq req.CreateLabRequest

	if err := c.ShouldBindJSON(&labReq); err != nil {
		c.JSON(400, gin.H{"error": "invalid json body"})
		return
	}

	lab, err := models.FindLab(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "lab not found"})
		return
	}

	lab.Name = labReq.Name
	lab.Local = labReq.Local
	lab.Acessible = labReq.Acessible
	lab.PcNumbers = labReq.PcNumbers
	lab.Status = models.LabStatus(labReq.Status)
	lab.Softwares = labReq.Softwares

	lab, err = models.UpdateLab(lab)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "error saving lab"})
		return
	}

	c.JSON(200, MapLabToResponse(lab))
}

func DeleteLab(c *gin.Context) {
	err := models.DeleteLab(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "error deleting lab"})
		return
	}

	c.JSON(204, nil)
}

func MapLabToResponse(lab models.Lab) interface{} {
	return gin.H{
		"id":        lab.ID.Hex(),
		"name":      lab.Name,
		"local":     lab.Local,
		"acessible": lab.Acessible,
		"pcNumbers": lab.PcNumbers,
		"status":    lab.Status,
		"softwares": lab.Softwares,
	}
}
