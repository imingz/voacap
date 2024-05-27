package service

import (
	"context"
	"time"
	"voacap/biz/dal/db"
	"voacap/biz/model/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type LinkService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewLinkService create LinkService service
func NewLinkService(ctx context.Context, c *app.RequestContext) *LinkService {
	return &LinkService{ctx: ctx, c: c}
}

// GetAntennas 获取所有天线信息
func (s *LinkService) GetLinks() ([]*common.Link, error) {
	links, err := db.QueryAllLink()
	if err != nil {
		return nil, err
	}

	result := make([]*common.Link, len(links))
	for i, link := range links {
		result[i] = &common.Link{
			CircuitReliability: link.CircuitReliability,
			Coefficient:        link.Coefficient,
			Date:               link.Date.Format(time.DateOnly),
			IxAntennaFbandMax:  link.IxAntennaFbandMax,
			IxAntennaFbandMin:  link.IxAntennaFbandMin,
			IxAntennaFile:      link.IxAntennaFile,
			IxAntennaID:        link.IxAntennaID,
			IxAntennaName:      link.IxAntennaName,
			IxPower:            link.IxPower,
			IxStationID:        link.IxStationID,
			IxStationLat:       link.IxStationLat,
			IxStationLng:       link.IxStationLng,
			IxStationName:      link.IxStationName,
			LinkID:             link.LinkID,
			LinkType:           link.LinkType,
			Noise:              link.Noise,
			RxAntennaFbandMax:  link.RxAntennaFbandMax,
			RxAntennaFbandMin:  link.RxAntennaFbandMin,
			RxAntennaFile:      link.RxAntennaFile,
			RxAntennaID:        link.RxAntennaID,
			RxAntennaName:      link.RxAntennaName,
			RxStationID:        link.RxStationID,
			RxStationLat:       link.RxStationLat,
			RxStationLng:       link.RxStationLng,
			RxStationName:      link.RxStationName,
			SNR:                link.SNR,
			SunspotNum:         link.SunspotNum,
			TimeType:           link.TimeType,
			TxAntennaFbandMax:  link.TxAntennaFbandMax,
			TxAntennaFbandMin:  link.TxAntennaFbandMin,
			TxAntennaFile:      link.TxAntennaFile,
			TxAntennaID:        link.TxAntennaID,
			TxAntennaName:      link.TxAntennaName,
			TxPower:            link.TxPower,
			TxStationID:        link.TxStationID,
			TxStationLat:       link.TxStationLat,
			TxStationLng:       link.TxStationLng,
			TxStationName:      link.TxStationName,
		}
	}

	return result, nil
}
