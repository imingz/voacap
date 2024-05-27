package db

import (
	"voacap/pkg/constants"
)

type Station struct {
	StationID  int64   `gorm:"primaryKey;column:stationID"`
	Slatitude  float64 `gorm:"column:SLatitude"`
	Slongitude float64 `gorm:"column:SLongitude"`
	Sname      string  `gorm:"column:SName"`
}

func (Station) TableName() string {
	return constants.StationTableName
}

// QueryAllStation 查询所有
func QueryAllStation() ([]*Station, error) {
	var stations []*Station
	if err := DB.Find(&stations).Error; err != nil {
		return nil, err
	}
	return stations, nil
}

// AddStation 添加
func AddStation(station *Station) error {
	return DB.Create(station).Error
}
