CREATE TABLE product_images (
    product_id      UUID NOT NULL,
    image_url       VARCHAR(200) NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (product_id, image_url)
);
