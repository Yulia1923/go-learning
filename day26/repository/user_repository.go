package repository

import (
	"go-learning/day26/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// 接收一个数据库对象，创建一个userrepository，并把数据库对象保存进去
// repository需要先获得数据库对象才能执行查询，这个函数统一完成初始化，避免每次使用都手动设置内部字段
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{ //创建一个UserRepository结构体值
		db: db,
	}
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
