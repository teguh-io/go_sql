package repositories

import (
	"database/sql"
	"gosql/models"
)

type EmployeeRepo interface {
	CreateEmployee(*models.Employee) error
	GetEmployees() ([]models.Employee, error)
	GetEmployeeByID(ID int) (*models.Employee, error)
	UpdateEmployeeByID(ID int, employee models.Employee) error
	DeleteEmployeeByID(ID int) error
}

type employeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (er *employeeRepo) CreateEmployee(employee *models.Employee) error {
	sqlQuery := `INSERT INTO employees(full_name, email, age, division) VALUES ($1, $2, $3, $4) Returning *`

	tx, err := er.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&employee.FullName, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (er *employeeRepo) GetEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	sqlQuery := `SELECT id, full_name, email, age, division FROM employees`

	stmt, err := er.db.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (er *employeeRepo) GetEmployeeByID(ID int) (*models.Employee, error) {
	var employee models.Employee
	sqlQuery := `SELECT * FROM employees WHERE id=$1`
	stmt, err := er.db.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(ID)

	err = rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		return nil, err
	}

	return &employee, nil

}

func (er *employeeRepo) UpdateEmployeeByID(ID int, employee models.Employee) error {
	sqlQuery := `UPDATE employees SET full_name=$1, email=$2, age=$3, division=$4`
	tx, err := er.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return nil
	}

	defer stmt.Close()

	_, err = stmt.Exec(&employee.FullName, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (er *employeeRepo) DeleteEmployeeByID(ID int) error {
	sqlQuery := `DELETE FROM employees WHERE id=$1`
	tx, err := er.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
