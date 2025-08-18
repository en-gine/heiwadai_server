-- Fix user_book table to use TIMESTAMPTZ for stay_from and stay_to columns
-- This ensures timezone information is preserved and prevents date shifting issues

ALTER TABLE user_book 
    ALTER COLUMN stay_from TYPE TIMESTAMPTZ USING stay_from AT TIME ZONE 'Asia/Tokyo',
    ALTER COLUMN stay_to TYPE TIMESTAMPTZ USING stay_to AT TIME ZONE 'Asia/Tokyo';

-- Also fix the book_plan_stay_date_info table which has similar issue
ALTER TABLE book_plan_stay_date_info
    ALTER COLUMN stay_date TYPE TIMESTAMPTZ USING stay_date AT TIME ZONE 'Asia/Tokyo';