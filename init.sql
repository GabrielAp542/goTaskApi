CREATE TABLE IF NOT EXISTS tasks (
    task_id SERIAL PRIMARY KEY,
    task_name VARCHAR(150) NOT NULL,
    compleated BOOLEAN
);