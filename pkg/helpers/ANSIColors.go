package helpers

import (
	"fmt"
	"os"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type ToBeColored struct {
	Expression string
}

func (t *ToBeColored) StdoutError() *ToBeColored {
	fmt.Printf(ErrorColor, t.Expression)
	fmt.Println()

	return t
}

func (t *ToBeColored) StdoutInfo() *ToBeColored {
	fmt.Printf(InfoColor, t.Expression)
	fmt.Println()

	return t
}

func (t *ToBeColored) StdoutWarning() *ToBeColored {
	fmt.Printf(WarningColor, t.Expression)
	fmt.Println()
	return t
}

func (t *ToBeColored) StdoutDebug() *ToBeColored {
	fmt.Printf(DebugColor, t.Expression)
	fmt.Println()
	return t
}

func (t *ToBeColored) StdoutNotice() *ToBeColored {
	fmt.Printf(NoticeColor, t.Expression)
	fmt.Println()
	return t
}

func (t *ToBeColored) StopExecution() {
	os.Exit(1)
}
