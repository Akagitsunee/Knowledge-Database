package repository

import (
	"fmt"
	"github.com/blockloop/scan"
	"kdb/model"
)

type EmployeeRepository struct {
}

func (repository *EmployeeRepository) GetAllEmployees() ([]model.Employee, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.Employee;")

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.Employee

	scan.Rows(&list, rows)

	return list, nil
}

func (repository *EmployeeRepository) GetEmployeeById(id string) (*model.Employee, error) {
	checkIfAlive()

	query := fmt.Sprintf("SELECT * FROM kdb.Employee WHERE kdb.Employee.EmployeeId =  %s;", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entry model.Employee

	scan.Row(&entry, rows)

	return &entry, nil
}

func (repository *EmployeeRepository) AddEmployee(employee *model.Employee) (*model.Employee, error) {
	checkIfAlive()

	query := fmt.Sprintf("INSERT INTO kdb.Employee([Name],[UserId]) VALUES('%s', '%s')", employee.Name, employee.UserId)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return employee, nil
}

func (repository *EmployeeRepository) DeleteEmployee(id string) error {
	checkIfAlive()

	query := fmt.Sprintf("UPDATE [kdb].[Employee] SET [Disabled] = 1 WHERE EmployeeId = %s", id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (repository *EmployeeRepository) ModifyEmployee(id string, employee *model.Employee) (*model.Employee, error) {
	checkIfAlive()

	query := fmt.Sprintf("UPDATE [kdb].[Employee] SET [Name] = '%s',[UserId] = '%s' WHERE [kdb].[Employee].EmployeeId = %s",
		employee.Name, employee.UserId, id)

	// Execute query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employee.EmployeeId = id
	return employee, nil
}

func (repository *EmployeeRepository) InitEmployeeRepository() EmployeeRepository {
	return EmployeeRepository{}
}
