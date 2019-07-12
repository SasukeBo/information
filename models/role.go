package models

type RoleStatus int

const (
	RDefault RoleStatus = iota // 初始状态
	RPublish                   // 发布状态
	RBlock                     // 屏蔽状态
	RDeleted                   // 删除状态
)

// Role 角色模型
type Role struct {
	Id     int        `orm:"auto;pk"`
	Name   string     `orm:"unique"`
	Status RoleStatus `orm:"default(0)"`
}

// Insert 创建一个角色
func (r *Role) Insert() error {
	if _, err := repo.Insert(r); err != nil {
		return err
	}
	return nil
}

// Delete 删除角色
func (r *Role) Delete() error {
	r.Status = RDeleted
	if _, err := repo.Update(r); err != nil {
		return err
	}
	return nil
}

// Update 更新角色
func (r *Role) Update(params map[string]interface{}) error {
	if err := repo.Read(r); err != nil {
		return err
	}
	if params["name"] != nil {
		r.Name = params["name"].(string)
	}
	if params["status"] != nil {
		r.Status = params["status"].(RoleStatus)
	}
	if _, err := repo.Update(r); err != nil {
		return err
	}
	return nil
}

// GetByID 根据ID获取角色
func (r *Role) GetByID() error {
	if err := repo.Read(r); err != nil {
		return err
	}
	return nil
}
