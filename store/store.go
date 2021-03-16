package store

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"simpleHttpServer/config"
	"time"
)

// mysql連線優化設定
type Conf struct {
	// 連線最大空閒數(長連線)
	// 每個連線都會消耗一定的內存，要依據機器容量與效能做決定
	MaxIdle int

	// 全部連線的最大值(長 or 短 連線總數)
	MaxOpen int

	// 每個連線如果閒置多久就應該從pool捨棄掉
	// 需要注意此值必須比mysql interactive_timeout與wait_timeout小一點
	// 因為mysql自動捨棄掉長連線時不會通知應用程式這時從pool拿出來的連線就
	// 可能沒用會導致此用此連線的sql操作失敗
	MaxLifetime time.Duration
}

type Context interface {
	CreateMember(*Member) error
}

type dataStore struct {
	*gorm.DB

	// database type
	driver string
}

// 根據config建立一個database store
func New(conf Conf) Context {
	db := &dataStore{
		DB:     Open(config.Conf.DB.Driver, conf),
		driver: config.Conf.DB.Driver,
	}

	return db
}

// 連結database
func Open(driver string, conf Conf) *gorm.DB {
	dbConfig := config.Conf.DB

	source := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&collation=%v&parseTime=true",
		dbConfig.User,      //dbConfig.User,
		dbConfig.Password,  //dbConfig.Password,
		dbConfig.Host,      //dbConfig.Host,
		dbConfig.Port,      //dbConfig.Port,
		dbConfig.Database,  //dbConfig.Database,
		dbConfig.Charset,   //dbConfig.Charset,
		dbConfig.Collation, //dbConfig.Collation,
	)

	db, err := gorm.Open(driver, source)
	if err != nil {
		panic(fmt.Sprintf("database connection failed-%s", err.Error()))
	}
	if config.Conf.App.Debug {
		db.LogMode(true)
	}

	d := db.CommonDB().(*sql.DB)

	// 默認不設定是2所以要大於2
	if conf.MaxIdle > 2 {
		d.SetMaxIdleConns(conf.MaxIdle)
		d.SetMaxOpenConns(conf.MaxOpen)
		d.SetConnMaxLifetime(conf.MaxLifetime)
	}

	return db
}
