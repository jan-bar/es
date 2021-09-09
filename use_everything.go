package es

import (
	"crypto/md5"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

/*
官方SDK地址
https://www.voidtools.com/zh-cn/support/everything/sdk/
*/
func init() {
	everythingPath, err := checkEverything()
	if err != nil {
		panic(err)
	}
	everything := syscall.NewLazyDLL(everythingPath)

	// write search state
	dllEverythingSetSearchW = everything.NewProc("Everything_SetSearchW")
	dllEverythingSetSearchA = everything.NewProc("Everything_SetSearchA")
	dllEverythingSetMatchPath = everything.NewProc("Everything_SetMatchPath")
	dllEverythingSetMatchCase = everything.NewProc("Everything_SetMatchCase")
	dllEverythingSetMatchWholeWord = everything.NewProc("Everything_SetMatchWholeWord")
	dllEverythingSetRegex = everything.NewProc("Everything_SetRegex")
	dllEverythingSetMax = everything.NewProc("Everything_SetMax")
	dllEverythingSetOffset = everything.NewProc("Everything_SetOffset")
	dllEverythingSetReplyWindow = everything.NewProc("Everything_SetReplyWindow")
	dllEverythingSetReplyID = everything.NewProc("Everything_SetReplyID")
	dllEverythingSetSort = everything.NewProc("Everything_SetSort")
	dllEverythingSetRequestFlags = everything.NewProc("Everything_SetRequestFlags")

	// read search state
	dllEverythingGetMatchPath = everything.NewProc("Everything_GetMatchPath")
	dllEverythingGetMatchCase = everything.NewProc("Everything_GetMatchCase")
	dllEverythingGetMatchWholeWord = everything.NewProc("Everything_GetMatchWholeWord")
	dllEverythingGetRegex = everything.NewProc("Everything_GetRegex")
	dllEverythingGetMax = everything.NewProc("Everything_GetMax")
	dllEverythingGetOffset = everything.NewProc("Everything_GetOffset")
	dllEverythingGetSearchA = everything.NewProc("Everything_GetSearchA")
	dllEverythingGetSearchW = everything.NewProc("Everything_GetSearchW")
	dllEverythingGetLastError = everything.NewProc("Everything_GetLastError")
	dllEverythingGetReplyWindow = everything.NewProc("Everything_GetReplyWindow")
	dllEverythingGetReplyID = everything.NewProc("Everything_GetReplyID")
	dllEverythingGetSort = everything.NewProc("Everything_GetSort")
	dllEverythingGetRequestFlags = everything.NewProc("Everything_GetRequestFlags")

	// other state
	dllEverythingQueryA = everything.NewProc("Everything_QueryA")
	dllEverythingQueryW = everything.NewProc("Everything_QueryW")
	dllEverythingIsQueryReply = everything.NewProc("Everything_IsQueryReply")
	dllEverythingSortResultsByPath = everything.NewProc("Everything_SortResultsByPath")
	dllEverythingReset = everything.NewProc("Everything_Reset")
	dllEverythingCleanUp = everything.NewProc("Everything_CleanUp")
	dllEverythingGetMajorVersion = everything.NewProc("Everything_GetMajorVersion")
	dllEverythingGetMinorVersion = everything.NewProc("Everything_GetMinorVersion")
	dllEverythingGetRevision = everything.NewProc("Everything_GetRevision")
	dllEverythingGetBuildNumber = everything.NewProc("Everything_GetBuildNumber")
	dllEverythingExit = everything.NewProc("Everything_Exit")
	dllEverythingIsDBLoaded = everything.NewProc("Everything_IsDBLoaded")
	dllEverythingIsAdmin = everything.NewProc("Everything_IsAdmin")
	dllEverythingIsAppData = everything.NewProc("Everything_IsAppData")
	dllEverythingRebuildDB = everything.NewProc("Everything_RebuildDB")
	dllEverythingUpdateAllFolderIndexes = everything.NewProc("Everything_UpdateAllFolderIndexes")
	dllEverythingSaveDB = everything.NewProc("Everything_SaveDB")
	dllEverythingSaveRunHistory = everything.NewProc("Everything_SaveRunHistory")
	dllEverythingDeleteRunHistory = everything.NewProc("Everything_DeleteRunHistory")
	dllEverythingGetTargetMachine = everything.NewProc("Everything_GetTargetMachine")
	dllEverythingGetRunCountFromFileNameW = everything.NewProc("Everything_GetRunCountFromFileNameW")
	dllEverythingGetRunCountFromFileNameA = everything.NewProc("Everything_GetRunCountFromFileNameA")
	dllEverythingSetRunCountFromFileNameW = everything.NewProc("Everything_SetRunCountFromFileNameW")
	dllEverythingSetRunCountFromFileNameA = everything.NewProc("Everything_SetRunCountFromFileNameA")
	dllEverythingIncRunCountFromFileNameW = everything.NewProc("Everything_IncRunCountFromFileNameW")
	dllEverythingIncRunCountFromFileNameA = everything.NewProc("Everything_IncRunCountFromFileNameA")

	// read result state
	dllEverythingGetNumFileResults = everything.NewProc("Everything_GetNumFileResults")
	dllEverythingGetNumFolderResults = everything.NewProc("Everything_GetNumFolderResults")
	dllEverythingGetNumResults = everything.NewProc("Everything_GetNumResults")
	dllEverythingGetTotFileResults = everything.NewProc("Everything_GetTotFileResults")
	dllEverythingGetTotFolderResults = everything.NewProc("Everything_GetTotFolderResults")
	dllEverythingGetTotResults = everything.NewProc("Everything_GetTotResults")
	dllEverythingIsVolumeResult = everything.NewProc("Everything_IsVolumeResult")
	dllEverythingIsFolderResult = everything.NewProc("Everything_IsFolderResult")
	dllEverythingIsFileResult = everything.NewProc("Everything_IsFileResult")
	dllEverythingGetResultFileNameW = everything.NewProc("Everything_GetResultFileNameW")
	dllEverythingGetResultFileNameA = everything.NewProc("Everything_GetResultFileNameA")
	dllEverythingGetResultPathW = everything.NewProc("Everything_GetResultPathW")
	dllEverythingGetResultPathA = everything.NewProc("Everything_GetResultPathA")
	dllEverythingGetResultFullPathNameA = everything.NewProc("Everything_GetResultFullPathNameA")
	dllEverythingGetResultFullPathNameW = everything.NewProc("Everything_GetResultFullPathNameW")
	dllEverythingGetResultListSort = everything.NewProc("Everything_GetResultListSort")
	dllEverythingGetResultListRequestFlags = everything.NewProc("Everything_GetResultListRequestFlags")
	dllEverythingGetResultExtensionW = everything.NewProc("Everything_GetResultExtensionW")
	dllEverythingGetResultExtensionA = everything.NewProc("Everything_GetResultExtensionA")
	dllEverythingGetResultSize = everything.NewProc("Everything_GetResultSize")
	dllEverythingGetResultDateCreated = everything.NewProc("Everything_GetResultDateCreated")
	dllEverythingGetResultDateModified = everything.NewProc("Everything_GetResultDateModified")
	dllEverythingGetResultDateAccessed = everything.NewProc("Everything_GetResultDateAccessed")
	dllEverythingGetResultAttributes = everything.NewProc("Everything_GetResultAttributes")
	dllEverythingGetResultFileListFileNameW = everything.NewProc("Everything_GetResultFileListFileNameW")
	dllEverythingGetResultFileListFileNameA = everything.NewProc("Everything_GetResultFileListFileNameA")
	dllEverythingGetResultRunCount = everything.NewProc("Everything_GetResultRunCount")
	dllEverythingGetResultDateRun = everything.NewProc("Everything_GetResultDateRun")
	dllEverythingGetResultDateRecentlyChanged = everything.NewProc("Everything_GetResultDateRecentlyChanged")
	dllEverythingGetResultHighlightedFileNameW = everything.NewProc("Everything_GetResultHighlightedFileNameW")
	dllEverythingGetResultHighlightedFileNameA = everything.NewProc("Everything_GetResultHighlightedFileNameA")
	dllEverythingGetResultHighlightedPathW = everything.NewProc("Everything_GetResultHighlightedPathW")
	dllEverythingGetResultHighlightedPathA = everything.NewProc("Everything_GetResultHighlightedPathA")
	dllEverythingGetResultHighlightedFullPathAndFileNameW = everything.NewProc("Everything_GetResultHighlightedFullPathAndFileNameW")
	dllEverythingGetResultHighlightedFullPathAndFileNameA = everything.NewProc("Everything_GetResultHighlightedFullPathAndFileNameA")
}

type EverythingErr uint8

func (e EverythingErr) Error() string {
	switch e {
	case EverythingOk:
		return "no error detected"
	case EverythingErrorMemory:
		return "out of memory."
	case EverythingErrorIpc:
		return "Everything search client is not running"
	case EverythingErrorRegisterClassEx:
		return "unable to register window class."
	case EverythingErrorCreateWindow:
		return "unable to create listening window"
	case EverythingErrorCreateThread:
		return "unable to create listening thread"
	case EverythingErrorInvalidIndex:
		return "invalid index"
	case EverythingErrorInvalidCall:
		return "invalid call"
	case EverythingErrorInvalidRequest:
		// 这种错误一般都是 EverythingSetRequestFlags 没有设置该字段要查询
		return "invalid request data, request data first."
	case EverythingErrorInvalidParameter:
		return "bad parameter."
	default:
		return "unknown exception"
	}
}

// go调用win32相关说明
// https://pkg.go.dev/github.com/0xrawsec/golang-win32/win32#pkg-index
const (
	EverythingOk                    EverythingErr = 0 // no error detected
	EverythingErrorMemory           EverythingErr = 1 // out of memory.
	EverythingErrorIpc              EverythingErr = 2 // Everything search client is not running
	EverythingErrorRegisterClassEx  EverythingErr = 3 // unable to register window class.
	EverythingErrorCreateWindow     EverythingErr = 4 // unable to create listening window
	EverythingErrorCreateThread     EverythingErr = 5 // unable to create listening thread
	EverythingErrorInvalidIndex     EverythingErr = 6 // invalid index
	EverythingErrorInvalidCall      EverythingErr = 7 // invalid call
	EverythingErrorInvalidRequest   EverythingErr = 8 // invalid request data, request data first.
	EverythingErrorInvalidParameter EverythingErr = 9 // bad parameter.

	EverythingSortNameAscending                 = 1
	EverythingSortNameDescending                = 2
	EverythingSortPathAscending                 = 3
	EverythingSortPathDescending                = 4
	EverythingSortSizeAscending                 = 5
	EverythingSortSizeDescending                = 6
	EverythingSortExtensionAscending            = 7
	EverythingSortExtensionDescending           = 8
	EverythingSortTypeNameAscending             = 9
	EverythingSortTypeNameDescending            = 10
	EverythingSortDateCreatedAscending          = 11
	EverythingSortDateCreatedDescending         = 12
	EverythingSortDateModifiedAscending         = 13
	EverythingSortDateModifiedDescending        = 14
	EverythingSortAttributesAscending           = 15
	EverythingSortAttributesDescending          = 16
	EverythingSortFileListFilenameAscending     = 17
	EverythingSortFileListFilenameDescending    = 18
	EverythingSortRunCountAscending             = 19
	EverythingSortRunCountDescending            = 20
	EverythingSortDateRecentlyChangedAscending  = 21
	EverythingSortDateRecentlyChangedDescending = 22
	EverythingSortDateAccessedAscending         = 23
	EverythingSortDateAccessedDescending        = 24
	EverythingSortDateRunAscending              = 25
	EverythingSortDateRunDescending             = 26

	EverythingRequestFileName                       = 0x00000001
	EverythingRequestPath                           = 0x00000002
	EverythingRequestFullPathAndFileName            = 0x00000004
	EverythingRequestExtension                      = 0x00000008
	EverythingRequestSize                           = 0x00000010
	EverythingRequestDateCreated                    = 0x00000020
	EverythingRequestDateModified                   = 0x00000040
	EverythingRequestDateAccessed                   = 0x00000080
	EverythingRequestAttributes                     = 0x00000100
	EverythingRequestFileListFileName               = 0x00000200
	EverythingRequestRunCount                       = 0x00000400
	EverythingRequestDateRun                        = 0x00000800
	EverythingRequestDateRecentlyChanged            = 0x00001000
	EverythingRequestHighlightedFileName            = 0x00002000
	EverythingRequestHighlightedPath                = 0x00004000
	EverythingRequestHighlightedFullPathAndFileName = 0x00008000

	EverythingTargetMachineX86 = 1
	EverythingTargetMachineX64 = 2
	EverythingTargetMachineArm = 3

	FullPathSize uint32 = 260 // 查阅资料,据说window最大路径长度为260
)

// everything.dll 不存在或者md5不对,则创建everything.dll
func checkEverything() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	everythingPath := filepath.Join(filepath.Dir(dir), "everything.dll")

	fr, err := os.Open(everythingPath)
	if err == nil {
		defer fr.Close()

		m := md5.New()
		_, err = io.Copy(m, fr)
		if err != nil {
			return "", err
		}

		same := true
		for i, v := range m.Sum(nil) {
			if everythingMd5[i] != v {
				same = false
				break
			}
		}
		if same {
			return everythingPath, nil
		}
	}

	err = os.WriteFile(everythingPath, everythingDll, 0666)
	if err != nil {
		return "", err
	}
	return everythingPath, nil
}

func checkErr(r1 uintptr, err error) error {
	if e, ok := err.(syscall.Errno); ok && e == 0 {
		if r1 == 0 {
			// 某些接口返回0时需要获取最后一次错误
			return EverythingGetLastError()
		}
		// 其他接口在调用时设置r1=1,避免走到上述逻辑
		return nil
	}
	return err
}

func uintPtrAToString(r uintptr) string {
	p := (*byte)(unsafe.Pointer(r))
	if p == nil {
		return ""
	}

	n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
	for *(*byte)(end) != 0 {
		end = unsafe.Add(end, add)
		n++
	}
	return string(unsafe.Slice(p, n))
}

func uintPtrWToString(r uintptr) string {
	p := (*uint16)(unsafe.Pointer(r))
	if p == nil {
		return ""
	}

	n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
	for *(*uint16)(end) != 0 {
		end = unsafe.Add(end, add)
		n++
	}
	return string(utf16.Decode(unsafe.Slice(p, n)))
}
