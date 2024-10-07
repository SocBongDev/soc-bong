CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    email TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    is_active BOOLEAN DEFAULT FALSE NOT NULL,
    verify_email BOOLEAN DEFAULT FALSE NOT NULL,
    dob DATE NOT NULL,
    phone_number TEXT NOT NULL,
    connection TEXT NOT NULL,
    password_hash TEXT NOT NULL,

    auth0_user_id TEXT NOT NULL,
    
    agency_id INTEGER NOT NULL,

    FOREIGN KEY (agency_id) REFERENCES agencies(id)
);

CREATE INDEX idx_users_email ON users(email);

CREATE TRIGGER update_timestamp_users AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
