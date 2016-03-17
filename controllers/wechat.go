package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
	"log"
	"net/http"
)

//
type WechatController struct {
	beego.Controller
}

var appId = beego.AppConfig.String("wxappid")
var appSecret = beego.AppConfig.String("wxappsecret")

func (this *WechatController) Get() {
	beego.BConfig.WebConfig.AutoRender = false
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	token := beego.AppConfig.String("wxtoken")

	sign := util.Sign(token, timestamp, nonce)
	echostr := this.GetString("echostr")
	if signature == sign {
		this.Ctx.WriteString(echostr)
	}
}
func (this *WechatController) Post() {

	beego.BConfig.WebConfig.AutoRender = false

	messageServeMux := mp.NewMessageServeMux()

	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler)
	//fmt.Println(request.MsgTypeText)

	mpServer := mp.NewDefaultServer(beego.AppConfig.String("wxorgId"), beego.AppConfig.String("wxtoken"), appId, nil, messageServeMux)

	mp.NewServerFrontend(mpServer, mp.ErrorHandlerFunc(ErrorHandler), nil)

}

func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	fmt.Println("我是问不能")
	text := request.GetText(r.MixedMsg)

	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	mp.WriteRawResponse(w, r, resp)
}
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}
