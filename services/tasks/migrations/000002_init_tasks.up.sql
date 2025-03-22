CREATE TABLE tasks
(
    uuid       UUID      NOT NULL DEFAULT gen_random_uuid(),
    name       VARCHAR   NOT NULL,
    text       VARCHAR   NOT NULL,
    status     VARCHAR   NOT NULL,
    created_by UUID      NOT NULL,
    started_at TIMESTAMP NOT NULL,
    ended_at   TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);


CREATE TRIGGER update_tasks_updated_at
    BEFORE UPDATE
    ON tasks
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
