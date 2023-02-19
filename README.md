# gosf

#### 介绍
顺丰接口的代码，本来想用gitee的，不知道啥原因发布的版本识别不出来，，这个待后面再去搞


#### 安装教程

```
go get github.com/wosyuyuda/gosf

client :=sf.F新建链接("QQB********pcUCKV", "HdOrk7Wu9uQiX********BBYV40emI7dN", "7551234567", V沙箱环境)
order := sf.SFOrder{
		IsReturnRoutelabel: 1,
		CargoDetails: []sf.CargoDetail{
			{Name: "苹果手机"},
		},
		ContactInfoList: []sf.ContactInfo{
			{ContactType: sf.V寄件人, Address: "十堰市丹江口市公园路155号", City: "十堰市", Contact: "张三丰",
				County: "武当山风景区", Mobile: "17888888888", Province: "湖北省"},
			{ContactType: sf.V收件人, Address: "湖北省襄阳市襄城区环城东路122号", City: "襄阳市", Contact: "郭襄阳",
				County: "襄城区", Mobile: "18963828829", Province: "湖北省"},
		},
	}

	//json.Unmarshal([]byte(body), &order)
	res, err := client.F下单(&order)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("res:", res)

	其它方法请直接使用
	client.F发送请求(服务code,请求数据)
	请求数据建议为str或者map结构，服务code请自行查看

```
下个版权将添加相应的结构体，优化使用功能

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
