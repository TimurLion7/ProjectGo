CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);


/*INSERT INTO tasks (task, is_done)
VALUES 
    ('task1', true),
    ('task2', true),
    ('task3', true),
    ('task4', true),
    ('task5', true)*/