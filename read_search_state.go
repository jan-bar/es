package es

import (
	"syscall"
)

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetMatchPath(void);
var dllEverythingGetMatchPath *syscall.LazyProc

func EverythingGetMatchPath() (bool, error) {
	r1, _, err := dllEverythingGetMatchPath.Call()
	return r1 == 1, checkErr(1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetMatchCase(void);
var dllEverythingGetMatchCase *syscall.LazyProc

func EverythingGetMatchCase() (bool, error) {
	r1, _, err := dllEverythingGetMatchCase.Call()
	return r1 == 1, checkErr(1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetMatchWholeWord(void);
var dllEverythingGetMatchWholeWord *syscall.LazyProc

func EverythingGetMatchWholeWord() (bool, error) {
	r1, _, err := dllEverythingGetMatchWholeWord.Call()
	return r1 == 1, checkErr(1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_GetRegex(void);
var dllEverythingGetRegex *syscall.LazyProc

func EverythingGetRegex() (bool, error) {
	r1, _, err := dllEverythingGetRegex.Call()
	return r1 == 1, checkErr(1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetMax(void);
var dllEverythingGetMax *syscall.LazyProc

func EverythingGetMax() (uint32, error) {
	r1, _, err := dllEverythingGetMax.Call()
	return uint32(r1), checkErr(1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetOffset(void);
var dllEverythingGetOffset *syscall.LazyProc

func EverythingGetOffset() (uint32, error) {
	r1, _, err := dllEverythingGetOffset.Call()
	return uint32(r1), checkErr(1, err)
}

// EVERYTHINGUSERAPI LPCSTR EVERYTHINGAPI Everything_GetSearchA(void);
var dllEverythingGetSearchA *syscall.LazyProc

func everythingGetSearchA() (string, error) {
	r1, _, err := dllEverythingGetSearchA.Call()
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrAToString(r1), nil
}

// EVERYTHINGUSERAPI LPCWSTR EVERYTHINGAPI Everything_GetSearchW(void);
var dllEverythingGetSearchW *syscall.LazyProc

func everythingGetSearchW() (string, error) {
	r1, _, err := dllEverythingGetSearchW.Call()
	if err = checkErr(r1, err); err != nil {
		return "", err
	}
	return uintPtrWToString(r1), nil
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetLastError(void);
// EverythingOk  跳转查看所有错误码
var dllEverythingGetLastError *syscall.LazyProc

func EverythingGetLastError() error {
	r1, _, err := dllEverythingGetLastError.Call()
	if e, ok := err.(syscall.Errno); ok && e == 0 {
		return EverythingErr(r1) // 返回得到的具体错误
	}
	return err
}

// EVERYTHINGUSERAPI HWND EVERYTHINGAPI Everything_GetReplyWindow(void);
var dllEverythingGetReplyWindow *syscall.LazyProc

func EverythingGetReplyWindow() (syscall.Handle, error) {
	r1, _, err := dllEverythingGetReplyWindow.Call()
	return syscall.Handle(r1), checkErr(1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetReplyID(void);
var dllEverythingGetReplyID *syscall.LazyProc

func EverythingGetReplyID() (uint32, error) {
	r1, _, err := dllEverythingGetReplyID.Call()
	return uint32(r1), checkErr(1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetSort(void); // Everything 1.4.1
// EverythingSortNameAscending  可以跳转到枚举位置,查看所有排序类型
var dllEverythingGetSort *syscall.LazyProc

func EverythingGetSort() (uint32, error) {
	r1, _, err := dllEverythingGetSort.Call()
	return uint32(r1), checkErr(1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetRequestFlags(void); // Everything 1.4.1
// EverythingRequestFileName  可以跳转到枚举位置,查看所有结果数据类型
var dllEverythingGetRequestFlags *syscall.LazyProc

func EverythingGetRequestFlags() (uint32, error) {
	r1, _, err := dllEverythingGetRequestFlags.Call()
	return uint32(r1), checkErr(1, err)
}
