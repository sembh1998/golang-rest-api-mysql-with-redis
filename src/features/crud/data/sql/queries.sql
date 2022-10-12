-- name: GetAllEmployees :many
SELECT 
    * 
FROM 
    employee e 
    left join department d on e.department_id = d.id;

-- name: GetEmployee :one
SELECT 
    * 
FROM 
    employee e 
    left join department d on e.department_id = d.id
WHERE 
    e.id = sqlc.arg(id);

-- name: CreateEmployee :execresult
INSERT INTO employee (name, department_id) VALUES (sqlc.arg(name), sqlc.arg(department_id));

-- name: UpdateEmployee :execresult
UPDATE employee SET name = sqlc.arg(name), department_id = sqlc.arg(department_id) WHERE id = sqlc.arg(id);

-- name: DeleteEmployee :execresult
DELETE FROM employee WHERE id = sqlc.arg(id);

-- name: GetAllDepartments :many
SELECT 
    *
FROM 
    department d;

-- name: GetDepartment :one
SELECT 
    *
FROM
    department d
WHERE
    id = sqlc.arg(id);

-- name: CreateDepartment :execresult
INSERT INTO department (name) VALUES (sqlc.arg(name));

-- name: UpdateDepartment :execresult
UPDATE department SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteDepartment :execresult
DELETE FROM department WHERE id = sqlc.arg(id);

-- name: GetDepartmentEmployees :many
SELECT 
    *
FROM
    employee e
WHERE
    department_id = sqlc.arg(id);

