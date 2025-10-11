-- migrate:up

INSERT INTO users
(username, email, password_hash, full_name, is_active) VALUES (
  'root',
  'root@mail.com',
  '$2a$10$.OAOVv7cVhHuwPOJgjK8beO5wM6r/R.ex7LIpgKHn4LJIwXl0iP0.', -- password = qweQWE123!@#
  'Superuser',
  true
);

INSERT INTO user_organization_roles (user_id, organization_id, role_id) VALUES (
  (SELECT id FROM users where username = 'root' LIMIT 1),
  (SELECT id FROM organizations where name = 'root' LIMIT 1),
  (SELECT id FROM roles where name = 'Admin' LIMIT 1)
)

-- migrate:down
DELETE FROM user_organization_roles WHERE 
user_id = (SELECT id FROM users WHERE username = 'root' LIMIT 1) AND 
organization_id = (SELECT id FROM organizations WHERE name = 'root' limit 1) AND
role_id = (SELECT id FROM roles WHERE name = 'Admin' LIMIT 1);

DELETE FROM users WHERE username = 'root';