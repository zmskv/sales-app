CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "User" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    is_created BOOLEAN DEFAULT NULL
);


CREATE TABLE "Sales_List" (
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    amount INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);