package v1

type Menus struct {
	Menus []MenuBaseSpec `json:"menu"`
}

type MenuBaseSpec struct {
	Id       string         `json:"id" gorm:"column:id"`
	Name     string         `json:"name" gorm:"column:name"`
	Parent   string         `json:"parent_id" gorm:"column:parent_id"`
	ZhName   string         `json:"zh_name" gorm:"column:zhName"`
	Type     int            `json:"type" gorm:"column:type"`
	Url      string         `json:"url" gorm:"column:url"`
	Children []MenuBaseSpec `json:"children" gorm:"-"`
	Icon     string         `json:"icon"`
}
