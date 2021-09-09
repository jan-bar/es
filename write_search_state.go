package es

import (
	"syscall"
	"unsafe"
)

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetSearchW(LPCWSTR lpString);
var dllEverythingSetSearchW *syscall.LazyProc

func everythingSetSearchW(lpString string) error {
	sp, err := syscall.UTF16PtrFromString(lpString)
	if err != nil {
		return err
	}
	_, _, err = dllEverythingSetSearchW.Call(uintptr(unsafe.Pointer(sp)))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetSearchA(LPCSTR lpString);
var dllEverythingSetSearchA *syscall.LazyProc

func everythingSetSearchA(lpString string) error {
	sp, err := syscall.BytePtrFromString(lpString)
	if err != nil {
		return err
	}
	_, _, err = dllEverythingSetSearchA.Call(uintptr(unsafe.Pointer(sp)))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetMatchPath(BOOL bEnable);
var dllEverythingSetMatchPath *syscall.LazyProc

func EverythingSetMatchPath(bEnable bool) error {
	b := 0
	if bEnable {
		b = 1
	}
	_, _, err := dllEverythingSetMatchPath.Call(uintptr(b))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetMatchCase(BOOL bEnable);
var dllEverythingSetMatchCase *syscall.LazyProc

func EverythingSetMatchCase(bEnable bool) error {
	b := 0
	if bEnable {
		b = 1
	}
	_, _, err := dllEverythingSetMatchCase.Call(uintptr(b))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetMatchWholeWord(BOOL bEnable);
var dllEverythingSetMatchWholeWord *syscall.LazyProc

func EverythingSetMatchWholeWord(bEnable bool) error {
	b := 0
	if bEnable {
		b = 1
	}
	_, _, err := dllEverythingSetMatchWholeWord.Call(uintptr(b))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetRegex(BOOL bEnable);
var dllEverythingSetRegex *syscall.LazyProc

func EverythingSetRegex(bEnable bool) error {
	b := 0
	if bEnable {
		b = 1
	}
	_, _, err := dllEverythingSetRegex.Call(uintptr(b))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetMax(DWORD dwMax);
var dllEverythingSetMax *syscall.LazyProc

func EverythingSetMax(dwMax uint32) error {
	_, _, err := dllEverythingSetMax.Call(uintptr(dwMax))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetOffset(DWORD dwOffset);
var dllEverythingSetOffset *syscall.LazyProc

func EverythingSetOffset(dwOffset uint32) error {
	_, _, err := dllEverythingSetOffset.Call(uintptr(dwOffset))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetReplyWindow(HWND hWnd);
var dllEverythingSetReplyWindow *syscall.LazyProc

func EverythingSetReplyWindow(hWnd syscall.Handle) error {
	_, _, err := dllEverythingSetReplyWindow.Call(uintptr(hWnd))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetReplyID(DWORD dwId);
var dllEverythingSetReplyID *syscall.LazyProc

func EverythingSetReplyID(dwId uint32) error {
	_, _, err := dllEverythingSetReplyID.Call(uintptr(dwId))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetSort(DWORD dwSort); // Everything 1.4.1
// EverythingSortNameAscending  可以跳转到枚举位置,查看所有排序类型
var dllEverythingSetSort *syscall.LazyProc

func EverythingSetSort(dwSort uint32) error {
	_, _, err := dllEverythingSetSort.Call(uintptr(dwSort))
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SetRequestFlags(DWORD dwRequestFlags); // Everything 1.4.1
// EverythingRequestFileName  可以跳转到枚举位置,查看所有结果数据类型
var dllEverythingSetRequestFlags *syscall.LazyProc

func EverythingSetRequestFlags(dwRequestFlags uint32) error {
	_, _, err := dllEverythingSetRequestFlags.Call(uintptr(dwRequestFlags))
	return checkErr(1, err)
}
