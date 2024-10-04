CREATE TABLE discounts
(
    uuid        UUID        NOT NULL DEFAULT gen_random_uuid(),
    created_by  UUID        NOT NULL,
    percent     FLOAT(2)    NOT NULL,
    started_at  TIMESTAMP   NOT NULL,
    ended_at    TIMESTAMP   NOT NULL,
    status      VARCHAR(32) NOT NULL,
    created_at  TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);

CREATE TABLE discounts_products
(
    discount_uuid UUID,
    product_uuid  UUID,
    PRIMARY KEY (discount_uuid, product_uuid),
    FOREIGN KEY (discount_uuid) REFERENCES discounts (uuid),
    FOREIGN KEY (product_uuid) REFERENCES products (uuid)
);

CREATE TRIGGER update_discounts_updated_at
    BEFORE UPDATE ON discounts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();
