package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/wechat"
)

//AppConf 应用程序全局配置
type AppConf struct {
	AppID          string `json:"appid" valid:"ascii,required"`
	AppSecret      string `json:"secret" valid:"ascii,required"`
	Token          string `json:"token" valid:"ascii"`
	EncodingAESKey string `json:"aes-key" valid:"ascii"`
	PayMchID       string `json:"mchid" valid:"ascii"` //支付 - 商户 ID
	ServeURL       string `json:"serve-url" valid:"ascii,required"`
	PayNotifyURL   string `json:"pay-notify-url" valid:"ascii"` //支付 - 接受微信支付结果通知的接口地址
	PayKey         string `json:"pay-key" valid:"ascii"`        //支付 - 商户后台设置的支付 key
}

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func bindConf(app *hydra.MicroApp) {
	app.Binder.API.SetMainConf(`{"address":":9999"}`)
	app.Binder.API.SetSubConf("app", `{
		"appid": "wx9e02ddcc88e13fd4",
		"secret": "6acb2bf99177524beba3d97d54df2de5",
		"token":"oTSvVuXdjb9Xx1FPi6bz",
		"aes-key": "D3mgxDexQDuqHm1MIWsyvhLMd3Y303cmf05JgjD9ZWY",
		"serve-url": "/"
	}`)
}

//bind 检查并缓存配置文件，初始化微信服务器用于接收微信通知
func bind(r *hydra.MicroApp) {
	bindConf(r)
	r.Initializing(func(c component.IContainer) error {

		//获取服务器配置
		var config AppConf
		if err := c.GetAppConf(&config); err != nil {
			return err
		}
		if b, err := govalidator.ValidateStruct(&config); !b {
			err = fmt.Errorf("app 配置文件有误:%v", err)
			return err
		}

		//创建微信处理服务
		ctx := &wechat.WContext{
			AppID:          config.AppID,
			AppSecret:      config.AppSecret,
			Token:          config.Token,
			EncodingAESKey: config.EncodingAESKey,
			PayMchID:       config.PayMchID,
			PayNotifyURL:   config.PayNotifyURL,
			PayKey:         config.PayKey,
		}
		r.Micro(config.ServeURL, wechat.NewSeverHandler(ctx, recvMessage))
		return nil
	})
}
