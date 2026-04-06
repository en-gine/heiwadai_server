UPDATE user_data SET prefecture = 0 WHERE prefecture IS NULL;
ALTER TABLE user_data ALTER COLUMN prefecture SET NOT NULL;
ALTER TABLE user_data ALTER COLUMN prefecture DROP DEFAULT;
