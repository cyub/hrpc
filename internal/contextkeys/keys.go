package contextkeys

type contextKey int

const (
	MethodNameKey contextKey = 1 + iota
	ServiceNameKey
	PackageNameKey
	StatusCodeKey
	RequestHeaderKey
	ResponseWriterKey
)
