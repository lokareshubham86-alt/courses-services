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

func GetCourseByID(id int) (models.Courses, error) {
	query := `select id, name , price from courses where id=?`

	var data models.Courses
	var err error
	rows := database.DB.QueryRow(query, id)
	err = rows.Scan(&data.ID, &data.Name, &data.Price)
	if err != nil {
		return data, err
	}
	return data, nil
}
