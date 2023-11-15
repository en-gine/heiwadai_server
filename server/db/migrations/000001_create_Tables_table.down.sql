
DROP TABLE IF EXISTS user_option;

DROP TABLE IF EXISTS checkin;

DROP TABLE IF EXISTS coupon_notices;

DROP TABLE IF EXISTS coupon_stores;

DROP TABLE IF EXISTS coupon_attached_user;

DROP TABLE IF EXISTS coupon;

DROP TABLE IF EXISTS banner;

DROP TABLE IF EXISTS post;

DROP TABLE IF EXISTS mail_magazine_log;

DROP TABLE IF EXISTS mail_magazine;

DROP TABLE IF EXISTS message;

DROP TABLE IF EXISTS user_book;

DROP TABLE IF EXISTS book_guest_data;

DROP TABLE IF EXISTS book_plan;

DROP TABLE IF EXISTS user_report;

DROP TABLE IF EXISTS admin_login_log;

DROP TABLE IF EXISTS admin;

DROP TABLE IF EXISTS stayable_store_info;

DROP TABLE IF EXISTS store;

DROP TABLE IF EXISTS user_data;

DROP TABLE IF EXISTS user_manager;

DROP trigger if exists on_auth_user_created on auth.users;

DROP function if exists public.handle_new_user;

DROP trigger if exists on_auth_user_updated on auth.users;

DROP function if exists public.handle_user_email_update;

DROP trigger if exists delete_public_user_trigger on auth.users;

DROP function if exists public.delete_public_user;

DROP FUNCTION IF EXISTS generate_booking_number();

DROP SEQUENCE IF EXISTS booking_number_sequence;

-- DROP TABLE IF EXISTS schema_migrations;
