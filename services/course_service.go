package services

import (
	"courses-services/models"
	"courses-services/repository"
)

func GetCourses() ([]models.Courses, error) {
	return repository.GetCourses()
}

func GetCourseByID(id int) (models.Courses, error) {
	return repository.GetCourseByID(id)
}
