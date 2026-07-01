CREATE SCHEMA messenger;

CREATE TABLE messenger.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW ()
);

CREATE TABLE messenger.chats (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE messenger.chats_participants (
    chat_id INTEGER REFERENCES messenger.chats(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES messenger.users(id) ON DELETE CASCADE,
    PRIMARY KEY (chat_id, user_id)
);

CREATE TABLE messenger.messages (
    id SERIAL PRIMARY KEY,
    chat_id INTEGER REFERENCES messenger.chats(id) ON DELETE CASCADE,
    sender_id INTEGER REFERENCES messenger.users(id) ON DELETE CASCADE,
    version BIGINT NOT NULL DEFAULT 1,
    message TEXT NOT NULL,
    is_delivered BOOLEAN NOT NULL DEFAULT FALSE,
    sent_at TIMESTAMPTZ DEFAULT NOW()
);
