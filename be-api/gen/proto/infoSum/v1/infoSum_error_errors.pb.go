// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package infoSumv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsGetInfosumError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_INFOSUM_ERROR.String() && e.Code == 501
}

func ErrorGetInfosumError(format string, args ...interface{}) *errors.Error {
	return errors.New(501, ErrorReason_GET_INFOSUM_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsDelInfosumError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DEL_INFOSUM_ERROR.String() && e.Code == 502
}

func ErrorDelInfosumError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, ErrorReason_DEL_INFOSUM_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsSaveInfosumError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SAVE_INFOSUM_ERROR.String() && e.Code == 503
}

func ErrorSaveInfosumError(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ErrorReason_SAVE_INFOSUM_ERROR.String(), fmt.Sprintf(format, args...))
}
