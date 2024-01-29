CREATE TABLE user_manager (
    id UUID PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    is_admin BOOLEAN NOT NULL DEFAULT false,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);


CREATE TABLE user_data (
    user_id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    first_name_kana VARCHAR NOT NULL,
    last_name_kana VARCHAR NOT NULL,
    company_name VARCHAR,
    birth_date TIMESTAMPTZ NULL,
    zip_code VARCHAR,
    prefecture int NOT NULL,
    city VARCHAR,
    address VARCHAR,
    tel VARCHAR,
    accept_mail BOOLEAN NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (user_id) REFERENCES user_manager (id) ON DELETE CASCADE
);

CREATE TABLE user_option (
    user_id UUID PRIMARY KEY,
    inner_note VARCHAR NOT NULL,
    is_black_customer BOOLEAN NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (user_id) REFERENCES user_data (user_id) ON DELETE CASCADE
);

CREATE TABLE store (
    id                UUID PRIMARY KEY,
    name              VARCHAR NOT NULL,
    branch_name       VARCHAR,
    zip_code          VARCHAR NOT NULL,
    address           VARCHAR NOT NULL,
    tel               VARCHAR NOT NULL,
    site_url          VARCHAR NOT NULL,
    stamp_image_url   VARCHAR NOT NULL,
    stayable          BOOLEAN NOT NULL,
    is_active         BOOLEAN NOT NULL,
    qr_code           UUID NOT NULL,
    un_limited_qr_code UUID NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);

CREATE TABLE stayable_store_info (
    store_id         UUID PRIMARY KEY,
    parking          VARCHAR NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    access_info      VARCHAR NOT NULL,
    rest_api_url     VARCHAR NOT NULL,
    booking_system_id VARCHAR NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (store_id) REFERENCES store (id)
);

CREATE TABLE admin (
    admin_id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    belong_to UUID NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (belong_to) REFERENCES store (id),
    FOREIGN KEY (admin_id) REFERENCES user_manager (id) ON DELETE CASCADE
);


CREATE TABLE checkin (
    id UUID PRIMARY KEY,
    store_id UUID NOT NULL,
    user_id UUID NOT NULL,
    check_in_at TIMESTAMPTZ NOT NULL,
    archive BOOLEAN NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (store_id) REFERENCES store (id),
    FOREIGN KEY (user_id) REFERENCES user_data (user_id)
);

CREATE TABLE coupon (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    coupon_type int NOT NULL,
    discount_amount INTEGER NOT NULL,
    expire_at TIMESTAMPTZ NOT NULL,
    is_combinationable BOOLEAN NOT NULL,
    coupon_status int NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);

CREATE TABLE coupon_notices (
	coupon_id UUID REFERENCES coupon(id) PRIMARY KEY DEFERRABLE INITIALLY DEFERRED, -- トランザクション外部キー評価遅延
	notice TEXT NOT NULL
);

CREATE TABLE coupon_stores (
	coupon_id UUID REFERENCES coupon(id) PRIMARY KEY DEFERRABLE INITIALLY DEFERRED, -- トランザクション外部キー評価遅延
	store_id UUID NOT NULL REFERENCES store(id)
);

CREATE TABLE coupon_attached_user (
	coupon_id UUID,
    used_at TIMESTAMPTZ,
    user_id UUID,
    FOREIGN KEY (coupon_id) REFERENCES coupon(id),
    FOREIGN KEY (user_id) REFERENCES user_data (user_id),
    PRIMARY KEY(coupon_id, user_id)
);


CREATE TABLE mail_magazine (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    author_id UUID NOT NULL,
    sent_at TIMESTAMPTZ,
    unsent_count int NOT NULL default 0,
    sent_count int NOT NULL default 0,
    target_prefectures INTEGER[] NULL default array[]::INTEGER[],
    mail_magazine_status int NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (author_id) REFERENCES admin (admin_id)
);

CREATE TABLE mail_magazine_log (
    id bigserial PRIMARY KEY,
    mail_magazine_id UUID NOT NULL,
    user_id UUID NOT NULL,
    email VARCHAR NOT NULL,
    sent_at TIMESTAMPTZ NULL, -- 送信日時　NULLの場合は未送信
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (mail_magazine_id) REFERENCES mail_magazine (id),
    FOREIGN KEY (user_id) REFERENCES user_data (user_id)
);


