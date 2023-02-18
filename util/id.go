package util

import (
	"fmt"
	"sync"
	"time"
)

const (
	idBits  uint8 = 2  // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点
	numBits uint8 = 10 // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位数 即每毫秒可生成 2^12-1=4096个唯一ID
	// 这里求最大值使用了位运算，-1 的二进制表示为 1 的补码，感兴趣的同学可以自己算算试试 -1 ^ (-1 << nodeBits) 这里是不是等于 1023
	idMax       int64 = -1 ^ (-1 << idBits)  // 节点ID的最大值，用于防止溢出
	numMax      int64 = -1 ^ (-1 << numBits) // 同上，用来表示生成id序号的最大值
	idtimeShift uint8 = idBits + numBits     // 时间戳向左的偏移量
	idShift     uint8 = numBits              // 节点ID向左的偏移量
	starttiime  int64 = 1629610098000        // 这个是我在写epoch这个变量时的时间戳(毫秒)，可以设置为当前时间
)

var (
	Idw *Id
)

// 定义一个woker工作节点所需要的基本参数
type Id struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录时间戳
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
	workerId  int64      // 该节点的ID
}

func init() {
	Idw = &Id{}
}

func (w *Id) F获取字符串id() (sid string) {
	id := w.F获取id()
	sid = fmt.Sprintf("%v", id)
	return
}

//不传值默认0产品,1订单，2账单，3店铺，4用户
func (w *Id) F获取id() int64 {
	//var ids []int64
	var num, timestamp int64
	w.mu.Lock()
	defer w.mu.Unlock() // 生成完成后记得 解锁 解锁 解锁

	num = w.number
	timestamp = w.timestamp

	now := time.Now().UnixNano() / 1e6 // 纳秒转10毫秒
	if timestamp == now {
		num++
		// 这里要判断，当前工作节点是否在1毫秒内已经生成numberMax个ID
		if num > numMax {
			// 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
			for now <= timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else if now > timestamp { //如果当前时间大于上次时间,初始化id
		// 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		num = 0
		w.timestamp = now
	} else { //如果时间回拨了,那么就用最大的机器id数-1.2<<8-1大概255,不管怎样机器255台机器部署应该是极限了吧
		w.workerId = int64(1<<idBits - 1)
	}
	// 第一段 now - epoch 为该算法目前已经奔跑了xxx毫秒
	// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	return int64((now-starttiime)<<idtimeShift | (w.workerId << idShift) | num)
}
