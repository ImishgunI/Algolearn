DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'execution_status') THEN
        CREATE TYPE execution_status AS ENUM ('pending', 'running', 'finished', 'failed');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS execution_sessions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INT NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    status execution_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ DEFAULT now()
);