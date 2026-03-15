package testdata

func notPrintfFuncAtAll() {}

func funcWithEllipsis(args ...interface{}) {}

func printfLikeButWithStrings(format string, args ...string) {}

func printfLikeButWithBadFormat(format int, args ...interface{}) {}

func prinfLikeFunc(format string, args ...interface{}) {}

func prinfLikeFuncWithReturnValue(format string, args ...interface{}) string {
	return ""
}
