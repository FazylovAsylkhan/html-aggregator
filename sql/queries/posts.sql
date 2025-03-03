-- name: CreatePost :one
INSERT INTO postsBaspana (
    id, 
    created_at, 
    updated_at, 
    title, 
    image, 
    cost_for_metr,
    address, 
    link_detail_post, 
    number_object, 
    count_access
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

