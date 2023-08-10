package common

type ListMeta struct {
	//TotalCount int64 `json:"totalCount,omitempty"`
	Pn    uint32 `json:"pn,omitempty"`
	PSize uint32 `json:"pSize,omitempty"`
}
