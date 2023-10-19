// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []Connection    `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *ConnectionPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ConnectionPaginatedResult) GetModels() []Connection {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *ConnectionPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
