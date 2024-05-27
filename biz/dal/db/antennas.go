package db

import (
	"fmt"
	"voacap/pkg/constants"
)

type Antennas struct {
	AntennaID int64  `json:"antennaID"`
	Aname     string `json:"Aname"`
	Afile     string `json:"Afile"`
	AfbandMin int64  `json:"AfbandMin"`
	AfbandMax int64  `json:"AfbandMax"`
}

func (Antennas) TableName() string {
	return constants.AntennaTableName
}

// QueryAllAntennas 查询所有
func QueryAllAntennas() ([]*Antennas, error) {
	var antennas []*Antennas
	if err := DB.Find(&antennas).Error; err != nil {
		fmt.Printf("QueryAllAntennas error: %v\n", err)
		return nil, err
	}
	return antennas, nil
}
