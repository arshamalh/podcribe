CREATE TABLE IF NOT EXISTS podcasts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    page_link VARCHAR,
    mp3_link VARCHAR,
    provider VARCHAR,
    mp3_path VARCHAR,
    wav_path VARCHAR,
    transcription_path VARCHAR,
    translation_path VARCHAR,
    referenced_count INTEGER DEFAULT 0,
    created_at TIMESTAMP
);