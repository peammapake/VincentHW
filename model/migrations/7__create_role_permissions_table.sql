CREATE TABLE role_permissions (
    role_id 	    UUID            NOT NULL, 
    permission 		VARCHAR(40) 	NOT NULL, 
    created_at      TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP,
    PRIMARY KEY (role_id,permission)
);