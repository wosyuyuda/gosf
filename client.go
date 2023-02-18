package gosf

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
	下单服务   = "EXP_RECE_CREATE_ORDER"
	查询服务   = "EXP_RECE_SEARCH_ORDER_RESP"
	下单     = "/std/service"
)

var (
	支持的服务 = []string{
		下单服务, 查询服务,
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
func (a *Client) F下单(body string) (err error) {

	data := map[string]interface{}{}
	json.Unmarshal([]byte(body), &data)

	by, err := a.F发送请求(下单服务, data)
	if err != nil {
		return
	}
	fmt.Println("数据是:", string(by))
	return
}

//body为结构体
func (a *Client) F发送请求(serverid string, body interface{}) (by []byte, err error) {
	switch serverid {
	case 下单服务, 查询服务:
	default:
		err = errors.New("未在服务列表")
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
	md5s := a.F生成签名(msgdata + ti + a.V检验码)
	//fmt.Println(md5s)
	myurl += "&msgDigest=" + md5s

	myurl += "&msgData=" + string(by)

	//return a.发送(myurl, "")
	url1 := 顺丰正式环境
	if a.V是否为沙箱环境 {
		url1 = 顺丰沙箱环境
	}
	by, err = util.F发送POST请求(url1+myurl, "")
	return
}

func (a *Client) F生成签名(str string) string {
	str = url.QueryEscape(str)
	md5Key := md5.New()
	md5Key.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(md5Key.Sum(nil))
}
