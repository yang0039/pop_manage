package mtproto

import (
	"reflect"

	"pop-api/baselib/logger"
)

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLAuthLogOut":                         RPCContextTuple{"/mtproto.RPCAuth/auth_logOut", func() interface{} { return new(Bool) }},
	"TLAuthResetAuthorizations":            RPCContextTuple{"/mtproto.RPCAuth/auth_resetAuthorizations", func() interface{} { return new(Bool) }},
	"TLAuthSendInvites":                    RPCContextTuple{"/mtproto.RPCAuth/auth_sendInvites", func() interface{} { return new(Bool) }},
	"TLAuthBindTempAuthKey":                RPCContextTuple{"/mtproto.RPCAuth/auth_bindTempAuthKey", func() interface{} { return new(Bool) }},
	"TLAuthCancelCode":                     RPCContextTuple{"/mtproto.RPCAuth/auth_cancelCode", func() interface{} { return new(Bool) }},
	"TLAuthDropTempAuthKeys":               RPCContextTuple{"/mtproto.RPCAuth/auth_dropTempAuthKeys", func() interface{} { return new(Bool) }},
	"TLAccountRegisterDevice":              RPCContextTuple{"/mtproto.RPCAccount/account_registerDevice", func() interface{} { return new(Bool) }},
	"TLAccountUnregisterDevice":            RPCContextTuple{"/mtproto.RPCAccount/account_unregisterDevice", func() interface{} { return new(Bool) }},
	"TLAccountUpdateNotifySettings":        RPCContextTuple{"/mtproto.RPCAccount/account_updateNotifySettings", func() interface{} { return new(Bool) }},
	"TLAccountResetNotifySettings":         RPCContextTuple{"/mtproto.RPCAccount/account_resetNotifySettings", func() interface{} { return new(Bool) }},
	"TLAccountUpdateStatus":                RPCContextTuple{"/mtproto.RPCAccount/account_updateStatus", func() interface{} { return new(Bool) }},
	"TLAccountReportPeer":                  RPCContextTuple{"/mtproto.RPCAccount/account_reportPeer", func() interface{} { return new(Bool) }},
	"TLAccountCheckUsername":               RPCContextTuple{"/mtproto.RPCAccount/account_checkUsername", func() interface{} { return new(Bool) }},
	"TLAccountDeleteAccount":               RPCContextTuple{"/mtproto.RPCAccount/account_deleteAccount", func() interface{} { return new(Bool) }},
	"TLAccountSetAccountTTL":               RPCContextTuple{"/mtproto.RPCAccount/account_setAccountTTL", func() interface{} { return new(Bool) }},
	"TLAccountUpdateDeviceLocked":          RPCContextTuple{"/mtproto.RPCAccount/account_updateDeviceLocked", func() interface{} { return new(Bool) }},
	"TLAccountResetAuthorization":          RPCContextTuple{"/mtproto.RPCAccount/account_resetAuthorization", func() interface{} { return new(Bool) }},
	"TLAccountUpdatePasswordSettings":      RPCContextTuple{"/mtproto.RPCAccount/account_updatePasswordSettings", func() interface{} { return new(Bool) }},
	"TLAccountConfirmPhone":                RPCContextTuple{"/mtproto.RPCAccount/account_confirmPhone", func() interface{} { return new(Bool) }},
	"TLContactsDeleteContacts":             RPCContextTuple{"/mtproto.RPCContacts/contacts_deleteContacts", func() interface{} { return new(Bool) }},
	"TLContactsBlock":                      RPCContextTuple{"/mtproto.RPCContacts/contacts_block", func() interface{} { return new(Bool) }},
	"TLContactsUnblock":                    RPCContextTuple{"/mtproto.RPCContacts/contacts_unblock", func() interface{} { return new(Bool) }},
	"TLContactsResetTopPeerRating":         RPCContextTuple{"/mtproto.RPCContacts/contacts_resetTopPeerRating", func() interface{} { return new(Bool) }},
	"TLContactsResetSaved":                 RPCContextTuple{"/mtproto.RPCContacts/contacts_resetSaved", func() interface{} { return new(Bool) }},
	"TLMessagesSetTyping":                  RPCContextTuple{"/mtproto.RPCMessages/messages_setTyping", func() interface{} { return new(Bool) }},
	"TLMessagesReportSpam":                 RPCContextTuple{"/mtproto.RPCMessages/messages_reportSpam", func() interface{} { return new(Bool) }},
	"TLMessagesHideReportSpam":             RPCContextTuple{"/mtproto.RPCMessages/messages_hideReportSpam", func() interface{} { return new(Bool) }},
	"TLMessagesDiscardEncryption":          RPCContextTuple{"/mtproto.RPCMessages/messages_discardEncryption", func() interface{} { return new(Bool) }},
	"TLMessagesSetEncryptedTyping":         RPCContextTuple{"/mtproto.RPCMessages/messages_setEncryptedTyping", func() interface{} { return new(Bool) }},
	"TLMessagesReadEncryptedHistory":       RPCContextTuple{"/mtproto.RPCMessages/messages_readEncryptedHistory", func() interface{} { return new(Bool) }},
	"TLMessagesReportEncryptedSpam":        RPCContextTuple{"/mtproto.RPCMessages/messages_reportEncryptedSpam", func() interface{} { return new(Bool) }},
	"TLMessagesUninstallStickerSet":        RPCContextTuple{"/mtproto.RPCMessages/messages_uninstallStickerSet", func() interface{} { return new(Bool) }},
	"TLMessagesEditChatAdmin":              RPCContextTuple{"/mtproto.RPCMessages/messages_editChatAdmin", func() interface{} { return new(Bool) }},
	"TLMessagesReorderStickerSets":         RPCContextTuple{"/mtproto.RPCMessages/messages_reorderStickerSets", func() interface{} { return new(Bool) }},
	"TLMessagesSaveGif":                    RPCContextTuple{"/mtproto.RPCMessages/messages_saveGif", func() interface{} { return new(Bool) }},
	"TLMessagesSetInlineBotResults":        RPCContextTuple{"/mtproto.RPCMessages/messages_setInlineBotResults", func() interface{} { return new(Bool) }},
	"TLMessagesEditInlineBotMessage":       RPCContextTuple{"/mtproto.RPCMessages/messages_editInlineBotMessage", func() interface{} { return new(Bool) }},
	"TLMessagesSetBotCallbackAnswer":       RPCContextTuple{"/mtproto.RPCMessages/messages_setBotCallbackAnswer", func() interface{} { return new(Bool) }},
	"TLMessagesSaveDraft":                  RPCContextTuple{"/mtproto.RPCMessages/messages_saveDraft", func() interface{} { return new(Bool) }},
	"TLMessagesReadFeaturedStickers":       RPCContextTuple{"/mtproto.RPCMessages/messages_readFeaturedStickers", func() interface{} { return new(Bool) }},
	"TLMessagesSaveRecentSticker":          RPCContextTuple{"/mtproto.RPCMessages/messages_saveRecentSticker", func() interface{} { return new(Bool) }},
	"TLMessagesClearRecentStickers":        RPCContextTuple{"/mtproto.RPCMessages/messages_clearRecentStickers", func() interface{} { return new(Bool) }},
	"TLMessagesSetInlineGameScore":         RPCContextTuple{"/mtproto.RPCMessages/messages_setInlineGameScore", func() interface{} { return new(Bool) }},
	"TLMessagesToggleDialogPin":            RPCContextTuple{"/mtproto.RPCMessages/messages_toggleDialogPin", func() interface{} { return new(Bool) }},
	"TLMessagesReorderPinnedDialogs":       RPCContextTuple{"/mtproto.RPCMessages/messages_reorderPinnedDialogs", func() interface{} { return new(Bool) }},
	"TLMessagesSetBotShippingResults":      RPCContextTuple{"/mtproto.RPCMessages/messages_setBotShippingResults", func() interface{} { return new(Bool) }},
	"TLMessagesSetBotPrecheckoutResults":   RPCContextTuple{"/mtproto.RPCMessages/messages_setBotPrecheckoutResults", func() interface{} { return new(Bool) }},
	"TLMessagesFaveSticker":                RPCContextTuple{"/mtproto.RPCMessages/messages_faveSticker", func() interface{} { return new(Bool) }},
	"TLUploadSaveFilePart":                 RPCContextTuple{"/mtproto.RPCUpload/upload_saveFilePart", func() interface{} { return new(Bool) }},
	"TLUploadSaveBigFilePart":              RPCContextTuple{"/mtproto.RPCUpload/upload_saveBigFilePart", func() interface{} { return new(Bool) }},
	"TLHelpSaveAppLog":                     RPCContextTuple{"/mtproto.RPCHelp/help_saveAppLog", func() interface{} { return new(Bool) }},
	"TLHelpSetBotUpdatesStatus":            RPCContextTuple{"/mtproto.RPCHelp/help_setBotUpdatesStatus", func() interface{} { return new(Bool) }},
	"TLChannelsReadHistory":                RPCContextTuple{"/mtproto.RPCChannels/channels_readHistory", func() interface{} { return new(Bool) }},
	"TLChannelsReportSpam":                 RPCContextTuple{"/mtproto.RPCChannels/channels_reportSpam", func() interface{} { return new(Bool) }},
	"TLChannelsEditAbout":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editAbout", func() interface{} { return new(Bool) }},
	"TLChannelsCheckUsername":              RPCContextTuple{"/mtproto.RPCChannels/channels_checkUsername", func() interface{} { return new(Bool) }},
	"TLChannelsUpdateUsername":             RPCContextTuple{"/mtproto.RPCChannels/channels_updateUsername", func() interface{} { return new(Bool) }},
	"TLChannelsSetStickers":                RPCContextTuple{"/mtproto.RPCChannels/channels_setStickers", func() interface{} { return new(Bool) }},
	"TLChannelsReadMessageContents":        RPCContextTuple{"/mtproto.RPCChannels/channels_readMessageContents", func() interface{} { return new(Bool) }},
	"TLBotsAnswerWebhookJSONQuery":         RPCContextTuple{"/mtproto.RPCBots/bots_answerWebhookJSONQuery", func() interface{} { return new(Bool) }},
	"TLPaymentsClearSavedInfo":             RPCContextTuple{"/mtproto.RPCPayments/payments_clearSavedInfo", func() interface{} { return new(Bool) }},
	"TLPhoneReceivedCall":                  RPCContextTuple{"/mtproto.RPCPhone/phone_receivedCall", func() interface{} { return new(Bool) }},
	"TLPhoneSaveCallDebug":                 RPCContextTuple{"/mtproto.RPCPhone/phone_saveCallDebug", func() interface{} { return new(Bool) }},
	"TLAuthCheckPhone":                     RPCContextTuple{"/mtproto.RPCAuth/auth_checkPhone", func() interface{} { return new(Auth_CheckedPhone) }},
	"TLAuthSendCode":                       RPCContextTuple{"/mtproto.RPCAuth/auth_sendCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthResendCode":                     RPCContextTuple{"/mtproto.RPCAuth/auth_resendCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAccountSendChangePhoneCode":         RPCContextTuple{"/mtproto.RPCAccount/account_sendChangePhoneCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAccountSendConfirmPhoneCode":        RPCContextTuple{"/mtproto.RPCAccount/account_sendConfirmPhoneCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthSignUp":                         RPCContextTuple{"/mtproto.RPCAuth/auth_signUp", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthSignIn":                         RPCContextTuple{"/mtproto.RPCAuth/auth_signIn", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthImportAuthorization":            RPCContextTuple{"/mtproto.RPCAuth/auth_importAuthorization", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthImportBotAuthorization":         RPCContextTuple{"/mtproto.RPCAuth/auth_importBotAuthorization", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthCheckPassword":                  RPCContextTuple{"/mtproto.RPCAuth/auth_checkPassword", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthRecoverPassword":                RPCContextTuple{"/mtproto.RPCAuth/auth_recoverPassword", func() interface{} { return new(Auth_Authorization) }},
	"TLAuthExportAuthorization":            RPCContextTuple{"/mtproto.RPCAuth/auth_exportAuthorization", func() interface{} { return new(Auth_ExportedAuthorization) }},
	"TLAuthRequestPasswordRecovery":        RPCContextTuple{"/mtproto.RPCAuth/auth_requestPasswordRecovery", func() interface{} { return new(Auth_PasswordRecovery) }},
	"TLAccountGetNotifySettings":           RPCContextTuple{"/mtproto.RPCAccount/account_getNotifySettings", func() interface{} { return new(PeerNotifySettings) }},
	"TLAccountUpdateProfile":               RPCContextTuple{"/mtproto.RPCAccount/account_updateProfile", func() interface{} { return new(User) }},
	"TLAccountUpdateUsername":              RPCContextTuple{"/mtproto.RPCAccount/account_updateUsername", func() interface{} { return new(User) }},
	"TLAccountChangePhone":                 RPCContextTuple{"/mtproto.RPCAccount/account_changePhone", func() interface{} { return new(User) }},
	"TLContactsImportCard":                 RPCContextTuple{"/mtproto.RPCContacts/contacts_importCard", func() interface{} { return new(User) }},
	"TLAccountGetPrivacy":                  RPCContextTuple{"/mtproto.RPCAccount/account_getPrivacy", func() interface{} { return new(Account_PrivacyRules) }},
	"TLAccountSetPrivacy":                  RPCContextTuple{"/mtproto.RPCAccount/account_setPrivacy", func() interface{} { return new(Account_PrivacyRules) }},
	"TLAccountGetAccountTTL":               RPCContextTuple{"/mtproto.RPCAccount/account_getAccountTTL", func() interface{} { return new(AccountDaysTTL) }},
	"TLAccountGetAuthorizations":           RPCContextTuple{"/mtproto.RPCAccount/account_getAuthorizations", func() interface{} { return new(Account_Authorizations) }},
	"TLAccountGetPassword":                 RPCContextTuple{"/mtproto.RPCAccount/account_getPassword", func() interface{} { return new(Account_Password) }},
	"TLAccountGetPasswordSettings":         RPCContextTuple{"/mtproto.RPCAccount/account_getPasswordSettings", func() interface{} { return new(Account_PasswordSettings) }},
	"TLAccountGetTmpPassword":              RPCContextTuple{"/mtproto.RPCAccount/account_getTmpPassword", func() interface{} { return new(Account_TmpPassword) }},
	"TLUsersGetFullUser":                   RPCContextTuple{"/mtproto.RPCUsers/users_getFullUser", func() interface{} { return new(UserFull) }},
	"TLContactsGetContacts":                RPCContextTuple{"/mtproto.RPCContacts/contacts_getContacts", func() interface{} { return new(Contacts_Contacts) }},
	"TLContactsGetContactsLayer70":         RPCContextTuple{"/mtproto.RPCContacts/contacts_getContactsLayer70", func() interface{} { return new(Contacts_Contacts) }},
	"TLContactsImportContacts":             RPCContextTuple{"/mtproto.RPCContacts/contacts_importContacts", func() interface{} { return new(Contacts_ImportedContacts) }},
	"TLContactsDeleteContact":              RPCContextTuple{"/mtproto.RPCContacts/contacts_deleteContact", func() interface{} { return new(Contacts_Link) }},
	"TLContactsGetBlocked":                 RPCContextTuple{"/mtproto.RPCContacts/contacts_getBlocked", func() interface{} { return new(Contacts_Blocked) }},
	"TLContactsSearch":                     RPCContextTuple{"/mtproto.RPCContacts/contacts_search", func() interface{} { return new(Contacts_Found) }},
	"TLContactsResolveUsername":            RPCContextTuple{"/mtproto.RPCContacts/contacts_resolveUsername", func() interface{} { return new(Contacts_ResolvedPeer) }},
	"TLContactsGetTopPeers":                RPCContextTuple{"/mtproto.RPCContacts/contacts_getTopPeers", func() interface{} { return new(Contacts_TopPeers) }},
	"TLMessagesGetMessages":                RPCContextTuple{"/mtproto.RPCMessages/messages_getMessages", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetHistory":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getHistory", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSearch":                     RPCContextTuple{"/mtproto.RPCMessages/messages_search", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSearchGlobal":               RPCContextTuple{"/mtproto.RPCMessages/messages_searchGlobal", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetUnreadMentions":          RPCContextTuple{"/mtproto.RPCMessages/messages_getUnreadMentions", func() interface{} { return new(Messages_Messages) }},
	"TLChannelsGetMessages":                RPCContextTuple{"/mtproto.RPCChannels/channels_getMessages", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetDialogs":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getDialogs", func() interface{} { return new(Messages_Dialogs) }},
	"TLMessagesReadHistory":                RPCContextTuple{"/mtproto.RPCMessages/messages_readHistory", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesDeleteMessages":             RPCContextTuple{"/mtproto.RPCMessages/messages_deleteMessages", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesReadMessageContents":        RPCContextTuple{"/mtproto.RPCMessages/messages_readMessageContents", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLChannelsDeleteMessages":             RPCContextTuple{"/mtproto.RPCChannels/channels_deleteMessages", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesDeleteHistory":              RPCContextTuple{"/mtproto.RPCMessages/messages_deleteHistory", func() interface{} { return new(Messages_AffectedHistory) }},
	"TLChannelsDeleteUserHistory":          RPCContextTuple{"/mtproto.RPCChannels/channels_deleteUserHistory", func() interface{} { return new(Messages_AffectedHistory) }},
	"TLMessagesSendMessage":                RPCContextTuple{"/mtproto.RPCMessages/messages_sendMessage", func() interface{} { return new(Updates) }},
	"TLMessagesSendMedia":                  RPCContextTuple{"/mtproto.RPCMessages/messages_sendMedia", func() interface{} { return new(Updates) }},
	"TLMessagesForwardMessages":            RPCContextTuple{"/mtproto.RPCMessages/messages_forwardMessages", func() interface{} { return new(Updates) }},
	"TLMessagesEditChatTitle":              RPCContextTuple{"/mtproto.RPCMessages/messages_editChatTitle", func() interface{} { return new(Updates) }},
	"TLMessagesEditChatPhoto":              RPCContextTuple{"/mtproto.RPCMessages/messages_editChatPhoto", func() interface{} { return new(Updates) }},
	"TLMessagesAddChatUser":                RPCContextTuple{"/mtproto.RPCMessages/messages_addChatUser", func() interface{} { return new(Updates) }},
	"TLMessagesDeleteChatUser":             RPCContextTuple{"/mtproto.RPCMessages/messages_deleteChatUser", func() interface{} { return new(Updates) }},
	"TLMessagesCreateChat":                 RPCContextTuple{"/mtproto.RPCMessages/messages_createChat", func() interface{} { return new(Updates) }},
	"TLMessagesForwardMessage":             RPCContextTuple{"/mtproto.RPCMessages/messages_forwardMessage", func() interface{} { return new(Updates) }},
	"TLMessagesImportChatInvite":           RPCContextTuple{"/mtproto.RPCMessages/messages_importChatInvite", func() interface{} { return new(Updates) }},
	"TLMessagesStartBot":                   RPCContextTuple{"/mtproto.RPCMessages/messages_startBot", func() interface{} { return new(Updates) }},
	"TLMessagesToggleChatAdmins":           RPCContextTuple{"/mtproto.RPCMessages/messages_toggleChatAdmins", func() interface{} { return new(Updates) }},
	"TLMessagesMigrateChat":                RPCContextTuple{"/mtproto.RPCMessages/messages_migrateChat", func() interface{} { return new(Updates) }},
	"TLMessagesSendInlineBotResult":        RPCContextTuple{"/mtproto.RPCMessages/messages_sendInlineBotResult", func() interface{} { return new(Updates) }},
	"TLMessagesEditMessage":                RPCContextTuple{"/mtproto.RPCMessages/messages_editMessage", func() interface{} { return new(Updates) }},
	"TLMessagesGetAllDrafts":               RPCContextTuple{"/mtproto.RPCMessages/messages_getAllDrafts", func() interface{} { return new(Updates) }},
	"TLMessagesSetGameScore":               RPCContextTuple{"/mtproto.RPCMessages/messages_setGameScore", func() interface{} { return new(Updates) }},
	"TLMessagesSendScreenshotNotification": RPCContextTuple{"/mtproto.RPCMessages/messages_sendScreenshotNotification", func() interface{} { return new(Updates) }},
	"TLHelpGetAppChangelog":                RPCContextTuple{"/mtproto.RPCHelp/help_getAppChangelog", func() interface{} { return new(Updates) }},
	"TLChannelsCreateChannel":              RPCContextTuple{"/mtproto.RPCChannels/channels_createChannel", func() interface{} { return new(Updates) }},
	"TLChannelsEditAdmin":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editAdmin", func() interface{} { return new(Updates) }},
	"TLChannelsEditTitle":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editTitle", func() interface{} { return new(Updates) }},
	"TLChannelsEditPhoto":                  RPCContextTuple{"/mtproto.RPCChannels/channels_editPhoto", func() interface{} { return new(Updates) }},
	"TLChannelsJoinChannel":                RPCContextTuple{"/mtproto.RPCChannels/channels_joinChannel", func() interface{} { return new(Updates) }},
	"TLChannelsLeaveChannel":               RPCContextTuple{"/mtproto.RPCChannels/channels_leaveChannel", func() interface{} { return new(Updates) }},
	"TLChannelsInviteToChannel":            RPCContextTuple{"/mtproto.RPCChannels/channels_inviteToChannel", func() interface{} { return new(Updates) }},
	"TLChannelsDeleteChannel":              RPCContextTuple{"/mtproto.RPCChannels/channels_deleteChannel", func() interface{} { return new(Updates) }},
	"TLChannelsToggleInvites":              RPCContextTuple{"/mtproto.RPCChannels/channels_toggleInvites", func() interface{} { return new(Updates) }},
	"TLChannelsToggleSignatures":           RPCContextTuple{"/mtproto.RPCChannels/channels_toggleSignatures", func() interface{} { return new(Updates) }},
	"TLChannelsUpdatePinnedMessage":        RPCContextTuple{"/mtproto.RPCChannels/channels_updatePinnedMessage", func() interface{} { return new(Updates) }},
	"TLChannelsEditBanned":                 RPCContextTuple{"/mtproto.RPCChannels/channels_editBanned", func() interface{} { return new(Updates) }},

	"TLChannelsEditBanned2": RPCContextTuple{"/mtproto.RPCChannels/channels_editBanned2", func() interface{} { return new(Updates) }},
	"TLChannelsGetBanned2":  RPCContextTuple{"/mtproto.RPCChannels/channels_getBanned2", func() interface{} { return new(Updates) }},

	"TLPhoneDiscardCall":                 RPCContextTuple{"/mtproto.RPCPhone/phone_discardCall", func() interface{} { return new(Updates) }},
	"TLPhoneSetCallRating":               RPCContextTuple{"/mtproto.RPCPhone/phone_setCallRating", func() interface{} { return new(Updates) }},
	"TLMessagesGetPeerSettings":          RPCContextTuple{"/mtproto.RPCMessages/messages_getPeerSettings", func() interface{} { return new(PeerSettings) }},
	"TLMessagesGetChats":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getChats", func() interface{} { return new(Messages_Chats) }},
	"TLMessagesGetCommonChats":           RPCContextTuple{"/mtproto.RPCMessages/messages_getCommonChats", func() interface{} { return new(Messages_Chats) }},
	"TLMessagesGetAllChats":              RPCContextTuple{"/mtproto.RPCMessages/messages_getAllChats", func() interface{} { return new(Messages_Chats) }},
	"TLChannelsGetChannels":              RPCContextTuple{"/mtproto.RPCChannels/channels_getChannels", func() interface{} { return new(Messages_Chats) }},
	"TLChannelsGetAdminedPublicChannels": RPCContextTuple{"/mtproto.RPCChannels/channels_getAdminedPublicChannels", func() interface{} { return new(Messages_Chats) }},
	"TLMessagesGetFullChat":              RPCContextTuple{"/mtproto.RPCMessages/messages_getFullChat", func() interface{} { return new(Messages_ChatFull) }},
	"TLChannelsGetFullChannel":           RPCContextTuple{"/mtproto.RPCChannels/channels_getFullChannel", func() interface{} { return new(Messages_ChatFull) }},
	"TLMessagesGetDhConfig":              RPCContextTuple{"/mtproto.RPCMessages/messages_getDhConfig", func() interface{} { return new(Messages_DhConfig) }},
	"TLMessagesRequestEncryption":        RPCContextTuple{"/mtproto.RPCMessages/messages_requestEncryption", func() interface{} { return new(EncryptedChat) }},
	"TLMessagesAcceptEncryption":         RPCContextTuple{"/mtproto.RPCMessages/messages_acceptEncryption", func() interface{} { return new(EncryptedChat) }},
	"TLMessagesSendEncrypted":            RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncrypted", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesSendEncryptedFile":        RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncryptedFile", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesSendEncryptedService":     RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncryptedService", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesGetAllStickers":           RPCContextTuple{"/mtproto.RPCMessages/messages_getAllStickers", func() interface{} { return new(Messages_AllStickers) }},
	"TLMessagesGetMaskStickers":          RPCContextTuple{"/mtproto.RPCMessages/messages_getMaskStickers", func() interface{} { return new(Messages_AllStickers) }},
	"TLMessagesGetWebPagePreview":        RPCContextTuple{"/mtproto.RPCMessages/messages_getWebPagePreview", func() interface{} { return new(MessageMedia) }},
	"TLMessagesUploadMedia":              RPCContextTuple{"/mtproto.RPCMessages/messages_uploadMedia", func() interface{} { return new(MessageMedia) }},
	"TLMessagesExportChatInvite":         RPCContextTuple{"/mtproto.RPCMessages/messages_exportChatInvite", func() interface{} { return new(ExportedChatInvite) }},
	"TLChannelsExportInvite":             RPCContextTuple{"/mtproto.RPCChannels/channels_exportInvite", func() interface{} { return new(ExportedChatInvite) }},
	"TLMessagesCheckChatInvite":          RPCContextTuple{"/mtproto.RPCMessages/messages_checkChatInvite", func() interface{} { return new(ChatInvite) }},
	"TLMessagesGetStickerSet":            RPCContextTuple{"/mtproto.RPCMessages/messages_getStickerSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersCreateStickerSet":         RPCContextTuple{"/mtproto.RPCStickers/stickers_createStickerSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersRemoveStickerFromSet":     RPCContextTuple{"/mtproto.RPCStickers/stickers_removeStickerFromSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersChangeStickerPosition":    RPCContextTuple{"/mtproto.RPCStickers/stickers_changeStickerPosition", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersAddStickerToSet":          RPCContextTuple{"/mtproto.RPCStickers/stickers_addStickerToSet", func() interface{} { return new(Messages_StickerSet) }},
	"TLMessagesInstallStickerSet":        RPCContextTuple{"/mtproto.RPCMessages/messages_installStickerSet", func() interface{} { return new(Messages_StickerSetInstallResult) }},
	"TLMessagesGetDocumentByHash":        RPCContextTuple{"/mtproto.RPCMessages/messages_getDocumentByHash", func() interface{} { return new(Document) }},
	"TLMessagesSearchGifs":               RPCContextTuple{"/mtproto.RPCMessages/messages_searchGifs", func() interface{} { return new(Messages_FoundGifs) }},
	"TLMessagesGetSavedGifs":             RPCContextTuple{"/mtproto.RPCMessages/messages_getSavedGifs", func() interface{} { return new(Messages_SavedGifs) }},
	"TLMessagesGetInlineBotResults":      RPCContextTuple{"/mtproto.RPCMessages/messages_getInlineBotResults", func() interface{} { return new(Messages_BotResults) }},
	"TLMessagesGetMessageEditData":       RPCContextTuple{"/mtproto.RPCMessages/messages_getMessageEditData", func() interface{} { return new(Messages_MessageEditData) }},
	"TLMessagesGetBotCallbackAnswer":     RPCContextTuple{"/mtproto.RPCMessages/messages_getBotCallbackAnswer", func() interface{} { return new(Messages_BotCallbackAnswer) }},
	"TLMessagesGetPeerDialogs":           RPCContextTuple{"/mtproto.RPCMessages/messages_getPeerDialogs", func() interface{} { return new(Messages_PeerDialogs) }},
	"TLMessagesGetPinnedDialogs":         RPCContextTuple{"/mtproto.RPCMessages/messages_getPinnedDialogs", func() interface{} { return new(Messages_PeerDialogs) }},
	"TLMessagesGetFeaturedStickers":      RPCContextTuple{"/mtproto.RPCMessages/messages_getFeaturedStickers", func() interface{} { return new(Messages_FeaturedStickers) }},
	"TLMessagesGetRecentStickers":        RPCContextTuple{"/mtproto.RPCMessages/messages_getRecentStickers", func() interface{} { return new(Messages_RecentStickers) }},
	"TLMessagesGetArchivedStickers":      RPCContextTuple{"/mtproto.RPCMessages/messages_getArchivedStickers", func() interface{} { return new(Messages_ArchivedStickers) }},
	"TLMessagesGetGameHighScores":        RPCContextTuple{"/mtproto.RPCMessages/messages_getGameHighScores", func() interface{} { return new(Messages_HighScores) }},
	"TLMessagesGetInlineGameHighScores":  RPCContextTuple{"/mtproto.RPCMessages/messages_getInlineGameHighScores", func() interface{} { return new(Messages_HighScores) }},
	"TLMessagesGetWebPage":               RPCContextTuple{"/mtproto.RPCMessages/messages_getWebPage", func() interface{} { return new(WebPage) }},
	"TLMessagesGetFavedStickers":         RPCContextTuple{"/mtproto.RPCMessages/messages_getFavedStickers", func() interface{} { return new(Messages_FavedStickers) }},
	"TLUpdatesGetState":                  RPCContextTuple{"/mtproto.RPCUpdates/updates_getState", func() interface{} { return new(Updates_State) }},
	"TLUpdatesGetDifference":             RPCContextTuple{"/mtproto.RPCUpdates/updates_getDifference", func() interface{} { return new(Updates_Difference) }},
	"TLUpdatesGetChannelDifference":      RPCContextTuple{"/mtproto.RPCUpdates/updates_getChannelDifference", func() interface{} { return new(Updates_ChannelDifference) }},
	"TLUpdatesGetChannelDifference57":    RPCContextTuple{"/mtproto.RPCUpdates/updates_getChannelDifference57", func() interface{} { return new(Updates_ChannelDifference) }},
	"TLPhotosUpdateProfilePhoto":         RPCContextTuple{"/mtproto.RPCPhotos/photos_updateProfilePhoto", func() interface{} { return new(UserProfilePhoto) }},
	"TLPhotosUploadProfilePhoto":         RPCContextTuple{"/mtproto.RPCPhotos/photos_uploadProfilePhoto", func() interface{} { return new(Photos_Photo) }},
	"TLPhotosGetUserPhotos":              RPCContextTuple{"/mtproto.RPCPhotos/photos_getUserPhotos", func() interface{} { return new(Photos_Photos) }},
	"TLUploadGetFile":                    RPCContextTuple{"/mtproto.RPCUpload/upload_getFile", func() interface{} { return new(Upload_File) }},
	"TLUploadGetWebFile":                 RPCContextTuple{"/mtproto.RPCUpload/upload_getWebFile", func() interface{} { return new(Upload_WebFile) }},
	"TLUploadGetCdnFile":                 RPCContextTuple{"/mtproto.RPCUpload/upload_getCdnFile", func() interface{} { return new(Upload_CdnFile) }},
	"TLHelpGetConfig":                    RPCContextTuple{"/mtproto.RPCHelp/help_getConfig", func() interface{} { return new(Config) }},

	"TLHelpGetConfig73":   RPCContextTuple{"/mtproto.RPCHelp/help_getConfig73", func() interface{} { return new(Config73) }},
	"TLHelpGetConfig82":   RPCContextTuple{"/mtproto.RPCHelp/help_getConfig82", func() interface{} { return new(Config82) }},
	"TLHelpGetWkConfig":   RPCContextTuple{"/mtproto.RPCHelp/help_getWkConfig", func() interface{} { return new(WkConfig) }},
	"TLHelpGetJsonConfig": RPCContextTuple{"/mtproto.RPCHelp/help_getJsonConfig", func() interface{} { return new(DataJSON) }},

	"TLMessagesSearch73":    RPCContextTuple{"/mtproto.RPCMessages/messages_search73", func() interface{} { return new(Messages_Messages) }},
	"TLAuthSendCode73":      RPCContextTuple{"/mtproto.RPCAuth/auth_sendCode73", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthSendCode82":      RPCContextTuple{"/mtproto.RPCAuth/auth_sendCode82", func() interface{} { return new(Auth_SentCode82) }},
	"TLAuthCheckInviteCode": RPCContextTuple{"/mtproto.RPCAuth/auth_checkInviteCode", func() interface{} { return new(Error) }},

	"TLAccountRegisterDevice73":     RPCContextTuple{"/mtproto.RPCAccount/account_registerDevice73", func() interface{} { return new(Bool) }},
	"TLHelpGetAppUpdate73":          RPCContextTuple{"/mtproto.RPCHelp/help_getAppUpdate73", func() interface{} { return new(Help_AppUpdate) }},
	"TLHelpGetInviteText73":         RPCContextTuple{"/mtproto.RPCHelp/help_getInviteText73", func() interface{} { return new(Help_InviteText) }},
	"TLPhotosUpdateProfilePhoto73":  RPCContextTuple{"/mtproto.RPCPhotos/photos_updateProfilePhoto73", func() interface{} { return new(UserProfilePhoto) }},
	"TLPhotosUploadProfilePhoto73":  RPCContextTuple{"/mtproto.RPCPhotos/photos_uploadProfilePhoto73", func() interface{} { return new(Photos_Photo) }},
	"TLMessagesSaveRecentSticker73": RPCContextTuple{"/mtproto.RPCMessages/messages_saveRecentSticker73", func() interface{} { return new(Bool) }},
	"TLMessagesReadHistory73":       RPCContextTuple{"/mtproto.RPCMessages/messages_readHistory73", func() interface{} { return new(Messages_AffectedMessages) }},
	"TLMessagesEditMessage73":       RPCContextTuple{"/mtproto.RPCMessages/messages_editMessage73", func() interface{} { return new(Updates) }},
	"TLHelpGetTermsOfService73":     RPCContextTuple{"/mtproto.RPCHelp/help_getTermsOfService73", func() interface{} { return new(Help_TermsOfService) }},

	"TLChannelsGetParticipants73": RPCContextTuple{"/mtproto.RPCChannels/channels_getParticipants73", func() interface{} { return new(Channels_ChannelParticipants) }},

	"TLHelpGetNearestDc":              RPCContextTuple{"/mtproto.RPCHelp/help_getNearestDc", func() interface{} { return new(NearestDc) }},
	"TLHelpGetAppUpdate":              RPCContextTuple{"/mtproto.RPCHelp/help_getAppUpdate", func() interface{} { return new(Help_AppUpdate) }},
	"TLHelpGetInviteText":             RPCContextTuple{"/mtproto.RPCHelp/help_getInviteText", func() interface{} { return new(Help_InviteText) }},
	"TLHelpGetSupport":                RPCContextTuple{"/mtproto.RPCHelp/help_getSupport", func() interface{} { return new(Help_Support) }},
	"TLHelpGetTermsOfService":         RPCContextTuple{"/mtproto.RPCHelp/help_getTermsOfService", func() interface{} { return new(Help_TermsOfService) }},
	"TLHelpGetCdnConfig":              RPCContextTuple{"/mtproto.RPCHelp/help_getCdnConfig", func() interface{} { return new(CdnConfig) }},
	"TLChannelsGetParticipants":       RPCContextTuple{"/mtproto.RPCChannels/channels_getParticipants", func() interface{} { return new(Channels_ChannelParticipants) }},
	"TLChannelsGetParticipant":        RPCContextTuple{"/mtproto.RPCChannels/channels_getParticipant", func() interface{} { return new(Channels_ChannelParticipant) }},
	"TLChannelsExportMessageLink":     RPCContextTuple{"/mtproto.RPCChannels/channels_exportMessageLink", func() interface{} { return new(ExportedMessageLink) }},
	"TLChannelsGetAdminLog":           RPCContextTuple{"/mtproto.RPCChannels/channels_getAdminLog", func() interface{} { return new(Channels_AdminLogResults) }},
	"TLBotsSendCustomRequest":         RPCContextTuple{"/mtproto.RPCBots/bots_sendCustomRequest", func() interface{} { return new(DataJSON) }},
	"TLPhoneGetCallConfig":            RPCContextTuple{"/mtproto.RPCPhone/phone_getCallConfig", func() interface{} { return new(DataJSON) }},
	"TLPaymentsGetPaymentForm":        RPCContextTuple{"/mtproto.RPCPayments/payments_getPaymentForm", func() interface{} { return new(Payments_PaymentForm) }},
	"TLPaymentsGetPaymentReceipt":     RPCContextTuple{"/mtproto.RPCPayments/payments_getPaymentReceipt", func() interface{} { return new(Payments_PaymentReceipt) }},
	"TLPaymentsValidateRequestedInfo": RPCContextTuple{"/mtproto.RPCPayments/payments_validateRequestedInfo", func() interface{} { return new(Payments_ValidatedRequestedInfo) }},
	"TLPaymentsSendPaymentForm":       RPCContextTuple{"/mtproto.RPCPayments/payments_sendPaymentForm", func() interface{} { return new(Payments_PaymentResult) }},
	"TLPaymentsGetSavedInfo":          RPCContextTuple{"/mtproto.RPCPayments/payments_getSavedInfo", func() interface{} { return new(Payments_SavedInfo) }},
	"TLPhoneRequestCall":              RPCContextTuple{"/mtproto.RPCPhone/phone_requestCall", func() interface{} { return new(Phone_PhoneCall) }},
	"TLPhoneAcceptCall":               RPCContextTuple{"/mtproto.RPCPhone/phone_acceptCall", func() interface{} { return new(Phone_PhoneCall) }},
	"TLPhoneConfirmCall":              RPCContextTuple{"/mtproto.RPCPhone/phone_confirmCall", func() interface{} { return new(Phone_PhoneCall) }},

	"TLPhone_WebrtcCreateOffer":  RPCContextTuple{"/mtproto.RPCPhone/phone_WebrtcCreateOffer", func() interface{} { return new(Bool) }},
	"TLPhone_WebrtcCreateAnswer": RPCContextTuple{"/mtproto.RPCPhone/phone_WebrtcCreateAnswer", func() interface{} { return new(Bool) }},
	"TLPhone_WebrtcAddCandidate": RPCContextTuple{"/mtproto.RPCPhone/phone_WebrtcAddCandidate", func() interface{} { return new(Bool) }},

	"TLLangpackGetLangPack":         RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLangPack", func() interface{} { return new(LangPackDifference) }},
	"TLLangpackGetDifference":       RPCContextTuple{"/mtproto.RPCLangpack/langpack_getDifference", func() interface{} { return new(LangPackDifference) }},
	"TLAccountGetWallPapers":        RPCContextTuple{"/mtproto.RPCAccount/account_getWallPapers", func() interface{} { return new(Vector_WallPaper) }},
	"TLUsersGetUsers":               RPCContextTuple{"/mtproto.RPCUsers/users_getUsers", func() interface{} { return new(Vector_User) }},
	"TLContactsGetStatuses":         RPCContextTuple{"/mtproto.RPCContacts/contacts_getStatuses", func() interface{} { return new(Vector_ContactStatus) }},
	"TLContactsExportCard":          RPCContextTuple{"/mtproto.RPCContacts/contacts_exportCard", func() interface{} { return new(VectorInt) }},
	"TLMessagesGetMessagesViews":    RPCContextTuple{"/mtproto.RPCMessages/messages_getMessagesViews", func() interface{} { return new(VectorInt) }},
	"TLMessagesReceivedMessages":    RPCContextTuple{"/mtproto.RPCMessages/messages_receivedMessages", func() interface{} { return new(Vector_ReceivedNotifyMessage) }},
	"TLMessagesReceivedQueue":       RPCContextTuple{"/mtproto.RPCMessages/messages_receivedQueue", func() interface{} { return new(VectorLong) }},
	"TLPhotosDeletePhotos":          RPCContextTuple{"/mtproto.RPCPhotos/photos_deletePhotos", func() interface{} { return new(VectorLong) }},
	"TLMessagesGetAttachedStickers": RPCContextTuple{"/mtproto.RPCMessages/messages_getAttachedStickers", func() interface{} { return new(Vector_StickerSetCovered) }},
	"TLUploadReuploadCdnFile":       RPCContextTuple{"/mtproto.RPCUpload/upload_reuploadCdnFile", func() interface{} { return new(Vector_CdnFileHash) }},
	"TLUploadGetCdnFileHashes":      RPCContextTuple{"/mtproto.RPCUpload/upload_getCdnFileHashes", func() interface{} { return new(Vector_CdnFileHash) }},
	"TLLangpackGetStrings":          RPCContextTuple{"/mtproto.RPCLangpack/langpack_getStrings", func() interface{} { return new(Vector_LangPackString) }},
	"TLLangpackGetLanguages":        RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLanguages", func() interface{} { return new(Vector_LangPackLanguage) }},

	"TLHelpGetScheme":              RPCContextTuple{"/mtproto.RPCHelp/help_getScheme", func() interface{} { return new(Scheme) }},
	"TLMessagesReadMentions":       RPCContextTuple{"/mtproto.RPCMessages/messages_readMentions", func() interface{} { return new(Messages_AffectedHistory) }},
	"TLMessagesGetRecentLocations": RPCContextTuple{"/mtproto.RPCMessages/messages_getRecentLocations", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSendMultiMedia":     RPCContextTuple{"/mtproto.RPCMessages/messages_sendMultiMedia", func() interface{} { return new(Updates) }},
	"TLHelpGetRecentMeUrls":        RPCContextTuple{"/mtproto.RPCHelp/help_getRecentMeUrls", func() interface{} { return new(Help_RecentMeUrls) }},

	"TLChannelsDeleteHistory":          RPCContextTuple{"/mtproto.RPCChannels/channels_deleteHistory", func() interface{} { return new(Bool) }},
	"TLChannelsTogglePreHistoryHidden": RPCContextTuple{"/mtproto.RPCChannels/channels_togglePreHistoryHidden", func() interface{} { return new(Updates) }},

	"TLApiAddAuthKey":         RPCContextTuple{"/mtproto.RPCApi/api_addAuthKey", func() interface{} { return new(Null) }},
	"TLApiInitConnectionLite": RPCContextTuple{"/mtproto.RPCApi/api_initConnectionLite", func() interface{} { return new(Null) }},
	"TLApiCoinRequest":        RPCContextTuple{"/mtproto.RPCApi/api_coinRequest", func() interface{} { return new(CoinMessage) }},

	"TLMessagesGetDialogs82":     RPCContextTuple{"/mtproto.RPCMessages/messages_getDialogs82", func() interface{} { return new(Messages_Dialogs) }},
	"TLMessagesGetHistory82":     RPCContextTuple{"/mtproto.RPCMessages/messages_getHistory82", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetMessages82":    RPCContextTuple{"/mtproto.RPCMessages/messages_getMessages82", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesSearch82":         RPCContextTuple{"/mtproto.RPCMessages/messages_search82", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetPeerDialogs82": RPCContextTuple{"/mtproto.RPCMessages/messages_getPeerDialogs82", func() interface{} { return new(Messages_PeerDialogs) }},
	"TLMessagesSendMedia82":      RPCContextTuple{"/mtproto.RPCMessages/messages_sendMedia82", func() interface{} { return new(Updates) }},

	"TLHelpGetTermsOfServiceUpdate": RPCContextTuple{"/mtproto.RPCHelp/help_getTermsOfServiceUpdate", func() interface{} { return new(Help_TermsOfServiceUpdate) }},
	"TLHelpGetDeepLinkInfo":         RPCContextTuple{"/mtproto.RPCHelp/help_getDeepLinkInfo", func() interface{} { return new(Help_DeepLinkInfo) }},
	"TLHelpAcceptTermsOfService":    RPCContextTuple{"/mtproto.RPCHelp/help_acceptTermsOfService", func() interface{} { return new(Bool) }},
	"TLHelpTest":                    RPCContextTuple{"/mtproto.RPCHelp/help_test", func() interface{} { return new(Bool) }},

	"TLUsersGetPointsHistory": RPCContextTuple{"/mtproto.RPCUsers/users_getPointsHistory", func() interface{} { return new(PointsHistoryResult) }},
	"TLUsersGetEnvelopeInfo":  RPCContextTuple{"/mtproto.RPCUsers/users_getEnvelopeInfo", func() interface{} { return new(EnvelopeInfo) }},
	"TLUsersGetEnvelopeState": RPCContextTuple{"/mtproto.RPCUsers/users_getEnvelopeState", func() interface{} { return new(EnvelopeState) }},
	"TLUsersGetPointsRank":    RPCContextTuple{"/mtproto.RPCUsers/users_getPointsRank", func() interface{} { return new(PointsRankResult) }},
	"TLUsersGetUserInfo":      RPCContextTuple{"/mtproto.RPCUsers/users_getUserInfo", func() interface{} { return new(UserInfo) }},

	"TLUsersCheckPasswd":               RPCContextTuple{"/mtproto.RPCUsers/users_checkPasswd", func() interface{} { return new(IntResult) }},
	"TLUsersSetPasswd":                 RPCContextTuple{"/mtproto.RPCUsers/users_setPasswd", func() interface{} { return new(IntResult) }},
	"TLUsersSendResetPasswdPhoneCode":  RPCContextTuple{"/mtproto.RPCUsers/users_sendResetPasswdPhoneCode", func() interface{} { return new(IntResult) }},
	"TLUsersCheckResetPasswdPhoneCode": RPCContextTuple{"/mtproto.RPCUsers/users_checkResetPasswdPhoneCode", func() interface{} { return new(IntResult) }},
	"TLUsersSetGender":                 RPCContextTuple{"/mtproto.RPCUsers/users_setGender", func() interface{} { return new(IntResult) }},
	"TLUsersSetUserInfo":               RPCContextTuple{"/mtproto.RPCUsers/users_setUserInfo", func() interface{} { return new(IntResult) }},

	"TLContactsApplyFriend":  RPCContextTuple{"/mtproto.RPCContacts/contacts_applyFriend", func() interface{} { return new(IntResult) }},
	"TLContactsAuditApply":   RPCContextTuple{"/mtproto.RPCContacts/contacts_auditApply", func() interface{} { return new(IntResult) }},
	"TLContactsGetApplyList": RPCContextTuple{"/mtproto.RPCContacts/contacts_getApplyList", func() interface{} { return new(DataJSON) }},
	"TLContactsSearchUser":   RPCContextTuple{"/mtproto.RPCContacts/contacts_searchUser", func() interface{} { return new(DataJSON) }},
	"TLContactsSearchSource": RPCContextTuple{"/mtproto.RPCContacts/contacts_searchSource", func() interface{} { return new(DataJSON) }},

	// layer114 115
	// "TLAuthSendCode114": RPCContextTuple{"/mtproto.RPCAuth/auth_sendCode114", func() interface{} { return new(Auth_SentCode) }},

	"TLAccountAcceptAuthorization":          RPCContextTuple{"/mtproto.RPCAccount/account_acceptAuthorization", func() interface{} { return new(Bool) }},
	"TLAccountCancelPasswordEmail":          RPCContextTuple{"/mtproto.RPCAccount/account_cancelPasswordEmail", func() interface{} { return new(Bool) }},
	"TLAccountConfirmPasswordEmail":         RPCContextTuple{"/mtproto.RPCAccount/account_confirmPasswordEmail", func() interface{} { return new(Bool) }},
	"TLAccountCreateTheme":                  RPCContextTuple{"/mtproto.RPCAccount/account_createTheme", func() interface{} { return new(Theme) }},
	"TLAccountDeleteSecureValue":            RPCContextTuple{"/mtproto.RPCAccount/account_deleteSecureValue", func() interface{} { return new(Bool) }},
	"TLAccountFinishTakeoutSession":         RPCContextTuple{"/mtproto.RPCAccount/account_finishTakeoutSession", func() interface{} { return new(Bool) }},
	"TLAccountGetAllSecureValues":           RPCContextTuple{"/mtproto.RPCAccount/account_getAllSecureValues", func() interface{} { return new(Vector_SecureValue) }},
	"TLAccountGetAuthorizationForm":         RPCContextTuple{"/mtproto.RPCAccount/account_getAuthorizationForm", func() interface{} { return new(Account_AuthorizationForm) }},
	"TLAccountGetAutoDownloadSettings":      RPCContextTuple{"/mtproto.RPCAccount/account_getAutoDownloadSettings", func() interface{} { return new(Account_AutoDownloadSettings) }},
	"TLAccountGetContactSignUpNotification": RPCContextTuple{"/mtproto.RPCAccount/account_getContactSignUpNotification", func() interface{} { return new(Bool) }},
	"TLAccountGetContentSettings":           RPCContextTuple{"/mtproto.RPCAccount/account_getContentSettings", func() interface{} { return new(Account_ContentSettings) }},
	"TLAccountGetMultiWallPapers":           RPCContextTuple{"/mtproto.RPCAccount/account_getMultiWallPapers", func() interface{} { return new(Vector_WallPaper) }},
	"TLAccountGetNotifyExceptions":          RPCContextTuple{"/mtproto.RPCAccount/account_getNotifyExceptions", func() interface{} { return new(Updates) }},
	"TLAccountGetPasswordSettings114":       RPCContextTuple{"/mtproto.RPCAccount/account_getPasswordSettings114", func() interface{} { return new(Account_PasswordSettings) }},

	"TLAccountGetSecureValue":    RPCContextTuple{"/mtproto.RPCAccount/account_getSecureValue", func() interface{} { return new(Vector_SecureValue) }},
	"TLAccountGetTheme":          RPCContextTuple{"/mtproto.RPCAccount/account_getTheme", func() interface{} { return new(Theme) }},
	"TLAccountGetThemes":         RPCContextTuple{"/mtproto.RPCAccount/account_getThemes", func() interface{} { return new(Account_Themes) }},
	"TLAccountGetTmpPassword114": RPCContextTuple{"/mtproto.RPCAccount/account_getTmpPassword114", func() interface{} { return new(Account_TmpPassword) }},
	"TLAccountGetWallPaper":      RPCContextTuple{"/mtproto.RPCAccount/account_getWallPaper", func() interface{} { return new(WallPaper) }},
	"TLAccountGetWallPapers114":  RPCContextTuple{"/mtproto.RPCAccount/account_getWallPapers114", func() interface{} { return new(Vector_WallPaper) }},

	"TLAccountGetWebAuthorizations": RPCContextTuple{"/mtproto.RPCAccount/account_getWebAuthorizations", func() interface{} { return new(Account_WebAuthorizations) }},
	"TLAccountInitTakeoutSession":   RPCContextTuple{"/mtproto.RPCAccount/account_initTakeoutSession", func() interface{} { return new(Account_Takeout) }},
	"TLAccountInstallTheme":         RPCContextTuple{"/mtproto.RPCAccount/account_installTheme", func() interface{} { return new(Bool) }},
	"TLAccountInstallWallPaper":     RPCContextTuple{"/mtproto.RPCAccount/account_installWallPaper", func() interface{} { return new(Bool) }},
	"TLAccountRegisterDevice114":    RPCContextTuple{"/mtproto.RPCAccount/account_registerDevice114", func() interface{} { return new(Bool) }},

	"TLAccountResendPasswordEmail":      RPCContextTuple{"/mtproto.RPCAccount/account_resendPasswordEmail", func() interface{} { return new(Bool) }},
	"TLAccountResetWallPapers":          RPCContextTuple{"/mtproto.RPCAccount/account_resetWallPapers", func() interface{} { return new(Bool) }},
	"TLAccountResetWebAuthorization":    RPCContextTuple{"/mtproto.RPCAccount/account_resetWebAuthorization", func() interface{} { return new(Bool) }},
	"TLAccountResetWebAuthorizations":   RPCContextTuple{"/mtproto.RPCAccount/account_resetWebAuthorizations", func() interface{} { return new(Bool) }},
	"TLAccountSaveAutoDownloadSettings": RPCContextTuple{"/mtproto.RPCAccount/account_saveAutoDownloadSettings", func() interface{} { return new(Bool) }},
	"TLAccountSaveSecureValue":          RPCContextTuple{"/mtproto.RPCAccount/account_saveSecureValue", func() interface{} { return new(SecureValue) }},
	"TLAccountSaveTheme":                RPCContextTuple{"/mtproto.RPCAccount/account_saveTheme", func() interface{} { return new(Bool) }},
	"TLAccountSaveWallPaper":            RPCContextTuple{"/mtproto.RPCAccount/account_saveWallPaper", func() interface{} { return new(Bool) }},
	"TLAccountSendChangePhoneCode114":   RPCContextTuple{"/mtproto.RPCAccount/account_sendChangePhoneCode114", func() interface{} { return new(Auth_SentCode) }},
	"TLAccountSendConfirmPhoneCode114":  RPCContextTuple{"/mtproto.RPCAccount/account_sendConfirmPhoneCode114", func() interface{} { return new(Auth_SentCode) }},

	"TLAccountSendVerifyEmailCode":          RPCContextTuple{"/mtproto.RPCAccount/account_sendVerifyEmailCode", func() interface{} { return new(Account_SentEmailCode) }},
	"TLAccountSendVerifyPhoneCode":          RPCContextTuple{"/mtproto.RPCAccount/account_sendVerifyPhoneCode", func() interface{} { return new(Auth_SentCode) }},
	"TLAccountSetContactSignUpNotification": RPCContextTuple{"/mtproto.RPCAccount/account_setContactSignUpNotification", func() interface{} { return new(Bool) }},
	"TLAccountSetContentSettings":           RPCContextTuple{"/mtproto.RPCAccount/account_setContentSettings", func() interface{} { return new(Bool) }},
	"TLAccountUnregisterDevice114":          RPCContextTuple{"/mtproto.RPCAccount/account_unregisterDevice114", func() interface{} { return new(Bool) }},
	"TLAccountUpdatePasswordSettings114":    RPCContextTuple{"/mtproto.RPCAccount/account_updatePasswordSettings114", func() interface{} { return new(Bool) }},

	"TLAccountUpdateTheme":     RPCContextTuple{"/mtproto.RPCAccount/account_updateTheme", func() interface{} { return new(Theme) }},
	"TLAccountUploadTheme":     RPCContextTuple{"/mtproto.RPCAccount/account_uploadTheme", func() interface{} { return new(Document) }},
	"TLAccountUploadWallPaper": RPCContextTuple{"/mtproto.RPCAccount/account_uploadWallPaper", func() interface{} { return new(WallPaper) }},
	"TLAccountVerifyEmail":     RPCContextTuple{"/mtproto.RPCAccount/account_verifyEmail", func() interface{} { return new(Bool) }},
	"TLAccountVerifyPhone":     RPCContextTuple{"/mtproto.RPCAccount/account_verifyPhone", func() interface{} { return new(Bool) }},

	"TLAuthAcceptLoginToken": RPCContextTuple{"/mtproto.RPCAuth/auth_acceptLoginToken", func() interface{} { return new(Authorization) }},
	"TLAuthCheckPassword114": RPCContextTuple{"/mtproto.RPCAuth/auth_checkPassword114", func() interface{} { return new(Auth_Authorization) }},

	"TLAuthExportLoginToken": RPCContextTuple{"/mtproto.RPCAuth/auth_exportLoginToken", func() interface{} { return new(Auth_LoginToken) }},
	"TLAuthImportLoginToken": RPCContextTuple{"/mtproto.RPCAuth/auth_importLoginToken", func() interface{} { return new(Auth_LoginToken) }},
	"TLAuthSendCode114":      RPCContextTuple{"/mtproto.RPCAuth/auth_sendCode114", func() interface{} { return new(Auth_SentCode) }},
	"TLAuthSignUp114":        RPCContextTuple{"/mtproto.RPCAuth/auth_signUp114", func() interface{} { return new(Auth_Authorization) }},

	"TLBotsSetBotCommands": RPCContextTuple{"/mtproto.RPCBots/bots_setBotCommands", func() interface{} { return new(Bool) }},

	"TLChannelsCreateChannel114": RPCContextTuple{"/mtproto.RPCChannels/channels_createChannel114", func() interface{} { return new(Updates) }},
	"TLChannelsEditAdmin114":     RPCContextTuple{"/mtproto.RPCChannels/channels_editAdmin114", func() interface{} { return new(Updates) }},
	"TLChannelsEditBanned114":    RPCContextTuple{"/mtproto.RPCChannels/channels_editBanned114", func() interface{} { return new(Updates) }},

	"TLChannelsEditCreator":                 RPCContextTuple{"/mtproto.RPCChannels/channels_editCreator", func() interface{} { return new(Updates) }},
	"TLChannelsEditLocation":                RPCContextTuple{"/mtproto.RPCChannels/channels_editLocation", func() interface{} { return new(Bool) }},
	"TLChannelsExportMessageLink114":        RPCContextTuple{"/mtproto.RPCChannels/channels_exportMessageLink114", func() interface{} { return new(ExportedMessageLink) }},
	"TLChannelsGetAdminedPublicChannels114": RPCContextTuple{"/mtproto.RPCChannels/channels_getAdminedPublicChannels114", func() interface{} { return new(Messages_Chats) }},

	"TLChannelsGetGroupsForDiscussion": RPCContextTuple{"/mtproto.RPCChannels/channels_getGroupsForDiscussion", func() interface{} { return new(Messages_Chats) }},
	"TLChannelsGetInactiveChannels":    RPCContextTuple{"/mtproto.RPCChannels/channels_getInactiveChannels", func() interface{} { return new(Messages_InactiveChats) }},
	"TLChannelsGetLeftChannels":        RPCContextTuple{"/mtproto.RPCChannels/channels_getLeftChannels", func() interface{} { return new(Messages_Chats) }},
	"TLChannelsGetMessages114":         RPCContextTuple{"/mtproto.RPCChannels/channels_getMessages114", func() interface{} { return new(Messages_Messages) }},

	"TLChannelsSetDiscussionGroup": RPCContextTuple{"/mtproto.RPCChannels/channels_setDiscussionGroup", func() interface{} { return new(Bool) }},
	"TLChannelsToggleSlowMode":     RPCContextTuple{"/mtproto.RPCChannels/channels_toggleSlowMode", func() interface{} { return new(Updates) }},

	"TLContactsAcceptContact":     RPCContextTuple{"/mtproto.RPCContacts/contacts_acceptContact", func() interface{} { return new(Updates) }},
	"TLContactsAddContact":        RPCContextTuple{"/mtproto.RPCContacts/contacts_addContact", func() interface{} { return new(Updates) }},
	"TLContactsDeleteByPhones":    RPCContextTuple{"/mtproto.RPCContacts/contacts_deleteByPhones", func() interface{} { return new(Bool) }},
	"TLContactsDeleteContacts114": RPCContextTuple{"/mtproto.RPCContacts/contacts_deleteContacts114", func() interface{} { return new(Updates) }},

	"TLContactsGetContactIDs":  RPCContextTuple{"/mtproto.RPCContacts/contacts_getContactIDs", func() interface{} { return new(VectorInt) }},
	"TLContactsGetLocated":     RPCContextTuple{"/mtproto.RPCContacts/contacts_getLocated", func() interface{} { return new(Updates) }},
	"TLContactsGetSaved":       RPCContextTuple{"/mtproto.RPCContacts/contacts_getSaved", func() interface{} { return new(Vector_SavedContact) }},
	"TLContactsToggleTopPeers": RPCContextTuple{"/mtproto.RPCContacts/contacts_toggleTopPeers", func() interface{} { return new(Bool) }},

	"TLFoldersDeleteFolder":    RPCContextTuple{"/mtproto.RPCFolders/folders_deleteFolder", func() interface{} { return new(Updates) }},
	"TLFoldersEditPeerFolders": RPCContextTuple{"/mtproto.RPCFolders/folders_editPeerFolders", func() interface{} { return new(Updates) }},

	"TLHelpEditUserInfo":    RPCContextTuple{"/mtproto.RPCHelp/help_editUserInfo", func() interface{} { return new(Help_UserInfo) }},
	"TLHelpGetAppConfig":    RPCContextTuple{"/mtproto.RPCHelp/help_getAppConfig", func() interface{} { return new(JSONValue) }},
	"TLHelpGetAppUpdate114": RPCContextTuple{"/mtproto.RPCHelp/help_getAppUpdate114", func() interface{} { return new(Help_AppUpdate) }},
	"TLHelpGetConfig114":    RPCContextTuple{"/mtproto.RPCHelp/help_getConfig114", func() interface{} { return new(Config114) }},

	"TLHelpGetPassportConfig": RPCContextTuple{"/mtproto.RPCHelp/help_getPassportConfig", func() interface{} { return new(Help_PassportConfig) }},
	"TLHelpGetPromoData":      RPCContextTuple{"/mtproto.RPCHelp/help_getPromoData", func() interface{} { return new(Help_PromoData) }},
	"TLHelpGetSupportName":    RPCContextTuple{"/mtproto.RPCHelp/help_getSupportName", func() interface{} { return new(Help_SupportName) }},
	"TLHelpGetUserInfo":       RPCContextTuple{"/mtproto.RPCHelp/help_getUserInfo", func() interface{} { return new(Help_UserInfo) }},
	"TLHelpHidePromoData":     RPCContextTuple{"/mtproto.RPCHelp/help_hidePromoData", func() interface{} { return new(Bool) }},

	"TLLangpackGetLangPack114":   RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLangPack114", func() interface{} { return new(LangPackDifference) }},
	"TLLangpackGetDifference114": RPCContextTuple{"/mtproto.RPCLangpack/langpack_getDifference114", func() interface{} { return new(LangPackDifference) }},
	"TLLangpackGetStrings114":    RPCContextTuple{"/mtproto.RPCLangpack/langpack_getStrings114", func() interface{} { return new(Vector_LangPackString) }},
	"TLLangpackGetLanguages114":  RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLanguages114", func() interface{} { return new(Vector_LangPackLanguage) }},
	"TLLangpackGetLanguage":      RPCContextTuple{"/mtproto.RPCLangpack/langpack_getLanguage", func() interface{} { return new(LangPackLanguage) }},

	"TLMessagesAcceptUrlAuth":               RPCContextTuple{"/mtproto.RPCMessages/messages_acceptUrlAuth", func() interface{} { return new(UrlAuthResult) }},
	"TLMessagesClearAllDrafts":              RPCContextTuple{"/mtproto.RPCMessages/messages_clearAllDrafts", func() interface{} { return new(Bool) }},
	"TLMessagesDeleteScheduledMessages":     RPCContextTuple{"/mtproto.RPCMessages/messages_deleteScheduledMessages", func() interface{} { return new(Updates) }},
	"TLMessagesEditChatAbout":               RPCContextTuple{"/mtproto.RPCMessages/messages_editChatAbout", func() interface{} { return new(Bool) }},
	"TLMessagesEditChatDefaultBannedRights": RPCContextTuple{"/mtproto.RPCMessages/messages_editChatDefaultBannedRights", func() interface{} { return new(Updates) }},
	"TLMessagesEditInlineBotMessage114":     RPCContextTuple{"/mtproto.RPCMessages/messages_editInlineBotMessage114", func() interface{} { return new(Bool) }},
	"TLMessagesEditMessage114":              RPCContextTuple{"/mtproto.RPCMessages/messages_editMessage114", func() interface{} { return new(Updates) }},
	"TLMessagesExportChatInvite114":         RPCContextTuple{"/mtproto.RPCMessages/messages_exportChatInvite114", func() interface{} { return new(ExportedChatInvite) }},
	"TLMessagesForwardMessages114":          RPCContextTuple{"/mtproto.RPCMessages/messages_forwardMessages114", func() interface{} { return new(Updates) }},

	"TLMessagesGetDialogFilters": RPCContextTuple{"/mtproto.RPCMessages/messages_getDialogFilters", func() interface{} { return new(Vector_DialogFilter) }},
	"TLMessagesGetDialogs114":    RPCContextTuple{"/mtproto.RPCMessages/messages_getDialogs114", func() interface{} { return new(Messages_Dialogs) }},

	"TLMessagesGetDialogUnreadMarks":       RPCContextTuple{"/mtproto.RPCMessages/messages_getDialogUnreadMarks", func() interface{} { return new(Vector_DialogPeer) }},
	"TLMessagesGetEmojiKeywords":           RPCContextTuple{"/mtproto.RPCMessages/messages_getEmojiKeywords", func() interface{} { return new(EmojiKeywordsDifference) }},
	"TLMessagesGetEmojiKeywordsDifference": RPCContextTuple{"/mtproto.RPCMessages/messages_getEmojiKeywordsDifference", func() interface{} { return new(EmojiKeywordsDifference) }},
	"TLMessagesGetEmojiKeywordsLanguages":  RPCContextTuple{"/mtproto.RPCMessages/messages_getEmojiKeywordsLanguages", func() interface{} { return new(Vector_EmojiLanguage) }},
	"TLMessagesGetEmojiURL":                RPCContextTuple{"/mtproto.RPCMessages/messages_getEmojiURL", func() interface{} { return new(EmojiURL) }},
	"TLMessagesGetOldFeaturedStickers":     RPCContextTuple{"/mtproto.RPCMessages/messages_getOldFeaturedStickers", func() interface{} { return new(Messages_FeaturedStickers) }},
	"TLMessagesGetOnlines":                 RPCContextTuple{"/mtproto.RPCMessages/messages_getOnlines", func() interface{} { return new(ChatOnlines) }},
	"TLMessagesGetPinnedDialogs114":        RPCContextTuple{"/mtproto.RPCMessages/messages_getPinnedDialogs114", func() interface{} { return new(Messages_PeerDialogs) }},

	"TLMessagesGetPollResults": RPCContextTuple{"/mtproto.RPCMessages/messages_getPollResults", func() interface{} { return new(Updates) }},
	"TLMessagesGetPollVotes":   RPCContextTuple{"/mtproto.RPCMessages/messages_getPollVotes", func() interface{} { return new(Messages_VotesList) }},
	//(TODO @work)114变成默认的了，，有空再解决这个BUG,先这样写
	"TLMessagesGetRecentLocations114": RPCContextTuple{"/mtproto.RPCMessages/messages_getRecentLocations114", func() interface{} { return new(Messages_Messages) }},

	"TLMessagesGetScheduledHistory":       RPCContextTuple{"/mtproto.RPCMessages/messages_getScheduledHistory", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetScheduledMessages":      RPCContextTuple{"/mtproto.RPCMessages/messages_getScheduledMessages", func() interface{} { return new(Messages_Messages) }},
	"TLMessagesGetSearchCounters":         RPCContextTuple{"/mtproto.RPCMessages/messages_getSearchCounters", func() interface{} { return new(Vector_Messages_SearchCounter) }},
	"TLMessagesGetSplitRanges":            RPCContextTuple{"/mtproto.RPCMessages/messages_getSplitRanges", func() interface{} { return new(Vector_MessageRange) }},
	"TLMessagesGetStatsURL":               RPCContextTuple{"/mtproto.RPCMessages/messages_getStatsURL", func() interface{} { return new(StatsURL) }},
	"TLMessagesGetStickers":               RPCContextTuple{"/mtproto.RPCMessages/messages_getStickers", func() interface{} { return new(Messages_Stickers) }},
	"TLMessagesGetSuggestedDialogFilters": RPCContextTuple{"/mtproto.RPCMessages/messages_getSuggestedDialogFilters", func() interface{} { return new(Vector_DialogFilterSuggested) }},
	"TLMessagesGetWebPagePreview114":      RPCContextTuple{"/mtproto.RPCMessages/messages_getWebPagePreview114", func() interface{} { return new(MessageMedia) }},

	"TLMessagesHidePeerSettingsBar":     RPCContextTuple{"/mtproto.RPCMessages/messages_hidePeerSettingsBar", func() interface{} { return new(Bool) }},
	"TLMessagesMarkDialogUnread":        RPCContextTuple{"/mtproto.RPCMessages/messages_markDialogUnread", func() interface{} { return new(Bool) }},
	"TLMessagesReorderPinnedDialogs114": RPCContextTuple{"/mtproto.RPCMessages/messages_reorderPinnedDialogs114", func() interface{} { return new(Bool) }},

	"TLMessagesReport":          RPCContextTuple{"/mtproto.RPCMessages/messages_report", func() interface{} { return new(Bool) }},
	"TLMessagesRequestUrlAuth":  RPCContextTuple{"/mtproto.RPCMessages/messages_requestUrlAuth", func() interface{} { return new(UrlAuthResult) }},
	"TLMessagesSearchGlobal114": RPCContextTuple{"/mtproto.RPCMessages/messages_searchGlobal114", func() interface{} { return new(Messages_Messages) }},

	"TLMessagesSearchStickerSets":      RPCContextTuple{"/mtproto.RPCMessages/messages_searchStickerSets", func() interface{} { return new(Messages_FoundStickerSets) }},
	"TLMessagesSendInlineBotResult114": RPCContextTuple{"/mtproto.RPCMessages/messages_sendInlineBotResult114", func() interface{} { return new(Updates) }},
	"TLMessagesSendMessage114":         RPCContextTuple{"/mtproto.RPCMessages/messages_sendMessage114", func() interface{} { return new(Updates) }},
	"TLMessagesSendMedia114":           RPCContextTuple{"/mtproto.RPCMessages/messages_sendMedia114", func() interface{} { return new(Updates) }},
	"TLMessagesSendMultiMedia114":      RPCContextTuple{"/mtproto.RPCMessages/messages_sendMultiMedia114", func() interface{} { return new(Updates) }},

	"TLMessagesSendScheduledMessages": RPCContextTuple{"/mtproto.RPCMessages/messages_sendScheduledMessages", func() interface{} { return new(Updates) }},
	"TLMessagesSendVote":              RPCContextTuple{"/mtproto.RPCMessages/messages_sendVote", func() interface{} { return new(Updates) }},
	"TLMessagesToggleDialogPin114":    RPCContextTuple{"/mtproto.RPCMessages/messages_toggleDialogPin114", func() interface{} { return new(Bool) }},

	"TLMessagesToggleStickerSets":        RPCContextTuple{"/mtproto.RPCMessages/messages_toggleStickerSets", func() interface{} { return new(Bool) }},
	"TLMessagesUpdateDialogFilter":       RPCContextTuple{"/mtproto.RPCMessages/messages_updateDialogFilter", func() interface{} { return new(Bool) }},
	"TLMessagesUpdateDialogFiltersOrder": RPCContextTuple{"/mtproto.RPCMessages/messages_updateDialogFiltersOrder", func() interface{} { return new(Bool) }},
	"TLMessagesUpdatePinnedMessage":      RPCContextTuple{"/mtproto.RPCMessages/messages_updatePinnedMessage", func() interface{} { return new(Updates) }},

	"TLPaymentsGetBankCardData": RPCContextTuple{"/mtproto.RPCPayments/payments_getBankCardData", func() interface{} { return new(Payments_BankCardData) }},

	"TLPhoneDiscardCall114":    RPCContextTuple{"/mtproto.RPCPhone/phone_discardCall114", func() interface{} { return new(Updates) }},
	"TLPhoneRequestCall114":    RPCContextTuple{"/mtproto.RPCPhone/phone_requestCall114", func() interface{} { return new(Phone_PhoneCall) }},
	"TLPhoneSetCallRating114":  RPCContextTuple{"/mtproto.RPCPhone/phone_setCallRating114", func() interface{} { return new(Updates) }},
	"TLPhoneSendSignalingData": RPCContextTuple{"/mtproto.RPCPhone/phone_sendSignalingData", func() interface{} { return new(Bool) }},

	"TLStatsGetBroadcastStats": RPCContextTuple{"/mtproto.RPCStats/stats_getBroadcastStats", func() interface{} { return new(Stats_BroadcastStats) }},
	"TLStatsLoadAsyncGraph":    RPCContextTuple{"/mtproto.RPCStats/stats_loadAsyncGraph", func() interface{} { return new(StatsGraph) }},

	"TLStickersCreateStickerSet114": RPCContextTuple{"/mtproto.RPCStickers/stickers_createStickerSet114", func() interface{} { return new(Messages_StickerSet) }},
	"TLStickersSetStickerSetThumb":  RPCContextTuple{"/mtproto.RPCStickers/stickers_setStickerSetThumb", func() interface{} { return new(Messages_StickerSet) }},

	"TLUsersSetSecureValueErrors": RPCContextTuple{"/mtproto.RPCUsers/users_setSecureValueErrors", func() interface{} { return new(Bool) }},

	// "TLWebrtcAddCandidate": RPCContextTuple{"/mtproto.RPCWebrtc/webrtc_addCandidate", func() interface{} { return new(Bool) }},
	// "TLWebrtcCreateAnswer": RPCContextTuple{"/mtproto.RPCWebrtc/webrtc_createAnswer", func() interface{} { return new(Bool) }},
	// "TLWebrtcCreateOffer": RPCContextTuple{"/mtproto.RPCWebrtc/webrtc_createOffer", func() interface{} { return new(Bool) }},

	"TLUploadGetCdnFileHashes114": RPCContextTuple{"/mtproto.RPCUpload/upload_getCdnFileHashes114", func() interface{} { return new(Vector_FileHash) }},
	"TLUploadGetFile114":          RPCContextTuple{"/mtproto.RPCUpload/upload_getFile114", func() interface{} { return new(Upload_File) }},
	"TLUploadReuploadCdnFile114":  RPCContextTuple{"/mtproto.RPCUpload/upload_reuploadCdnFile114", func() interface{} { return new(Vector_FileHash) }},

	// layer 117
	"TLAccountGetGlobalPrivacySettings": RPCContextTuple{"/mtproto.RPCAccount/account_getGlobalPrivacySettings", func() interface{} { return new(GlobalPrivacySettings) }},
	"TLAccountSetGlobalPrivacySettings": RPCContextTuple{"/mtproto.RPCAccount/account_setGlobalPrivacySettings", func() interface{} { return new(GlobalPrivacySettings) }},
	"TLHelpDismissSuggestion":           RPCContextTuple{"/mtproto.RPCHelp/help_dismissSuggestion", func() interface{} { return new(Bool) }},

	"TLPhotosUploadProfilePhoto117": RPCContextTuple{"/mtproto.RPCPhotos/photos_uploadProfilePhoto117", func() interface{} { return new(Photos_Photo) }},
	"TLPhotosUpdateProfilePhoto117": RPCContextTuple{"/mtproto.RPCPhotos/photos_updateProfilePhoto117", func() interface{} { return new(Photos_Photo) }},

	"TLStatsGetMegagroupStats": RPCContextTuple{"/mtproto.RPCStats/stats_getMegagroupStats", func() interface{} { return new(Stats_MegagroupStats) }},

	// layer 118
	"TLHelpGetCountriesList": RPCContextTuple{"/mtproto.RPCHelp/help_getCountriesList", func() interface{} { return new(Help_CountriesList) }},

	"TLMessagesGetBotCallbackAnswer118": RPCContextTuple{"/mtproto.RPCMessages/messages_getBotCallbackAnswer118", func() interface{} { return new(Messages_BotCallbackAnswer) }},
	"TLMessagesSendEncrypted118":        RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncrypted118", func() interface{} { return new(Messages_SentEncryptedMessage) }},
	"TLMessagesSendEncryptedFile118":    RPCContextTuple{"/mtproto.RPCMessages/messages_sendEncryptedFile118", func() interface{} { return new(Messages_SentEncryptedMessage) }},

	"TLLangpackGetDifference118": RPCContextTuple{"/mtproto.RPCLangpack/langpack_getDifference118", func() interface{} { return new(LangPackDifference) }},
}

func FindRPCContextTuple(t interface{}, layerId int32) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	rt_name := rt.Name()

	if layerId == 73 {
		if rt_name == "TLHelpGetConfig" {
			rt_name = "TLHelpGetConfig73"
			logger.LogSugar.Warnf("FindRPCContextTuple layerId:%d, rt_name:%s:", layerId, rt_name)
		}
	} else if layerId == 82 {
		if rt_name == "TLHelpGetConfig" {
			rt_name = "TLHelpGetConfig82"
			logger.LogSugar.Warnf("FindRPCContextTuple layerId:%d, rt_name:%s:", layerId, rt_name)
		} else if rt_name == "TLAuthSendCode73" {
			rt_name = "TLAuthSendCode82"
			logger.LogSugar.Warnf("FindRPCContextTuple layerId:%d, rt_name:%s:", layerId, rt_name)
		}
	} else if layerId >= 114 {
		if rt_name == "TLHelpGetConfig" {
			rt_name = "TLHelpGetConfig114"
			logger.LogSugar.Warnf("FindRPCContextTuple layerId:%d, rt_name:%s:", layerId, rt_name)
		}
	}

	m, ok := rpcContextRegisters[rt_name]
	if !ok {
		logger.LogSugar.Errorf("Can't find name: %s", rt_name)
		return nil
	}
	return &m
}
