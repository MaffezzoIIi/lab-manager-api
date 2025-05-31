package controller

import (
	"fmt"
	"lab-manager-api/controller/model/req"
	"lab-manager-api/models/lab"

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

	new_lab, err := lab.NewLab(labReq.Name, labReq.Local, labReq.Acessible,
		labReq.PcNumbers, lab.LabStatus(labReq.Status), labReq.Softwares)

	if err != nil {
		c.JSON(500, gin.H{"error": "error creating lab"})
		return
	}

	new_lab, err = lab.SaveLab(new_lab)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, MapLabToResponse(new_lab))
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
	lab, err := lab.FindLab(c.Param("id"))
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
	labs, err := lab.FindAllLabs()
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

	found_lab, err := lab.FindLab(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "lab not found"})
		return
	}

	found_lab.Name = labReq.Name
	found_lab.Local = labReq.Local
	found_lab.Acessible = labReq.Acessible
	found_lab.PcNumbers = labReq.PcNumbers
	found_lab.Status = lab.LabStatus(labReq.Status)
	found_lab.Softwares = labReq.Softwares

	found_lab, err = lab.UpdateLab(found_lab)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "error saving lab"})
		return
	}

	c.JSON(200, MapLabToResponse(found_lab))
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
	err := lab.DeleteLab(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": "error deleting lab"})
		return
	}

	c.JSON(204, nil)
}

func MapLabToResponse(lab lab.Lab) interface{} {
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
