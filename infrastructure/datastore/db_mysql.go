package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yoshinorihisakawa/sample-api-hoop/conf"
)

func NewMySqlDB() *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		conf.C.Database.User,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Dbname,
	)

	conn, err := gorm.Open("mysql", connectionString)

	if nil != err {
		panic(err)
	}

	// 応答確認
	err = conn.DB().Ping()
	if nil != err {
		panic(err)
	}
	// sqlログの詳細を出力
	conn.LogMode(true)

	// DBのエンジンを設定
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
