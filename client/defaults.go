package client

// see https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
const (
	// Request context
	HeaderUserAgent = "User-Agent"
	// Content types that are acceptable for the response
	HeaderAccept = "Accept"
)

// Message body information
const (
	// The MIME type of the body of the request (used with POST and PUT requests)
	ContentType = "Content-Type"
	ContentEncoding = "Content-Encoding"
	ContentLength= "Content-Length"
)

// values
const (
	UserAgent = "tz-goreadme-client/1.0"
	DefaultAcceptType = "application/json, text/plain, */*"
	ContentTypeURLParameters = "application/x-www-form-urlencoded;charset=utf-8"
	ContentTypeJSON = "application/json;charset=utf-8"
)

// default transformers
var (
DefaultRequestTransformers []RequestTransformer
DefaultResponseTransformers []ReponseTransformer
)

func init()  {
	DefaultResponseTransformers = []ReponseTransformer{}
	DefaultRequestTransformers = []RequestTransformer{
		basicRequestTransformer,
	}
}