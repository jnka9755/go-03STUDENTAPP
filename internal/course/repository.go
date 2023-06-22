package course

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *Course) error
		GetAll(filters Filters, offset, limit int) ([]Course, error)
		Get(id string) (*Course, error)
		Delete(id string) error
		Update(id string, course *UpdateCourse) error
		Count(filters Filters) (int, error)
	}

	repository struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepository(l *log.Logger, db *gorm.DB) Repository {

	return &repository{
		db:  db,
		log: l,
	}
}

func (r *repository) Create(course *Course) error {

	if err := r.db.Create(course).Error; err != nil {
		r.log.Println("Repository ->", err)
		return err
	}

	r.log.Println("Repository -> Create user with id: ", course.ID)

	return nil
}

func (r *repository) GetAll(filters Filters, offset, limit int) ([]Course, error) {

	r.log.Println("GetAll course Repository")
	var courses []Course

	tx := r.db.Model(&courses)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)

	result := tx.Order("created_at desc").Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}

func (r *repository) Get(id string) (*Course, error) {

	r.log.Println("Get course Repository")
	course := Course{ID: id}

	if err := r.db.First(&course).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *repository) Delete(id string) error {

	r.log.Println("Delete course Repository")
	course := Course{ID: id}

	if err := r.db.Delete(&course).Error; err != nil {
		return err
	}

	r.log.Println("Repository -> Delete course with id: ", course.ID)

	return nil
}

func (r *repository) Update(id string, course *UpdateCourse) error {

	r.log.Println("Udate course Repository")

	values := make(map[string]interface{})

	if course.Name != nil {
		values["name"] = *course.Name
	}

	if course.StartDate != nil {
		values["start_date"] = *course.StartDate
	}

	if course.EndDate != nil {
		values["end_date"] = *course.EndDate
	}

	if err := r.db.Model(&Course{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Count(filters Filters) (int, error) {

	var count int64
	tx := r.db.Model(Course{})
	tx = applyFilters(tx, filters)

	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name)
	}

	return tx
}
