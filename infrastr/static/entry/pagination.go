package entry

// Pagination 分页
type Pagination struct {
	Page      int   `json:"page"`      // 当前页码
	PageSize  int   `json:"pageSize"`  // 每页显示条数
	Total     int64 `json:"total"`     // 数据总条数
	TotalPage int   `json:"totalPage"` // 总
}
