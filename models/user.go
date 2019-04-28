package models

type User struct {
	Id   string `gorm:"type:varchar(32) UNSIGNED COMMENT '用户ID' AUTO_INCREMENT;PRIMARY_KEY;NOT NULL;" json:"id"`
	Name string `gorm:"type:varchar(255) COMMENT '名称';NOT NULL;" json:"name"`
}
