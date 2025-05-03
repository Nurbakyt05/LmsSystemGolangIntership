-- +goose Up
-- +goose StatementBegin

-- Курс
INSERT INTO courses (name, description)
VALUES
    ('Golang Developer', 'Learn how to build scalable backend systems with Go'),
    ('Python Developer', 'Master Python for data science and backend development');

-- Главы
INSERT INTO chapters (name, description, "order", course_id)
VALUES
    ('Control Structures', 'if, else, switch in Go', 1, 1),
    ('Data Types', 'basic types, slices, maps', 2, 1),
    ('Python Basics', 'variables, loops, conditions', 1, 2);

-- Уроки
INSERT INTO lessons (name, description, content, "order", chapter_id)
VALUES
    ('If-Else Statement in Go', 'Understanding conditionals', 'In Go, the if statement...', 1, 1),
    ('Switch in Go', 'Using switch-case', 'Switch lets you replace multiple if-else blocks...', 2, 1),
    ('Slices in Go', 'Dynamic arrays', 'Slices are references to arrays...', 1, 2),
    ('Variables in Python', 'Declaring and using variables', 'In Python, variables are created when first assigned...', 1, 3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM lessons;
DELETE FROM chapters;
DELETE FROM courses;
-- +goose StatementEnd
