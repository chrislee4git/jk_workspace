package db

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestMyDB_ExecSql(t *testing.T) {
	//初始化本地数据库

	db,err := NewMyDB("localhost:3306", "root","",100,3)
	if err != nil {
		fmt.Errorf("%v", errors.Wrap(err,"Init MyDB error!"))
		panic("Init DB Error!")
	}
	_, err = db.ExecSql("insert into  addresses (`address`, `hash`) values ('APL-QB56-NZBY-RHJD-CADDK', '') ")
	if err != nil {
		fmt.Errorf("%v", errors.Wrap(err," MyDB select Error,main line 23 !"))
		if errors.Is(err, sql.ErrNoRows)  {
			fmt.Println("MyDB select Error, ErrNoRows!")
		}else{
			fmt.Sprintf("%v",errors.Cause(err).Error())
		}
		return
	}
}

func TestMyDB_Select(t *testing.T) {
	db,err := NewMyDB("localhost:3306", "root","",100,3)
	if err != nil {
		fmt.Errorf("%v", errors.Wrap(err,"Init MyDB error!"))
		panic("Init DB Error!")
	}

	rows, err := db.Select("select pub,type from addresses where address = 'APL-QB56-NZBY-RHJD-CADDK' ")
	//rows, err := db.ExecSql("insert into  addresses values ('APL-QB56-NZBY-RHJD-CADDK') ")
	if err != nil {
		fmt.Errorf("%v", errors.Wrap(err," MyDB select Error,mydb_test.go line 40 !"))
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
			fmt.Errorf("%v", errors.Wrap(err," MyDB Scan Error,mydb_test.go line 57 !"))
		}

	}

	fmt.Printf("Select sucess ! pubkey = %s , type = %s", pubkey, txtype)
	return
}
