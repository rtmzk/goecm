package mysql

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"gorm.io/gorm"
)

type menu struct {
	db *gorm.DB
}

func newMenu(ds *datastore) *menu {
	return &menu{
		ds.db,
	}
}

func (m *menu) Get(ctx context.Context, opts metav1.ListOptions) ([]v1.MenuBaseSpec, error) {
	return m.getMenuTree()
}

func (m *menu) getTreemapMenu() (map[string][]v1.MenuBaseSpec, error) {
	var allMenus []v1.MenuBaseSpec
	var treeMap = make(map[string][]v1.MenuBaseSpec)
	err := m.db.Find(&allMenus).Error

	for _, v := range allMenus {
		treeMap[v.Parent] = append(treeMap[v.Parent], v)
	}

	return treeMap, err
}

func (m *menu) getMenuTree() ([]v1.MenuBaseSpec, error) {
	menuTree, err := m.getTreemapMenu()
	menus := menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func getChildrenList(menu *v1.MenuBaseSpec, treeMap map[string][]v1.MenuBaseSpec) (err error) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}

	return err
}
