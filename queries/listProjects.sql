-- $1 = account_id
SELECT * FROM projects
WHERE ($1::uuid IS NULL OR account_id = $1::uuid); -- type cast is necessary so that pgx doesn't fail to infer it
