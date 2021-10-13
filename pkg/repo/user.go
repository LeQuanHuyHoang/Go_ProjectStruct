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
}

func (h *Repo) CheckEmail(email string) (*model.User, error) {
	//check user exist
	//query email in db
	rs := &model.User{}
	if err := h.Postgres.Where("email=?", email).First(rs).Error; err != nil {
		//case1: err not found => email not exist
		//case2: db err => can't connect db
		return nil, err
	}
	return rs, nil
}

func (h *Repo) CreateUser(user *model.User) (*model.User, error) {
	if err := h.Postgres.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
