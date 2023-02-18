package model

type SFOrder struct {
	//XMLName           xml.Name `json:"Order"`
	OrderId            string        `json:"orderId"`            //订单id
	Language           string        `json:"language"`           //语言
	MonthlyCard        string        `json:"monthlyCard"`        //月结账号
	ParcelQty          int           `json:"parcelQty"`          //包裹数，一个包裹对应一个运单号；若包裹数大于1，则返回一个母运单号和N-1个子运单号
	PayMethod          int           `json:"payMethod"`          //付款方式，支持以下值： 1:寄方付 2:收方付 3:第三方付
	TotalLength        int           `json:"totalLength"`        //物品长,非必填
	TotalWeight        int           `json:"totalWeight"`        //物品宽,非必填
	TotalHeight        int           `json:"totalHeight"`        //物品高,非必填
	IsOneselfPickup    int           `json:"isOneselfPickup"`    //快件自取，支持以下值： 1：客户同意快件自取 0：客户不同意快件自取
	ExpressTypeId      int           `json:"expressTypeId"`      //快件产品类别， 支持附录 《快件产品类别表》 的产品编码值，仅可使用与顺丰销售约定的快件产品
	IsReturnRoutelabel int           `json:"isReturnRoutelabel"` //是否返回路由标签： 默认1， 1：返回路由标签， 0：不返回；除部分特殊用户外，其余用户都默认返回
	ContactInfoList    []ContactInfo `json:"contactInfoList"`    //收件人信息,要两个
	CargoDetails       []CargoDetail `json:"CargoDetails"`       //物品信息
	CargoDesc          string        `json:"cargoDesc"`          //托寄物品描述，非必填
}

type ContactInfo struct {
	Address     string `json:"address"`     //地址
	Province    string `json:"province"`    //省
	City        string `json:"city"`        //市
	County      string `json:"county"`      //区
	Mobile      string `json:"mobile"`      //手机
	Contact     string `json:"contact"`     //人物
	ContactType int    `json:"contactType"` //地址类型： 1，寄件方信息 2，到件方信息
}

//托寄物品信息
type CargoDetail struct {
	Name string `json:"name"` //市
}

type Response struct {
	ApiErrorMsg   string `json:"apiErrorMsg"` //返回提示
	ApiResponseID string `json:"apiResponseID"`
	ApiResultCode string `json:"apiResultCode"`
	ApiResultData string `json:"apiResultData"`
}

type ResultData struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

type OrderResponse struct {
	ResultData
	MsgData interface{} `json:"msgData"`
}
