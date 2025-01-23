-- name: CreateURL :exec
INSERT INTO urls (
    original_url,
    short_code,
    is_custom,
    expired_at
) VALUES (
     ?,?,?,?
);

-- name: GetInsertedURL :one
SELECT * FROM urls WHERE id = LAST_INSERT_ID();


-- name: IsShortCodeAvailable :one
SELECT NOT EXISTS(
    SELECT 1 FROM urls
    WHERE short_code = ?
) AS is_available;

-- name: GetURLByShortCode :one
SELECT * FROM urls
WHERE short_code = ?
AND expired_at > CURRENT_TIMESTAMP();
