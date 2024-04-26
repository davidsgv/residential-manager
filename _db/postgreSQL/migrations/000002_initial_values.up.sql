/*
    script: initial values
    author: DavidSG
    description: populate initial values for the app
*/

INSERT INTO permit (id, name, description)
VALUES 
    (1 << 0, 'Create Admin', 'Puede crear administradores'),
    (1 << 1, 'Create Watchman', 'Puede crear vigilantes'),
    (1 << 2, 'Create Apartment Admin', 'Puede crear apartamentos'),
    (1 << 3, 'Create Resident', 'Puede crear residentes'),
    (1 << 4, 'Query Users', 'Puede consultar usuarios'),
    (1 << 5, 'Query Permissions', 'Puede consultar Permisos'),
    (1 << 6, 'Query Roles', 'Puede consultar Roles'),
    (1 << 7, 'Update Users', 'Puede actualizar usuarios');
    -- (1 << 5, "", "")
    -- (1 << 6, "", "")
    -- (1 << 7, "", "")
    -- (1 << 8, "", "")

INSERT INTO rol(name)
VALUES ('Administrador'), ('Vigilante'), ('Encargado'), ('Residente');

INSERT INTO permit_rol(rol_id, permit_id)
VALUES (1,1),(1,2),(1,4),(1,8),(1,16),(1,32),(1,64),(1,128);

INSERT INTO block (block) 
VALUES ('A');

INSERT INTO apartment (number, block_id)
VALUES ('101', 1);

INSERT INTO user_account (uuid, mail, password, verified, apartment_id, rol_id, token, token_expire)
VALUES ('32dc1674-6e29-4ba3-a190-b484b2b9fab1', 'example@gmail.com', 
    '313233a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a',
    true, 1, 1, 'first user', now());