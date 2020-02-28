CREATE TABLE User (
  id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
  login  TEXT,
  password TEXT,
  role INTEGER
);

CREATE TABLE History (
    id  INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    dates TEXT,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES User(id)
);

CREATE TABLE Commands (
    id  INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    dates TEXT,
    commands TEXT,
    command_id INTEGER,
    FOREIGN KEY(command_id) REFERENCES User(id),
    FOREIGN KEY(command_id) REFERENCES History(id)
);
