// Code generated by hertz generator.

package station

import (
	"context"

	"voacap/internal/pkg/utils"
	station "voacap/internal/voacap/biz/model/station"
	service "voacap/internal/voacap/biz/service/station"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetStations .
// @router /stations/getStations [GET]
func GetStations(ctx context.Context, c *app.RequestContext) {
	var err error
	var req station.GetStationsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, station.GetStationsResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	data, err := service.NewStationService(ctx, c).GetStations()

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, station.GetStationsResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
		Data: data,
	})
}

// AddStation .
// @router /stations/addStation [POST]
func AddStation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req station.AddStationRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, station.AddStationResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	err = service.NewStationService(ctx, c).AddStation(&req)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, station.AddStationResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
	})
}

// UpdateStationById .
// @router /stations/updateStationById [POST]
func UpdateStationById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req station.UpdateStationByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, station.UpdateStationByIdResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	err = service.NewStationService(ctx, c).UpdateStationById(&req)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, station.UpdateStationByIdResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
	})
}

// DeleteStationById .
// @router /stations/deleteStationById [POST]
func DeleteStationById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req station.DeleteStationByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, station.DeleteStationByIdResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	err = service.NewStationService(ctx, c).DeleteStationById(&req)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, station.DeleteStationByIdResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
	})
}