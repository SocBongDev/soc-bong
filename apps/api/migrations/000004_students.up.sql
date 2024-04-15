CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    enrolled_at DATE,
    dob DATE,
    gender BOOLEAN,
    ethnic TEXT,
    birth_place TEXT,
    temp_address TEXT,
    permanent_address_province TEXT,
    permanent_address_district TEXT,
    permanent_address_commune TEXT,
    
    agency_id INTEGER NOT NULL,
    class_id INTEGER NOT NULL,

    FOREIGN KEY (agency_id) REFERENCES agencies (id)
    FOREIGN KEY (class_id) REFERENCES classes (id)
);

CREATE TRIGGER update_timestamp_students AFTER UPDATE ON students
FOR EACH ROW
BEGIN
    UPDATE students SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
