-- create table snacks (
--   id serial,
--   sname text,
--   sdesc text,
--   category text,
--   simage text,
--   healthScore int,
--   PRIMARY KEY (id)
-- );

-- create table snackSightings (
--   snackid int,
--   CONSTRAINT fk_snack_id FOREIGN KEY (snackId) REFERENCES snacks(id),
--   sightTime timestamp,
--   sightLocation text,
--   CONSTRAINT pk_sighting PRIMARY KEY (snackId, sightTime, sightLocation)
-- );


-- INSERT INTO snacks (sname, sdesc, category, simage, healthScore) VALUES ('KitKat', 'Chocolate wafer', 'Chocolate bar', NULL, 0);
-- INSERT INTO snacks (sname, sdesc, category, simage, healthScore) VALUES ('Smoked almond', 'Salted smoked almonds', 'Nuts', NULL, 0);
-- INSERT INTO snacks (sname, sdesc, category, simage, healthScore) VALUES ('Apple', 'A round fruit', 'Fruit', NULL, 0);


create table snackSightings (
  sname text,
  simage text NOT NULL DEFAULT '',
  sightTime timestamp,
  sightLocation text,
  sighter text NOT NULL DEFAULT '',
  CONSTRAINT pk_sighting PRIMARY KEY (sname, sightTime, sightLocation)
);