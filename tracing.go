package tracing

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	blocks "github.com/ipfs/go-block-format"
	cid "github.com/ipfs/go-cid"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

// Span starts a new span using the standard IPFS tracing conventions.
func Span(ctx context.Context, componentName string, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return otel.Tracer("").Start(ctx, fmt.Sprintf("%s.%s", componentName, spanName), opts...)
}

// SpanWithStringAttribute is a helper function to assist the common pattern of starting a new span
// with a single string attribute
func SpanWithStringAttribute(ctx context.Context, componentName string, spanName string, k string, v string) (context.Context, trace.Span) {
	return Span(ctx, componentName, spanName, trace.WithAttributes(attribute.String(k, v)))
}

// SpanWithIntAttribute is a helper function to assist the common pattern of starting a new span
// with a single int attribute
func SpanWithIntAttribute(ctx context.Context, componentName string, spanName string, k string, v int) (context.Context, trace.Span) {
	return Span(ctx, componentName, spanName, trace.WithAttributes(attribute.Int(k, v)))
}

// SpanWithPathAttribute is a helper function to assist the common pattern of starting a new span
// with a single path attribute
func SpanWithPathAttribute(ctx context.Context, componentName string, spanName string, p path.Path) (context.Context, trace.Span) {
	ctx, span := Span(ctx, componentName, spanName)
	if span.IsRecording() {
		span.SetAttributes(attribute.String("path", p.String()))
	}
	return ctx, span
}

// SpanWithCidAttribute is a helper function to assist the common pattern of starting a new span
// with a single cid attribute
func SpanWithCidAttribute(ctx context.Context, componentName string, spanName string, c cid.Cid) (context.Context, trace.Span) {
	ctx, span := Span(ctx, componentName, spanName)
	if span.IsRecording() {
		span.SetAttributes(attribute.String("cid", c.String()))
	}
	return ctx, span
}

// SpanWithCidListAttribute is a helper function to assist the common pattern of starting a new span
// with a list of cids as an attribute
func SpanWithCidListAttribute(ctx context.Context, componentName string, spanName string, cs []cid.Cid) (context.Context, trace.Span) {
	ctx, span := Span(ctx, componentName, spanName)
	if span.IsRecording() {
		var value string
		if len(cs) == 0 {
			value = "empty list"
		} else {
			max := 3
			if max > len(cs) {
				max = len(cs)
			}

			cids := make([]string, max)
			for i := range cids {
				cids[i] = cs[i].String()
			}

			value = strings.Join(cids, ",")

			if max < len(cs) {
				value += fmt.Sprintf(" and %d more", len(cs)-max)
			}
		}
		span.SetAttributes(attribute.String("cids", value))
	}
	return ctx, span
}

// SpanWithBlockAttribute is a helper function to assist the common pattern of starting a new span
// with a single block attribute
func SpanWithBlockAttribute(ctx context.Context, componentName string, spanName string, b blocks.Block) (context.Context, trace.Span) {
	ctx, span := Span(ctx, componentName, spanName)
	if span.IsRecording() {
		span.SetAttributes(attribute.String("block", b.Cid().String()))
	}
	return ctx, span
}

// SpanWithBlockListAttribute is a helper function to assist the common pattern of starting a new span
// with a single attribute containing a list of blocks
func SpanWithBlockListAttribute(ctx context.Context, componentName string, spanName string, bs []blocks.Block) (context.Context, trace.Span) {
	ctx, span := Span(ctx, componentName, spanName)
	if span.IsRecording() {
		var value string
		if len(bs) == 0 {
			value = "empty list"
		} else {
			max := 3
			if max > len(bs) {
				max = len(bs)
			}

			cids := make([]string, max)
			for i := range cids {
				cids[i] = bs[i].Cid().String()
			}

			value = strings.Join(cids, ",")

			if max < len(bs) {
				value += fmt.Sprintf(" and %d more", len(bs)-max)
			}
			span.SetAttributes(attribute.String("blocks", value))
		}
	}
	return ctx, span
}
