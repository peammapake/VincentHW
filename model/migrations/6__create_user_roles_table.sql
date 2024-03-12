CREATE TABLE user_roles (
    id 				UUID DEFAULT uuid_generate_v4() PRIMARY KEY, 
    role_name 		VARCHAR(20) 	NOT NULL, 
    created_at      TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP
);