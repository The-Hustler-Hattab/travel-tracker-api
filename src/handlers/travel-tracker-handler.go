package handlers

import (
	"net/http"
	"strconv"
	"tax-travel-tracker/src/db/models"
	"tax-travel-tracker/src/db/repository"
	"github.com/gin-gonic/gin"
)

// CreateTravelRecordHandler godoc
// @Summary Create a new travel record
// @Description Add a new record to the travel tracker
// @Tags TravelTracker
// @Accept json
// @Produce json
// @Param record body models.TaxTravelTracker true "Travel Record"
// @Success 201 {object} models.TaxTravelTracker
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /travel-tracker [post]
func CreateTravelRecordHandler(c *gin.Context) {
	var record models.TaxTravelTracker
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateTravelRecord(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}

	c.JSON(http.StatusCreated, record)
}

// GetTravelRecordHandler godoc
// @Summary Get a travel record by ID
// @Description Fetch a single travel record using its ID
// @Tags TravelTracker
// @Produce json
// @Param id path int true "Travel Record ID"
// @Success 200 {object} models.TaxTravelTracker
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /travel-tracker/{id} [get]
func GetTravelRecordHandler(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	record, err := repository.GetTravelRecordByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// UpdateTravelRecordHandler godoc
// @Summary Update an existing travel record
// @Description Modify an existing record in the travel tracker
// @Tags TravelTracker
// @Accept json
// @Produce json
// @Param id path int true "Travel Record ID"
// @Param record body models.TaxTravelTracker true "Updated Travel Record"
// @Success 200 {object} models.TaxTravelTracker
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /travel-tracker/{id} [put]
func UpdateTravelRecordHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedRecord models.TaxTravelTracker
	if err := c.ShouldBindJSON(&updatedRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	record, err := repository.GetTravelRecordByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	record.TravelFrom = updatedRecord.TravelFrom
	record.TravelTo = updatedRecord.TravelTo
	record.TravelDistance = updatedRecord.TravelDistance
	record.EstimatedTaxDeductions = updatedRecord.EstimatedTaxDeductions
	record.CreatedBy = updatedRecord.CreatedBy
	record.TravelTime = updatedRecord.TravelTime

	if err := repository.UpdateTravelRecord(record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// DeleteTravelRecordHandler godoc
// @Summary Delete a travel record
// @Description Remove a record from the travel tracker
// @Tags TravelTracker
// @Param id path int true "Travel Record ID"
// @Success 204
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /travel-tracker/{id} [delete]
func DeleteTravelRecordHandler(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := repository.DeleteTravelRecord(intID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}


// GetAllTravelRecordsHandler godoc
// @Summary Get all travel records
// @Description Retrieve all records from the travel tracker
// @Tags TravelTracker
// @Produce json
// @Success 200 {array} models.TaxTravelTracker
// @Failure 500 {object} gin.H
// @Router /travel-tracker [get]
func GetAllTravelRecordsHandler(c *gin.Context) {
	records, err := repository.GetAllTravelRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve records"})
		return
	}

	c.JSON(http.StatusOK, records)
}
