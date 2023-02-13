package repository

import "golang-unit-test/entity"

type CategoryRepository interface {
	FindByID(ID int) *entity.Category
}
