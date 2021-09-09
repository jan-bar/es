//go:build ASCII
// +build ASCII
package es

// 对应C里面的UNICODE宏定义未设置,这里是指定ASCII时使用
var (
	EverythingSetSearch                               = everythingSetSearchA
	EverythingGetSearch                               = everythingGetSearchA
	EverythingQuery                                   = everythingQueryA
	EverythingGetResultFileName                       = everythingGetResultFileNameA
	EverythingGetResultPath                           = everythingGetResultPathA
	EverythingGetResultFullPathName                   = everythingGetResultFullPathNameA
	EverythingGetResultExtension                      = everythingGetResultExtensionA
	EverythingGetResultFileListFileName               = everythingGetResultFileListFileNameA
	EverythingGetResultHighlightedFileName            = everythingGetResultHighlightedFileNameA
	EverythingGetResultHighlightedPath                = everythingGetResultHighlightedPathA
	EverythingGetResultHighlightedFullPathAndFileName = everythingGetResultHighlightedFullPathAndFileNameA
	EverythingGetRunCountFromFileName                 = everythingGetRunCountFromFileNameA
	EverythingSetRunCountFromFileName                 = everythingSetRunCountFromFileNameA
	EverythingIncRunCountFromFileName                 = everythingIncRunCountFromFileNameA
)
