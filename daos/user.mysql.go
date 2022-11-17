package daos

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type (
	MySQLOptions struct {
		ConnectionString string
	}
	iDbConnectFactory interface {
		Create() *sql.DB
	}
	dbConnectFactory struct {
		options MySQLOptions
		logger  *log.Logger
	}
	mysqlUserDAO struct {
		factory iDbConnectFactory
		logger  *log.Logger
	}
)

func newConnectionFactory(options MySQLOptions, logger *log.Logger) iDbConnectFactory {
	return dbConnectFactory{
		options: options,
		logger:  logger,
	}
}

func (factory dbConnectFactory) Create() *sql.DB {
	var db, err = sql.Open("mysql", factory.options.ConnectionString)
	if err != nil {
		factory.logger.Fatalln(err.Error())
	}

	return db
}

func newMySQLUserDAO(factory iDbConnectFactory, logger *log.Logger) *IUserDAO {
	var result IUserDAO = &mysqlUserDAO{
		factory: factory,
		logger:  logger,
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

	_ = db.Close()

}

func (mysql *mysqlUserDAO) Update(po UserPO) {
	const SQL string = "UPDATE `user` SET `Name` = ?, `Age` = ? WHERE Id = ?"

	var db = mysql.factory.Create()

	var _, err = db.Exec(SQL, po.Id, po.Name, po.Age)
	if err != nil {
		mysql.logger.Println(err.Error())
	}

	_ = db.Close()
}

func (mysql *mysqlUserDAO) Delete(id string) {

	const SQL string = "DELETE FROM `user` WHERE `Id` = ?"

	var db = mysql.factory.Create()

	var _, err = db.Exec(SQL, id)
	if err != nil {
		mysql.logger.Println(err.Error())
	}

	_ = db.Close()
}

func (mysql *mysqlUserDAO) FindById(id string) UserPO {

	const SQL string = "SELECT Id, Name, Age FROM `user` WHERE `Id` = ?"

	var db = mysql.factory.Create()
	var po UserPO
	var row = db.QueryRow(SQL, id)
	_ = row.Scan(&po.Id, &po.Name, po.Age)

	_ = db.Close()

	return po
}

func (mysql *mysqlUserDAO) FindAll() []UserPO {

	const SQL string = "SELECT Id, Name, Age FROM `user`"

	var db = mysql.factory.Create()

	var rows, err = db.Query(SQL)

	if err != nil {
		_ = db.Close()
		mysql.logger.Println(err.Error())
		return nil
	}

	var result []UserPO
	for rows.Next() {
		var po UserPO
		_ = rows.Scan(&po.Id, &po.Name, &po.Age)
		result = append(result, po)
	}

	_ = db.Close()
	return result
}
