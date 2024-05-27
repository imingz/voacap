package db

import (
	"time"
	"voacap/pkg/constants"
)

type Link struct {
	LinkID             int64     `gorm:"primaryKey;column:linkID"`
	LinkType           string    `gorm:"column:linkType;not null;comment:'链路类型'"`
	Date               time.Time `gorm:"column:date;not null;comment:'日期'"`
	TimeType           string    `gorm:"column:timeType;not null;comment:'时间类型'"`
	Coefficient        string    `gorm:"column:coefficient;not null;comment:'电离层系数'"`
	SunspotNum         int64     `gorm:"column:sunspotNum;not null;comment:'太阳黑子数'"`
	CircuitReliability float64   `gorm:"column:circuitReliability;not null;comment:'电路可靠度'"`
	SNR                float64   `gorm:"column:SNR;not null;comment:'所需信噪比'"`
	Noise              float64   `gorm:"column:noise;not null;comment:'接收点噪声'"`
	TxStationID        int64     `gorm:"column:txStationID;not null;comment:'发射站点ID'"`
	TxStationName      string    `gorm:"column:txStationName;null;default:null"`
	TxStationLat       float64   `gorm:"column:txStationLat;not null"`
	TxStationLng       float64   `gorm:"column:txStationLng;not null"`
	TxAntennaID        int64     `gorm:"column:txAntennaID;not null;comment:'发射天线ID'"`
	TxAntennaName      string    `gorm:"column:txAntennaName;not null"`
	TxAntennaFile      string    `gorm:"column:txAntennaFile;not null"`
	TxAntennaFbandMin  int64     `gorm:"column:txAntennaFbandMin;not null"`
	TxAntennaFbandMax  int64     `gorm:"column:txAntennaFbandMax;not null"`
	TxPower            float64   `gorm:"column:txPower;not null;comment:'发射功率'"`
	RxStationID        int64     `gorm:"column:rxStationID;not null;comment:'接收站点ID'"`
	RxStationName      string    `gorm:"column:rxStationName;null;default:null"`
	RxStationLat       float64   `gorm:"column:rxStationLat;not null"`
	RxStationLng       float64   `gorm:"column:rxStationLng;not null"`
	RxAntennaID        int64     `gorm:"column:rxAntennaID;not null;comment:'接收天线ID'"`
	RxAntennaName      string    `gorm:"column:rxAntennaName;not null"`
	RxAntennaFile      string    `gorm:"column:rxAntennaFile;not null"`
	RxAntennaFbandMin  int64     `gorm:"column:rxAntennaFbandMin;not null"`
	RxAntennaFbandMax  int64     `gorm:"column:rxAntennaFbandMax;not null"`
	IxStationID        int64     `gorm:"column:ixStationID;null;default:null;comment:'干扰站点ID'"`
	IxStationName      string    `gorm:"column:ixStationName;null;default:null"`
	IxStationLat       float64   `gorm:"column:ixStationLat;null;default:null"`
	IxStationLng       float64   `gorm:"column:ixStationLng;null;default:null"`
	IxAntennaID        int64     `gorm:"column:ixAntennaID;null;default:null;comment:'干扰天线ID'"`
	IxAntennaName      string    `gorm:"column:ixAntennaName;null;default:null"`
	IxAntennaFile      string    `gorm:"column:ixAntennaFile;null;default:null"`
	IxAntennaFbandMin  int64     `gorm:"column:ixAntennaFbandMin;null;default:null"`
	IxAntennaFbandMax  int64     `gorm:"column:ixAntennaFbandMax;null;default:null"`
	IxPower            float64   `gorm:"column:ixPower;null;default:null;comment:'干扰信号发射功率'"`
}

func (Link) TableName() string {
	return constants.LinkTableName
}

// QueryAllLink 查询所有
func QueryAllLink() ([]*Link, error) {
	var links []*Link
	if err := DB.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

// CreateLink 创建
func CreateLink(link *Link) error {
	return DB.Create(link).Error
}
