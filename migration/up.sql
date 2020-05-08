-- create schema
CREATE TABLE IF NOT EXISTS accounts(
	id UUID PRIMARY KEY,
	name VARCHAR(512) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users(
	id UUID PRIMARY KEY,
	default_account_id UUID DEFAULT NULL REFERENCES accounts(id),
	name VARCHAR(512) NOT NULL,
	email VARCHAR(320) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS projects(
	id UUID PRIMARY KEY,
	account_id UUID NOT NULL REFERENCES accounts(id),
	name VARCHAR(512) NOT NULL,
	description TEXT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_accounts(
	user_id UUID NOT NULL REFERENCES users(id),
	account_id UUID NOT NULL REFERENCES accounts(id),
	default_project_id UUID REFERENCES projects(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(user_id, account_id)
);

CREATE TABLE IF NOT EXISTS task_statuses(
	id UUID PRIMARY KEY,
	account_id UUID NOT NULL REFERENCES accounts(id),
	name VARCHAR(256) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks(
	id UUID PRIMARY KEY,
	title VARCHAR(512) NOT NULL,
	description TEXT,
	project_id UUID NOT NULL REFERENCES projects(id),
	reporter_id UUID NOT NULL REFERENCES users(id),
	status_id UUID NOT NULL REFERENCES task_statuses(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS assigned_users(
	task_id UUID NOT NULL REFERENCES tasks(id),
	user_id UUID NOT NULL REFERENCES users(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- load some data
INSERT INTO accounts (id, name)
VALUES (
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'wrkhub test'
),
(
	'c3dd6b60-553c-4995-abda-c1573a720d73',
	'Test Account'
) ON CONFLICT DO NOTHING;

INSERT INTO users (id, name, email, default_account_id)
VALUES (
	'd163193e-6f6a-4a71-92a1-c76d3148559a',
	'Test User',
	'test@watdo.app',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac'
) ON CONFLICT DO NOTHING;

INSERT INTO projects (id, account_id, name, description)
VALUES (
	'93305d12-6186-4002-a180-1d93ea9f74cb',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'Test Project',
	'A test project'
) ON CONFLICT DO NOTHING;

INSERT INTO task_statuses (id, account_id, name)
VALUES (
	'95836ddf-a578-4282-8db1-10a04abfd220',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'ToDo'
) ON CONFLICT DO NOTHING;

INSERT INTO task_statuses (id, account_id, name)
VALUES (
	'88a75370-e482-4c14-9181-80ffe238dd60',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'In Progress'
) ON CONFLICT DO NOTHING;

INSERT INTO task_statuses (id, account_id, name)
VALUES ('22451d49-c372-4df6-b629-8cd248bf584b',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'Done'
) ON CONFLICT DO NOTHING;

INSERT INTO user_accounts (user_id, account_id, default_project_id)
VALUES (
	'd163193e-6f6a-4a71-92a1-c76d3148559a',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'93305d12-6186-4002-a180-1d93ea9f74cb'
),
(
	'd163193e-6f6a-4a71-92a1-c76d3148559a',
	'c3dd6b60-553c-4995-abda-c1573a720d73',
	NULL
) ON CONFLICT DO NOTHING;

INSERT INTO task_statuses (id, account_id, name)
VALUES (
	'2470e4eb-9702-4769-b2e3-c3c410e6f258',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'Todo'
), (
	'ac04d1b9-64a5-474c-9517-531cbc246b42',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'In Progress'
), (
	'b5082c3b-ba21-4bec-8e39-ef0d0c68cb6a',
	'5f2d9a95-b335-493f-a7ed-8dbd81027bac',
	'Done'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (id, title, description, project_id, reporter_id, status_id)
VALUES (
	'6ee8cf36-ee06-457a-b7f3-a96eecacdd9b',
	'Test Task 1',
	'Just a simple test task',
	'93305d12-6186-4002-a180-1d93ea9f74cb',
	'd163193e-6f6a-4a71-92a1-c76d3148559a',
	'2470e4eb-9702-4769-b2e3-c3c410e6f258'
), (
	'0b74ec3c-f6df-4b5d-b5ac-e4a013e0004f',
	'Test Task 2',
	'Another simple task',
	'93305d12-6186-4002-a180-1d93ea9f74cb',
	'd163193e-6f6a-4a71-92a1-c76d3148559a',
	'2470e4eb-9702-4769-b2e3-c3c410e6f258'
), (
	'adf970d9-f30c-462a-9db9-59c52c06e30b',
	'Test Task 3',
	'This task is kind of pointless',
	'93305d12-6186-4002-a180-1d93ea9f74cb',
	'd163193e-6f6a-4a71-92a1-c76d3148559a',
	'ac04d1b9-64a5-474c-9517-531cbc246b42'
) ON CONFLICT DO NOTHING;
