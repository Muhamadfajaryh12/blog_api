package dto

type TagDTO struct {
	Tag string `form:"tag" json:"tag" binding:"required" example:"Sports"`
}

type TagResponseDTO struct {
	ID  uint64 `json:"id"`
	Tag string `json:"tag"`
}

type TagDetailResponseDTO struct {
	ID    uint64            `json:"id"`
	Tag   string            `json:"tag"`
	Blogs []BlogResponseDTO `json:"blogs"`
}