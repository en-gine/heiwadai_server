CREATE TABLE "user" (
    id UUID PRIMARY KEY,
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
    mail VARCHAR NOT NULL,
    accept_mail BOOLEAN NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);

CREATE TABLE user_option (
    id UUID PRIMARY KEY,
    inner_note VARCHAR NOT NULL,
    is_black_customer BOOLEAN NOT NULL,
    user_id UUID,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE TABLE store (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    branch_name VARCHAR,
    address VARCHAR NOT NULL,
	tel VARCHAR NOT NULL,
	parking VARCHAR NOT NULL,
	access_info VARCHAR NOT NULL,
    is_active BOOLEAN NOT NULL,
    stamp_image_url VARCHAR,
    stayable BOOLEAN NOT NULL,
    qr_code UUID NOT NULL,
    un_limited_qr_code UUID NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);

CREATE TABLE admin (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    belong_to UUID,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (belong_to) REFERENCES store (id)
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
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE TABLE coupon (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    coupon_type int NOT NULL,
    discount_amount INTEGER NOT NULL,
    expire_at TIMESTAMPTZ NOT NULL,
    is_combinationable BOOLEAN NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    coupon_status int NOT NULL
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
    FOREIGN KEY (user_id) REFERENCES "user" (id),
    PRIMARY KEY(coupon_id, user_id)
);

CREATE TABLE banner (
    id UUID PRIMARY KEY,
    image_url VARCHAR NOT NULL,
    url VARCHAR NOT NULL,
    status int NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now()
);

CREATE TABLE post (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    author UUID,
    post_date TIMESTAMPTZ NOT NULL,
    post_status int NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (author) REFERENCES admin (id)
);

CREATE TABLE mail_magazine (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    author UUID,
    sent_at TIMESTAMPTZ,
    sent_count int default 0,
    mail_magazine_status int NOT NULL,
    create_at TIMESTAMPTZ NOT NULL default now(),
    update_at TIMESTAMPTZ NOT NULL default now(),
    FOREIGN KEY (author) REFERENCES admin (id)
);
--- userが作成されるたびに、userテーブルにもidとemailをinsertする
create or replace function public.handle_new_user() 
returns trigger as $$
begin
  insert into public.user (id, email)
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
  update public.user
  set email = new.email
  where id = new.id;
  
  return new;
end;
$$ language plpgsql security definer;

create trigger on_auth_user_updated
  after update of email on auth.users
  for each row execute procedure public.handle_user_email_update();