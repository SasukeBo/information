INSERT INTO privilege(priv_name, priv_type, status) values
	('user_r', 0, 0),
	('user_w', 0, 0),
	('user_login_r', 0, 0),
	('user_login_w', 0, 0),
	('user_extend_r', 0, 0),
	('user_extend_w', 0, 0),
	('privilege_r', 0, 0),
	('privilege_w', 0, 0),
	('role_r', 0, 0),
	('role_w', 0, 0),
	('role_priv_r', 0, 0),
	('role_priv_w', 0, 0),
	('device_r', 0, 0),
	('device_w', 0, 0),
	('device_charge_r', 0, 0),
	('device_charge_w', 0, 0),
	('device_charge_ability_r', 0, 0),
	('device_charge_ability_w', 0, 0),
	('device_param_r', 0, 0),
	('device_param_w', 0, 0),
	('device_param_value_r', 0, 0),
	('device_param_value_w', 0, 0),
	('device_status_log_r', 0, 0),
	('device_status_log_w', 0, 0);

INSERT INTO role (role_name, status, created_at, updated_at) values ('admin', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

CREATE OR REPLACE FUNCTION insert_privs_of_admin_role()
RETURNS VOID AS $$
DECLARE
    rec INTEGER;
BEGIN
    FOR rec IN
	SELECT id FROM privilege
    LOOP
   	INSERT INTO role_priv (role_id, privilege_id) values (1, rec);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
select insert_privs_of_admin_role();
