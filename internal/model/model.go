package model

import (
	"fmt"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunModel == "debug" {
		// 	启用debug
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// 使用钩子函数处理公共字段

// BeforeCreate 创建前
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Unix()
	m.CreatedOn = uint32(now)
	m.ModifiedOn = uint32(now)
	return nil
}

func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.ModifiedOn = uint32(time.Now().Unix())
	return nil
}

func (m *Model) BeforeDelete(tx *gorm.DB) error {
	// 查询表是否存在 DeleteOn 和 IsDel 字段
	hasDeleteOnAndIsDel := hasDeleteOnAndIsDelFields(tx, tx.Statement.Schema.Table)
	if hasDeleteOnAndIsDel {
		// 如果存在这两个字段，则执行 UPDATE 操作更新这两个字段的值
		m.DeletedOn = uint32(time.Now().Unix())
		m.IsDel = 1
		return tx.Model(m).Updates(map[string]interface{}{"DeleteOn": m.DeletedOn, "IsDel": m.IsDel}).Error
	} else {
		// 否则，执行 DELETE 操作
		return tx.Delete(m).Error
	}
}

// 查询表是否存在 DeleteOn 和 IsDel 字段
func hasDeleteOnAndIsDelFields(db *gorm.DB, table string) bool {
	// 查询字段信息
	var columns []string
	db.Raw(fmt.Sprintf("PRAGMA table_info(%s);", table)).Pluck("name", &columns)

	// 检查字段是否存在
	hasDeleteOn := false
	hasIsDel := false
	for _, column := range columns {
		if column == "DeleteOn" {
			hasDeleteOn = true
		}
		if column == "IsDel" {
			hasIsDel = true
		}
	}

	return hasDeleteOn && hasIsDel
}

// 新增行为的回调
//func updateTimeStampForCreateCallback(db *gorm.DB) {
//	if db.Error != nil {
//		nowTime := time.Now().Unix()
//
//		//获取当前是否包含所需的字段。
//		if createTimeField, ok := db.Statement.Schema.FieldsByName["CreateOn"]; ok {
//			if createTimeField.NotNull { //如果字段不为空
//				db.Statement.SetColumn("CreateOn", nowTime)
//			}
//		}
//		if modifyTimeField, ok := db.Statement.Schema.FieldsByName["ModifiedOn"]; ok {
//			if modifyTimeField.NotNull {
//				db.Statement.SetColumn("ModifiedOn", nowTime)
//			}
//		}
//	}
//}

// 更新行为的回调
//func updateTimeStampForUpdateCallback(db *gorm.DB) {
//	// 获取当前设置了标识gorm:update_column的字段属性
//	if _, ok := db.Statement.Get("gorm:update_column"); !ok {
//		db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
//	}
//}

// 删除行为的回调
//func deleteCallback(db *gorm.DB) {
//	if db.Error != nil {
//		var extraOption string
//
//		if str, ok := db.Statement.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		// 查找是否包含某个字段
//		deletedOnField, hasDeletedOnField := db.Statement.Schema.FieldsByName["DeletedOn"]
//		isDelField, hasIsDelFiedld := db.Statement.Schema.FieldsByName["IsDel"]
//		if !db.Statement.Unscoped && hasDeletedOnField && hasIsDelFiedld {
//			now := time.Now().Unix()
//			db.Statement.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v, %v=%v%v%v",
//				db.Statement.qu
//				))
//		}
//
//	}
//}
