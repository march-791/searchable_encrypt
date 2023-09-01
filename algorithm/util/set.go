/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) Beijing Advanced Innovation Center for Future Blockchain and Privacy Computing (未来区块链与隐私计算高精尖创新中心).
All rights reserved.

SPDX-License-Identifier: Apache-2.0

*/

package util

//定义1个set结构体 内部主要是使用了map
type set struct {
	elements map[interface{}]bool
}

//NewSet interface{}表示任意类型
//NewSet ...表示任意参数个数
//NewSet 初始化一个set 并加入给定的参数
func NewSet(items ...interface{}) Set {
	st := set{
		elements: make(map[interface{}]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

//Set 定义一个Set接口
type Set interface {
	Add(item interface{}) bool //添加一个元素
	Delete(item interface{})   //删除给定元素
	Len() int                  //返回set大小
	GetItems() []interface{}   //返回set中的元素
	In(item interface{}) bool  //判断元素item是否在set中

}

//注意：下面是set指针类型实现了Set接口  而不是set
/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

func (st *set) Add(item interface{}) bool {
	if st.elements[item] {
		return false
	}
	st.elements[item] = true
	return true
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

func (st *set) Delete(item interface{}) {
	delete(st.elements, item) //调用的是map的删除方法
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

func (st *set) Len() int {
	return len(st.elements)
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

func (st *set) GetItems() []interface{} {
	keys := make([]interface{}, 0, len(st.elements)) //创建一个初始大小为0 容量为sz的切片
	for key := range st.elements {
		keys = append(keys, key)
	}
	return keys //返回切片
}

/**
 * @Author peng
 * @Description
 * @Date 16:31 2022/6/28
 * @Param
 * @return
 **/

func (st *set) In(item interface{}) bool {
	if _, in := st.elements[item]; in { //使用了map中的判断一个元素是否存在的方法 in为true则存在
		return true
	}
	return false
}
