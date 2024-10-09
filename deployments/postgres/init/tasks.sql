CREATE DATABASE todoapp_db;

\c todoapp_db

DROP TABLE IF EXISTS task_status;
DROP TABLE IF EXISTS tasks;

CREATE TABLE task_status (
    id integer PRIMARY KEY NOT NULL UNIQUE,
    name varchar(64) NOT NULL
);

INSERT INTO
    task_status (
        id,
        name
    )
VALUES
    (
        0,
        'Pending'
    ),
    (
        1,
        'In Progress'
    ),
    (
        2,
        'Completed'
    );

CREATE TABLE tasks (
    id uuid PRIMARY KEY NOT NULL UNIQUE,
    name varchar(64) NOT NULL,
    description varchar(256) NOT NULL,
    status_id integer NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

INSERT INTO
    tasks (
        id,
        name,
        description,
        status_id,
        created_at,
        updated_at
    )
VALUES
    (
        'b81240b0-7122-4d06-bdb2-8bcf512d6c63',
        'Task One',
        'This is the first task',
        1,
        '2024-09-27 10:00:00',
        '2024-09-27 10:00:00'
    ),
    (
        'fad796a1-e0ed-4ee5-9f88-9b7258d35ae9',
        'Task Two',
        'This is the second task',
        2,
        '2024-09-27 10:10:00',
        '2024-09-27 10:15:00'
    ),
    (
        '07aaadbc-8967-406f-aebd-58b289377aef',
        'Task Three',
        'This is the third task',
        1,
        '2024-09-27 10:20:00',
        '2024-09-27 10:25:00'
    ),
    (
        '8b119430-438b-40d1-a28d-3d11d6afcfba',
        'Task Four',
        'This is the fourth task',
        3,
        '2024-09-27 10:30:00',
        '2024-09-27 10:35:00'
    ),
    (
        'b5e93ba4-ce33-4e36-83cb-c71177464a25',
        'Task Five',
        'This is the fifth task',
        2,
        '2024-09-27 10:40:00',
        '2024-09-27 10:45:00'
    );
