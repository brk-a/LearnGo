CREATE TABLE IF NOT EXISTS `accounts` (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    owner VARCHAR NOT NULL,
    balance INTEGER NOT NULL DEFAULT 0,
    currency VARCHAR NOT NULL DEFAULT 'KES',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
);