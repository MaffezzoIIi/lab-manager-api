package controller

import (
	"fmt"
	"lab-manager-api/controller/model/req"
	"lab-manager-api/models"

	"github.com/gin-gonic/gin"
)

// CreateLab godoc
// @Summary Create a new lab
// @Description Create a new lab
// @Tags labs
// @Accept  json
// @Produce  json
// @Param lab body req.CreateLabRequest true "Lab object that needs to be created"
// @Success 201 {object} res.CreateLabResponse
// @Failure 400 {object} rest_err.RestErr
// @Router /api/v1/labs/create [post]
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

	lab, err = models.SaveLab(lab)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, MapLabToResponse(lab))
}

// GetLab godoc
// @Summary Retrieve a lab by ID
// @Description Fetches a lab by its unique identifier
// @Tags labs
// @Accept  json
// @Produce  json
// @Param id path string true "Lab ID"
// @Success 200 {object} res.CreateLabResponse
// @Failure 404 {object} rest_err.RestErr
// @Router /api/v1/labs/{id} [get]
func GetLab(c *gin.Context) {
	lab, err := models.FindLab(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "lab not found"})
		return
	}

	c.JSON(200, MapLabToResponse(lab))
}

// GetLabs godoc
// @Summary Retrieve all labs
// @Description Fetches a list of all labs
// @Tags labs
// @Accept  json
// @Produce  json
// @Success 200 {array} res.CreateLabResponse
// @Failure 500 {object} rest_err.RestErr
// @Router /api/v1/labs [get]
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

// UpdateLab godoc
// @Summary Update an existing lab
// @Description Updates the details of a specific lab by ID
// @Tags labs
// @Accept  json
// @Produce  json
// @Param id path string true "Lab ID"
// @Param lab body req.CreateLabRequest true "Lab object with updated details"
// @Success 200 {object} res.CreateLabResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /api/v1/labs/{id} [put]
func UpdateLab(c *gin.Context) {
	var labReq req.CreateLabRequest

	if err := c.ShouldBindJSON(&labReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

// DeleteLab godoc
// @Summary Delete a lab
// @Description Deletes a specific lab by ID
// @Tags labs
// @Accept  json
// @Produce  json
// @Param id path string true "Lab ID"
// @Success 204 {object} nil
// @Failure 500 {object} rest_err.RestErr
// @Router /api/v1/labs/{id} [delete]
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
		"id":        lab.ID,
		"name":      lab.Name,
		"local":     lab.Local,
		"acessible": lab.Acessible,
		"pcNumbers": lab.PcNumbers,
		"status":    lab.Status,
		"softwares": lab.Softwares,
	}
}
