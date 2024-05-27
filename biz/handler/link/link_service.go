// Code generated by hertz generator.

package link

import (
	"context"

	link "voacap/biz/model/link"
	service "voacap/biz/service/link"
	"voacap/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetLinks .
// @router /links/getLinks [GET]
func GetLinks(ctx context.Context, c *app.RequestContext) {
	var err error
	var req link.GetLinksRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, link.GetLinksResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	data, err := service.NewLinkService(ctx, c).GetLinks()

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, link.GetLinksResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
		Data: data,
	})
}

// AddLink .
// @router /links/addLink [POST]
func AddLink(ctx context.Context, c *app.RequestContext) {
	var err error
	var req link.AddLinkRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, link.AddLinkResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	err = service.NewLinkService(ctx, c).AddLink(&req)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, link.AddLinkResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
	})
}

// UpdateLinkById .
// @router /links/updateLinkById [POST]
func UpdateLinkById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req link.UpdateLinkByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, link.UpdateLinkByIdResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	err = service.NewLinkService(ctx, c).UpdateLinkById(&req)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, link.UpdateLinkByIdResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
	})
}
