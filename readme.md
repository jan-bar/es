## 介绍
[官方SDK地址](https://www.voidtools.com/zh-cn/support/everything/sdk/)

1. 本项目会将官方dll编译到可执行程序中,运行时无需考虑dll问题。
2. 根据官方介绍,使用SDK前需要运行`everything`程序。
3. 执行`go build -tag ASCII`时编译`ascii`相关接口,否则编译`unicode`接口。
4. 可以参考`examples/search.go`示例代码。
5. 本项目主要是我用来学习window下go调用dll的方法,包含传参和返回值处理。
6. 相关代码大部分都是研究go源码得到的启发。
7. 特别注意,我的代码里面有几个使用了`go1.17`才有的特性。


下面是示例代码：
```go
package main

import (
	"fmt"
	"os"

	"github.com/jan-bar/es"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage:%s test.txt\n", os.Args[0])
		return
	}

	err := es.EverythingSetSearch(os.Args[1])
	if err != nil {
		panic(err)
	}
	err = es.EverythingSetMax(5)
	if err != nil {
		panic(err)
	}

	// 设置好需要查询的内容,不然后续遍历时可能报错
	err = es.EverythingSetRequestFlags(es.EverythingRequestFileName | es.EverythingRequestPath |
		es.EverythingRequestDateCreated | es.EverythingRequestDateModified | es.EverythingRequestDateAccessed |
		es.EverythingRequestSize)
	if err != nil {
		panic(err)
	}

	// 定好排序规则
	err = es.EverythingSetSort(es.EverythingSortDateModifiedAscending)
	if err != nil {
		panic(err)
	}

	// 开始查询
	fmt.Println("EverythingQuery:", es.EverythingQuery(true))

	// 得到查询结果个数
	num, err := es.EverythingGetNumResults()
	fmt.Println("EverythingGetNumResults", num, err)

	for i := uint32(0); i < num; i++ {
		fmt.Println("---------------------------------------------------")
		s, err := es.EverythingGetResultSize(i)
		fmt.Printf("FileSize:[%d],%v\n", s, err)
		p, err := es.EverythingGetResultFullPathName(i)
		fmt.Printf("FullPathName:[%s],%v\n", p, err)
		p, err = es.EverythingGetResultFileName(i)
		fmt.Printf("FileName:[%s],%v\n", p, err)
		t, err := es.EverythingGetResultDateCreated(i)
		fmt.Printf("DateCreated:[%s],%v\n", t, err)
		t, err = es.EverythingGetResultDateModified(i)
		fmt.Printf("DateModified:[%s],%v\n", t, err)
		t, err = es.EverythingGetResultDateAccessed(i)
		fmt.Printf("DateAccessed:[%s],%v\n", t, err)
	}
}
```