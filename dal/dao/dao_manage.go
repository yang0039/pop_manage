package dao

import (
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dao/mysql_dao"
)

var daoList *MysqlDAOList

var Master *sqlx.DB

type MysqlDAOList struct {
	// common
	CommonDAO            *mysql_dao.CommonDAO
	UserDAO              *mysql_dao.UserDAO
	ChatDAO              *mysql_dao.ChatDAO
	ChatParticipantDAO   *mysql_dao.ChatParticipantDAO
	CallDAO              *mysql_dao.CallDAO
	UserMsgboxDAO        *mysql_dao.UserMsgboxDAO
	NoteDAO              *mysql_dao.NoteDAO
	PhoneTransactionsDAO *mysql_dao.PhoneTransactionsDAO
	BannedDAO            *mysql_dao.BannedDAO
	MenuDAO              *mysql_dao.MenuDAO
	PermissionsDAO       *mysql_dao.PermissionsDAO
	PermissionsUrlDAO    *mysql_dao.PermissionsUrlDAO
	RoleDAO              *mysql_dao.RoleDAO
	AccountDAO           *mysql_dao.AccountDAO
	LoginLogDAO          *mysql_dao.LoginLogDAO
	ChannelMsgRowDAO     *mysql_dao.ChannelMsgRowDAO
	RawMessageRowDAO     *mysql_dao.RawMessageRowDAO
	UserMsgRowDAO        *mysql_dao.UserMsgRowDAO
	AllowIpDAO           *mysql_dao.AllowIpDAO
	ReportDAO            *mysql_dao.ReportDAO
	PhotoDAO             *mysql_dao.PhotoDAO
	DocumentDAO          *mysql_dao.DocumentDAO
	LabelDAO             *mysql_dao.LabelDAO
	PeerStatusDAO        *mysql_dao.PeerStatusDAO
	FilesDAO             *mysql_dao.FilesDAO
	BannedInfoDAO        *mysql_dao.BannedInfoDAO
	UserLogsDAO          *mysql_dao.UserLogsDAO
	RequestRecoreDAO     *mysql_dao.RequestRecoreDAO
	UserOperaDAO         *mysql_dao.UserOperaDAO
}

func InstallMysqlDAOManager(db *sqlx.DB, dbindex int) {
	mysql_dao.Dbindex = dbindex
	Master = db

	daoList = &MysqlDAOList{}
	daoList.CommonDAO = mysql_dao.NewCommonDAO(db)
	daoList.UserDAO = mysql_dao.NewUserDAO(db)
	daoList.ChatDAO = mysql_dao.NewChatDAO(db)
	daoList.ChatParticipantDAO = mysql_dao.NewChatParticipartDAO(db)
	daoList.CallDAO = mysql_dao.NewCallDAO(db)
	daoList.UserMsgboxDAO = mysql_dao.NewUserMsgboxDAO(db)
	daoList.NoteDAO = mysql_dao.NewNoteDAO(db)
	daoList.PhoneTransactionsDAO = mysql_dao.NewPhoneTransactionsDAO(db)
	daoList.BannedDAO = mysql_dao.NewBannedDAO(db)
	daoList.MenuDAO = mysql_dao.NewMenuDAO(db)
	daoList.PermissionsDAO = mysql_dao.NewPermissionsDAO(db)
	daoList.PermissionsUrlDAO = mysql_dao.NewPermissionsUrlDAO(db)
	daoList.RoleDAO = mysql_dao.NewRoleDAO(db)
	daoList.AccountDAO = mysql_dao.NewAccountDAO(db)
	daoList.LoginLogDAO = mysql_dao.NewLoginLogDAO(db)
	daoList.ChannelMsgRowDAO = mysql_dao.NewChannelMsgRowDAO(db)
	daoList.RawMessageRowDAO = mysql_dao.NewRawMessageRowDAO(db)
	daoList.UserMsgRowDAO = mysql_dao.NewUserMsgRowDAO(db)
	daoList.AllowIpDAO = mysql_dao.NewAllowIpDAO(db)
	daoList.ReportDAO = mysql_dao.NewReportDAO(db)
	daoList.PhotoDAO = mysql_dao.NewPhotoDAO(db)
	daoList.DocumentDAO = mysql_dao.NewDocumentDAO(db)
	daoList.LabelDAO = mysql_dao.NewLabelDAO(db)
	daoList.PeerStatusDAO = mysql_dao.NewPeerStatusDAO(db)
	daoList.FilesDAO = mysql_dao.NewFilesDAO(db)
	daoList.BannedInfoDAO = mysql_dao.NewBannedInfoDAO(db)
	daoList.UserLogsDAO = mysql_dao.NewUserLogsDAO(db)
	daoList.RequestRecoreDAO = mysql_dao.NewRequestRecoreDAO(db)
	daoList.UserOperaDAO = mysql_dao.NewUserOperaDAO(db)
}

