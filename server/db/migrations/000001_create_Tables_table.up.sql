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
    birth_date TIMESTAMPTZ NOT NULL,
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
    store_id UUID,
    user_id UUID,
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
	coupon_id UUID REFERENCES coupon(id) PRIMARY KEY,
	notice TEXT NOT NULL
);

CREATE TABLE coupon_stores (
	coupon_id UUID REFERENCES coupon(id) PRIMARY KEY,
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
    sent_count int default NULL,
    mail_magazine_status int NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (author_id) REFERENCES admin (admin_id)
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