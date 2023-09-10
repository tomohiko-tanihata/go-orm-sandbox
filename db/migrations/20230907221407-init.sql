-- +migrate Up
CREATE TABLE grades
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE classes
(
    id       SERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    grade_id INT REFERENCES grades (id)
);
CREATE TABLE students
(
    id       SERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    class_id INT REFERENCES classes (id)
);
-- +migrate Down
DROP TABLE students;
DROP TABLE classes;
DROP TABLE grades;
