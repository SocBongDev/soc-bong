CREATE TABLE IF NOT EXISTS registrations (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    is_processed BOOLEAN DEFAULT FALSE NOT NULL,
    note TEXT,
    parent_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    student_class TEXT NOT NULL,
    student_dob DATE NOT NULL,
    student_name TEXT NOT NULL
);
