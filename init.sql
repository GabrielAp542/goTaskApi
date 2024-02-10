-- Insertar datos de ejemplo en la tabla de usuarios (users)
INSERT INTO Users (user_id, username) VALUES
(1, 'Usuario 1'),
(2, 'Usuario 2'),
(3, 'Usuario 3');

-- Insertar datos de ejemplo en la tabla de tareas (tasks)
INSERT INTO tasks (task_id, user_id, task_name) VALUES
(1, 1, 'Tarea 1 del Usuario 1'),
(2, 1, 'Tarea 2 del Usuario 1'),
(3, 2, 'Tarea 1 del Usuario 2'),
(4, 2, 'Tarea 2 del Usuario 2'),
(5, 3, 'Tarea 1 del Usuario 3'),
(6, 3, 'Tarea 2 del Usuario 3');