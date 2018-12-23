//go:generate go-bindata -ignore=\.go -pkg=gtpl -o=bindata.go ./...

package gtpl

// Templates to be used in the response of suggester
var (
	Greeting         = NewFormatTemplate("response/greeting.md", MustAsset)
	NewcomerGreeting = NewStringTemplate("response/newcomer_greeting.md", MustAsset)
	Praise           = NewFormatTemplate("response/praise.md", MustAsset)
	Todo             = NewFormatTemplate("response/todo.md", MustAsset)
	Improvement      = NewFormatTemplate("response/improvement.md", MustAsset)
	Comment          = NewFormatTemplate("response/comment.md", MustAsset)
	Tip              = NewFormatTemplate("response/tip.md", MustAsset)
	Questions        = NewStringTemplate("response/questions.md", MustAsset)

	Compile       = NewStringTemplate("tools/compile.md", MustAsset)
	PassTests     = NewStringTemplate("tools/pass_tests.md", MustAsset)
	RaceCondition = NewStringTemplate("tools/race_condition.md", MustAsset)
	NotLinted     = NewStringTemplate("tools/not_linted.md", MustAsset)
	NotVetted     = NewStringTemplate("tools/not_vetted.md", MustAsset)
	NotFormatted  = NewStringTemplate("tools/not_formatted.md", MustAsset)

	Benchmarking = NewStringTemplate("topic/benchmarking.md", MustAsset)
	Regex        = NewStringTemplate("topic/regex.md", MustAsset)
	Hints        = NewStringTemplate("topic/hints.md", MustAsset)
	Tips         = NewStringTemplateSlice("tips/", MustAsset)
)
