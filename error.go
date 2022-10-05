package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrCode int32

const (
	//all
	ErrCode_NoQueryCondition ErrCode = 1001
	ErrCode_NotImplemented   ErrCode = 1002
	ErrCode_Unknown          ErrCode = 1003
	ErrCode_Internal         ErrCode = 1004

	//member
	ErrCode_AccountRegistered      ErrCode = 2001
	ErrCode_AccountNotRegistered   ErrCode = 2002
	ErrCode_PhoneRegistered        ErrCode = 2003
	ErrCode_PhoneNotRegistered     ErrCode = 2004
	ErrCode_MailRegistered         ErrCode = 2005
	ErrCode_MailNotRegistered      ErrCode = 2006
	ErrCode_PasswordFormatMismatch ErrCode = 2007
	ErrCode_PasswordMismatch       ErrCode = 2008
	ErrCode_InvalidCountry         ErrCode = 2009
	ErrCode_PhoneVerifiedMismatch  ErrCode = 2010
	ErrCode_MailVerifiedMismatch   ErrCode = 2011
)

var (
	//all
	ErrNoQueryCondition = status.Error(codes.Code(ErrCode_NoQueryCondition), "no query condition")
	ErrNotImplemented   = status.Error(codes.Code(ErrCode_NotImplemented), "not implemented")

	//member
	ErrAccountRegistered      = status.Error(codes.Code(ErrCode_AccountRegistered), "account registered")
	ErrAccountNotRegistered   = status.Error(codes.Code(ErrCode_AccountNotRegistered), "account not registered")
	ErrPhoneRegistered        = status.Error(codes.Code(ErrCode_PhoneRegistered), "phone number registered")
	ErrPhoneNotRegistered     = status.Error(codes.Code(ErrCode_PhoneNotRegistered), "phone number not registered")
	ErrMailRegistered         = status.Error(codes.Code(ErrCode_MailRegistered), "mail registered")
	ErrMailNotRegistered      = status.Error(codes.Code(ErrCode_MailNotRegistered), "mail not registered")
	ErrPasswordFormatMismatch = status.Error(codes.Code(ErrCode_PasswordFormatMismatch), "password format mismatch")
	ErrPasswordMismatch       = status.Error(codes.Code(ErrCode_PasswordMismatch), "password mismatch")
	ErrInvalidCountry         = status.Error(codes.Code(ErrCode_InvalidCountry), "invalid country")
	ErrPhoneVerifiedMismatch  = status.Error(codes.Code(ErrCode_PhoneVerifiedMismatch), "phone verified mismatch")
	ErrMailVerifiedMismatch   = status.Error(codes.Code(ErrCode_MailVerifiedMismatch), "mail verified mismatch")
)
