create table users (
  username text primary key,
  password text
);

create table user_sessions (
  token text primary key,
  username text,
  expiry timestamp,
  CONSTRAINT fk_session_user FOREIGN KEY (username) REFERENCES users(username)
);