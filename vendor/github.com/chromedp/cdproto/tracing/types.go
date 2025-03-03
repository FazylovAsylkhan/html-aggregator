package tracing

import (
	"fmt"
	"strings"
)

// Code generated by cdproto-gen. DO NOT EDIT.

// MemoryDumpConfig configuration for memory dump. Used only when
// "memory-infra" category is enabled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-MemoryDumpConfig
type MemoryDumpConfig struct{}

// TraceConfig [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-TraceConfig
type TraceConfig struct {
	RecordMode           RecordMode        `json:"recordMode,omitempty,omitzero"`           // Controls how the trace buffer stores data.
	TraceBufferSizeInKb  float64           `json:"traceBufferSizeInKb,omitempty,omitzero"`  // Size of the trace buffer in kilobytes. If not specified or zero is passed, a default value of 200 MB would be used.
	EnableSampling       bool              `json:"enableSampling,omitempty,omitzero"`       // Turns on JavaScript stack sampling.
	EnableSystrace       bool              `json:"enableSystrace,omitempty,omitzero"`       // Turns on system tracing.
	EnableArgumentFilter bool              `json:"enableArgumentFilter,omitempty,omitzero"` // Turns on argument filter.
	IncludedCategories   []string          `json:"includedCategories,omitempty,omitzero"`   // Included category filters.
	ExcludedCategories   []string          `json:"excludedCategories,omitempty,omitzero"`   // Excluded category filters.
	SyntheticDelays      []string          `json:"syntheticDelays,omitempty,omitzero"`      // Configuration to synthesize the delays in tracing.
	MemoryDumpConfig     *MemoryDumpConfig `json:"memoryDumpConfig,omitempty,omitzero"`     // Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}

// StreamFormat data format of a trace. Can be either the legacy JSON format
// or the protocol buffer format. Note that the JSON format will be deprecated
// soon.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-StreamFormat
type StreamFormat string

// String returns the StreamFormat as string value.
func (t StreamFormat) String() string {
	return string(t)
}

// StreamFormat values.
const (
	StreamFormatJSON  StreamFormat = "json"
	StreamFormatProto StreamFormat = "proto"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *StreamFormat) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch StreamFormat(s) {
	case StreamFormatJSON:
		*t = StreamFormatJSON
	case StreamFormatProto:
		*t = StreamFormatProto
	default:
		return fmt.Errorf("unknown StreamFormat value: %v", s)
	}
	return nil
}

// StreamCompression compression type to use for traces returned via streams.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-StreamCompression
type StreamCompression string

// String returns the StreamCompression as string value.
func (t StreamCompression) String() string {
	return string(t)
}

// StreamCompression values.
const (
	StreamCompressionNone StreamCompression = "none"
	StreamCompressionGzip StreamCompression = "gzip"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *StreamCompression) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch StreamCompression(s) {
	case StreamCompressionNone:
		*t = StreamCompressionNone
	case StreamCompressionGzip:
		*t = StreamCompressionGzip
	default:
		return fmt.Errorf("unknown StreamCompression value: %v", s)
	}
	return nil
}

// MemoryDumpLevelOfDetail details exposed when memory request explicitly
// declared. Keep consistent with memory_dump_request_args.h and
// memory_instrumentation.mojom.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-MemoryDumpLevelOfDetail
type MemoryDumpLevelOfDetail string

// String returns the MemoryDumpLevelOfDetail as string value.
func (t MemoryDumpLevelOfDetail) String() string {
	return string(t)
}

// MemoryDumpLevelOfDetail values.
const (
	MemoryDumpLevelOfDetailBackground MemoryDumpLevelOfDetail = "background"
	MemoryDumpLevelOfDetailLight      MemoryDumpLevelOfDetail = "light"
	MemoryDumpLevelOfDetailDetailed   MemoryDumpLevelOfDetail = "detailed"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *MemoryDumpLevelOfDetail) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch MemoryDumpLevelOfDetail(s) {
	case MemoryDumpLevelOfDetailBackground:
		*t = MemoryDumpLevelOfDetailBackground
	case MemoryDumpLevelOfDetailLight:
		*t = MemoryDumpLevelOfDetailLight
	case MemoryDumpLevelOfDetailDetailed:
		*t = MemoryDumpLevelOfDetailDetailed
	default:
		return fmt.Errorf("unknown MemoryDumpLevelOfDetail value: %v", s)
	}
	return nil
}

// Backend backend type to use for tracing. chrome uses the Chrome-integrated
// tracing service and is supported on all platforms. system is only supported
// on Chrome OS and uses the Perfetto system tracing service. auto chooses
// system when the perfettoConfig provided to Tracing.start specifies at least
// one non-Chrome data source; otherwise uses chrome.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-TracingBackend
type Backend string

// String returns the Backend as string value.
func (t Backend) String() string {
	return string(t)
}

// Backend values.
const (
	BackendAuto   Backend = "auto"
	BackendChrome Backend = "chrome"
	BackendSystem Backend = "system"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *Backend) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch Backend(s) {
	case BackendAuto:
		*t = BackendAuto
	case BackendChrome:
		*t = BackendChrome
	case BackendSystem:
		*t = BackendSystem
	default:
		return fmt.Errorf("unknown Backend value: %v", s)
	}
	return nil
}

// RecordMode controls how the trace buffer stores data.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#type-TraceConfig
type RecordMode string

// String returns the RecordMode as string value.
func (t RecordMode) String() string {
	return string(t)
}

// RecordMode values.
const (
	RecordModeRecordUntilFull        RecordMode = "recordUntilFull"
	RecordModeRecordContinuously     RecordMode = "recordContinuously"
	RecordModeRecordAsMuchAsPossible RecordMode = "recordAsMuchAsPossible"
	RecordModeEchoToConsole          RecordMode = "echoToConsole"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *RecordMode) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch RecordMode(s) {
	case RecordModeRecordUntilFull:
		*t = RecordModeRecordUntilFull
	case RecordModeRecordContinuously:
		*t = RecordModeRecordContinuously
	case RecordModeRecordAsMuchAsPossible:
		*t = RecordModeRecordAsMuchAsPossible
	case RecordModeEchoToConsole:
		*t = RecordModeEchoToConsole
	default:
		return fmt.Errorf("unknown RecordMode value: %v", s)
	}
	return nil
}

// TransferMode whether to report trace events as series of dataCollected
// events or to save trace to a stream (defaults to ReportEvents).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#method-start
type TransferMode string

// String returns the TransferMode as string value.
func (t TransferMode) String() string {
	return string(t)
}

// TransferMode values.
const (
	TransferModeReportEvents   TransferMode = "ReportEvents"
	TransferModeReturnAsStream TransferMode = "ReturnAsStream"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *TransferMode) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch TransferMode(s) {
	case TransferModeReportEvents:
		*t = TransferModeReportEvents
	case TransferModeReturnAsStream:
		*t = TransferModeReturnAsStream
	default:
		return fmt.Errorf("unknown TransferMode value: %v", s)
	}
	return nil
}
