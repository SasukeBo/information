INSERT INTO privilege(sign, priv_type, name) values
  ('device_c', 0, '可以创建设备'),

  ('admin_user_r', 1, '该管理员可以读用户账号信息'),
  ('admin_user_u', 1, '该管理员可以修改用户账号信息'),
  ('admin_user_d', 1, '该管理员可以删除用户账号'),
  ('admin_user_login_r', 1, '该管理员可以查看用户登录记录'),
  ('admin_user_extend_r', 1, '该管理员可以查看用户资料'),
  ('admin_user_extend_w', 1, '该管理员可以修改用户资料'),
  ('admin_role_c', 1, '该管理员可以创建角色'),
  ('admin_role_u', 1, '该管理员可以修改角色信息'),
  ('admin_role_d', 1, '该管理员可以删除角色'),
  ('admin_role_priv_w', 1, '该管理员可以为角色分配权限'),
  ('admin_device_r', 1, '该管理员可以查看任何设备'),
  ('admin_device_u', 1, '该管理员可以修改任何设备'),
  ('admin_device_d', 1, '该管理员可以删除任何设备'),
  ('device_u', 2, '该设备负责人可以更新设备信息'),
  ('device_d', 2, '该设备负责人可以删除设备'),
  ('device_charge_c', 2, '该设备负责人可以增加设备负责人'),
  ('device_charge_d', 2, '该设备负责人可以删除设备负责人'),
  ('device_charge_u', 2, '该设备负责人可以修改设备负责人'),
  ('device_charge_ability_c', 2, '该设备负责人可以增加负责人的权限'),
  ('device_charge_ability_d', 2, '该设备负责人可以删除负责人的权限'),
  ('device_param_c', 2, '该设备负责人可以增加设备检测参数'),
  ('device_param_u', 2, '该设备负责人可以修改设备检测参数'),
  ('device_param_d', 2, '该设备负责人可以删除设备检测参数');

INSERT INTO role (role_name, status, created_at, updated_at) values ('default', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO role (role_name, status, created_at, updated_at, is_admin) values ('admin', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true);

CREATE OR REPLACE FUNCTION insert_privs_of_admin_role()
RETURNS VOID AS $$
DECLARE
    rec INTEGER;
BEGIN
    FOR rec IN
  SELECT id FROM privilege WHERE priv_type != '2'
    LOOP
   	INSERT INTO role_priv (role_id, privilege_id) values ((SELECT id FROM role WHERE role_name = 'admin'), rec);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
select insert_privs_of_admin_role();
