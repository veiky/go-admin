package sms

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/page"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/chart_legend"
	"github.com/GoAdminGroup/themes/adminlte/components/description"
	"github.com/GoAdminGroup/themes/adminlte/components/infobox"
	"github.com/GoAdminGroup/themes/adminlte/components/productlist"
	"github.com/GoAdminGroup/themes/adminlte/components/progress_group"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"html/template"
)

func (e *SMS)setConf(ctx *context.Context) {
	accessKeyId := ctx.Get()
	accessKeySecret := ctx.Get()
	phoneNumbers := ctx.Get()
	signName := ctx.Get()
	templateCode := ctx.Get()
	templateParam := ctx.Get()
	conn := e.Base.Conn
	_, err := conn.Exec("INSERT INTO sms_config() VALUE(?)", accessKeyId, accessKeySecret)
	if err != nil {
		// todo
	}
	// response client

	// build service
	smsClient := aliyunsmsclient.New(gatewayUrl)
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	fmt.Println("Got raw response from server:", string(result.RawResponse))
	if err != nil {
		panic("Failed to send Message: " + err.Error())
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	if result.IsSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
	} else {
		fmt.Println("Failed to send a SMS:", resultJson)
	}
}

func (e *SMS)confList(ctx *context.Context) {
	conn := e.Base.Conn
	re, err := conn.Query("SELECT distinct service FROM sms_config")
	if err != nil {
		// todo
	}
	// response client
}

func (e *SMS)getConf(ctx *context.Context) {
	id := ctx.Get()
	conn := e.Base.Conn
	re, err := conn.Query("SELECT distinct service FROM sms_config WHERE id=?", id)
	if err != nil {
		// todo
	}
	// response client
}
