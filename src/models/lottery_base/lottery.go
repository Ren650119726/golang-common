package lottery_base

import (
	"gorm.io/gorm"
)

//彩种表
type Lottery struct {
	Id           int64  `gorm:"type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	ClassId      int64  `gorm:"type:BIGINT(20);NOT NULL"`                          //分类id
	Name         string `gorm:"type:VARCHAR(50);NOT NULL"`                         //名称
	Code         string `gorm:"type:VARCHAR(50);NOT NULL"`                         //编码
	PeriodType   int8   `gorm:"type:TINYINT(1);"`                                  //时期类型；1-1分钟1期
	Disable      int8   `gorm:"type:TINYINT(1);"`                                  //禁用1-禁用、0-启用
	IsSelf       int8   `gorm:"type:TINYINT(1);NOT NULL"`                          //官彩或私彩;1-私彩;0-管彩
	Remarks      string `gorm:"type:VARCHAR(255);"`                                //备注
	DisplayOrder int32  `gorm:"type:INT(11);NOT NULL"`                             //排序
	CreatedAt    int64  `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime;"` //创建时间
	UpdatedAt    int64  `gorm:"column:update_time;type:TIMESTAMP;autoUpdateTime;"` //最后更新时间
	DeletedAt    gorm.DeletedAt
}
