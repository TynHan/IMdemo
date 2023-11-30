package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Get User List
//
//	@Tags		User Service
//	@Success	200	{string}	json{"code", "message"}
//	@Router		/user/getlist [get]
func GetUserList(ctx *gin.Context) {
	data := models.GetUserList()
	// time.Sleep(time.Second * 2)
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	},
}

func SendMsg(c *gin.Context) {
	fmt.Println("Arrive send msg")
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error found in seding msg, ", err.Error())
		return
	}
	defer func(ws *websocket.Conn) {
		err1 := ws.Close()
		if err1 != nil {
			fmt.Println("Error found in closing ws, ", err.Error())
		}
	}(ws)
	MsgHandler(ws, c)

}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
    for{
        fmt.Println("Arrive MsgHandler")
        msg, err := utils.Subscribe(c, utils.PublicKey)
        if err != nil {
            fmt.Println("Error found in handling msg, ", err.Error())
            return
        }
        fmt.Println("Arrive MsgHandler2")
        t := time.Now().Format("2006-01-02 15:04:05")
        fmt.Println("Arrive MsgHandler3")
        m := fmt.Sprintf("[ws][%s]:%s", t, msg)
        fmt.Println("Arrive MsgHandler4")
        fmt.Println("msg get:", m)
        err = ws.WriteMessage(1, []byte(m))
        if err != nil {
            fmt.Println("Error found in writing msg to websocket, ", err.Error())
            return
        }
    }
}
