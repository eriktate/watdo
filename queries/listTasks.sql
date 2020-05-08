SELECT *
FROM tasks
WHERE ($1::uuid IS NULL OR project_id = $1::uuid); -- type cast is necessary so that pgx doesn't fail to infer it
