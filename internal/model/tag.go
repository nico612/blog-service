package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

//对标签模块的模型操作进行封装，并且只与实体产生关系

/**
gorm相关操作
Model: 指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰。
Where：设置筛选条件，接受 map，struct 或 string 作为条件。
Offset：偏移量，用于指定开始返回记录之前要跳过的记录数。
Limit：限制检索的记录数。
Find：查找符合筛选条件的记录。
Updates：更新所选字段。
Delete：删除数据。
Count：统计行为，用于统计模型的记录数。
*/

func (t Tag) Count(db *gorm.DB) (int64, error) {

	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Crate(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 更新整个对象需要用updates, 而更新某个字段使用update
// db.Model(&Tag{}).Where("id = ?", t.ID).Update("name", "newName")
// GORM 中使用 struct 类型传入进行更新时，GORM 是不会对值为零值的字段进行变更。
// 这又是为什么呢，我们可以猜想，更根本的原因是因为在识别这个结构体中的这个字段值时，很难判定是真的是零值，还是外部传入恰好是该类型的零值.
func (t Tag) Update(db *gorm.DB, values interface{}) error {

	return db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, t.IsDel).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND  is_del = ?", t.ID, 0).Delete(&t).Error
}
