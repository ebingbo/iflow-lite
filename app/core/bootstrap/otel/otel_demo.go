package otel

import "iflow-lite/core/otel"

func InitOTEL() {
	otel.InitTracer()
	otel.InitMetric()
}
