CREATE TABLE discounts
(
    uuid       UUID        NOT NULL DEFAULT gen_random_uuid(),
    created_by UUID        NOT NULL,
    percent    REAL        NOT NULL,
    started_at TIMESTAMP   NOT NULL,
    ended_at   TIMESTAMP   NOT NULL,
    status     VARCHAR NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);

CREATE
    OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
        = NOW();
    RETURN NEW;
END;
$$
    language 'plpgsql';

CREATE TRIGGER update_discounts_updated_at
    BEFORE UPDATE
    ON discounts
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
