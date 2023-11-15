package sigexit

// https://tldp.org/LDP/abs/html/exitcodes.html
var (
	StandardExitCodes = map[int]string{
		1:   "catchall for general errors",
		2:   "misuse of shell builtins",
		126: "command invoked cannot execute",
		127: "command not found",
		128: "invalid argument to exit",
		255: "exit status out of range",
	}

	// See /usr/include/sysexits.h
	SystematizeExitCodes = map[int]SymbolDescription{
		64: {"EX_USAGE", "command line usage error"},
		65: {"EX_DATAERR", "data format error"},
		66: {"EX_NOINPUT", "cannot open input"},
		67: {"EX_NOUSER ", "addressee unknown"},
		68: {"EX_NOHOST ", "host name unknown"},
		69: {"EX_UNAVAILABLE", "service unavailable"},
		70: {"EX_SOFTWARE", "internal software error"},
		71: {"EX_OSERR", "system error (e.g., can't fork)"},
		72: {"EX_OSFILE ", "critical OS file missing"},
		73: {"EX_CANTCREAT  ", "can't create (user) output file"},
		74: {"EX_IOERR", "input/output error"},
		75: {"EX_TEMPFAIL", "temp failure; user is invited to retry"},
		76: {"EX_PROTOCOL", "remote error in protocol"},
		77: {"EX_NOPERM ", "permission denied"},
		78: {"EX_CONFIG ", "configuration error"},
	}
)
