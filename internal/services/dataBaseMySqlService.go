package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/vugsk/CurrencyExchangerProjectGoLang/internal/models"
)

type DataBaseService struct {
	models.InterfaceDataBase
	name string
	db   *sql.DB
}

// private

func (dbs DataBaseService) error(message string) error {
	return fmt.Errorf(message)
}

// public

func (dbs *DataBaseService) Connect(userName, password, dataBaseName string) error {
	var err error
	dbs.name = dataBaseName
	dbs.db, err = sql.Open("mysql", userName+":"+password+"@tcp(localhost:3306)/"+dataBaseName)
	if err != nil {
		return err
	}

	return dbs.db.Ping()
}

func (dbs DataBaseService) Close() error {
	return dbs.db.Close()
}

func (dbs DataBaseService) Insert(tableName string, obj interface{}) error {
	if obj == nil {
		return dbs.error("obj is nil")
	}

	if tableName == "" {
		return dbs.error("tableName is empty")
	}

	switch value := obj.(type) {
	case models.RequestRegistration:
		_, err := dbs.db.Exec(fmt.Sprintf("insert into %s.%s (id_profile, name, login, email, id_telegram, password) values(?,?,?,?,?,?)", dbs.name, tableName),
			GeneratingAnIdProfile(value),
			" ",
			value.Login,
			value.Email,
			" ",
			value.Password)
		return err
	}

	return dbs.error("not found -> struct")
}

func (dbs DataBaseService) Update(tableName string, obj interface{}) error {

	return nil
}

func (dbs DataBaseService) Delete(tableName string, idProfile string) error {
	return nil
}

func (dbs DataBaseService) Read(tableName string, idProfile string) (err error, obj interface{}) {

	return nil, nil
}

func (dbs DataBaseService) ReadAll(tableName string) (objs []interface{}, err error) {
	return nil, nil
}

func (dbs DataBaseService) Count(tableName string) (count uint, err error) {
	return 0, nil
}

func (dbs DataBaseService) IsEmpty(tableName string) bool {
	return false
}

func (dbs DataBaseService) Ping() error {
	return dbs.db.Ping()
}

func (dbs DataBaseService) Migrate() error {
	return nil
}

func (dbs DataBaseService) IsConnected() bool {
	return false
}
