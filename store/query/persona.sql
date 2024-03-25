
-- name: CreatePersona :one
INSERT INTO "persona"(
    nombre,
    ocupacion,
    edad
)values ($1,$2,$3) returning *;