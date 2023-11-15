package sigexit

import "fmt"

func GetExitStatus(number int) string {
	code := number
	if code > 255 {
		code = number % 256
	}

	if status, ok := StandardExitCodes[code]; ok {
		return fmt.Sprintf("exit with non-zero code %d (%s)", code, status)
	}

	if code > 128 && code <= 165 {
		signum := code - 128
		if sig, ok := StandardSignals[code-128]; ok {
			return fmt.Sprintf("interrupt by signal %d (%s)", signum, sig.Symbol)
		} else {
			return fmt.Sprintf("interrupt by signal %d (UNKNOWN)", signum)
		}
	}

	if status, ok := SystematizeExitCodes[code]; ok {
		return fmt.Sprintf("exit with non-zero code %d (%s)", code, status)
	}

	return fmt.Sprintf("exit with non-zero code %d (UNKNOWN)", number)
}
