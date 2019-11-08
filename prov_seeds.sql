INSERT INTO privilege(sign, priv_type, name) values
  ('device_c', 0, '创建设备'),

  ('admin_user_r', 1, '读用户账号信息'),
  ('admin_user_u', 1, '修改用户账号信息'),
  ('admin_user_d', 1, '删除用户账号'),
  ('admin_user_login_r', 1, '查看用户登录记录'),
  ('admin_user_extend_r', 1, '查看用户资料'),
  ('admin_user_extend_w', 1, '修改用户资料'),
  ('admin_role_c', 1, '创建角色'),
  ('admin_role_u', 1, '修改角色信息'),
  ('admin_role_d', 1, '删除角色'),
  ('admin_role_priv_w', 1, '分配角色权限'),
  ('admin_device_r', 1, '查看任何设备'),
  ('admin_device_u', 1, '修改任何设备'),
  ('admin_device_d', 1, '删除任何设备');

INSERT INTO role (role_name, status, created_at, updated_at) values ('default', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO role (role_name, status, created_at, updated_at, is_admin) values ('admin', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true);

CREATE OR REPLACE FUNCTION insert_privs_of_admin_role()
RETURNS VOID AS $$
DECLARE
    rec INTEGER;
BEGIN
    FOR rec IN
  SELECT id FROM privilege
    LOOP
   	INSERT INTO role_priv (role_id, privilege_id) values ((SELECT id FROM role WHERE role_name = 'admin'), rec);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
select insert_privs_of_admin_role();
