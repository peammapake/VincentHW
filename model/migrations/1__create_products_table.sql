CREATE TABLE products (
    id                      UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    code                    VARCHAR(20) UNIQUE NOT NULL,
    name                    VARCHAR(50) NOT NULL,
    description             VARCHAR(200),
    highlight_cover_img_url VARCHAR(200),
    brand_id                UUID,
    category_id             UUID,
    net_price               NUMERIC(15, 2),
    gross_price             NUMERIC(15, 2),
    stock                   INTEGER DEFAULT 0 NOT NULL,
    created_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP
);