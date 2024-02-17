package gum

/*
#include <frida-gumjs.h>
*/
import "C"

type GumOS = C.GumOS

const (
	GUM_OS_WINDOWS = GumOS(0)
	GUM_OS_MACOS   = GumOS(1)
	GUM_OS_LINUX   = GumOS(2)
	GUM_OS_IOS     = GumOS(3)
	GUM_OS_WATCHOS = GumOS(4)
	GUM_OS_TVOS    = GumOS(5)
	GUM_OS_ANDROID = GumOS(6)
	GUM_OS_FREEBSD = GumOS(7)
	GUM_OS_QNX     = GumOS(8)
)

type GumThreadState = C.GumThreadState

const (
	GUM_THREAD_RUNNING         = GumThreadState(1)
	GUM_THREAD_STOPPED         = GumThreadState(2)
	GUM_THREAD_WAITING         = GumThreadState(3)
	GUM_THREAD_UNINTERRUPTIBLE = GumThreadState(4)
	GUM_THREAD_HALTED          = GumThreadState(5)
)

type GumPageProtection = C.GumPageProtection

const (
	GUM_PAGE_NO_ACCESS = GumPageProtection(0)
	GUM_PAGE_READ      = GumPageProtection((1 << 0))
	GUM_PAGE_WRITE     = GumPageProtection((1 << 1))
	GUM_PAGE_EXECUTE   = GumPageProtection((1 << 2))
)

type GumSymbolType = C.GumSymbolType

const (
	/* Common */
	GUM_SYMBOL_UNKNOWN = GumSymbolType(0)
	GUM_SYMBOL_SECTION = GumSymbolType(1)

	/* Mach-O */
	GUM_SYMBOL_UNDEFINED          = GumSymbolType(2)
	GUM_SYMBOL_ABSOLUTE           = GumSymbolType(3)
	GUM_SYMBOL_PREBOUND_UNDEFINED = GumSymbolType(4)
	GUM_SYMBOL_INDIRECT           = GumSymbolType(5)

	/* ELF */
	GUM_SYMBOL_OBJECT   = GumSymbolType(6)
	GUM_SYMBOL_FUNCTION = GumSymbolType(7)
	GUM_SYMBOL_FILE     = GumSymbolType(8)
	GUM_SYMBOL_COMMON   = GumSymbolType(9)
	GUM_SYMBOL_TLS      = GumSymbolType(10)
)

type GumExportType = C.GumExportType

const (
	GUM_EXPORT_FUNCTION = GumExportType(1)
	GUM_EXPORT_VARIABLE = GumExportType(2)
)
