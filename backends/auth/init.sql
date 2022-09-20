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

create table snacks (
  snackId serial NOT NULL,
  snackName text,
  snackDesc text,
  snackCat text,
  snackPic text,
  healthScore int,
  PRIMARY KEY (snackId)
);

create table snackSighting (
  snack_id text,
  CONSTRAINT fk_snack_id FOREIGN KEY (snack_id) REFERENCES snacks(snack_id),
  sightTime dateTime,
  sightLocation text,
  sightEstDuration dateTime,
  CONSTRAINT pk_sighting PRIMARY KEY (snack_id, sightTime, sightLocation)
);