CREATE TABLE users (
    id 				UUID DEFAULT uuid_generate_v4() PRIMARY KEY, 
    first_name 		VARCHAR(50) 	NOT NULL, 
    last_name 		VARCHAR(50) 	NOT NULL, 
    username 		VARCHAR(30) 	NOT NULL,
	"password"  	VARCHAR(120)  	NOT NULL,
	role_id			UUID 			NOT NULL,
    created_at      TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP 		DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP
);