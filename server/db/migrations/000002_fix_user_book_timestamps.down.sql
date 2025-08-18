-- Revert user_book table to use TIMESTAMP for stay_from and stay_to columns
ALTER TABLE user_book 
    ALTER COLUMN stay_from TYPE TIMESTAMP,
    ALTER COLUMN stay_to TYPE TIMESTAMP;

-- Revert book_plan_stay_date_info table
ALTER TABLE book_plan_stay_date_info
    ALTER COLUMN stay_date TYPE TIMESTAMP;