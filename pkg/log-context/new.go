package log_context

import (
	"context"
	"fmt"
)

var MetadataKey = &Metadata{}

type CtxMetadataOpts string

var (
	OverwriteSetMetadata = CtxMetadataOpts("overwrite")
)

type Metadata struct {
	KeyPairs map[string]interface{}
}

type LogContext struct {
	ctx          context.Context // do not change value in context
	trackingData map[string]interface{}
}

func New(ctx context.Context) *LogContext {
	lCtx := &LogContext{
		ctx:          ctx,
		trackingData: make(map[string]interface{}),
	}
	value := ctx.Value(MetadataKey)
	if value == nil {
		return lCtx
	}

	meta, ok := value.(*Metadata)
	if ok {
		lCtx.trackingData = meta.KeyPairs
	}
	return lCtx
}

func (l *LogContext) GetCtx() context.Context {
	return l.ctx
}

func (l *LogContext) GetTrackingMessage() string {
	if len(l.trackingData) == 0 {
		return ""
	}

	msg := ""
	for key, value := range l.trackingData {
		if msg == "" {
			msg += fmt.Sprintf("%v: %v", key, value) // no comma
			continue
		}
		msg += fmt.Sprintf(", %v: %v", key, value)
	}

	return msg
}

func (l *LogContext) SetTrackingData(key string, value interface{}, opts ...CtxMetadataOpts) {
	l.trackingData[key] = value
	for _, opt := range opts {
		switch opt {
		case OverwriteSetMetadata:
			l.ctx = context.WithValue(l.ctx, MetadataKey, &Metadata{
				l.trackingData,
			})
		}
	}
}

func (l *LogContext) Clone() *LogContext {
	newMap := make(map[string]interface{})
	for key, value := range l.trackingData {
		newMap[key] = value
	}

	return &LogContext{
		ctx:          l.ctx,
		trackingData: newMap,
	}
}
