package util

import (
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

type Filter struct {
	Page     int
	PageSize int
	Sort     string
}

func NewFilter(page int, pageSize int, sort string) Filter {
	return Filter{page, pageSize, sort}
}

func (f Filter) Limit() int {
	return f.PageSize
}

func (f Filter) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f Filter) FormatSort() string {
	if strings.HasPrefix(f.Sort, "-") {
		return strings.ToLower(f.Sort)
	}
	return strings.ToLower(f.Sort)
}

func (f Filter) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filter) ValidateSort(safeList []string) bool {
	return slices.Contains(safeList, f.FormatSort())
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `jEmptyson:"total_records,omitempty"`
}

func (f *Filter) CalculateMetadata(totalRecords int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  f.Page,
		PageSize:     f.PageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(f.PageSize))),
		TotalRecords: totalRecords,
	}
}
