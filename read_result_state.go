package es

import (
	"syscall"
	"time"
	"unsafe"
)

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetNumFileResults(void);
var dllEverythingGetNumFileResults *syscall.LazyProc

func EverythingGetNumFileResults() (uint32, error) {
	r1, _, err := dllEverythingGetNumFileResults.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetNumFolderResults(void);
var dllEverythingGetNumFolderResults *syscall.LazyProc

func EverythingGetNumFolderResults() (uint32, error) {
	r1, _, err := dllEverythingGetNumFolderResults.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetNumResults(void);
var dllEverythingGetNumResults *syscall.LazyProc

func EverythingGetNumResults() (uint32, error) {
	r1, _, err := dllEverythingGetNumResults.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetTotFileResults(void);
var dllEverythingGetTotFileResults *syscall.LazyProc

func EverythingGetTotFileResults() (uint32, error) {
	r1, _, err := dllEverythingGetTotFileResults.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetTotFolderResults(void);
var dllEverythingGetTotFolderResults *syscall.LazyProc

func EverythingGetTotFolderResults() (uint32, error) {
	r1, _, err := dllEverythingGetTotFolderResults.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetTotResults(void);
var dllEverythingGetTotResults *syscall.LazyProc

func EverythingGetTotResults() (uint32, error) {
	r1, _, err := dllEverythingGetTotResults.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsVolumeResult(DWORD dwIndex);
var dllEverythingIsVolumeResult *syscall.LazyProc

func EverythingIsVolumeResult(dwIndex uint32) error {
	r1, _, err := dllEverythingIsVolumeResult.Call(uintptr(dwIndex))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsFolderResult(DWORD dwIndex);
var dllEverythingIsFolderResult *syscall.LazyProc

func EverythingIsFolderResult(dwIndex uint32) error {
	r1, _, err := dllEverythingIsFolderResult.Call(uintptr(dwIndex))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsFileResult(DWORD dwIndex);
var dllEverythingIsFileResult *syscall.LazyProc

func EverythingIsFileResult(dwIndex uint32) error {
	r1, _, err := dllEverythingIsFileResult.Call(uintptr(dwIndex))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultFileNameW(DWORD dwIndex);
var dllEverythingGetResultFileNameW *syscall.LazyProc

func everythingGetResultFileNameW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultFileNameW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultFileNameA(DWORD dwIndex);
var dllEverythingGetResultFileNameA *syscall.LazyProc

func everythingGetResultFileNameA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultFileNameA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultPathW(DWORD dwIndex);
var dllEverythingGetResultPathW *syscall.LazyProc

func everythingGetResultPathW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultPathW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultPathA(DWORD dwIndex);
var dllEverythingGetResultPathA *syscall.LazyProc

func everythingGetResultPathA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultPathA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetResultFullPathNameA(DWORD dwIndex,LPSTR buf,DWORD bufSize);
// TODO 该方法获取路径包含中文时会存在问题,最好还是用 everythingGetResultFullPathNameW ,go源码只用了W相关方法
var dllEverythingGetResultFullPathNameA *syscall.LazyProc

func everythingGetResultFullPathNameA(dwIndex uint32, size ...uint32) (string, error) {
	bufSize := FullPathSize
	if len(size) > 0 {
		bufSize = size[0]
	}
	buf := make([]byte, bufSize)

	r1, _, err := dllEverythingGetResultFullPathNameA.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(bufSize))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return string(buf[:r1]), nil
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetResultFullPathNameW(DWORD dwIndex,LPWSTR wBuf,DWORD wBuf_size_in_wChars);
var dllEverythingGetResultFullPathNameW *syscall.LazyProc

func everythingGetResultFullPathNameW(dwIndex uint32, size ...uint32) (string, error) {
	bufSize := FullPathSize
	if len(size) > 0 {
		bufSize = size[0]
	}
	buf := make([]uint16, bufSize)

	r1, _, err := dllEverythingGetResultFullPathNameW.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(bufSize))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return syscall.UTF16ToString(buf), nil
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetResultListSort(void); // Everything 1.4.1
// EverythingSortNameAscending  可以跳转到枚举位置,查看所有排序类型
var dllEverythingGetResultListSort *syscall.LazyProc

func EverythingGetResultListSort() (uint32, error) {
	r1, _, err := dllEverythingGetResultListSort.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetResultListRequestFlags(void); // Everything 1.4.1
// EverythingRequestFileName  可以跳转到枚举位置,查看所有结果数据类型
var dllEverythingGetResultListRequestFlags *syscall.LazyProc

func EverythingGetResultListRequestFlags() (uint32, error) {
	r1, _, err := dllEverythingGetResultListRequestFlags.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultExtensionW(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultExtensionW *syscall.LazyProc

func everythingGetResultExtensionW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultExtensionW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultExtensionA(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultExtensionA *syscall.LazyProc

func everythingGetResultExtensionA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultExtensionA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetResultSize(DWORD dwIndex,LARGE_INTEGER *lpSize); // Everything 1.4.1
var dllEverythingGetResultSize *syscall.LazyProc

func EverythingGetResultSize(dwIndex uint32) (int64, error) {
	var lpSize int64 // 只要占8个字节就行
	r1, _, err := dllEverythingGetResultSize.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpSize)))
	if err = checkErr(r1, err); err != nil {
		return 0, err
	}
	return lpSize, nil
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetResultDateCreated(DWORD dwIndex,FILETIME *lpDateCreated); // Everything 1.4.1
var dllEverythingGetResultDateCreated *syscall.LazyProc

func EverythingGetResultDateCreated(dwIndex uint32) (time.Time, error) {
	var lpDateCreated syscall.Filetime
	r1, _, err := dllEverythingGetResultDateCreated.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpDateCreated)))
	if err = checkErr(r1, err); err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, lpDateCreated.Nanoseconds()), nil
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetResultDateModified(DWORD dwIndex,FILETIME *lpDateModified); // Everything 1.4.1
var dllEverythingGetResultDateModified *syscall.LazyProc

func EverythingGetResultDateModified(dwIndex uint32) (time.Time, error) {
	var lpDateModified syscall.Filetime
	r1, _, err := dllEverythingGetResultDateModified.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpDateModified)))
	if err = checkErr(r1, err); err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, lpDateModified.Nanoseconds()), nil
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetResultDateAccessed(DWORD dwIndex,FILETIME *lpDateAccessed); // Everything 1.4.1
var dllEverythingGetResultDateAccessed *syscall.LazyProc

func EverythingGetResultDateAccessed(dwIndex uint32) (time.Time, error) {
	var lpDateAccessed syscall.Filetime
	r1, _, err := dllEverythingGetResultDateAccessed.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpDateAccessed)))
	if err = checkErr(r1, err); err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, lpDateAccessed.Nanoseconds()), nil
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetResultAttributes(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultAttributes *syscall.LazyProc

func EverythingGetResultAttributes(dwIndex uint32) (uint32, error) {
	r1, _, err := dllEverythingGetResultAttributes.Call(uintptr(dwIndex))
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultFileListFileNameW(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultFileListFileNameW *syscall.LazyProc

func everythingGetResultFileListFileNameW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultFileListFileNameW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultFileListFileNameA(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultFileListFileNameA *syscall.LazyProc

func everythingGetResultFileListFileNameA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultFileListFileNameA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetResultRunCount(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultRunCount *syscall.LazyProc

func EverythingGetResultRunCount(dwIndex uint32) (uint32, error) {
	r1, _, err := dllEverythingGetResultRunCount.Call(uintptr(dwIndex))
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetResultDateRun(DWORD dwIndex,FILETIME *lpDateRun);
var dllEverythingGetResultDateRun *syscall.LazyProc

func EverythingGetResultDateRun(dwIndex uint32) (time.Time, error) {
	var lpDateRun syscall.Filetime
	r1, _, err := dllEverythingGetResultDateRun.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpDateRun)))
	if err = checkErr(r1, err); err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, lpDateRun.Nanoseconds()), nil
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetResultDateRecentlyChanged(DWORD dwIndex,FILETIME *lpDateRecentlyChanged);
var dllEverythingGetResultDateRecentlyChanged *syscall.LazyProc

func EverythingGetResultDateRecentlyChanged(dwIndex uint32) (time.Time, error) {
	var lpDateRecentlyChanged syscall.Filetime
	r1, _, err := dllEverythingGetResultDateRecentlyChanged.Call(uintptr(dwIndex),
		uintptr(unsafe.Pointer(&lpDateRecentlyChanged)))
	if err = checkErr(r1, err); err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, lpDateRecentlyChanged.Nanoseconds()), nil
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultHighlightedFileNameW(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultHighlightedFileNameW *syscall.LazyProc

func everythingGetResultHighlightedFileNameW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultHighlightedFileNameW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultHighlightedFileNameA(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultHighlightedFileNameA *syscall.LazyProc

func everythingGetResultHighlightedFileNameA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultHighlightedFileNameA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultHighlightedPathW(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultHighlightedPathW *syscall.LazyProc

func everythingGetResultHighlightedPathW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultHighlightedPathW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultHighlightedPathA(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultHighlightedPathA *syscall.LazyProc

func everythingGetResultHighlightedPathA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultHighlightedPathA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetResultHighlightedFullPathAndFileNameW(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultHighlightedFullPathAndFileNameW *syscall.LazyProc

func everythingGetResultHighlightedFullPathAndFileNameW(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultHighlightedFullPathAndFileNameW.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetResultHighlightedFullPathAndFileNameA(DWORD dwIndex); // Everything 1.4.1
var dllEverythingGetResultHighlightedFullPathAndFileNameA *syscall.LazyProc

func everythingGetResultHighlightedFullPathAndFileNameA(dwIndex uint32) (string, error) {
	r1, _, err := dllEverythingGetResultHighlightedFullPathAndFileNameA.Call(uintptr(dwIndex))
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}
