package trace

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	apitrace "go.opentelemetry.io/otel/api/trace"
)

func TestSerialize(t *testing.T) {
	span := SpanData{
		SpanContext: apitrace.SpanContext{
			TraceID: apitrace.ID([16]byte{1}),
			SpanID:  apitrace.SpanID([8]byte{2}),
		},
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}

	b := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(b)
	err := encoder.Encode(span)
	assert.NoError(t, err)

	outSpan := SpanData{}
	original := b.String()
	decoder := json.NewDecoder(b)
	err = decoder.Decode(&outSpan)
	assert.NoError(t, err)

	b2 := bytes.NewBuffer(nil)
	encoder2 := json.NewEncoder(b2)
	err = encoder2.Encode(outSpan)
	assert.NoError(t, err)
	assert.Equal(t, original, b2.String())
}
