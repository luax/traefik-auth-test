DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
  id INTEGER PRIMARY KEY,
  message TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);
INSERT INTO messages (id, message, created_at) VALUES (1, 'Hello world', now());
INSERT INTO messages (id, message, created_at) VALUES (2, 'Foo bar', now());
INSERT INTO messages (id, message, created_at) VALUES (3, 'whaaat', now());
