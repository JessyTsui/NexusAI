package model

import (
	"nexus-ai/common"
	"time"

	"gorm.io/gorm"
)

// 模型组信息，包括模型组ID、名称、描述、配置和价格系数等
type ModelGroup struct {
	ModelGroupID          string `gorm:"column:model_group_id;type:char(36);primaryKey;default:(UUID())" json:"model_group_id"` // 模型组唯一标识
	ModelGroupName        string `gorm:"column:model_group_name;size:100;not null" json:"model_group_name"`                     // 模型组名称
	ModelGroupDescription string `gorm:"column:model_group_description;type:text" json:"model_group_description"`               // 模型组描述

	ModelGroupPriceFactor common.JSON `gorm:"column:model_group_price_factor;type:json" json:"model_group_price_factor"` // 模型组价格系数
	ModelGroupOptions     common.JSON `gorm:"column:model_group_options;type:json" json:"model_group_options"`           // 模型组配置

	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`                          // 模型组创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"` // 模型组信息最后更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                                                      // 模型组删除时间
}

func (ModelGroup) TableName() string {
	return "model_groups"
}