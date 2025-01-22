-- name: CreateURL :exec
INSERT INTO urls (
    original_url,
    short_code,
    is_custom,
    expired_at
) VALUES (
    $1, $2, $3, $4
);

-- name: GetInsertedURL :one
SELECT * FROM urls WHERE id = LAST_INSERT_ID();
