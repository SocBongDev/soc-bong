CREATE TABLE IF NOT EXISTS parents (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    parent_name TEXT NOT NULL,
    parent_dob DATE,
    parent_gender BOOLEAN,
    phone_number TEXT,
    zalo TEXT,
    occupation TEXT,    
    lanlord TEXT,
    roi TEXT,
    parent_birth_place TEXT,
    res_registration TEXT,
    
    student_id INTEGER NOT NULL,
    
    FOREIGN KEY (student_id) REFERENCES students (id),
    UNIQUE(student_id, parent_name, parent_gender)
);

CREATE TRIGGER update_timestamp_parents AFTER UPDATE ON parents
FOR EACH ROW
BEGIN
    UPDATE parents SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
