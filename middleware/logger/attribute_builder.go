package logger

import (
	"log/slog"
	"time"
)

type AttributeBuilder struct {
	attrs []slog.Attr
}

func NewAttributeBuilder() *AttributeBuilder {
	return &AttributeBuilder{}
}

func (b *AttributeBuilder) WithString(key string, value string) *AttributeBuilder {
	b.attrs = append(b.attrs, slog.String(key, value))
	return b
}

func (b *AttributeBuilder) WithInt(key string, value int) *AttributeBuilder {
	b.attrs = append(b.attrs, slog.Int(key, value))
	return b
}

func (b *AttributeBuilder) WithDuration(key string, value time.Duration) *AttributeBuilder {
	b.attrs = append(b.attrs, slog.Duration(key, value))
	return b
}

func (b *AttributeBuilder) WithBool(key string, value bool) *AttributeBuilder {
	b.attrs = append(b.attrs, slog.Bool(key, value))
	return b
}

func (b *AttributeBuilder) WithMap(key string, values map[string]string) *AttributeBuilder {
	for k, v := range values {
		b.attrs = append(b.attrs, slog.String(k, v))
	}
	return b
}

func (b *AttributeBuilder) Build() []slog.Attr {
	return b.attrs
}
