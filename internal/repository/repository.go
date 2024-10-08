package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}

func (r *Repository[T]) SoftDelete(entity *T) error {
	return r.DB.Model(entity).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *Repository[T]) CountById(id any) (int64, error) {
	var total int64
	err := r.DB.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindById(entity *T, id any) error {
	return r.DB.Where("id = ?", id).Take(entity).Error
}

func (r *Repository[T]) FindAll(entities *[]T) error {
	return r.DB.Find(entities).Error
}

func (r *Repository[T]) FindByField(entities *[]T, field string, value any) error {
	return r.DB.Where(field+" = ?", value).Find(entities).Error
}
