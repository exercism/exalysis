//go:generate go-bindata -ignore=\.go -pkg=gtpl -o=bindata.go ./...

package gtpl

//Templates to be used in the response of suggester
var (
	Greeting         = NewFormatTemplate("response/greeting.md", MustAsset)
	NewcomerGreeting = NewStringTemplate("response/newcomer_greeting.md", MustAsset)
	Praise           = NewFormatTemplate("response/praise.md", MustAsset)
	Todo             = NewFormatTemplate("response/todo.md", MustAsset)
	Improvement      = NewFormatTemplate("response/improvement.md", MustAsset)
	Comment          = NewFormatTemplate("response/comment.md", MustAsset)
	Tip              = NewFormatTemplate("response/tip.md", MustAsset)
	Questions        = NewStringTemplate("response/questions.md", MustAsset)

	Compile      = NewStringTemplate("tools/compile.md", MustAsset)
	PassTests    = NewStringTemplate("tools/pass_tests.md", MustAsset)
	NotLinted    = NewStringTemplate("tools/not_linted.md", MustAsset)
	NotFormatted = NewStringTemplate("tools/not_formatted.md", MustAsset)

	Benchmarking = NewStringTemplate("topic/benchmarking.md", MustAsset)
	Regex        = NewStringTemplate("topic/regex.md", MustAsset)

	Tips = []StringTemplate{
		NewStringTemplate("tips/advanced_tests.md", MustAsset),
		NewStringTemplate("tips/awesome_go.md", MustAsset),
		NewStringTemplate("tips/channels.md", MustAsset),
		NewStringTemplate("tips/community.md", MustAsset),
		NewStringTemplate("tips/concurrency.md", MustAsset),
		NewStringTemplate("tips/conduct.md", MustAsset),
		NewStringTemplate("tips/effective.md", MustAsset),
		NewStringTemplate("tips/errors.md", MustAsset),
		NewStringTemplate("tips/font.md", MustAsset),
		NewStringTemplate("tips/magic.md", MustAsset),
		NewStringTemplate("tips/modules.md", MustAsset),
		NewStringTemplate("tips/proverbs.md", MustAsset),
		NewStringTemplate("tips/review.md", MustAsset),
		NewStringTemplate("tips/simplicity.md", MustAsset),
		NewStringTemplate("tips/tdd.md", MustAsset),
		NewStringTemplate("tips/tests.md", MustAsset),
		NewStringTemplate("tips/tour.md", MustAsset),
	}
)
