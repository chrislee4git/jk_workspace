package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	mydb "mydb/db"
)
var db *mydb.MyDB

func init()  {
	//初始化本地数据库
	var err error
	db,err = mydb.NewMyDB("localhost:3306", "root","",100,3)
	if err != nil {
		fmt.Errorf("%v", errors.Wrap(err,"Init MyDB error!"))
		panic("Init DB Error!")
	}
}
/*
*	sql Dao层返回错误可能是sql.ErrNoRows  ,此时不应该wrap这个错误,因为调用dao层接口的地方有很多，在这边Wrap没有意义，应该把错误抛到上层，在上层调用的地方去Warp,
*	可以携带更多的信息，打印堆栈信息等。
*   同时，可以用pkg/errors 中error.Is()接口去判断底层错误是否是sql.ErrNoRows，也可以打出根音errors.Cause(err).Error
 */
func main() {
	rows, err := db.Select("select pub,type from addresses where address = 'APL-QB56-NZBY-RHJD-CADDK' ")
	//rows, err := db.ExecSql("insert into  addresses values ('APL-QB56-NZBY-RHJD-CADDK') ")
	if err != nil {
		fmt.Errorf("%v", errors.Wrap(err," MyDB select Error,main line 23 !"))
		if errors.Is(err, sql.ErrNoRows)  {
			fmt.Println("MyDB select Error, ErrNoRows!")
		}else{
			fmt.Sprintf("%v",errors.Cause(err).Error())
		}
		return
	}
	defer rows.Close()

	var (
		pubkey 	string
		txtype   string
	)
	for rows.Next() {
		err := rows.Scan(&pubkey, &txtype)
		if err != nil {
			fmt.Errorf("%v", errors.Wrap(err," MyDB Scan Error,main line 42 !"))
		}

	}

	fmt.Printf("Select sucess ! pubkey = %s , type = %s", pubkey, txtype)
	return
}
