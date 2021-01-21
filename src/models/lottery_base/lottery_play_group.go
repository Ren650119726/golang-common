package lottery_base

import "gorm.io/gorm"

//玩法分组表
type LotteryPlayGroup struct {
	Id           int64  `gorm:"type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	ClassId      int64  `gorm:"type:BIGINT(20);NOT NULL"`  //分类id
	Name         string `gorm:"type:VARCHAR(50);NOT NULL"` //名称
	Code         string `gorm:"type:VARCHAR(50);"`         //编码
	Model        int8   `gorm:"type:TINYINT(1);NOT NULL"`
	DisplayOrder int32  `gorm:"type:INT(11);NOT NULL"`
	Remarks      string `gorm:"type:VARCHAR(255);"`
	CreatedAt    int64  `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime;"`
	UpdatedAt    int64  `gorm:"column:update_time;type:TIMESTAMP;autoUpdateTime;"`
	DeletedAt    gorm.DeletedAt
}
