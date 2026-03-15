-- Ejecuta esto en tu SQL Server (EvaluacionGoland)
CREATE TABLE game_library (
    id INT IDENTITY(1,1) PRIMARY KEY,
    rawg_id INTEGER NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    platform VARCHAR(100),
    cover_url TEXT,
    personal_note TEXT,
    personal_score INTEGER,
    status VARCHAR(20),
    added_at DATETIME DEFAULT GETDATE()
);