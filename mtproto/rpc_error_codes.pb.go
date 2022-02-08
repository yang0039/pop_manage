// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc_error_codes.proto

package mtproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TLRpcErrorCodes int32

const (
	//	Error handling
	//
	//	There will be errors when working with the API, and they must be correctly handled on the client.
	//
	//	An error is characterized by several parameters:
	//
	//	Error Code
	//	Similar to HTTP status. Contains information on the type of error that occurred: for example,
	//	a data input error, privacy error, or server error. This is a required parameter.
	//
	//	Error Type
	//	A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED.
	//	This is an optional parameter.
	//
	//	Error Description
	//	May contain more detailed information on the error and how to resolve it,
	//	for example: authorization required, use auth.* methods. Please note that the description text is subject to change,
	//	one should avoid tying application logic to these messages. This is an optional parameter.
	//
	//	Error Constructors
	//	There should be a way to handle errors that are returned in rpc_error constructors.
	//
	//	If an error constructor does not differentiate between type and description
	//	but instead contains a single field called error_message (as in the example above),
	//	it must be split into 2 components, for example, using the following regular expression: /^([A-Z_0-9]+)(: (.+))?/.
	//
	// Protobuf3 enum first is 0
	TLRpcErrorCodes_ERROR_CODE_OK TLRpcErrorCodes = 0
	// The request must be repeated, but directed to a different data center.
	//
	// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
	// PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
	// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
	// USER_MIGRATE_X: the user whose identity is being used to execute
	// 				   queries is associated with a different data center (for registration)
	//
	// In all these cases, the error description’s string literal
	// 		contains the number of the data center (instead of the X) to which the repeated query must be sent.
	// More information about redirects between data centers »
	//
	TLRpcErrorCodes_FILE_MIGRATE_X    TLRpcErrorCodes = 303000
	TLRpcErrorCodes_PHONE_MIGRATE_X   TLRpcErrorCodes = 303001
	TLRpcErrorCodes_NETWORK_MIGRATE_X TLRpcErrorCodes = 303002
	TLRpcErrorCodes_USER_MIGRATE_X    TLRpcErrorCodes = 303003
	TLRpcErrorCodes_ERROR_SEE_OTHER   TLRpcErrorCodes = 303
	// The query contains errors. In the event that a request was created using a form
	// and contains user generated data,
	// the user should be notified that the data must be corrected before the query is repeated.
	//
	//
	// Examples of Errors:
	//	FIRSTNAME_INVALID: The first name is invalid
	//	LASTNAME_INVALID: The last name is invalid
	//	PHONE_NUMBER_INVALID: The phone number is invalid
	//	PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
	//	PHONE_CODE_EMPTY: phone_code is missing
	//	PHONE_CODE_EXPIRED: The confirmation code has expired
	//	API_ID_INVALID: The api_id/api_hash combination is invalid
	//	PHONE_NUMBER_OCCUPIED: The phone number is already in use
	//	PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
	//	USERS_TOO_FEW: Not enough users (to create a chat, for example)
	//	USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
	//	TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
	//	FILE_PART_INVALID: The file part number is invalid
	//	FILE_PARTS_INVALID: The number of file parts is invalid
	//	FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage
	//	MD5_CHECKSUM_INVALID: The MD5 checksums do not match
	//	PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
	//	FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
	//	FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
	//	MSG_WAIT_FAILED: A waiting call returned an error
	//
	TLRpcErrorCodes_FIRSTNAME_INVALID        TLRpcErrorCodes = 400000
	TLRpcErrorCodes_LASTNAME_INVALID         TLRpcErrorCodes = 400001
	TLRpcErrorCodes_PHONE_NUMBER_INVALID     TLRpcErrorCodes = 400002
	TLRpcErrorCodes_PHONE_CODE_HASH_EMPTY    TLRpcErrorCodes = 400003
	TLRpcErrorCodes_PHONE_CODE_EMPTY         TLRpcErrorCodes = 400004
	TLRpcErrorCodes_PHONE_CODE_EXPIRED       TLRpcErrorCodes = 400005
	TLRpcErrorCodes_API_ID_INVALID           TLRpcErrorCodes = 400006
	TLRpcErrorCodes_PHONE_NUMBER_OCCUPIED    TLRpcErrorCodes = 400007
	TLRpcErrorCodes_PHONE_NUMBER_UNOCCUPIED  TLRpcErrorCodes = 400008
	TLRpcErrorCodes_USERS_TOO_FEW            TLRpcErrorCodes = 400009
	TLRpcErrorCodes_USERS_TOO_MUCH           TLRpcErrorCodes = 400010
	TLRpcErrorCodes_TYPE_CONSTRUCTOR_INVALID TLRpcErrorCodes = 400011
	TLRpcErrorCodes_FILE_PART_INVALID        TLRpcErrorCodes = 400012
	TLRpcErrorCodes_FILE_PART_X_MISSING      TLRpcErrorCodes = 400013
	TLRpcErrorCodes_MD5_CHECKSUM_INVALID     TLRpcErrorCodes = 400014
	TLRpcErrorCodes_PHOTO_INVALID_DIMENSIONS TLRpcErrorCodes = 400015
	TLRpcErrorCodes_FIELD_NAME_INVALID       TLRpcErrorCodes = 400016
	TLRpcErrorCodes_FIELD_NAME_EMPTY         TLRpcErrorCodes = 400017
	TLRpcErrorCodes_MSG_WAIT_FAILED          TLRpcErrorCodes = 400018
	// android client code:
	//    if (error.code == 400 && "PARTICIPANT_VERSION_OUTDATED".equals(error.text)) {
	//        callFailed(VoIPController.ERROR_PEER_OUTDATED);
	TLRpcErrorCodes_PARTICIPANT_VERSION_OUTDATED TLRpcErrorCodes = 400019
	//
	TLRpcErrorCodes_PHONE_CODE_INVALID          TLRpcErrorCodes = 400020
	TLRpcErrorCodes_PHONE_INVITE_CODE_INCORRECT TLRpcErrorCodes = 400021
	TLRpcErrorCodes_PHONE_INVITE_CODE_INVALID   TLRpcErrorCodes = 400022
	TLRpcErrorCodes_PHONE_NUMBER_BANNED         TLRpcErrorCodes = 400030
	TLRpcErrorCodes_SESSION_PASSWORD_NEEDED     TLRpcErrorCodes = 400040
	// password
	TLRpcErrorCodes_CODE_INVALID          TLRpcErrorCodes = 400050
	TLRpcErrorCodes_PASSWORD_HASH_INVALID TLRpcErrorCodes = 400051
	TLRpcErrorCodes_NEW_PASSWORD_BAD      TLRpcErrorCodes = 400052
	TLRpcErrorCodes_NEW_SALT_INVALID      TLRpcErrorCodes = 400053
	TLRpcErrorCodes_EMAIL_INVALID         TLRpcErrorCodes = 400054
	TLRpcErrorCodes_EMAIL_UNCONFIRMED     TLRpcErrorCodes = 400055
	TLRpcErrorCodes_PASSWORD_CHECK_FAILED TLRpcErrorCodes = 400056
	TLRpcErrorCodes_PASSWORD_BAD_FORMAT   TLRpcErrorCodes = 400057
	TLRpcErrorCodes_EMAIL_UNCONFIRMED_6   TLRpcErrorCodes = 400058
	// username
	TLRpcErrorCodes_USERNAME_INVALID      TLRpcErrorCodes = 400060
	TLRpcErrorCodes_USERNAME_OCCUPIED     TLRpcErrorCodes = 400061
	TLRpcErrorCodes_USERNAMES_UNAVAILABLE TLRpcErrorCodes = 400062
	// chat
	TLRpcErrorCodes_CHAT_ID_INVALID         TLRpcErrorCodes = 400070
	TLRpcErrorCodes_CHAT_NOT_MODIFIED       TLRpcErrorCodes = 400071
	TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS  TLRpcErrorCodes = 400072
	TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION TLRpcErrorCodes = 400073
	//
	TLRpcErrorCodes_BAD_REQUEST TLRpcErrorCodes = 400
	// There was an unauthorized attempt to use functionality available only to authorized users.
	//
	// Examples of Errors:
	//	AUTH_KEY_UNREGISTERED: The key is not registered in the system
	//	AUTH_KEY_INVALID: The key is invalid
	//	USER_DEACTIVATED: The user has been deleted/deactivated
	//	SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	//	SESSION_EXPIRED: The authorization has expired
	//	ACTIVE_USER_REQUIRED: The method is only available to already activated users
	//	AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	//
	TLRpcErrorCodes_AUTH_KEY_UNREGISTERED TLRpcErrorCodes = 401000
	TLRpcErrorCodes_AUTH_KEY_INVALID      TLRpcErrorCodes = 401001
	TLRpcErrorCodes_USER_DEACTIVATED      TLRpcErrorCodes = 401002
	TLRpcErrorCodes_SESSION_REVOKED       TLRpcErrorCodes = 401003
	TLRpcErrorCodes_SESSION_EXPIRED       TLRpcErrorCodes = 401004
	TLRpcErrorCodes_ACTIVE_USER_REQUIRED  TLRpcErrorCodes = 401005
	TLRpcErrorCodes_AUTH_KEY_PERM_EMPTY   TLRpcErrorCodes = 401006
	// Only a small portion of the API methods are available to unauthorized users:
	//
	// auth.sendCode
	// auth.sendCall
	// auth.checkPhone
	// auth.signUp
	// auth.signIn
	// auth.importAuthorization
	// help.getConfig
	// help.getNearestDc
	//
	// Other methods will result in an error: 401 UNAUTHORIZED.
	TLRpcErrorCodes_UNAUTHORIZED TLRpcErrorCodes = 401
	// Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	//
	//
	// android client code:
	//    } else if(error.code==403 && "USER_PRIVACY_RESTRICTED".equals(error.text)){
	//        callFailed(VoIPController.ERROR_PRIVACY);
	TLRpcErrorCodes_USER_PRIVACY_RESTRICTED TLRpcErrorCodes = 403000
	TLRpcErrorCodes_FORBIDDEN               TLRpcErrorCodes = 403
	// 406
	// android client code:
	// }else if(error.code==406){
	//     callFailed(VoIPController.ERROR_LOCALIZED);
	TLRpcErrorCodes_ERROR_LOCALIZED TLRpcErrorCodes = 406000
	TLRpcErrorCodes_LOCALIZED       TLRpcErrorCodes = 406
	// The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded.
	// For example, in an attempt to request a large number of text messages (SMS) for the same phone number.
	//
	// Error Example:
	// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	//
	TLRpcErrorCodes_FLOOD_WAIT_X TLRpcErrorCodes = 420000
	// PEER_FLOOD
	// FLOOD_WAIT
	TLRpcErrorCodes_FLOOD TLRpcErrorCodes = 420
	// An internal server error occurred while a request was being processed;
	// for example, there was a disruption while accessing a database or file storage.
	//
	// If a client receives a 500 error, or you believe this error should not have occurred,
	// please collect as much information as possible about the query and error and send it to the developers.
	TLRpcErrorCodes_INTERNAL              TLRpcErrorCodes = 500
	TLRpcErrorCodes_INTERNAL_SERVER_ERROR TLRpcErrorCodes = 500000
	// If a server returns an error with a code other than the ones listed above,
	// it may be considered the same as a 500 error and treated as an internal server error.
	//
	TLRpcErrorCodes_OTHER TLRpcErrorCodes = 501
	//    // OFFSET_INVALID
	//    // RETRY_LIMIT
	//    // FILE_TOKEN_INVALID
	//    // REQUEST_TOKEN_INVALID
	//
	//    // CHANNEL_PRIVATE
	//    // CHANNEL_PUBLIC_GROUP_NA
	//    // USER_BANNED_IN_CHANNEL
	//
	//
	//    // MESSAGE_NOT_MODIFIED
	//
	//    // USERS_TOO_MUCH
	//
	//    // -1000
	//
	//    /////////////////////////////////////////////////////////////
	//     // android client code:
	//       } else if (request instanceof TLRPC.TL_auth_resendCode) {
	//        if (error.text.contains("PHONE_NUMBER_INVALID")) {
	//            showSimpleAlert(fragment, LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	//        } else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
	//            showSimpleAlert(fragment, LocaleController.getString("InvalidCode", R.string.InvalidCode));
	//        } else if (error.text.contains("PHONE_CODE_EXPIRED")) {
	//            showSimpleAlert(fragment, LocaleController.getString("CodeExpired", R.string.CodeExpired));
	//        } else if (error.text.startsWith("FLOOD_WAIT")) {
	//            showSimpleAlert(fragment, LocaleController.getString("FloodWait", R.string.FloodWait));
	//        } else if (error.code != -1000) {
	//            showSimpleAlert(fragment, LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
	//        }
	//
	//     /////////////////////////////////////////////////////////////
	//        } else if (request instanceof TLRPC.TL_updateUserName) {
	//            switch (error.text) {
	//                case "USERNAME_INVALID":
	//                    showSimpleAlert(fragment, LocaleController.getString("UsernameInvalid", R.string.UsernameInvalid));
	//                    break;
	//                case "USERNAME_OCCUPIED":
	//                    showSimpleAlert(fragment, LocaleController.getString("UsernameInUse", R.string.UsernameInUse));
	//                    break;
	//                case "USERNAMES_UNAVAILABLE":
	//                    showSimpleAlert(fragment, LocaleController.getString("FeatureUnavailable", R.string.FeatureUnavailable));
	//                    break;
	//                default:
	//                    showSimpleAlert(fragment, LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred));
	//                    break;
	//            }
	//
	//     /////////////////////////////////////////////////////////////
	//            } else if (request instanceof TLRPC.TL_payments_sendPaymentForm) {
	//            switch (error.text) {
	//                case "BOT_PRECHECKOUT_FAILED":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentPrecheckoutFailed", R.string.PaymentPrecheckoutFailed));
	//                    break;
	//                case "PAYMENT_FAILED":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentFailed", R.string.PaymentFailed));
	//                    break;
	//                default:
	//                    showSimpleToast(fragment, error.text);
	//                    break;
	//            }
	//        } else if (request instanceof TLRPC.TL_payments_validateRequestedInfo) {
	//            switch (error.text) {
	//                case "SHIPPING_NOT_AVAILABLE":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentNoShippingMethod", R.string.PaymentNoShippingMethod));
	//                    break;
	//                default:
	//                    showSimpleToast(fragment, error.text);
	//                    break;
	//            }
	//        }
	//
	//     /////////////////////////////////////////////////////////////
	//
	//        } else {
	//            if (error.text.equals("2FA_RECENT_CONFIRM")) {
	//                needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ResetAccountCancelledAlert", R.string.ResetAccountCancelledAlert));
	//            } else if (error.text.startsWith("2FA_CONFIRM_WAIT_")) {
	//                Bundle params = new Bundle();
	//                params.putString("phoneFormated", requestPhone);
	//                params.putString("phoneHash", phoneHash);
	//                params.putString("code", phoneCode);
	//                params.putInt("startTime", ConnectionsManager.getInstance().getCurrentTime());
	//                params.putInt("waitTime", Utilities.parseInt(error.text.replace("2FA_CONFIRM_WAIT_", "")));
	//                setPage(8, true, params, false);
	//            } else {
	TLRpcErrorCodes_OTHER2 TLRpcErrorCodes = 502
	// db error
	TLRpcErrorCodes_DBERR      TLRpcErrorCodes = 600
	TLRpcErrorCodes_DBERR_SQL  TLRpcErrorCodes = 600000
	TLRpcErrorCodes_DBERR_CONN TLRpcErrorCodes = 600001
	// db error
	TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
)

