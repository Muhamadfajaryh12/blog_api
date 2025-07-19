package dto

type DashboardDTO struct {
	CountView    int64        `json:"count_view"`
	CountBlog    int64        `json:"count_blog"`
	CountComment int64        `json:"count_comment"`
	CoutViewWeek []ViewDayDTO `json:"view_week"`
}

type ViewDayDTO struct {
	Date      string `json:"date"`
	CountView int64  `json:"count_view"`
}

type CommentDayDTO struct {
	Date         string `json:"date"`
	CountComment int64  `json:"count_comment"`
}