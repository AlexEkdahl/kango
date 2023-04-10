-- init.sql

CREATE TABLE IF NOT EXISTS task (
    id BIGSERIAL PRIMARY KEY,
    status INTEGER NOT NULL,
    subject VARCHAR(255) NOT NULL,
    description TEXT
);


CREATE TABLE IF NOT EXISTS board (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS task_board (
    task_id BIGINT NOT NULL,
    board_id BIGINT NOT NULL,
    PRIMARY KEY (task_id, board_id),
    FOREIGN KEY (task_id) REFERENCES task (id) ON DELETE CASCADE,
    FOREIGN KEY (board_id) REFERENCES board (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_task_id ON task (id);
CREATE INDEX IF NOT EXISTS idx_board_id ON board (id);
CREATE INDEX IF NOT EXISTS idx_task_board_task_id ON task_board (task_id);
CREATE INDEX IF NOT EXISTS idx_task_board_board_id ON task_board (board_id);
