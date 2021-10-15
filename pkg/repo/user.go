package repo

import (
	_ "database/sql"
	"project-struct/pkg/model"

	"gorm.io/gorm"
)

type Repo struct {
	Postgres *gorm.DB
	//Mongodb databasemigrationservice.MongoDbSettings
	//Mysql sql.DB
}

func NewRepo(pg *gorm.DB) IRepo {
	repo := &Repo{
		Postgres: pg,
	}
	return repo
}

type IRepo interface {
	CheckEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	ViewProfile(userID string) (*model.User, error)
	UpdateUser(userID, newpass string) (*model.User, error)
	Delete(userID string) error
}

func (r *Repo) CheckEmail(email string) (*model.User, error) {
	//check user exist
	//query email in db
	rs := &model.User{}
	if err := r.Postgres.Where("email=?", email).First(rs).Error; err != nil {
		//case1: err not found => email not exist
		//case2: db err => can't connect db
		return nil, err
	}
	return rs, nil
}

func (r *Repo) CreateUser(user *model.User) (*model.User, error) {
	if err := r.Postgres.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repo) ViewProfile(userID string) (*model.User, error) {
	rs := &model.User{}
	if err := r.Postgres.Where("ID=?", userID).First(rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *Repo) UpdateUser(userID, newpass string) (*model.User, error) {
	rs := &model.User{}
	if err := r.Postgres.Model(&rs).Where("ID = ?", userID).Update("Password", newpass).Error; err != nil {
		return nil, err
	}
	/* 	if err := r.Postgres.Where("ID=?", userID).First(rs).Error; err != nil {
		return nil, err
	} */
	return rs, nil
}

func (r *Repo) Delete(userID string) error {
	rs := &model.User{}
	if err := r.Postgres.Where("ID = ?", userID).Delete(&rs).Error; err != nil {
		return err
	}
	/* 	if err := r.Postgres.Where("ID=?", userID).First(rs).Error; err != nil {
		return nil, err
	} */
	return nil
}
