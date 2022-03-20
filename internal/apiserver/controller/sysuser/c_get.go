package sysuser

import (
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) GetSysUserById(c *gin.Context) {
	var idInfo schema.ById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}
	api, err := u.srv.GetSysUserById(c.Request.Context(), idInfo.ID)
	response.Ok(c, err, api)
}

func (u *Controller) GetSysUserList(c *gin.Context) {
	var pageInfo schema.SearchSysUser
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.Ok(c, err, nil)
		return
	}

	list, total, err := u.srv.GetSysUserList(c.Request.Context(), pageInfo)
	response.PageOk(c, err, list, total, pageInfo.Page, pageInfo.PageSize)
}
