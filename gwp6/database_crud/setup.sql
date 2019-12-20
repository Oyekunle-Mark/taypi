-- drop the tables first should they exist
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments;

-- Create posts table
CREATE TABLE posts
(
    id SERIAL PRIMARY KEY,
    content TEXT,
    author VARCHAR(255)
);

-- Create comments table
CREATE TABLE comments
(
    id SERIAL PRIMARY KEY,
    content TEXT,
    author VARCHAR(255),
    post_id INTEGER REFERENCES posts(id)
);
