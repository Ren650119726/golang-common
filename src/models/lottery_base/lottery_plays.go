package lottery_base

import (
	"gorm.io/gorm"
)

//玩法
type LotteryPlays struct {
	Id           int64   `gorm:"type:BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	ClassId      int64   `gorm:"type:BIGINT(20);NOT NULL"`
	GroupId      int64   `gorm:"type:BIGINT(20);NOT NULL"`
	TypeId       int64   `gorm:"type:BIGINT(20);"`
	Name         string  `gorm:"type:VARCHAR(50);NOT NULL"`
	Code         string  `gorm:"type:VARCHAR(50);NOT NULL"`
	Odds         float64 `gorm:"type:DECIMAL(13, 2);"`
	DisplayOrder int32   `gorm:"type:INT(11);NOT NULL"`
	Remarks      string  `gorm:"type:VARCHAR(255);"`
	CreatedAt    int64   `gorm:"column:create_time;type:TIMESTAMP;autoCreateTime;"`
	UpdatedAt    int64   `gorm:"column:update_time;type:TIMESTAMP;autoUpdateTime;"`
	DeletedAt    gorm.DeletedAt
}
