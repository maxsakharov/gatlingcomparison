package models

import (
	"strings"
	"time"
)

type ReportType int

const (
	REPORT_TYPE_HTML ReportType = iota
	REPORT_TYPE_LOG
	REPORT_TYPE_UNKNOWN
)

var types = map[ReportType]string{
	REPORT_TYPE_HTML:    "HTML",
	REPORT_TYPE_LOG:     "LOG",
	REPORT_TYPE_UNKNOWN: "UNKNOWN",
}

func (rt ReportType) String() string {
	return types[rt]
}

func ReportTypeForName(name string) ReportType {
	for i, t := range types {
		if strings.ToUpper(name) == strings.ToUpper(t) {
			return i
		}
	}
	return REPORT_TYPE_UNKNOWN
}

type Report struct {
	Reports *[]Report
}

type ReportRequest struct {
	Name        string
	RequestName string
	Ok          int
	Ko          int
	Rps         int
	Min         int
	Max         int
	Pct50       float32
	Pct75       float32
	Pct90       float32
	Pct95       float32
	Pct99       float32
}

type ReportMetadata struct {
	Name       string
	DateUpload time.Time
	Type       ReportType
}

type ReportMetadataList struct {
	Reports []ReportMetadata
}
