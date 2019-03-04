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
	db.Model(&Post{}).Where(filters).Offset((page - 1)*pageSize).Limit(pageSize).Find(&posts)
	return
}

func CreatePost(post Post) (err error) {
	result := db.Create(&post)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func EditPost(id int, data interface{}) bool {
	db.Model(&Post{}).Where("id = ?", id).Update(data)
	return true

}
