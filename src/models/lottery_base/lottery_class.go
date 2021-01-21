package lottery_base

import (
	"gorm.io/gorm"
)

//彩种类型表
type LotteryClass struct {
	Id           int64  `gorm:"type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	Name         string `gorm:"type:VARCHAR(50);NOT NULL"`                         //分类名称
	Code         string `gorm:"type:VARCHAR(50);NOT NULL"`                         //分类编码
	Remarks      string `gorm:"type:VARCHAR(255);"`                                //备注
	DisplayOrder int32  `gorm:"type:INT(11);NOT NULL"`                             //排序
	CreatedAt    int64  `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime;"` //创建时间
	UpdatedAt    int64  `gorm:"column:update_time;type:TIMESTAMP;autoUpdateTime;"` //最后更新时间
	DeletedAt    gorm.DeletedAt
}