CREATE TABLE message (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    author_id UUID NOT NULL,
    display_date TIMESTAMPTZ NOT NULL default now(),
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (author_id) REFERENCES admin (admin_id)
);

CREATE TABLE user_report (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    user_id UUID NOT NULL,
    user_name VARCHAR NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (user_id) REFERENCES user_data (user_id)
);

CREATE TABLE book_guest_data (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    first_name_kana VARCHAR(255) NOT NULL,
    last_name_kana VARCHAR(255) NOT NULL,
    company_name VARCHAR(255),
    zip_code VARCHAR(7),
    prefecture INTEGER NULL,
    city VARCHAR(255),
    address TEXT,
    tel VARCHAR(15),
    mail VARCHAR(255) NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);

CREATE TABLE book_plan (
    id UUID PRIMARY KEY,
    plan_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    image_url VARCHAR(255) NOT NULL,
    room_type INTEGER NOT NULL,
    meal_type_morning BOOLEAN NOT NULL,
    meal_type_dinner BOOLEAN NOT NULL,
    smoke_type INTEGER NOT NULL,
    overview TEXT NOT NULL,
    store_id UUID NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (store_id) REFERENCES store(id)
);

CREATE TABLE user_book (
    id UUID PRIMARY KEY,
    tl_booking_number VARCHAR(255) NOT NULL,
    stay_from TIMESTAMP NOT NULL,
    stay_to TIMESTAMP NOT NULL,
    adult INTEGER NOT NULL CHECK (adult >= 0),
    child INTEGER NOT NULL CHECK (child >= 0),
    room_count INTEGER NOT NULL CHECK (room_count > 0),
    check_in_time VARCHAR NOT NULL,
    total_cost INTEGER NOT NULL CHECK (total_cost >= 0),
    guest_data_id UUID NOT NULL,
    book_plan_id UUID NOT NULL,
    book_user_id UUID NOT NULL,
    note TEXT,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (book_user_id) REFERENCES user_data(user_id),
    FOREIGN KEY (guest_data_id) REFERENCES book_guest_data(id) ON DELETE CASCADE,
    FOREIGN KEY (book_plan_id) REFERENCES book_plan(id) ON DELETE CASCADE
);
--- userが作成されるたびに、userテーブルにもidとemailをinsertする
create or replace function public.handle_new_user() 
returns trigger as $$
begin
  insert into public.user_manager (id, email)
  values (new.id, new.email);
  return new;
end;
$$ language plpgsql security definer;

create trigger on_auth_user_created
  after insert on auth.users
  for each row execute procedure public.handle_new_user();

--- userが更新されるたびに、userテーブルのemailを更新する
create or replace function public.handle_user_email_update() 
returns trigger as $$
begin
  update public.user_manager
  set email = new.email
  where id = new.id;
  
  return new;
end;
$$ language plpgsql security definer;

create trigger on_auth_user_updated
  after update of email on auth.users
  for each row execute procedure public.handle_user_email_update();

-- userが削除されるときに、userテーブルのデータも削除する
CREATE OR REPLACE FUNCTION delete_public_user()
RETURNS TRIGGER AS $$
BEGIN
  DELETE FROM public.user_manager WHERE id = OLD.id;
  RETURN OLD;
END;
$$ language plpgsql security definer;

CREATE TRIGGER delete_public_user_trigger
AFTER DELETE ON auth.users
FOR EACH ROW EXECUTE FUNCTION delete_public_user();

-- 予約用シーケンスを作成
CREATE SEQUENCE booking_number_sequence;

CREATE OR REPLACE FUNCTION generate_booking_number()
RETURNS TEXT AS $$
DECLARE
  current_date_str VARCHAR(8);
  sequence_number BIGINT;
BEGIN
  -- 現在の日付を'YYYYMMDD'形式の文字列として取得
  current_date_str := TO_CHAR(NOW(), 'YYYYMMDD');
  
  -- シーケンスを使用して連番を取得 (シーケンスは別途作成)
  sequence_number := NEXTVAL('booking_number_sequence');
  
  -- 日付の文字列と9桁の0埋めされた連番を結合
  RETURN current_date_str || LPAD(sequence_number::TEXT, 9, '0');
END;
$$ LANGUAGE plpgsql;