var TLRpcErrorCodes_name = map[int32]string{
	0:      "ERROR_CODE_OK",
	303000: "FILE_MIGRATE_X",
	303001: "PHONE_MIGRATE_X",
	303002: "NETWORK_MIGRATE_X",
	303003: "USER_MIGRATE_X",
	303:    "ERROR_SEE_OTHER",
	400000: "FIRSTNAME_INVALID",
	400001: "LASTNAME_INVALID",
	400002: "PHONE_NUMBER_INVALID",
	400003: "PHONE_CODE_HASH_EMPTY",
	400004: "PHONE_CODE_EMPTY",
	400005: "PHONE_CODE_EXPIRED",
	400006: "API_ID_INVALID",
	400007: "PHONE_NUMBER_OCCUPIED",
	400008: "PHONE_NUMBER_UNOCCUPIED",
	400009: "USERS_TOO_FEW",
	400010: "USERS_TOO_MUCH",
	400011: "TYPE_CONSTRUCTOR_INVALID",
	400012: "FILE_PART_INVALID",
	400013: "FILE_PART_X_MISSING",
	400014: "MD5_CHECKSUM_INVALID",
	400015: "PHOTO_INVALID_DIMENSIONS",
	400016: "FIELD_NAME_INVALID",
	400017: "FIELD_NAME_EMPTY",
	400018: "MSG_WAIT_FAILED",
	400019: "PARTICIPANT_VERSION_OUTDATED",
	400020: "PHONE_CODE_INVALID",
	400021: "PHONE_INVITE_CODE_INCORRECT",
	400022: "PHONE_INVITE_CODE_INVALID",
	400030: "PHONE_NUMBER_BANNED",
	400040: "SESSION_PASSWORD_NEEDED",
	400050: "CODE_INVALID",
	400051: "PASSWORD_HASH_INVALID",
	400052: "NEW_PASSWORD_BAD",
	400053: "NEW_SALT_INVALID",
	400054: "EMAIL_INVALID",
	400055: "EMAIL_UNCONFIRMED",
	400056: "PASSWORD_CHECK_FAILED",
	400057: "PASSWORD_BAD_FORMAT",
	400058: "EMAIL_UNCONFIRMED_6",
	400060: "USERNAME_INVALID",
	400061: "USERNAME_OCCUPIED",
	400062: "USERNAMES_UNAVAILABLE",
	400070: "CHAT_ID_INVALID",
	400071: "CHAT_NOT_MODIFIED",
	400072: "PARTICIPANT_NOT_EXISTS",
	400073: "NO_EDIT_CHAT_PERMISSION",
	400:    "BAD_REQUEST",
	401000: "AUTH_KEY_UNREGISTERED",
	401001: "AUTH_KEY_INVALID",
	401002: "USER_DEACTIVATED",
	401003: "SESSION_REVOKED",
	401004: "SESSION_EXPIRED",
	401005: "ACTIVE_USER_REQUIRED",
	401006: "AUTH_KEY_PERM_EMPTY",
	401:    "UNAUTHORIZED",
	403000: "USER_PRIVACY_RESTRICTED",
	403:    "FORBIDDEN",
	406000: "ERROR_LOCALIZED",
	406:    "LOCALIZED",
	420000: "FLOOD_WAIT_X",
	420:    "FLOOD",
	500:    "INTERNAL",
	500000: "INTERNAL_SERVER_ERROR",
	501:    "OTHER",
	502:    "OTHER2",
	600:    "DBERR",
	600000: "DBERR_SQL",
	600001: "DBERR_CONN",
	700:    "NOTRETURN_CLIENT",
}

