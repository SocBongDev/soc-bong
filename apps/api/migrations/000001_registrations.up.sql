CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    is_processed BOOLEAN DEFAULT FALSE NOT NULL,
    note TEXT,
    parent_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    student_class TEXT NOT NULL,
    student_dob DATE NOT NULL,
    student_name TEXT NOT NULL,

    agency_id INTEGER NOT NULL DEFAULT 1,
    UNIQUE(parent_name, phone_number, student_class, student_dob, student_name)
);

CREATE TRIGGER update_timestamp_registrations AFTER UPDATE ON registrations
FOR EACH ROW
BEGIN
    UPDATE registrations SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
