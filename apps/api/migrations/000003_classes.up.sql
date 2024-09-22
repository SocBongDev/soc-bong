CREATE TABLE IF NOT EXISTS classes (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    name TEXT NOT NULL,
    grade TEXT CHECK(grade IN ('buds', 'seed', 'leaf')) NOT NULL,

    agency_id INTEGER NOT NULL,
    teacher_id TEXT NOT NULL,

    FOREIGN KEY (agency_id) REFERENCES agencies (id)
    UNIQUE(name, grade, agency_id, teacher_id)
);

CREATE TRIGGER update_timestamp_classes AFTER UPDATE ON classes
FOR EACH ROW
BEGIN
    UPDATE classes SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
