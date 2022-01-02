package api

import (
	"github.com/gin-gonic/gin"
	"kdb/httputil"
	"kdb/model"
	"kdb/repository"
	"net/http"
)

var employeeRepository repository.EmployeeRepository

func init() {
	employeeRepository = employeeRepository.InitEmployeeRepository()
}

// GetAllEmployees godoc
// @Summary      GetAllEmployees
// @Description  GetAllEmployees
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Employee
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /employee [get]
func GetAllEmployees(c *gin.Context) {
	employees, err := employeeRepository.GetAllEmployees()

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GetEmployeeById godoc
// @Summary      GetEmployeeById
// @Description  GetEmployeeById
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        id path int  true  "Id of employee"
// @Success      200  {object}   model.Employee
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /employee/{id} [get]
func GetEmployeeById(c *gin.Context) {
	id := c.Param("id")

	employees, err := employeeRepository.GetEmployeeById(id)

	if employees.EmployeeId == "0" || employees.EmployeeId == "" {
		c.AbortWithStatus(404)
		return
	}

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, employees)
}

// CreateEmployee godoc
// @Summary      CreateEmployee
// @Description  CreateEmployee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        employee body model.Employee true "Create employee"
// @Success      200  {object}  model.Employee
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /employee [post]
func CreateEmployee(c *gin.Context) {
	var e model.Employee

	c.ShouldBindJSON(&e)

	employees, err := employeeRepository.AddEmployee(&e)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, employees)
}

// UpdateEmployee godoc
// @Summary      UpdateEmployee
// @Description  UpdateEmployee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param id path int true "ID of the employee to be updated"
// @Param employee body model.Employee true "Updated employee"
// @Success      200  {array}   model.Employee
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /employee/{id} [put]
func UpdateEmployee(c *gin.Context) {
	var e model.Employee
	id := c.Param("id")

	c.ShouldBindJSON(&e)

	employees, err := employeeRepository.ModifyEmployee(id, &e)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

// DeleteEmployee godoc
// @Summary      DeleteEmployee
// @Description  DeleteEmployee
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        id path string true "ID of the employee to be deleted"
// @Success 	 204 "No Content"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /employee/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := employeeRepository.DeleteEmployee(id)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, id)
}
