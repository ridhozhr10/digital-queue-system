-- migrate:up
ALTER TABLE organizations ADD COLUMN is_root BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE permissions DROP COLUMN created_at;
ALTER TABLE permissions DROP COLUMN updated_at;

INSERT INTO organizations (name, is_root) VALUES ('root', TRUE);

INSERT INTO roles (name) VALUES ('Admin');

INSERT INTO permissions (name) VALUES
('user:create'),
('user:view'),
('user:update'),
('user:delete');

-- migrate:down
DELETE FROM permissions WHERE name IN (
    'user:create', 'user:view', 'user:update', 'user:delete'
);
DELETE FROM roles WHERE name = 'Admin';
DELETE FROM organizations WHERE name = 'root';
ALTER TABLE organizations DROP COLUMN is_root;
ALTER TABLE permissions ADD COLUMN created_at timestamp DEFAULT (now());
ALTER TABLE permissions ADD COLUMN updated_at timestamp DEFAULT (now());

