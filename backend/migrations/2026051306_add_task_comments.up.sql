CREATE TABLE IF NOT EXISTS coopera.task_comments (
    id         SERIAL PRIMARY KEY,
    task_id    INT          NOT NULL REFERENCES coopera.tasks(id) ON DELETE CASCADE,
    user_id    INT          NOT NULL REFERENCES coopera.users(id) ON DELETE CASCADE,
    username   VARCHAR(100) NOT NULL,
    text       TEXT         NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_task_comments_task_id ON coopera.task_comments(task_id);
