package service

import "battery-analysis-platform/pkg/jd"

type JsonServicer interface {
	Do() (*jd.Response, error)
}

type FileServicer interface {
	Do() (string, error)
}
