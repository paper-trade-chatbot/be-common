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
	ErrCode_NoRequiredParam  ErrCode = 1005
	ErrCode_InvalidParam     ErrCode = 1006
	ErrCode_NoPermission     ErrCode = 1007

	//api
	ErrCode_APIRequestTooMany ErrCode = 2001

	//member
	ErrCode_AccountRegistered      ErrCode = 3001
	ErrCode_AccountNotRegistered   ErrCode = 3002
	ErrCode_PhoneRegistered        ErrCode = 3003
	ErrCode_PhoneNotRegistered     ErrCode = 3004
	ErrCode_MailRegistered         ErrCode = 3005
	ErrCode_MailNotRegistered      ErrCode = 3006
	ErrCode_PasswordFormatMismatch ErrCode = 3007
	ErrCode_PasswordMismatch       ErrCode = 3008
	ErrCode_InvalidCountry         ErrCode = 3009
	ErrCode_PhoneVerifiedMismatch  ErrCode = 3010
	ErrCode_MailVerifiedMismatch   ErrCode = 3011
	ErrCode_WrongAccountFormat     ErrCode = 3012
	ErrCode_WrongPasswordFormat    ErrCode = 3013
	ErrCode_WrongMailFormat        ErrCode = 3014
	ErrCode_WrongPhoneFormat       ErrCode = 3015

	//auth
	ErrCode_TokenExpired   ErrCode = 4001
	ErrCode_MemberDisabled ErrCode = 4002
	ErrCode_MemberDerived  ErrCode = 4003

	//product
	ErrCode_NoSuchProduct ErrCode = 5001

	//candle

	//quote

	//wallet
	ErrCode_UpdateWalletInterrupted ErrCode = 8001
	ErrCode_NoSuchWallet            ErrCode = 8002
	ErrCode_NoSuchTransactionRecord ErrCode = 8003
	ErrCode_TransactionNotSuccess   ErrCode = 8004
	ErrCode_NoSuchMember            ErrCode = 8005
	ErrCode_InsufficientBalance     ErrCode = 8006

	//order
	ErrCode_NoSuchOrder     ErrCode = 10001
	ErrCode_OrderNotPending ErrCode = 10002

	//position
	ErrCode_InvalidCloseAmount ErrCode = 11001
	ErrCode_NoSuchPosition     ErrCode = 11002
)

var (
	//all
	ErrNoQueryCondition = status.Error(codes.Code(ErrCode_NoQueryCondition), "no query condition")
	ErrNotImplemented   = status.Error(codes.Code(ErrCode_NotImplemented), "not implemented")
	ErrUnknown          = status.Error(codes.Code(ErrCode_Unknown), "unknown error")
	ErrInternal         = status.Error(codes.Code(ErrCode_Internal), "internal error")
	ErrNoRequiredParam  = status.Error(codes.Code(ErrCode_NoRequiredParam), "no required parameter")
	ErrInvalidParam     = status.Error(codes.Code(ErrCode_InvalidParam), "invalid parameter")
	ErrNoPermission     = status.Error(codes.Code(ErrCode_NoPermission), "no permission")

	//api
	ErrAPIRequestTooMany = status.Error(codes.Code(ErrCode_APIRequestTooMany), "api request too many")

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
	ErrWrongAccountFormat     = status.Error(codes.Code(ErrCode_WrongAccountFormat), "wrong account format")
	ErrWrongPasswordFormat    = status.Error(codes.Code(ErrCode_WrongPasswordFormat), "wrong password format")
	ErrWrongMailFormat        = status.Error(codes.Code(ErrCode_WrongMailFormat), "wrong mail format")
	ErrWrongPhoneFormat       = status.Error(codes.Code(ErrCode_WrongPhoneFormat), "wrong phone format")

	//auth
	ErrTokenExpired   = status.Error(codes.Code(ErrCode_TokenExpired), "token expired")
	ErrMemberDisabled = status.Error(codes.Code(ErrCode_MemberDisabled), "member disabled")
	ErrMemberDerived  = status.Error(codes.Code(ErrCode_MemberDerived), "member derived")

	//product
	ErrNoSuchProduct = status.Error(codes.Code(ErrCode_NoSuchProduct), "no such product")

	//candle

	//quote

	//wallet
	ErrUpdateWalletInterrupted = status.Error(codes.Code(ErrCode_UpdateWalletInterrupted), "interrupted when updating wallet")
	ErrNoSuchWallet            = status.Error(codes.Code(ErrCode_NoSuchWallet), "no such wallet")
	ErrNoSuchTransactionRecord = status.Error(codes.Code(ErrCode_NoSuchTransactionRecord), "no transaction record")
	ErrTransactionNotSuccess   = status.Error(codes.Code(ErrCode_TransactionNotSuccess), "this transaction is not successful")
	ErrNoSuchMember            = status.Error(codes.Code(ErrCode_NoSuchMember), "no such member")
	ErrInsufficientBalance     = status.Error(codes.Code(ErrCode_InsufficientBalance), "insufficient balance")

	//order
	ErrNoSuchOrder     = status.Error(codes.Code(ErrCode_NoSuchOrder), "no such order")
	ErrOrderNotPending = status.Error(codes.Code(ErrCode_OrderNotPending), "order not pending")

	//position
	ErrInvalidCloseAmount = status.Error(codes.Code(ErrCode_InvalidCloseAmount), "invalid close amount")
	ErrNoSuchPosition     = status.Error(codes.Code(ErrCode_NoSuchPosition), "no such position")
)
