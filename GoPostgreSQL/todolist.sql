
/*
listar tablas \dt
listar usuario \d
*/
DROP DATABASE IF EXISTS postgres;
/*CREATE DATABASE postgres;*/
DROP TABLE todo_list;
CREATE TABLE todo_list (
  id SERIAL PRIMARY KEY,
  task VARCHAR(255) NOT NULL,
  status BOOLEAN DEFAULT false NOT NULL,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  due_date DATE
);

INSERT INTO todo_list (task, status, due_date) VALUES ('Lavar la ropa', false, '2022-12-31');
INSERT INTO todo_list (task, status, due_date) VALUES ('Comprar comida', false, '2022-04-25');
INSERT INTO todo_list (task, status, due_date) VALUES ('Llamar a mi madre', false, '2022-02-24');
INSERT INTO todo_list (task, status, due_date) VALUES ('Estudiar para el examen', false, '2022-01-26');
INSERT INTO todo_list (task, status, due_date) VALUES ('Limpiar la casa', false, '2022-08-07');
