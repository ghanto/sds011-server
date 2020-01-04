CREATE TABLE smog_entries (
  id          serial    PRIMARY KEY,
  pm10        numeric   NOT NULL,
  pm25        numeric   NOT NULL,
  temperature numeric   NOT NULL,
  humidity    numeric   NOT NULL,
  created_at  timestamp NOT NULL
);
