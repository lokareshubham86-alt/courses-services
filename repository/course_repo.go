package repository

import (
	"courses-services/database"
	"courses-services/models"
)

func GetCourses() ([]models.Courses, error) {

	var structtype []models.Courses
	query := `SELECT id,name,price FROM courses`
	rows, err := database.DB.Query(query)

	if err != nil {
		return structtype, err
	}
	for rows.Next() {
		var val models.Courses
		err := rows.Scan(&val.ID, &val.Name, &val.Price)
		if err != nil {
			return structtype, err
		}
		structtype = append(structtype, val)
	}
	return structtype, nil
}
