package daos

import (
	"database/sql"
	"log"
)

type mysqlUserDAO struct {
	factory iDbConnectFactory
	logger  *log.Logger
}

type (
	MySQLOptions struct {
		ConnectionString string
	}
	iDbConnectFactory interface {
		Create() *sql.DB
	}
)

type dbConnectFactory struct {
	options MySQLOptions
	logger  *log.Logger
}

func newConnectionFactory(options MySQLOptions) iDbConnectFactory {
	return dbConnectFactory{
		options: options,
		logger:  log.Default(),
	}
}

func (factory dbConnectFactory) Create() *sql.DB {
	var db, err = sql.Open("mysql", factory.options.ConnectionString)
	if err != nil {
		factory.logger.Println(err.Error())
	}

	return db
}

func newMySQLUserDAO(factory iDbConnectFactory) *IUserDAO {
	var result IUserDAO = &mysqlUserDAO{
		factory: factory,
		logger:  log.Default(),
	}

	return &result
}

func (mysql *mysqlUserDAO) Add(po UserPO) {
	var db = mysql.factory.Create()

	const SQL string = "INSERT INTO `user` (`Id`,`Name`,`Age`) VALUES (?,?,?)"

	var _, err = db.Exec(SQL, po.Id, po.Name, po.Age)
	if err != nil {
		println(err.Error())
	}

	db.Close()

}

func (mysql *mysqlUserDAO) Update(po UserPO) {

	const SQL string = "UPDATE `user` SET `Name` = ?, `Age` = ? WHERE Id = ?"

	var db = mysql.factory.Create()

	var _, err = db.Exec(SQL, po.Id, po.Name, po.Age)
	if err != nil {
		mysql.logger.Println(err.Error())
	}

	db.Close()
}

func (mysql *mysqlUserDAO) Delete(id string) {

	const SQL string = "DELETE FROM `user` WHERE `Id` = ?"

	var db = mysql.factory.Create()

	var _, err = db.Exec(SQL, id)
	if err != nil {
		mysql.logger.Println(err.Error())
	}

	db.Close()
}

func (mysql *mysqlUserDAO) FindById(id string) UserPO {

	const SQL string = "SELECT Id, Name, Age FROM `user` WHERE `Id` = ?"

	var db = mysql.factory.Create()
	var po UserPO
	var row = db.QueryRow(SQL, id)
	row.Scan(&po.Id, &po.Name, po.Age)

	db.Close()

	return po
}

func (mysql *mysqlUserDAO) FindAll() []UserPO {

	const SQL string = "SELECT Id, Name, Age FROM `user`"

	var db = mysql.factory.Create()

	var rows, err = db.Query(SQL)

	if err != nil {
		db.Close()
		mysql.logger.Println(err.Error())
		return nil
	}

	var result []UserPO
	for rows.Next() {
		var po UserPO
		rows.Scan(&po.Id, &po.Name, &po.Age)
		result = append(result, po)
	}

	db.Close()
	return result
}
