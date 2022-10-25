package pagination

import (
	"github.com/paper-trade-chatbot/be-proto/general"
)

func NewPagination(pageSize int32) *general.Pagination {
	return &general.Pagination{
		Page:     1,
		PageSize: pageSize,
	}
}

func NextPagination(paginationInfo *general.PaginationInfo) *general.Pagination {

	if paginationInfo.GetCurrentPage() >= paginationInfo.GetTotalPages() {
		return nil
	}

	return &general.Pagination{
		Page:     paginationInfo.GetNextPage(),
		PageSize: paginationInfo.GetPageSize(),
	}
}

//SetPaginationDto set pagination dto
func SetPaginationDto(page int32, pageSize int32, count int32, index int32) *general.PaginationInfo {

	var nextPage, previousPage, totalPages int32
	if index+pageSize < count {
		nextPage = page + 1
	} else {
		nextPage = page
	}

	if page > 1 {
		previousPage = page - 1
	} else {
		previousPage = page
	}

	if count%pageSize != 0 {
		totalPages = count/pageSize + 1
	} else {
		totalPages = count / pageSize
	}

	return &general.PaginationInfo{
		CurrentPage:  page,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		TotalPages:   totalPages,
		TotalRows:    count,
	}
}

func GetOffsetAndLimit(pagination *general.Pagination) (offset int, limit int) {
	offset = int((pagination.Page - 1) * pagination.PageSize)
	limit = int(pagination.PageSize)
	return
}
