CREATE PROCEDURE sp_create_user(
    uuid VARCHAR, 
    mail VARCHAR,
    password VARCHAR,
    verified boolean,
    rol_param VARCHAR,
    block_param VARCHAR,
    number_param VARCHAR,
    token_param VARCHAR,
    token_expire_param TIMESTAMP
)
language plpgsql    
as $$
declare
    block_id_var block.id%type;
    apartment_id_var apartment.id%type;
    rol_id_var rol.id%type;
begin
    --search block id
    SELECT id
	INTO block_id_var
	FROM block
	WHERE block = block_param;
	
	if not found then
        INSERT INTO block (block)
		VALUES (block_param)
		RETURNING id INTO block_id_var;
    end if;

    --search apartment id
    SELECT id
	INTO apartment_id_var
	FROM apartment
	WHERE number = number_param
        AND block_id = block_id_var;

    if not found then
        INSERT INTO apartment (number, block_id)
		VALUES (number_param, block_id_var)
		RETURNING id INTO apartment_id_var;
    end if;

    --search rol
    SELECT id
    INTO rol_id_var
    FROM rol
    WHERE name = rol_param;

    INSERT INTO user_account (uuid, mail, password, verified, apartment_id, rol_id, token, token_expire)
    VALUES (uuid, mail, password, verified, apartment_id_var, rol_id_var, token_param, token_expire_param);
end;$$;