-- +migrate Up
CREATE TYPE shopping_categories as ENUM('food', 'necessity');

CREATE TABLE shopping_items (
  id SERIAL PRIMARY KEY,
  owner_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  category shopping_categories NOT NULL,
  picked BOOLEAN NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE shopping_items;