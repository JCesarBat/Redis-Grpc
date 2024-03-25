
-- name: CreatePersona :one
INSERT INTO "persona"(
    nombre,
    ocupacion
)values ($1,$2) returning *;