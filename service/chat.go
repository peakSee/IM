package service

import (
	"IM/helper"
	"IM/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func ChatList(c *gin.Context) {
	roomIdentity := c.Query("room_identity")
	if roomIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "房间号不能为空",
		})
		return
	}
	//判断用户是否属于该房间
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	_, err := models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, roomIdentity)
	if err != nil {
		log.Printf("[ChatList ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "非法访问",
		})
		return
	}
	pageIndex, _ := strconv.ParseInt(c.Query("page_index"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 32)
	skip := (pageIndex - 1) * pageSize
	//聊天记录查询
	data, err := models.GetMessageListByRoomIdentity(roomIdentity, &pageSize, &skip)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "数据加载成功",
		"data": gin.H{
			"list": data,
		},
	})
}
