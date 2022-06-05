package metrics

import (
	"context"
	"go-kit/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
)

var _ metrics.Histogram = (*histogram)(nil)

type histogram struct {
	his syncint64.Histogram
	lvs []attribute.KeyValue
}

// NewHistogram new a prometheus histogram and returns Histogram.
func NewHistogram(h syncint64.Histogram) metrics.Histogram {
	return &histogram{
		his: h,
	}
}

func (h histogram) With(lvs []attribute.KeyValue) metrics.Histogram {
	return &histogram{
		his: h.his,
		lvs: lvs,
	}
}

func (h histogram) Record(ctx context.Context, incr int64) {
	h.his.Record(ctx, incr, h.lvs...)
}
