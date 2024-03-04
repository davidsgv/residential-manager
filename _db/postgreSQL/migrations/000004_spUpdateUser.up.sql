CREATE PROCEDURE sp_update_user(
    uuid_param VARCHAR,
    rol_param VARCHAR,
    block_param VARCHAR,
    number_param VARCHAR
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

    UPDATE user_account
    SET rol_id = rol_id_var, apartment_id = apartment_id_var
    WHERE uuid = uuid_param;
end;$$;