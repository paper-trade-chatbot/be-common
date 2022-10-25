package pagination

import (
	"errors"
	"fmt"

	"github.com/paper-trade-chatbot/be-proto/general"
)

type HasPagination interface {
	GetPagination() *general.Pagination
}

type HasPaginationInfo interface {
	GetPaginationInfo() *general.PaginationInfo
}

func NewPagination(pageSize int32) *general.Pagination {
	return &general.Pagination{
		Page:     1,
		PageSize: pageSize,
	}
}

func NextPagination(pInfo *general.PaginationInfo) *general.Pagination {

	if pInfo == nil || pInfo.GetCurrentPage() >= pInfo.GetTotalPages() {
		return nil
	}

	return &general.Pagination{
		Page:     pInfo.GetNextPage(),
		PageSize: pInfo.GetPageSize(),
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

func IteratePage[Request HasPagination, Response HasPaginationInfo](pageSize int32, req Request, iterateFunc func(Request) (Response, error)) ([]Response, error) {

	responses := []Response{}

	pagination := NewPagination(pageSize)
	req.GetPagination().Page = pagination.Page
	req.GetPagination().PageSize = pagination.PageSize
	var paginationInfo *general.PaginationInfo

	for ok := true; ok; ok = (NextPagination(paginationInfo) != nil) {

		res, err := iterateFunc(req)
		if err != nil {
			nestedErr := errors.New(fmt.Sprintf("IteratePage page(%d) err: %v", req.GetPagination().Page, err))
			return nil, nestedErr
		}

		responses = append(responses, res)

		paginationInfo = res.GetPaginationInfo()

		nextPagination := NextPagination(res.GetPaginationInfo())
		pagination.Page = nextPagination.Page
		pagination.PageSize = nextPagination.PageSize
	}

	return responses, nil
}
