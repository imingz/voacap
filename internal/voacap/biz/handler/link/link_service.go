// Code generated by hertz generator.

package link

import (
	"context"
	"voacap/internal/pkg/utils"
	link "voacap/internal/voacap/biz/model/link"
	service "voacap/internal/voacap/biz/service/link"
	"voacap/pkg/errno"

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
		resp := errno.ParamErr.WithMessage(err.Error())
		c.JSON(consts.StatusOK, link.GetLinksResponse{
			Code: resp.ErrCode,
			Msg:  resp.ErrMsg,
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
		resp := errno.ParamErr.WithMessage(err.Error())
		c.JSON(consts.StatusOK, link.AddLinkResponse{
			Code: resp.ErrCode,
			Msg:  resp.ErrMsg,
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
		resp := errno.ParamErr.WithMessage(err.Error())
		c.JSON(consts.StatusOK, link.UpdateLinkByIdResponse{
			Code: resp.ErrCode,
			Msg:  resp.ErrMsg,
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

// DeleteLinkById .
// @router /links/deleteLinkById [POST]
func DeleteLinkById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req link.DeleteLinkByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := errno.ParamErr.WithMessage(err.Error())
		c.JSON(consts.StatusOK, link.DeleteLinkByIdResponse{
			Code: resp.ErrCode,
			Msg:  resp.ErrMsg,
		})
		return
	}

	err = service.NewLinkService(ctx, c).DeleteLinkById(&req)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, link.DeleteLinkByIdResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
	})
}

// WriteLink2File .
// @router /links/writeLink2File [POST]
func WriteLink2File(ctx context.Context, c *app.RequestContext) {
	var err error
	var req link.WriteLink2FileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := errno.ParamErr.WithMessage(err.Error())
		c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
			Code: resp.ErrCode,
			Msg:  resp.ErrMsg,
		})
		return
	}

	s := service.NewLinkService(ctx, c)
	err = s.WriteLink2File(&req, true)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	err = s.CalculateLink()
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	transLink, err := s.GetLinkResult()
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	var interferLink []*link.CalculateResult
	var sir []*link.SIR

	if req.LinkType == "干扰" {
		err = s.WriteLink2File(&req, false)
		if err != nil {
			resp := utils.BuildBaseResp(err)
			c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
				Code: resp.StatusCode,
				Msg:  resp.StatusMsg,
			})
			return
		}

		err = s.CalculateLink()
		if err != nil {
			resp := utils.BuildBaseResp(err)
			c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
				Code: resp.StatusCode,
				Msg:  resp.StatusMsg,
			})
			return
		}

		interferLink, err = s.GetLinkResult()
		if err != nil {
			resp := utils.BuildBaseResp(err)
			c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
				Code: resp.StatusCode,
				Msg:  resp.StatusMsg,
			})
			return
		}

		for i := 0; i < len(transLink); i++ {
			SIR := make([]float64, 2, 31) // 初始化容量为31，前2个元素为0.0

			for j := 2; j < 31; j++ {
				dbuInterfer := interferLink[i].DBU[j]
				dbuTrans := transLink[i].DBU[j]
				sir := 0.0
				if dbuInterfer != 0 {
					sir = dbuInterfer - dbuTrans
				}
				SIR = append(SIR, sir)
			}
			sir = append(sir, &link.SIR{GMT: transLink[i].GMT, SIR: SIR})
		}
	}

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, link.WriteLink2FileResponse{
		Code: resp.StatusCode,
		Msg:  resp.StatusMsg,
		Data: &link.WriteLink2FileData{
			TransLink:    transLink,
			InterferLink: interferLink,
			SIR:          sir,
		},
	})
}
