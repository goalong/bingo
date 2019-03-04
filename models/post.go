package models

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	UserId int `json:"user_id"`
	Date string `json:"date"`
	LikeCount int `json:"like_count"`

}

func GetPosts(page int, pageSize int, filters map[string]interface{}) (posts []Post) {
	db.Where(filters).Offset((page - 1)*pageSize).Limit(pageSize).Find(&posts)
	return

}
