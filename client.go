package sf

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/wosyuyuda/gosf/util"
)

type Client struct {
	V用户id    string //用户id
	V检验码     string //校验码
	V月结账号    string
	V是否为沙箱环境 bool
}

const (
	V沙箱环境  = 1
	V正式环境  = 0
	顺丰正式环境 = "https://sfapi.sf-express.com"
	顺丰沙箱环境 = "https://sfapi-sbox.sf-express.com"
	下单     = "/std/service"

	V下单服务        = "EXP_RECE_CREATE_ORDER"
	V查询服务        = "EXP_RECE_SEARCH_ORDER_RESP"
	V订单取消接口      = "EXP_RECE_UPDATE_ORDER"
	V路由查询接口      = "EXP_RECE_SEARCH_ROUTES"
	V子单号申请接口     = "EXP_RECE_GET_SUB_MAILNO"
	V清单运费查询接口    = "EXP_RECE_QUERY_SFWAYBILL"
	V仓配退货下单接口    = "EXP_RECE_CREATE_REVERSE_ORDER"
	V仓配退货消单接口    = "EXP_RECE_CANCEL_REVERSE_ORDER"
	V派件通知接口      = "EXP_RECE_DELIVERY_NOTICE"
	V截单转寄退回接口    = "EXP_RECE_WANTED_INTERCEPT"
	V换货下单接口      = "EXP_RECE_CREATE_EXCHANGE_ORDER"
	V预下单接口       = "EXP_RECE_PRE_ORDER"
	V顺丰会员地址簿查询接口 = "COM_RECE_QUERY_ADDRESS_BOOK"
)

var (
	支持的服务 = []string{
		V下单服务, V查询服务, V订单取消接口, V路由查询接口, V子单号申请接口,
		V清单运费查询接口, V仓配退货下单接口, V仓配退货消单接口,
		V派件通知接口, V截单转寄退回接口, V换货下单接口, V预下单接口,
		V顺丰会员地址簿查询接口,
	}
)

func F新建链接(用户id, 检验码, 月结账号 string, 沙箱环境 ...int) *Client {
	为沙箱 := len(沙箱环境) == 1 && 沙箱环境[0] == V沙箱环境
	client := &Client{
		V用户id:    用户id,
		V检验码:     检验码,
		V月结账号:    月结账号,
		V是否为沙箱环境: 为沙箱,
	}
	return client
}

//下单前有些参数要检验一下最好。。。
func (a *Client) F下单(body string) (by []byte, err error) {

	data := map[string]interface{}{}
	json.Unmarshal([]byte(body), &data)

	by, err = a.F发送请求(V下单服务, data)
	if err != nil {
		return
	}
	//fmt.Println("数据是:", string(by))
	return
}

//body为结构体
func (a *Client) F发送请求(serverid string, body interface{}) (by []byte, err error) {
	是否包含 := false
	for _, v := range 支持的服务 {
		if v == serverid {
			是否包含 = true
		}
	}
	if !是否包含 {
		err = errors.New("暂未包含此服务")
		return
	}
	myurl := 下单

	myurl += "?partnerID=" + a.V用户id
	myurl += "&requestID=" + util.Idw.F获取字符串id()
	myurl += "&serviceCode=" + serverid
	//获取时间戳并拼接
	ti := fmt.Sprintf("%v", time.Now().Unix()*1000)
	myurl += "&timestamp=" + ti
	//把结构体转成json字符串，再生成签名
	by, _ = json.Marshal(body)
	msgdata := string(by)

	//fmt.Println(md5s)
	myurl += "&msgDigest=" + a.F生成签名(msgdata+ti+a.V检验码)

	myurl += "&msgData=" + msgdata

	//return a.发送(myurl, "")
	环境地址 := 顺丰正式环境
	if a.V是否为沙箱环境 {
		环境地址 = 顺丰沙箱环境
	}
	by, err = util.F发送POST请求(环境地址+myurl, "")
	return
}

func (a *Client) F生成签名(str string) string {
	str = url.QueryEscape(str)
	md5Key := md5.New()
	md5Key.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(md5Key.Sum(nil))
}
