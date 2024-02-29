CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid () PRIMARY KEY,
    email TEXT UNIQUE,
    phone_number TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id UUID DEFAULT gen_random_uuid ()  PRIMARY KEY,
    name VARCHAR(200) NOT NULL CHECK (name != ''),
    price DECIMAL CHECK (price >= 0),
    created_at  TIMESTAMP without time zone default (now() at time zone 'utc'),
    updated_at   TIMESTAMP without time zone default (now() at time zone 'utc')
);

CREATE TABLE IF NOT EXISTS subscriptions (
    user_id UUID REFERENCES users(id),
    product_id UUID REFERENCES products(id),
    PRIMARY KEY (user_id, product_id)
);