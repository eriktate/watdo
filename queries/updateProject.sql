UPDATE
	projects
SET
	name = :name,
	description = :description
WHERE
	id = :id;
