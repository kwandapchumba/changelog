-- name: AddOtp :one
INSERT INTO otp (otp, email, expiry)
VALUES ($1, $2, $3)
ON CONFLICT (email) DO UPDATE SET otp = EXCLUDED.otp, expiry = EXCLUDED.expiry
RETURNING *;

-- name: GetOtp :one
SELECT * FROM otp WHERE otp = $1 LIMIT 1;

-- name: UpdateOtp :exec
UPDATE otp SET verified = 'true' WHERE otp = $1;