package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/muve-go/go-crud/dao"
	"github.com/lakshanwd/muve-go/go-crud/repo"
)

//BookingGet - handle booking get requests
func BookingGet(c *gin.Context) {
	from, err := strconv.Atoi(c.Query("from"))
	if err != nil {
		from = 0
	}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		count = 100
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		if list, err := repo.GetBooking(int64(id)); err == nil {
			c.JSON(http.StatusOK, convertListToArray(list))
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	} else {
		if list, err := repo.GetBookings(int64(from), count); err == nil {
			c.JSON(http.StatusOK, convertListToArray(list))
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

//BookingPost - handle booking post requests
func BookingPost(c *gin.Context) {
	var booking dao.Booking
	if err := c.ShouldBindJSON(&booking); err == nil {
		if insertedID, err := repo.InsertBooking(&booking); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": true, "id": insertedID})
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//BookingPut - handle booking put requests
func BookingPut(c *gin.Context) {
	var booking dao.Booking
	if err := c.ShouldBindJSON(&booking); err == nil {
		if id, err := strconv.Atoi(c.Param("id")); err == nil {
			if rowsEfected, err := repo.UpdateBooking(&booking, id); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": true, "rowsEfected": rowsEfected})
			} else {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//BookingDelete - handle booking delete requests
func BookingDelete(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		if rowsDeleted, err := repo.RemoveBooking(id); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": true, "rowsDeleted": rowsDeleted})
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
