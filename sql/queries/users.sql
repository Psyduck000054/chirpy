-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
select * from users
where id = $1;

-- name: EditUserInfo :one
update users set
hashed_password = $2,
email = $3,
updated_at = now()
where id = $1
returning *;

-- name: AddChirpyRed :one
update users set is_chirpy_red = true where id = $1
returning *;