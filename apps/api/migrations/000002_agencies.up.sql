CREATE TABLE IF NOT EXISTS agencies (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    name TEXT UNIQUE NOT NULL,
    address TEXT UNIQUE NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TRIGGER update_timestamp_agencies AFTER UPDATE ON agencies
FOR EACH ROW
BEGIN
    UPDATE agencies SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
