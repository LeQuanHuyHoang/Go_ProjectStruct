package repo

import (
	"crawl-data/pkg/model"
	_ "database/sql"

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
