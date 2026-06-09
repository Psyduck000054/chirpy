-- name: SaveChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: RetrieveAllChirps_Asc :many
select * from chirps order by created_at asc;
-- name: RetrieveAllChirps_Desc :many
select * from chirps order by created_at desc;


-- name: RetrieveChirpsByAuthorID_Asc :many
select * from chirps where user_id = $1 order by created_at asc;
-- name: RetrieveChirpsByAuthorID_Desc :many
select * from chirps where user_id = $1 order by created_at desc;

-- name: GetChirp :one
select * from chirps where id = $1;

-- name: DeleteChirp :exec
delete from chirps where id = $1;