var TLRpcErrorCodes_value = map[string]int32{
	"ERROR_CODE_OK":                0,
	"FILE_MIGRATE_X":               303000,
	"PHONE_MIGRATE_X":              303001,
	"NETWORK_MIGRATE_X":            303002,
	"USER_MIGRATE_X":               303003,
	"ERROR_SEE_OTHER":              303,
	"FIRSTNAME_INVALID":            400000,
	"LASTNAME_INVALID":             400001,
	"PHONE_NUMBER_INVALID":         400002,
	"PHONE_CODE_HASH_EMPTY":        400003,
	"PHONE_CODE_EMPTY":             400004,
	"PHONE_CODE_EXPIRED":           400005,
	"API_ID_INVALID":               400006,
	"PHONE_NUMBER_OCCUPIED":        400007,
	"PHONE_NUMBER_UNOCCUPIED":      400008,
	"USERS_TOO_FEW":                400009,
	"USERS_TOO_MUCH":               400010,
	"TYPE_CONSTRUCTOR_INVALID":     400011,
	"FILE_PART_INVALID":            400012,
	"FILE_PART_X_MISSING":          400013,
	"MD5_CHECKSUM_INVALID":         400014,
	"PHOTO_INVALID_DIMENSIONS":     400015,
	"FIELD_NAME_INVALID":           400016,
	"FIELD_NAME_EMPTY":             400017,
	"MSG_WAIT_FAILED":              400018,
	"PARTICIPANT_VERSION_OUTDATED": 400019,
	"PHONE_CODE_INVALID":           400020,
	"PHONE_INVITE_CODE_INCORRECT":  400021,
	"PHONE_INVITE_CODE_INVALID":    400022,
	"PHONE_NUMBER_BANNED":          400030,
	"SESSION_PASSWORD_NEEDED":      400040,
	"CODE_INVALID":                 400050,
	"PASSWORD_HASH_INVALID":        400051,
	"NEW_PASSWORD_BAD":             400052,
	"NEW_SALT_INVALID":             400053,
	"EMAIL_INVALID":                400054,
	"EMAIL_UNCONFIRMED":            400055,
	"PASSWORD_CHECK_FAILED":        400056,
	"PASSWORD_BAD_FORMAT":          400057,
	"EMAIL_UNCONFIRMED_6":          400058,
	"USERNAME_INVALID":             400060,
	"USERNAME_OCCUPIED":            400061,
	"USERNAMES_UNAVAILABLE":        400062,
	"CHAT_ID_INVALID":              400070,
	"CHAT_NOT_MODIFIED":            400071,
	"PARTICIPANT_NOT_EXISTS":       400072,
	"NO_EDIT_CHAT_PERMISSION":      400073,
	"BAD_REQUEST":                  400,
	"AUTH_KEY_UNREGISTERED":        401000,
	"AUTH_KEY_INVALID":             401001,
	"USER_DEACTIVATED":             401002,
	"SESSION_REVOKED":              401003,
	"SESSION_EXPIRED":              401004,
	"ACTIVE_USER_REQUIRED":         401005,
	"AUTH_KEY_PERM_EMPTY":          401006,
	"UNAUTHORIZED":                 401,
	"USER_PRIVACY_RESTRICTED":      403000,
	"FORBIDDEN":                    403,
	"ERROR_LOCALIZED":              406000,
	"LOCALIZED":                    406,
	"FLOOD_WAIT_X":                 420000,
	"FLOOD":                        420,
	"INTERNAL":                     500,
	"INTERNAL_SERVER_ERROR":        500000,
	"OTHER":                        501,
	"OTHER2":                       502,
	"DBERR":                        600,
	"DBERR_SQL":                    600000,
	"DBERR_CONN":                   600001,
	"NOTRETURN_CLIENT":             700,
}

