ALTER TABLE coopera.tasks
    ADD COLUMN tags     TEXT[]      NOT NULL DEFAULT '{}',
    ADD COLUMN priority VARCHAR(20) NOT NULL DEFAULT 'low';
