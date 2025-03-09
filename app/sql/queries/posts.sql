-- name: CreatePost :one
INSERT INTO postsBaspana (
    id, 
    created_at, 
    updated_at, 
    date_publication,
    title, 
    image, 
    cost_for_metr,
    address, 
    link_detail_post, 
    number_object, 
    count_access,
    region
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: GetPost :one
SELECT *
FROM postsBaspana
WHERE number_object = $1
AND link_detail_post = $2;