-- Создание пользователя
CREATE USER backend_user WITH PASSWORD 'password';

-- Создание базы данных
CREATE DATABASE backend_db OWNER backend_user;

-- Подключение к созданной базе данных
\c backend_db;

-- Создание таблицы для задач
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    script_name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Создание таблицы для результатов
CREATE TABLE results (
    id SERIAL PRIMARY KEY,
    task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
    data TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Создание таблицы для пользователей (если нужна аутентификация)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Предоставление прав пользователю на таблицы
GRANT ALL PRIVILEGES ON DATABASE backend_db TO backend_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO backend_user;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO backend_user;