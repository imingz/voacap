package db

import (
	"fmt"
	"voacap/pkg/constants"
)

type Antenna struct {
	AntennaID int64 `gorm:"column:antennaID"`
	Aname     string
	Afile     string
	AfbandMin int64 `gorm:"column:AfbandMin"`
	AfbandMax int64 `gorm:"column:AfbandMax"`
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

func AddAntenna(antenna Antenna) error {
	return DB.Create(&antenna).Error
}
