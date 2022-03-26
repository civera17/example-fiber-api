package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/civera17/fintech-test/models"
	"gorm.io/gorm"
)

const (
	callBackBeforeName = "custom:before"
	callBackAfterName  = "custom:after"
	startTime          = "_start_time"
)

// Plugin GORM plugin interface
type Plugin interface {
	Name() string
	Initialize(*gorm.DB) error
}

// SqlInfoPlugin implements GORM plugin interface , registers callbacs before and after 
// 		all types of queries to inject a query that writes to db info about query
type SqlInfoPlugin struct{}

func (op *SqlInfoPlugin) Name() string {
	return "sqlInfoPlugin"
}

func (op *SqlInfoPlugin) Initialize(db *gorm.DB) (err error) {
	// //  Prior to the start
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// //  After the end
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &SqlInfoPlugin{}

func before(db *gorm.DB) {
	ConnectCallbackFreeDb()
	db.InstanceSet(startTime, time.Now())
}

func after(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	sqlInfo := new(models.QueryInfo)
	sqlInfo.SQL = sql
	sqlInfo.CostSeconds = fmt.Sprintf("%f", time.Since(ts).Seconds())
	sqlInfo.Type = strings.Split(sql, " ")[0]

	CBDB.Db.Create(&sqlInfo)

	fmt.Println("Created info about sql")
}
