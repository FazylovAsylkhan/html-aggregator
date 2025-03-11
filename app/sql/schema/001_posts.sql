-- +goose Up
CREATE TABLE postsBaspana (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    date_publication TIMESTAMP NOT NULL,
    title TEXT NOT NULL, 
    image TEXT NOT NULL, 
    cost_for_metr INT CHECK (cost_for_metr >= 0) NOT NULL,
    address TEXT NOT NULL, 
    link_detail_post TEXT NOT NULL, 
    number_object  INT CHECK (number_object >= 0) NOT NULL,
    count_access INT CHECK (count_access >= 0) NOT NULL,
    region TEXT NOT NULL
);

-- +goose Down
DROP TABLE postsBaspana;