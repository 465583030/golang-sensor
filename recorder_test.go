package instana_test

import (
	"testing"

	"github.com/instana/golang-sensor"
	ext "github.com/opentracing/opentracing-go/ext"
	"github.com/stretchr/testify/assert"
)

func TestRecorderBasics(t *testing.T) {
	opts := instana.Options{LogLevel: instana.Debug}
	recorder := instana.NewTestRecorder()
	tracer := instana.NewTracerWithEverything(&opts, recorder)

	span := tracer.StartSpan("http-client")
	span.SetTag(string(ext.SpanKind), "exit")
	span.SetTag("http.status", 200)
	span.SetTag("http.url", "https://www.instana.com/product/")
	span.SetTag(string(ext.HTTPMethod), "GET")
	span.Finish()

	// Validate GetSpans
	spans := recorder.GetSpans()
	assert.Equal(t, 1, len(spans))

	// Validate Reset & GetSpans Result
	recorder.Reset()
	spans = recorder.GetSpans()
	assert.Equal(t, 0, len(spans))
}
