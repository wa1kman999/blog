package model

import "gorm.io/gorm"

type Label struct {
	gorm.Model
	Name        string    `json:"name" gorm:"comment:标签名称"`
	Alias       string    `json:"alias" gorm:"comment:标签别名"`
	Description string    `json:"description" gorm:"type:text;comment:标签描述"`
	Article     []Article `json:"article" gorm:"comment:文章"` //一个标签有多篇文章
}
