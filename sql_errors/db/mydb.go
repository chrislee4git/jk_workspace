package db

import (
	"database/sql"
	//"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

type MyDB struct {
	SqlDb 		*sql.DB				//mysql连接池
	Url   		string				//连接信息url
	UserName 	string				//用户名
	PassWd 		string				//密码
	MaxOpenConns	int				//连接池最多保存多少个数据库连接
	MaxIdleConns	int				//连接池中保持的最大空闲连接的数量

}

func NewMyDB(url string, user string,pw  string,maxopenconns, maxidleconns int) (*MyDB, error)  {
	db,err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/aplwallet?timeout=10s&&loc=Local&parseTime=true&allowOldPasswords=1",user, pw, url))
	if err != nil {
		return nil, err
	}
	//defer  db.Close()

	db.SetMaxOpenConns(maxopenconns)
	db.SetMaxIdleConns(maxidleconns)
	return &MyDB{
		SqlDb: db,
		Url: url,
		UserName: user,
		PassWd: pw,
		MaxOpenConns: maxopenconns,
		MaxIdleConns: maxidleconns,
	},nil
}

func (db *MyDB) Test() error  {
	return db.SqlDb.Ping()
}

func (db *MyDB) Select (sqlstr string) (*sql.Rows, error) {
	return  db.SqlDb.Query(sqlstr)
}

func (db *MyDB) ExecSql (sqlstr string) (sql.Result, error) {
	return db.SqlDb.Exec(sqlstr)
}

//sql执行返回sql.ErrNoRows 错误
//func IsErrNoRows(err error) bool{
//	return errors.Is(err, sql.ErrNoRows)
//}

