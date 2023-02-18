# gosf

#### 介绍
顺丰接口的代码，本来想用gitee的，不知道啥原因发布的版本识别不出来，，这个待后面再去搞


#### 安装教程

```
go get github.com/wosyuyuda/gosf

client :=sf.F新建链接("QQBXX1pcUCKV", "HdOrk7Wu9uQiXnjXO6kaBBYV40emI7dN", "7551234567", V沙箱环境)
client.F下单(body)


var (
	body = `{
		"cargoDetails": [{
			"amount": 308.0,
			"count": 1.0,
			"name": "君宝牌地毯",
			"unit": "个",
			"volume": 0.0,
			"weight": 0.1
		}],
		"contactInfoList": [{
			"address": "十堰市丹江口市公园路155号",
			"city": "十堰市",
			"company": "清雅轩保健品专营店",
			"contact": "张三丰",
			"contactType": 1,
			"county": "武当山风景区",
			"mobile": "17006805888",
			"province": "湖北省"
		}, {
			"address": "湖北省襄阳市襄城区环城东路122号",
			"city": "襄阳市",
			"contact": "郭襄阳",
			"county": "襄城区",
			"contactType": 2,
			"mobile": "18963828829",
			"province": "湖北省"
		}],
		"customsInfo": {},
		"expressTypeId": 1,
		"isReturnRoutelabel": 1,
		"extraInfoList": [],
		"isOneselfPickup": 0,
		"language": "zh-CN",
		"monthlyCard": "7551234567",
		"orderId": "QIAO-2020069-005",
		"parcelQty": 1,
		"payMethod": 1,
		"totalWeight": 6
	}`
)

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
