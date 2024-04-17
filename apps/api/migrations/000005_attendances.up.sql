CREATE TABLE IF NOT EXISTS attendances (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    attended_at DATE NOT NULL,
    attended_status INTEGER CHECK(attended_status IN (0, 1, 2, 3, 4)) DEFAULT 0 NOT NULL,
    class_id INTEGER NOT NULL,
    student_id INTEGER NOT NULL,

    FOREIGN KEY (class_id) REFERENCES classes (id),
    FOREIGN KEY (student_id) REFERENCES students (id),
    UNIQUE(attended_at, attended_status, student_id)
);

CREATE INDEX class_id_checked_at_idx 
ON attendances (class_id, attended_at);

CREATE TRIGGER update_timestamp_attendances AFTER UPDATE ON attendances
FOR EACH ROW
BEGIN
    UPDATE attendances SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
