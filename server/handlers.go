package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studentAPI/database"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome Home"})
}

func Create(c *gin.Context) {
	var student database.Student
	//bind data to struct
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, Response{Success: false, Message: http.StatusText(400)})
		return
	}
	sid, err := database.AddStudent(student)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, Response{Success: false, Message: http.StatusText(422)})
		return
	}
	c.JSON(http.StatusCreated, Response{Success: true, Data: sid})
}

func FindAll(c *gin.Context){
	results, err := database.FetchStudents()

	if err != nil {
		c.JSON(http.StatusExpectationFailed, Response{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Success: true, Data: results})
}

func FindOneById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, Response{Success: false, Message: "Invalid StudentID"})
		return
	}

	student, err2 := database.StudentByID(id);
	if  err2 != nil {
		c.JSON(http.StatusNotFound, Response{Success: false, Message: http.StatusText(404)})
		return
	}
	c.JSON(http.StatusOK, Response{Success: true, Data: student})
}