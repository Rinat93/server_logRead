CREATE TABLE IF NOT EXISTS User (
  id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  login  TEXT,
  password TEXT,
  dates_reg TEXT
);

CREATE TABLE IF NOT EXISTS History (
    id  INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    dates TEXT,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES User(id)
);

CREATE TABLE  IF NOT EXISTS Commands (
    id  INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    dates TEXT,
    commands TEXT,
    command_id INTEGER,
    FOREIGN KEY(command_id) REFERENCES History(id)
);