func (x TLRpcErrorCodes) String() string {
	return proto.EnumName(TLRpcErrorCodes_name, int32(x))
}

func (TLRpcErrorCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7668e0f73d89aa3b, []int{0}
}

func init() {
	proto.RegisterEnum("mtproto.TLRpcErrorCodes", TLRpcErrorCodes_name, TLRpcErrorCodes_value)
}

func init() { proto.RegisterFile("rpc_error_codes.proto", fileDescriptor_7668e0f73d89aa3b) }

var fileDescriptor_7668e0f73d89aa3b = []byte{
	// 997 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0xcd, 0x6f, 0x1c, 0x35,
	0x14, 0x4f, 0xba, 0xb4, 0x50, 0xb7, 0x69, 0x1c, 0xb7, 0x69, 0x9d, 0xb6, 0x2c, 0x02, 0xf5, 0xc4,
	0x81, 0x03, 0x08, 0xee, 0xde, 0xf1, 0xdb, 0x8c, 0x95, 0x19, 0x7b, 0x6a, 0x7b, 0x36, 0x09, 0x17,
	0x4b, 0x84, 0x1e, 0x51, 0xa2, 0xa5, 0x7f, 0x00, 0xdf, 0xdf, 0x1f, 0x29, 0x85, 0x8a, 0x0f, 0x09,
	0xf5, 0xc0, 0x81, 0x1b, 0x12, 0x30, 0xd0, 0xa2, 0x85, 0x00, 0x07, 0x28, 0x12, 0x52, 0xcb, 0xad,
	0x47, 0x94, 0x5c, 0xf8, 0x16, 0x07, 0xe0, 0x8c, 0x6c, 0x67, 0x66, 0x67, 0xdb, 0xde, 0x66, 0x7f,
	0xbf, 0xe7, 0xf7, 0x7e, 0xef, 0xbd, 0x9f, 0xbd, 0x68, 0x7e, 0xb8, 0xb1, 0xe6, 0xce, 0x0c, 0x87,
	0xeb, 0x43, 0xb7, 0xb6, 0xfe, 0xe8, 0x99, 0xc7, 0xef, 0xdb, 0x18, 0xae, 0x9f, 0x5d, 0x27, 0xb7,
	0x3f, 0x76, 0x36, 0x7c, 0xdc, 0x7b, 0x79, 0x06, 0xcd, 0xda, 0x4c, 0x6f, 0xac, 0x81, 0x8f, 0x49,
	0x7c, 0x08, 0x99, 0x43, 0x33, 0xa0, 0xb5, 0xd2, 0x2e, 0x51, 0x1c, 0x9c, 0x5a, 0xc2, 0x53, 0xe4,
	0x08, 0x3a, 0xd4, 0x17, 0x19, 0xb8, 0x5c, 0x2c, 0x6a, 0x66, 0xc1, 0xad, 0xe0, 0xb7, 0xb7, 0x08,
	0x99, 0x47, 0xb3, 0x45, 0xaa, 0x64, 0x1b, 0x7e, 0x67, 0x8b, 0x90, 0x63, 0x68, 0x4e, 0x82, 0x5d,
	0x56, 0x7a, 0xa9, 0x45, 0xbc, 0xbb, 0x45, 0x7c, 0x96, 0xd2, 0x80, 0x6e, 0xa1, 0xef, 0x05, 0x74,
	0x36, 0x96, 0x33, 0x00, 0x4e, 0xd9, 0x14, 0x34, 0xfe, 0x68, 0x8f, 0x4f, 0xd2, 0x17, 0xda, 0x58,
	0xc9, 0x72, 0x70, 0x42, 0x0e, 0x58, 0x26, 0x38, 0x7e, 0xa2, 0xa2, 0xe4, 0x28, 0xc2, 0x19, 0xbb,
	0x01, 0x7f, 0xb2, 0xa2, 0xe4, 0x38, 0x3a, 0x12, 0xc5, 0xc8, 0x32, 0xef, 0x81, 0x6e, 0xb8, 0xa7,
	0x2a, 0x4a, 0x4e, 0xa0, 0xf9, 0xc8, 0x85, 0x8e, 0x52, 0x66, 0x52, 0x07, 0x79, 0x61, 0x57, 0xf1,
	0xd3, 0x31, 0x61, 0x8b, 0x8c, 0xf8, 0x33, 0x15, 0x25, 0x14, 0x91, 0x36, 0xbe, 0x52, 0x08, 0x0d,
	0x1c, 0x3f, 0x5b, 0x51, 0xdf, 0x07, 0x2b, 0x84, 0x13, 0xbc, 0x29, 0xf2, 0x5c, 0xbb, 0xc8, 0xae,
	0x00, 0x95, 0x24, 0x65, 0x21, 0x80, 0xe3, 0xe7, 0x2b, 0x4a, 0xee, 0x44, 0xc7, 0x26, 0xc8, 0x52,
	0x36, 0xf4, 0x0b, 0x15, 0x25, 0x87, 0xd1, 0x8c, 0x9f, 0x8c, 0x71, 0x56, 0x29, 0xd7, 0x87, 0x65,
	0xfc, 0x62, 0x2c, 0x33, 0x06, 0xf3, 0x32, 0x49, 0xf1, 0x4b, 0x15, 0x25, 0x5d, 0x44, 0xed, 0x6a,
	0xe1, 0x55, 0x49, 0x63, 0x75, 0x99, 0x58, 0x35, 0xee, 0xf5, 0xe5, 0x8a, 0xc6, 0xc1, 0x65, 0xe0,
	0x0a, 0xa6, 0x6d, 0x43, 0xbc, 0x52, 0x51, 0xb2, 0x80, 0x0e, 0x8f, 0x89, 0x15, 0x97, 0x0b, 0x63,
	0x84, 0x5c, 0xc4, 0xaf, 0xc6, 0xd9, 0xe5, 0xfc, 0x41, 0x97, 0xa4, 0x90, 0x2c, 0x99, 0x32, 0x6f,
	0x8e, 0xbd, 0x16, 0xeb, 0x15, 0xa9, 0xb2, 0xaa, 0x06, 0x1d, 0x17, 0x39, 0x48, 0x23, 0x94, 0x34,
	0xf8, 0xf5, 0x38, 0xa6, 0xbe, 0x80, 0x8c, 0xbb, 0x89, 0x8d, 0x6c, 0xc6, 0xc1, 0xb6, 0x98, 0x38,
	0xd8, 0x73, 0x15, 0xf5, 0xb6, 0xc9, 0xcd, 0xa2, 0x5b, 0x66, 0xc2, 0xba, 0x3e, 0x13, 0x19, 0x70,
	0xfc, 0x46, 0x45, 0xc9, 0x3d, 0xe8, 0xa4, 0x97, 0x26, 0x12, 0x51, 0x30, 0x69, 0xdd, 0x00, 0xb4,
	0x2f, 0xe2, 0x54, 0x69, 0x39, 0xb3, 0xc0, 0xf1, 0xf9, 0x9b, 0x76, 0x52, 0x17, 0x7b, 0xb3, 0xa2,
	0xe4, 0x6e, 0x74, 0x22, 0x32, 0x42, 0x0e, 0x84, 0x6d, 0x02, 0x12, 0xa5, 0x35, 0x24, 0x16, 0xbf,
	0x55, 0x51, 0x72, 0x17, 0x5a, 0xb8, 0x55, 0x48, 0xcc, 0x71, 0x21, 0x4e, 0x68, 0x62, 0x49, 0x3d,
	0x26, 0x25, 0x70, 0xfc, 0x7e, 0xdc, 0x9f, 0x01, 0x13, 0x04, 0x15, 0xcc, 0x98, 0x65, 0xa5, 0xb9,
	0x93, 0x00, 0x1c, 0x38, 0xfe, 0xb0, 0xa2, 0x84, 0xa0, 0x83, 0x13, 0xd9, 0x3e, 0xde, 0xf5, 0x43,
	0x1d, 0x1a, 0x2c, 0x57, 0x93, 0x9f, 0xc4, 0xd9, 0x48, 0x58, 0x1e, 0xe7, 0xea, 0x31, 0x8e, 0x3f,
	0x1d, 0xe3, 0x86, 0x65, 0xe3, 0xe5, 0x55, 0xd1, 0x20, 0x90, 0x33, 0x91, 0x35, 0xe0, 0x67, 0x71,
	0xd5, 0x11, 0x2c, 0x65, 0xa2, 0x64, 0x5f, 0xe8, 0x1c, 0x38, 0xfe, 0xfc, 0x86, 0xd2, 0x61, 0xa9,
	0xf5, 0x9c, 0x2f, 0xed, 0x76, 0xd9, 0x2a, 0xeb, 0xfa, 0x4a, 0xe7, 0xcc, 0xe2, 0xcb, 0x91, 0xba,
	0x29, 0xa1, 0x7b, 0x08, 0x7f, 0x11, 0x85, 0x79, 0x33, 0x4e, 0x2c, 0x79, 0x14, 0x35, 0x34, 0x78,
	0x63, 0xe9, 0x2f, 0xa3, 0x86, 0x9a, 0x30, 0xae, 0x94, 0x6c, 0xc0, 0x44, 0xc6, 0x7a, 0x19, 0xe0,
	0xaf, 0xa2, 0x05, 0x92, 0x94, 0xd9, 0xf6, 0x15, 0xfa, 0x2e, 0x26, 0x0b, 0xb0, 0x54, 0xd6, 0xe5,
	0x8a, 0x8b, 0xbe, 0x4f, 0xf6, 0x7d, 0x45, 0xc9, 0x49, 0x74, 0xb4, 0xed, 0x0d, 0xcf, 0xc3, 0x8a,
	0x30, 0xd6, 0xe0, 0x2b, 0x71, 0x39, 0x52, 0x39, 0xe0, 0xc2, 0xba, 0x70, 0xbc, 0x00, 0x1d, 0xdc,
	0xad, 0x24, 0xfe, 0xa1, 0xa2, 0x04, 0xa3, 0x03, 0xbe, 0x4f, 0x0d, 0xa7, 0x4b, 0x30, 0x16, 0x6f,
	0x76, 0xbc, 0x36, 0x56, 0xda, 0xd4, 0x2d, 0xc1, 0xaa, 0x2b, 0xa5, 0x86, 0x45, 0x61, 0x2c, 0xf8,
	0xdb, 0xfd, 0xcb, 0x28, 0x74, 0xda, 0x90, 0xb5, 0xb8, 0x5f, 0x47, 0xcd, 0x04, 0x1c, 0x07, 0x96,
	0x58, 0x31, 0x08, 0x9e, 0xfc, 0x6d, 0x14, 0x7a, 0xa9, 0xad, 0xa1, 0x61, 0xa0, 0x96, 0x80, 0xe3,
	0xdf, 0x27, 0xe1, 0xfa, 0xed, 0xf8, 0x63, 0x14, 0xae, 0x5a, 0x38, 0x0e, 0x2e, 0x24, 0xf3, 0xa2,
	0x02, 0xf7, 0xe7, 0x28, 0x8c, 0xbf, 0xa9, 0xec, 0x7b, 0xd8, 0xbd, 0x33, 0x7f, 0x8d, 0x28, 0x99,
	0x43, 0x07, 0x4b, 0xe9, 0x49, 0xa5, 0xc5, 0xc3, 0xc0, 0xf1, 0xb9, 0x8e, 0xef, 0x3a, 0xa4, 0x28,
	0xb4, 0x18, 0xb0, 0x64, 0xd5, 0x69, 0x30, 0x56, 0x8b, 0xc4, 0xcb, 0xba, 0xf4, 0x23, 0x25, 0x87,
	0xd0, 0xfe, 0xbe, 0xd2, 0x3d, 0xc1, 0x39, 0x48, 0x7c, 0xbe, 0xe3, 0xf5, 0xc4, 0x67, 0x36, 0x53,
	0x09, 0xcb, 0x42, 0x92, 0xbf, 0x77, 0x42, 0xd8, 0x18, 0xb8, 0xd0, 0xf1, 0x4e, 0xee, 0x67, 0x4a,
	0xf1, 0x78, 0x3d, 0x57, 0xf0, 0xc5, 0x9f, 0x16, 0x08, 0x42, 0x7b, 0x03, 0x86, 0x3f, 0xe8, 0x90,
	0x19, 0x74, 0x87, 0x90, 0xd6, 0xef, 0x35, 0xc3, 0xff, 0x84, 0x49, 0xd6, 0x3f, 0x9d, 0x01, 0x3d,
	0x00, 0xed, 0x42, 0x15, 0x7c, 0xf1, 0xdb, 0xae, 0x3f, 0x17, 0xdf, 0xf3, 0x7f, 0x3b, 0xe4, 0x00,
	0xda, 0x17, 0xbe, 0xef, 0xc7, 0xff, 0x75, 0x3c, 0xc1, 0x7b, 0xa0, 0x35, 0xbe, 0x7e, 0x1b, 0x99,
	0x45, 0xfb, 0xc3, 0xb7, 0x33, 0xa7, 0x33, 0xfc, 0xf5, 0xd5, 0x53, 0x04, 0x23, 0x14, 0x81, 0x44,
	0x49, 0x89, 0xbf, 0xb9, 0x7a, 0x8a, 0xcc, 0x23, 0x2c, 0x95, 0xd5, 0x60, 0x4b, 0x2d, 0x5d, 0x92,
	0x09, 0x90, 0x16, 0x8f, 0xf6, 0xf6, 0x8e, 0x5f, 0xd9, 0xee, 0x4e, 0x5f, 0xdb, 0xee, 0x4e, 0xff,
	0xbc, 0xdd, 0x9d, 0xde, 0xdc, 0xe9, 0x4e, 0x5d, 0xdb, 0xe9, 0x4e, 0x5d, 0xdf, 0xe9, 0x4e, 0xa5,
	0x7b, 0x1e, 0xd9, 0x17, 0xfe, 0xde, 0x1e, 0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x07, 0xb7,
	0xa1, 0x00, 0x07, 0x00, 0x00,
}