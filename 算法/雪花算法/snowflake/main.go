package main

import (
	"fmt"
	"sync"
	"time"
)

type SnowFlake struct {
	//-------------锁-------------------//
	//互斥锁
	Mutex sync.Mutex
	//----------外部输入的值-------------//
	MachineId int64
	//-----------每部分位移的位移---------//
	//时间戳的位移
	TimeStampShift int64
	//机器id的位移
	MachineIdShift int64
	//-----------每部分的最大值----------//
	TimeStampMax int64
	OrderMax     int64
	//-----------内部值------------------//
	//序列
	order int64
	//上次的时间错
	LastTimeStamp int64
}

//构造方法
func getSnowHost(MachineId int64, TimeStampBit int64, MachineIdBit int64) (*SnowFlake, error) {
	//----------位移-------------------------------------//
	//时间戳位移
	TimeStampShift := 63 - TimeStampBit
	//机器id位移
	MachineIdShift := TimeStampShift - MachineIdBit
	if TimeStampShift < 0 || MachineIdShift < 0 {
		return nil, fmt.Errorf("bit位出错")
	}
	//--------------最值---------------------------------//
	//时间戳最大值
	TimeStampMax := int64(1<<TimeStampBit) - 1
	//机器Id最大值
	MachineIdMax := int64(1<<MachineIdBit) - 1
	//序列最大值
	OrderMax := int64(1<<(63-TimeStampBit-MachineIdBit)) - 1

	if MachineId > MachineIdMax {
		return nil, fmt.Errorf("机器id超出范围")
	}

	return &SnowFlake{MachineId: MachineId, TimeStampShift: TimeStampShift, MachineIdShift: MachineIdShift, TimeStampMax: TimeStampMax, OrderMax: OrderMax}, nil
}

func (s *SnowFlake) getUuid() (int64, error) {
	//加锁
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	//获取时间戳
	timeStamp := time.Now().UnixMilli()
	//校验时间戳范围
	if timeStamp > s.TimeStampMax {
		return 0, fmt.Errorf("时间戳超出范围")
	}
	//时间戳不同情况下的处理方式
	if timeStamp < s.LastTimeStamp {
		return 0, fmt.Errorf("时间回拨")
	} else if timeStamp == s.LastTimeStamp {
		s.order++
	} else {
		s.order = 0
	}
	//暂存时间戳
	s.LastTimeStamp = timeStamp
	//获取uuid并返回
	return timeStamp<<s.TimeStampShift | s.MachineId<<s.MachineIdShift | s.order, nil
}

func main() {
	obj, err := getSnowHost(1, 41, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj)
	fmt.Println(obj.getUuid())
}