func GetCommonDAO() (dao *mysql_dao.CommonDAO) {
	return daoList.CommonDAO
}

func GetUserDAO() (dao *mysql_dao.UserDAO) {
	return daoList.UserDAO
}

func GetChatDAO() (dao *mysql_dao.ChatDAO) {
	return daoList.ChatDAO
}

func GetChatParticipantDAO() (dao *mysql_dao.ChatParticipantDAO) {
	return daoList.ChatParticipantDAO
}

func GetCallDAO() (dao *mysql_dao.CallDAO) {
	return daoList.CallDAO
}

func GetUserMsgboxDAO() (dao *mysql_dao.UserMsgboxDAO) {
	return daoList.UserMsgboxDAO
}

func GetNoteDAO() (dao *mysql_dao.NoteDAO) {
	return daoList.NoteDAO
}

func GetPhoneTransactionsDAO() (dao *mysql_dao.PhoneTransactionsDAO) {
	return daoList.PhoneTransactionsDAO
}

func GetBannedDAO() (dao *mysql_dao.BannedDAO) {
	return daoList.BannedDAO
}

func GetMenuDAO() (dao *mysql_dao.MenuDAO) {
	return daoList.MenuDAO
}

func GetPermissionsDAO() (dao *mysql_dao.PermissionsDAO) {
	return daoList.PermissionsDAO
}

func GetPermissionsUrlDAO() (dao *mysql_dao.PermissionsUrlDAO) {
	return daoList.PermissionsUrlDAO
}

func GetRoleDAO() (dao *mysql_dao.RoleDAO) {
	return daoList.RoleDAO
}

func GetAccountDAO() (dao *mysql_dao.AccountDAO) {
	return daoList.AccountDAO
}

func GetLoginLogDAO() (dao *mysql_dao.LoginLogDAO) {
	return daoList.LoginLogDAO
}

func GetChannelMsgRowDAO() (dao *mysql_dao.ChannelMsgRowDAO) {
	return daoList.ChannelMsgRowDAO
}

func GetRawMessageRowDAO() (dao *mysql_dao.RawMessageRowDAO) {
	return daoList.RawMessageRowDAO
}

func GetUserMsgRowDAO() (dao *mysql_dao.UserMsgRowDAO) {
	return daoList.UserMsgRowDAO
}

func GetAllowIpDAO() (dao *mysql_dao.AllowIpDAO) {
	return daoList.AllowIpDAO
}

func GetReportDAO() (dao *mysql_dao.ReportDAO) {
	return daoList.ReportDAO
}

func GetPhotoDAO() (dao *mysql_dao.PhotoDAO) {
	return daoList.PhotoDAO
}

func GetDocumentDAO() (dao *mysql_dao.DocumentDAO) {
	return daoList.DocumentDAO
}

func GetLabelDAO() (dao *mysql_dao.LabelDAO) {
	return daoList.LabelDAO
}

func GetPeerStatusDAO() (dao *mysql_dao.PeerStatusDAO) {
	return daoList.PeerStatusDAO
}

func GetFilesDAO() (dao *mysql_dao.FilesDAO) {
	return daoList.FilesDAO
}

func GetBannedInfoDAO() (dao *mysql_dao.BannedInfoDAO) {
	return daoList.BannedInfoDAO
}

func GetUserLogsDAO() (dao *mysql_dao.UserLogsDAO) {
	return daoList.UserLogsDAO
}

func GetRequestRecoreDAO() (dao *mysql_dao.RequestRecoreDAO) {
	return daoList.RequestRecoreDAO
}

func GetUserOperaDAO() (dao *mysql_dao.UserOperaDAO) {
	return daoList.UserOperaDAO
}
