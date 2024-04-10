package dao

import (
	"config_tools/app/request"
	"config_tools/config"
	"fmt"

	"github.com/kennycch/gotools/general"
)

type Account struct {
	Model
	Account    string `gorm:"type:varchar(100);comment:账号" json:"account"`
	Password   string `gorm:"type:varchar(100);comment:密码" json:"password"`
	IsSupper   uint8  `gorm:"comment:是否超级管理账号：0：否，1：是" json:"isSupper"`
	IsActivate uint8  `gorm:"comment:是否激活：0：否，1：是" json:"isActivate"`
}

type AccountList struct {
	Id         uint32 `json:"id"`
	Account    string `gorm:"type:varchar(100);comment:账号" json:"account"`
	IsSupper   uint8  `gorm:"comment:是否超级管理账号：0：否，1：是" json:"isSupper"`
	IsActivate uint8  `gorm:"comment:是否激活：0：否，1：是" json:"isActivate"`
}

func init() {
	tables = append(tables, createTable{
		table:   &Account{},
		comment: "账号表",
	})
}

// dataFilling 数据填充
func (a *Account) dataFilling() {
	a = &Account{
		Account: "admin",
		Password: general.Md5(fmt.Sprintf("%s._123456_.%s",
			config.Sign.SignKey,
			config.Jwt.SecretKey)),
		IsSupper:   1,
		IsActivate: 1,
	}
	DB.FirstOrCreate(a, Account{Account: "admin"})
}

// GetAccountByAccount 根据账号获取账号
func GetAccountByAccount(account *Account) error {
	return DB.Where("account = ?", account.Account).
		First(account).
		Error
}

// GetAccountByAccountPassword 根据账号密码获取账号
func GetAccountByAccountPassword(account *Account) error {
	return DB.Where("account = ? AND password = ?", account.Account, account.Password).
		First(account).
		Error
}

// GetAccountById 根据Id获取账号
func GetAccountById(account *Account) error {
	return DB.Where("id = ?", account.Model.Id).
		First(account).
		Error
}

// CreateAccount 创建账号
func CreateAccount(account *Account) error {
	return DB.Create(account).
		Error
}

// GetAccountList 获取账号列表
func GetAccountList(req *request.AccountListRequest) ([]*Account, int64) {
	list, count := []*Account{}, int64(0)
	quest := DB.Model(&Account{}).
		Select([]string{
			"id",
			"account",
			"is_supper",
			"is_activate",
		})
	if req.Account != "" {
		quest.Where("account like ?", "%"+req.Account+"%")
	}
	if req.Status != nil {
		quest.Where("status = ?", *req.Status)
	}
	quest.Count(&count)
	quest.Offset(GetOffset(req.Page, req.PageSize)).
		Limit(req.PageSize).
		Order("id desc").
		Find(&list)
	return list, count
}

// AccountUpdate 账号更新
func AccountUpdate(id uint32, updates map[string]interface{}) error {
	return DB.Model(&Account{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}
