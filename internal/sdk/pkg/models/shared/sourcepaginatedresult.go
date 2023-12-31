// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SourcePaginatedResult - List of sources
type SourcePaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []Source        `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}
