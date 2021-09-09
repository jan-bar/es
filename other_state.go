package es

import (
	"syscall"
	"unsafe"
)

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_QueryA(BOOL bWait);
var dllEverythingQueryA *syscall.LazyProc

func everythingQueryA(bWait bool) error {
	b := 0
	if bWait {
		b = 1
	}
	r1, _, err := dllEverythingQueryA.Call(uintptr(b))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_QueryW(BOOL bWait);
var dllEverythingQueryW *syscall.LazyProc

func everythingQueryW(bWait bool) error {
	b := 0
	if bWait {
		b = 1
	}
	r1, _, err := dllEverythingQueryW.Call(uintptr(b))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsQueryReply(UINT message,WPARAM wParam, LPARAM lParam,DWORD dwId);
var dllEverythingIsQueryReply *syscall.LazyProc

func EverythingIsQueryReply(message uint, wParam, lParam uintptr, dwId uint32) error {
	r1, _, err := dllEverythingIsQueryReply.Call(uintptr(message),
		wParam, lParam, uintptr(dwId))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_SortResultsByPath(void);
var dllEverythingSortResultsByPath *syscall.LazyProc

func EverythingSortResultsByPath() error {
	_, _, err := dllEverythingSortResultsByPath.Call()
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_Reset(void);
var dllEverythingReset *syscall.LazyProc

func EverythingReset() error {
	_, _, err := dllEverythingReset.Call()
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI void EVERYTHINGAPI Everything_CleanUp(void);
var dllEverythingCleanUp *syscall.LazyProc

func EverythingCleanUp() error {
	_, _, err := dllEverythingCleanUp.Call()
	return checkErr(1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetMajorVersion(void);
var dllEverythingGetMajorVersion *syscall.LazyProc

func EverythingGetMajorVersion() (uint32, error) {
	r1, _, err := dllEverythingGetMajorVersion.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetMinorVersion(void);
var dllEverythingGetMinorVersion *syscall.LazyProc

func EverythingGetMinorVersion() (uint32, error) {
	r1, _, err := dllEverythingGetMinorVersion.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetRevision(void);
var dllEverythingGetRevision *syscall.LazyProc

func EverythingGetRevision() (uint32, error) {
	r1, _, err := dllEverythingGetRevision.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetBuildNumber(void);
var dllEverythingGetBuildNumber *syscall.LazyProc

func EverythingGetBuildNumber() (uint32, error) {
	r1, _, err := dllEverythingGetBuildNumber.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_Exit(void);
var dllEverythingExit *syscall.LazyProc

func EverythingExit() error {
	r1, _, err := dllEverythingExit.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsDBLoaded(void); // Everything 1.4.1
var dllEverythingIsDBLoaded *syscall.LazyProc

func EverythingIsDBLoaded() error {
	r1, _, err := dllEverythingIsDBLoaded.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsAdmin(void); // Everything 1.4.1
var dllEverythingIsAdmin *syscall.LazyProc

func EverythingIsAdmin() error {
	r1, _, err := dllEverythingIsAdmin.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_IsAppData(void); // Everything 1.4.1
var dllEverythingIsAppData *syscall.LazyProc

func EverythingIsAppData() error {
	r1, _, err := dllEverythingIsAppData.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_RebuildDB(void); // Everything 1.4.1
var dllEverythingRebuildDB *syscall.LazyProc

func EverythingRebuildDB() error {
	r1, _, err := dllEverythingRebuildDB.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_UpdateAllFolderIndexes(void); // Everything 1.4.1
var dllEverythingUpdateAllFolderIndexes *syscall.LazyProc

func EverythingUpdateAllFolderIndexes() error {
	r1, _, err := dllEverythingUpdateAllFolderIndexes.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_SaveDB(void); // Everything 1.4.1
var dllEverythingSaveDB *syscall.LazyProc

func EverythingSaveDB() error {
	r1, _, err := dllEverythingSaveDB.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_SaveRunHistory(void); // Everything 1.4.1
var dllEverythingSaveRunHistory *syscall.LazyProc

func EverythingSaveRunHistory() error {
	r1, _, err := dllEverythingSaveRunHistory.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_DeleteRunHistory(void); // Everything 1.4.1
var dllEverythingDeleteRunHistory *syscall.LazyProc

func EverythingDeleteRunHistory() error {
	r1, _, err := dllEverythingDeleteRunHistory.Call()
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetTargetMachine(void); // Everything 1.4.1
var dllEverythingGetTargetMachine *syscall.LazyProc

func EverythingGetTargetMachine() (uint32, error) {
	r1, _, err := dllEverythingGetTargetMachine.Call()
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetRunCountFromFileNameW(LPCWSTR lpFileName); // Everything 1.4.1
var dllEverythingGetRunCountFromFileNameW *syscall.LazyProc

func everythingGetRunCountFromFileNameW(lpFileName string) (uint32, error) {
	sp, err := syscall.UTF16PtrFromString(lpFileName)
	if err != nil {
		return 0, err
	}
	r1, _, err := dllEverythingGetRunCountFromFileNameW.Call(uintptr(unsafe.Pointer(sp)))
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_GetRunCountFromFileNameA(LPCSTR lpFileName); // Everything 1.4.1
var dllEverythingGetRunCountFromFileNameA *syscall.LazyProc

func everythingGetRunCountFromFileNameA(lpFileName string) (uint32, error) {
	sp, err := syscall.BytePtrFromString(lpFileName)
	if err != nil {
		return 0, err
	}
	r1, _, err := dllEverythingGetRunCountFromFileNameA.Call(uintptr(unsafe.Pointer(sp)))
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_SetRunCountFromFileNameW(LPCWSTR lpFileName,DWORD dwRunCount); // Everything 1.4.1
var dllEverythingSetRunCountFromFileNameW *syscall.LazyProc

func everythingSetRunCountFromFileNameW(lpFileName string, dwRunCount uint32) error {
	sp, err := syscall.UTF16PtrFromString(lpFileName)
	if err != nil {
		return err
	}
	r1, _, err := dllEverythingSetRunCountFromFileNameW.Call(
		uintptr(unsafe.Pointer(sp)), uintptr(dwRunCount))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI BOOL EVERYTHINGAPI Everything_SetRunCountFromFileNameA(LPCSTR lpFileName,DWORD dwRunCount); // Everything 1.4.1
var dllEverythingSetRunCountFromFileNameA *syscall.LazyProc

func everythingSetRunCountFromFileNameA(lpFileName string, dwRunCount uint32) error {
	sp, err := syscall.BytePtrFromString(lpFileName)
	if err != nil {
		return err
	}
	r1, _, err := dllEverythingSetRunCountFromFileNameA.Call(
		uintptr(unsafe.Pointer(sp)), uintptr(dwRunCount))
	return checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_IncRunCountFromFileNameW(LPCWSTR lpFileName); // Everything 1.4.1
var dllEverythingIncRunCountFromFileNameW *syscall.LazyProc

func everythingIncRunCountFromFileNameW(lpFileName string) (uint32, error) {
	sp, err := syscall.UTF16PtrFromString(lpFileName)
	if err != nil {
		return 0, err
	}
	r1, _, err := dllEverythingIncRunCountFromFileNameW.Call(uintptr(unsafe.Pointer(sp)))
	return uint32(r1), checkErr(r1, err)
}

// EVERYTHINGUSERAPI DWORD EVERYTHINGAPI Everything_IncRunCountFromFileNameA(LPCSTR lpFileName); // Everything 1.4.1
var dllEverythingIncRunCountFromFileNameA *syscall.LazyProc

func everythingIncRunCountFromFileNameA(lpFileName string) (uint32, error) {
	sp, err := syscall.BytePtrFromString(lpFileName)
	if err != nil {
		return 0, err
	}
	r1, _, err := dllEverythingIncRunCountFromFileNameA.Call(uintptr(unsafe.Pointer(sp)))
	return uint32(r1), checkErr(r1, err)
}
