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
  snackName text,
  snackDesc text,
  snackCat text,
  snackPic text,
  healthScore int,
  PRIMARY KEY (snackName)
);

create table snackSighting (
  snackName text,
  CONSTRAINT fk_snack_name FOREIGN KEY (snackName) REFERENCES snacks(snackName),
  sightTime timestamp,
  sightLocation text,
  sightEstDuration timestamp,
  CONSTRAINT pk_sighting PRIMARY KEY (snackName, sightTime, sightLocation)
);