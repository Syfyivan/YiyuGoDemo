package service

import (
	"awesomeProject1/pkg/models"
	"gorm.io/gorm"
	"math"
)

type PaginationResult struct {
	Records    []models.Diary // 记录数据
	Total      int64          // 总记录数
	Page       int            // 当前页码
	PageSize   int            // 每页记录数
	TotalPages int            // 总页数
}

func Paginate(db *gorm.DB, page, pageSize int, model interface{}) (*PaginationResult, error) {
	var count int64
	var err error

	// 构建基础查询
	query := db.Model(model)

	// 获取总数
	err = query.Count(&count).Error
	if err != nil {
		return nil, err
	}

	// 计算总页数和偏移量
	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))
	offset := (page - 1) * pageSize

	// 执行分页查询
	var records []models.Diary
	err = query.Offset(offset).Limit(pageSize).Find(&records).Error
	if err != nil {
		return nil, err
	}

	// 创建分页响应
	return &PaginationResult{
		Records:    records,
		Total:      count,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
