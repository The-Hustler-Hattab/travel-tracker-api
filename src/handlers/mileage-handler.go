package handlers

import (
	"net/http"
	"strconv"
	"tax-travel-tracker/src/db/repository"
	"github.com/gin-gonic/gin"
)

// GetAllMileageRecordsHandler retrieves all mileage records.
// @Summary Retrieve all mileage records
// @Description Get a list of all mileage records.
// @Tags Mileage
// @Accept json
// @Produce json
// @Success 200 {array} models.IrsStandardMileageRate
// @Failure 500 {object} gin.H
// @Router /mileage [get]
func GetAllMileageRecordsHandler(c *gin.Context) {
	records, err := repository.GetAllMileageRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve records"})
		return
	}

	c.JSON(http.StatusOK, records)
}


// GetMilageRecordHandler retrieves a mileage record by its ID.
// @Summary Retrieve mileage record
// @Description Get a mileage record by its unique ID.
// @Tags Mileage
// @Accept json
// @Produce json
// @Param id path int true "Mileage Record ID"
// @Success 200 {object} models.IrsStandardMileageRate
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /mileage/{id} [get]
func GetMilageRecordHandler(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	record, err := repository.GetMilageRecordByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}