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
  snackId serial,
  snackName text,
  snackDesc text,
  snackCat text,
  snackPic text,
  healthScore int,
  PRIMARY KEY (snackId)
);

create table snackSighting (
  snackId int,
  CONSTRAINT fk_snack_id FOREIGN KEY (snackId) REFERENCES snacks(snackId),
  sightTime timestamp,
  sightLocation text,
  sightEstDuration timestamp,
  CONSTRAINT pk_sighting PRIMARY KEY (snackId, sightTime, sightLocation)
);