-- UP
CREATE TABLE IF NOT EXISTS example_table (id SERIAL PRIMARY KEY, name VARCHAR(255));

-- DOWN
DROP TABLE IF EXISTS example_table;
