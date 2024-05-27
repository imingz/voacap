package db

import (
	"fmt"
	"voacap/pkg/constants"
)

type Antenna struct {
	AntennaID int64  `json:"antennaID" gorm:"primaryKey"`
	Aname     string `json:"Aname"`
	Afile     string `json:"Afile"`
	AfbandMin int64  `json:"AfbandMin" gorm:"column:AfbandMin"`
	AfbandMax int64  `json:"AfbandMax" gorm:"column:AfbandMax"`
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
