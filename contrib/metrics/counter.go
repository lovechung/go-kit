package metrics

import (
	"context"
	"github.com/lovechung/go-kit/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
)

var _ metrics.Counter = (*counter)(nil)

type counter struct {
	cnt syncint64.Counter
	lvs []attribute.KeyValue
}

func NewCounter(c syncint64.Counter) metrics.Counter {
	return &counter{
		cnt: c,
	}
}

func (c *counter) With(lvs []attribute.KeyValue) metrics.Counter {
	return &counter{
		cnt: c.cnt,
		lvs: lvs,
	}
}

func (c *counter) Add(ctx context.Context, incr int64) {
	c.cnt.Add(ctx, incr, c.lvs...)
}
