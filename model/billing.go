package model

import (
	"time"

	"nexus-ai/common"

	"gorm.io/gorm"
)

// 用户的账单信息，包括计费周期、金额和支付状态等
type Billing struct {
	BillingID       string `gorm:"column:billing_id;type:char(36);primaryKey;default:(UUID())" json:"billing_id"`      // 账单ID
	UserID          string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"` // 用户ID
	BillingNo       string `gorm:"column:billing_no;size:64;uniqueIndex;not null" json:"billing_no"`                   // 账单编号
	BillingType     string `gorm:"column:billing_type;size:20;index;not null" json:"billing_type"`                     // 账单类型(消费/退款)
	BillingCycle    string `gorm:"column:billing_cycle;size:20;not null" json:"billing_cycle"`                         // 账单周期(一次性/日/周/月)
	Remark          string `gorm:"column:remark;size:255" json:"remark"`                                               // 账单备注
	BillingCurrency string `gorm:"column:billing_currency;size:10;not null" json:"billing_currency"`                   // 账单币种 usd cny hkd ...
	BillingStatus   int8   `gorm:"column:billing_status;index;not null;default:1" json:"billing_status"`               // 账单状态(1:未出账 2:已出账 3:已支付 4:已逾期)

	BillingMoney   common.JSON `gorm:"column:billing_money;type:json" json:"billing_money"`     // 账单金额 单位 分
	BillingDetails common.JSON `gorm:"column:billing_details;type:json" json:"billing_details"` // 账单明细(各模型用量统计)
	QuotaDetails   common.JSON `gorm:"column:quota_details;type:json" json:"quota_details"`     // 配额使用明细
	PayTime        *time.Time  `gorm:"column:pay_time" json:"pay_time"`                         // 支付时间
	StartTime      time.Time   `gorm:"column:start_time;not null" json:"start_time"`            // 账单开始时间
	EndTime        time.Time   `gorm:"column:end_time;not null" json:"end_time"`                // 账单结束时间
	DueTime        time.Time   `gorm:"column:due_time;not null" json:"due_time"`                // 账单到期时间

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                                                      // 删除时间
}

// TableName 表名
func (Billing) TableName() string {
	return "billings"
}
