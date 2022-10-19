package pagination

import (
	"github.com/paper-trade-chatbot/be-proto/general"
)

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
