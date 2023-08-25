
DROP TABLE IF EXISTS user_option;

DROP TABLE IF EXISTS checkin;

DROP TABLE IF EXISTS coupon_notices;

DROP TABLE IF EXISTS coupon_stores;

DROP TABLE IF EXISTS coupon_attached_user;

DROP TABLE IF EXISTS coupon;

DROP TABLE IF EXISTS banner;

DROP TABLE IF EXISTS post;

DROP TABLE IF EXISTS mail_magazine;

DROP TABLE IF EXISTS admin;

DROP TABLE IF EXISTS stayable_store_info;

DROP TABLE IF EXISTS store;

DROP TABLE IF EXISTS "user";

DROP trigger if exists on_auth_user_created on auth.users;

DROP function if exists public.handle_new_user;

DROP trigger if exists on_auth_user_updated on auth.users;

DROP function if exists public.handle_user_email_update;

-- DROP TABLE IF EXISTS schema_migrations;
