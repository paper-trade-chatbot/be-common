package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrCode int32

const (
	//member
	ErrCode_AccountRegistered      ErrCode = 1001
	ErrCode_AccountNotRegistered   ErrCode = 1002
	ErrCode_PhoneRegistered        ErrCode = 1003
	ErrCode_PhoneNotRegistered     ErrCode = 1010
	ErrCode_MailRegistered         ErrCode = 1012
	ErrCode_MailNotRegistered      ErrCode = 1011
	ErrCode_PasswordFormatMismatch ErrCode = 1004
	ErrCode_PasswordMismatch       ErrCode = 1005
	ErrCode_InvalidCountry         ErrCode = 1013
	ErrCode_PhoneVerifiedMismatch  ErrCode = 1007
	ErrCode_MailVerifiedMismatch   ErrCode = 1009
)

var (
	//member
	ErrAccountRegistered      = status.Error(codes.Code(ErrCode_AccountRegistered), "account registered")
	ErrAccountNotRegistered   = status.Error(codes.Code(ErrCode_AccountNotRegistered), "account not registered")
	ErrPhoneRegistered        = status.Error(codes.Code(ErrCode_PhoneRegistered), "phone number registered")
	ErrPasswordFormatMismatch = status.Error(codes.Code(ErrCode_PasswordFormatMismatch), "password format mismatch")
	ErrPasswordMismatch       = status.Error(codes.Code(ErrCode_PasswordMismatch), "password mismatch")
	ErrPhoneVerifiedMismatch  = status.Error(codes.Code(ErrCode_PhoneVerifiedMismatch), "phone verified mismatch")
	ErrMailVerifiedMismatch   = status.Error(codes.Code(ErrCode_MailVerifiedMismatch), "mail verified mismatch")
	ErrPhoneNotRegistered     = status.Error(codes.Code(ErrCode_PhoneNotRegistered), "phone number not registered")
	ErrMailNotRegistered      = status.Error(codes.Code(ErrCode_MailNotRegistered), "mail not registered")
	ErrMailRegistered         = status.Error(codes.Code(ErrCode_MailRegistered), "mail registered")
	ErrInvalidCountry         = status.Error(codes.Code(ErrCode_InvalidCountry), "invalid country")
)
