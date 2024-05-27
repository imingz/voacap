package service

import (
	"context"
	"time"
	"voacap/biz/dal/db"
	"voacap/biz/model/common"
	"voacap/biz/model/link"

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

// AddLink 添加链路信息
func (s *LinkService) AddLink(req *link.AddLinkRequest) error {
	date, err := time.Parse(time.DateOnly, req.SysConfig.Date)
	if err != nil {
		return err
	}

	var linkType string

	if req.InterferStation.TransStation.StationID == -1 {
		linkType = "干扰"
	} else {
		linkType = "通信"
	}

	l := &db.Link{
		LinkType: linkType,

		// SysConfig
		Date:               date,
		TimeType:           req.SysConfig.TimeType,
		Coefficient:        req.SysConfig.Coefficient,
		SunspotNum:         req.SysConfig.SunspotNum,
		CircuitReliability: req.SysConfig.CircuitReliability,
		SNR:                req.SysConfig.SNR,
		Noise:              req.SysConfig.Noise,

		// TransStation
		TxStationID:   req.TransStation.TransStation.StationID,
		TxStationLat:  req.TransStation.TransStation.Slatitude,
		TxStationLng:  req.TransStation.TransStation.Slongitude,
		TxStationName: req.TransStation.TransStation.Sname,

		TxAntennaID:       req.TransStation.TransAntenna.AntennaID,
		TxAntennaName:     req.TransStation.TransAntenna.Aname,
		TxAntennaFile:     req.TransStation.TransAntenna.Afile,
		TxAntennaFbandMin: req.TransStation.TransAntenna.AfbandMin,
		TxAntennaFbandMax: req.TransStation.TransAntenna.AfbandMax,

		TxPower: req.TransStation.TransPower,

		// RecvStation
		RxStationID:   req.RecvStation.RecvStation.StationID,
		RxStationLat:  req.RecvStation.RecvStation.Slatitude,
		RxStationLng:  req.RecvStation.RecvStation.Slongitude,
		RxStationName: req.RecvStation.RecvStation.Sname,

		RxAntennaID:       req.RecvStation.RecvAntenna.AntennaID,
		RxAntennaName:     req.RecvStation.RecvAntenna.Aname,
		RxAntennaFile:     req.RecvStation.RecvAntenna.Afile,
		RxAntennaFbandMin: req.RecvStation.RecvAntenna.AfbandMin,
		RxAntennaFbandMax: req.RecvStation.RecvAntenna.AfbandMax,

		// InterferStation
		IxStationID:   req.InterferStation.TransStation.StationID,
		IxStationLat:  req.InterferStation.TransStation.Slatitude,
		IxStationLng:  req.InterferStation.TransStation.Slongitude,
		IxStationName: req.InterferStation.TransStation.Sname,

		IxAntennaID:       req.InterferStation.TransAntenna.AntennaID,
		IxAntennaName:     req.InterferStation.TransAntenna.Aname,
		IxAntennaFile:     req.InterferStation.TransAntenna.Afile,
		IxAntennaFbandMin: req.InterferStation.TransAntenna.AfbandMin,
		IxAntennaFbandMax: req.InterferStation.TransAntenna.AfbandMax,

		IxPower: req.InterferStation.TransPower,
	}

	return db.CreateLink(l)
}
