package tracer

import (
	"github.com/w3c/distributed-tracing/tests/api"
	"github.com/w3c/distributed-tracing/tests/internal/random"
)

// Tracer is a reference implementation of api.Tracer
type Tracer struct {
	//
}

// New creates a new Tracer.
func New() *Tracer {
	return &Tracer{}
}

func (t *Tracer) StartSpan(tc api.TraceContext) api.Span {
	// TODO tc.TraceState should take priority
	traceID, parentSpanID, sampled, err := tc.ParseTraceParent()
	if traceID == "" {
		traceID = random.New64BitID() + random.New64BitID()
	}
	if err != nil {
		sampled = true // TODO how should this be decided? read from Request?
	}
	return &Span{
		traceID:       traceID,
		spanID:        random.New64BitID(),
		parentSpanID:  parentSpanID,
		correlationID: "", // TODO should depend on the participation mode
		sampled:       sampled,
		traceState:    tc.TraceState,
	}
}
