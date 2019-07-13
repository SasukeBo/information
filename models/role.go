package models

// Role 角色模型
type Role struct {
	ID     int        `orm:"auto;pk;column(id)"`
	Name   string     `orm:"unique"`
	Status BaseStatus `orm:"default(0)"`
}

// Insert 创建一个角色
func (r *Role) Insert() error {
	_, err := repo.Insert(r)
	return err
}

// Delete 删除角色
func (r *Role) Delete() error {
	r.Status = Deleted
	_, err := repo.Update(r)
	return err
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
		r.Status = params["status"].(BaseStatus)
	}
	_, err := repo.Update(r)
	return err
}

// GetByID 根据ID获取角色
func (r *Role) GetByID() error {
	return repo.Read(r)
}

// GetByName 根据Name获取角色
func (r *Role) GetByName() error {
	return repo.Read(r, "Name")
}
