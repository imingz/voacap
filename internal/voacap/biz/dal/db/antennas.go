package db

import (
	"fmt"
	"voacap/internal/pkg/constants"
)

type Antenna struct {
	AntennaID int64  `gorm:"primaryKey;column:antennaID"`
	Aname     string `gorm:"column:Aname"`
	Afile     string `gorm:"column:Afile"`
	AfbandMin int64  `gorm:"column:AfbandMin"`
	AfbandMax int64  `gorm:"column:AfbandMax"`
}

func (Antenna) TableName() string {
	return constants.AntennaTableName
}

// QueryAllAntennas 查询所有
func QueryAllAntennas() ([]*Antenna, error) {
	var antennas []*Antenna
	if err := DB.Find(&antennas).Error; err != nil {
		fmt.Printf("QueryAllAntennas error: %v\n", err)
		return nil, err
	}
	return antennas, nil
}

// AddAntenna 添加
func AddAntenna(antenna Antenna) error {
	return DB.Create(&antenna).Error
}

// UpdateAntennaById 更新
func UpdateAntennaById(antenna Antenna) error {
	return DB.Save(&antenna).Error
}

// DeleteAntennaById 删除
func DeleteAntennaById(antennaID int64) error {
	return DB.Delete(&Antenna{}, antennaID).Error
}
