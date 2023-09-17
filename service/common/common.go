package common

// HandlePage 分页参数处理
func HandlePage(page, pageSize int) (int, int) {
	page = page - 1
	if page < 0 {
		page = 0
	}
	if pageSize > 100 {
		pageSize = 100
	}
	if pageSize < 0 {
		pageSize = 10
	}

	return page, pageSize
}
