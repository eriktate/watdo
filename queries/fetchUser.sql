SELECT
	u.*,
	ua.account_id,
	ua.default_project_id,
	a.name AS account_name
FROM users AS u
LEFT JOIN user_accounts AS ua ON ua.user_id = u.id
LEFT JOIN accounts AS a ON a.id = ua.account_id
WHERE
	u.id = $1;
