//go:build !ASCII
// +build !ASCII
package es

// 对应C里面的UNICODE宏定义启用,这里是默认使用
var (
	EverythingSetSearch                               = everythingSetSearchW
	EverythingGetSearch                               = everythingGetSearchW
	EverythingQuery                                   = everythingQueryW
	EverythingGetResultFileName                       = everythingGetResultFileNameW
	EverythingGetResultPath                           = everythingGetResultPathW
	EverythingGetResultFullPathName                   = everythingGetResultFullPathNameW
	EverythingGetResultExtension                      = everythingGetResultExtensionW
	EverythingGetResultFileListFileName               = everythingGetResultFileListFileNameW
	EverythingGetResultHighlightedFileName            = everythingGetResultHighlightedFileNameW
	EverythingGetResultHighlightedPath                = everythingGetResultHighlightedPathW
	EverythingGetResultHighlightedFullPathAndFileName = everythingGetResultHighlightedFullPathAndFileNameW
	EverythingGetRunCountFromFileName                 = everythingGetRunCountFromFileNameW
	EverythingSetRunCountFromFileName                 = everythingSetRunCountFromFileNameW
	EverythingIncRunCountFromFileName                 = everythingIncRunCountFromFileNameW
)
