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
