INSERT INTO privilege(sign, priv_type, name) values
	('user_r', 0, '用户账号读权限'),
	('user_w', 0, '用户账号写权限'),
	('user_login_r', 0, '用户登录记录读权限'),
	('user_extend_r', 0, '用户资料读权限'),
	('user_extend_w', 0, '用户资料写权限'),
	('role_r', 0, '角色信息读权限'),
	('role_w', 0, '角色信息写权限'),
	('device_r', 0, '设备读权限'),
	('device_w', 0, '设备写权限'),
	('device_charge_r', 1, '当前设备负责人信息读权限'),
  ('device_charge_w', 1, '当前设备负责人信息写权限'),
	('device_param_r', 1, '当前设备参数读权限'),
	('device_param_w', 1, '当前设备参数写权限');

INSERT INTO role (role_name, status, created_at, updated_at) values ('default', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO role (role_name, status, created_at, updated_at) values ('admin', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

CREATE OR REPLACE FUNCTION insert_privs_of_admin_role()
RETURNS VOID AS $$
DECLARE
    rec INTEGER;
BEGIN
    FOR rec IN
	SELECT id FROM privilege WHERE priv_type = '0'
    LOOP
   	INSERT INTO role_priv (role_id, privilege_id) values (1, rec);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
select insert_privs_of_admin_role();
