CREATE TABLE users (
  id uuid PRIMARY KEY,
  name varchar(20) NOT NULL,
  picture varchar(100) NOT NULL,
  description text,
  mail varchar(40) NOT NULL,
  createdAt TIMESTAMP DEFAULT now()
);

CREATE TABLE user_followers (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    follower_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, follower_id)
);

CREATE INDEX idx_user_followers_follower_id ON user_followers(follower_id);

CREATE TABLE posts (
  id uuid PRIMARY KEY,
  comment text NOT NULL,
  postedBy uuid,
  createdAt TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_user FOREIGN KEY(postedBy) REFERENCES users(id)
);

CREATE TABLE comments (
  id uuid PRIMARY KEY,
  content text NOT NULL,
  postId uuid,
  commentedBy uuid,
  createdAt TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_post FOREIGN KEY(postId) REFERENCES posts(id),
  CONSTRAINT fk_user FOREIGN KEY(commentedBy) REFERENCES users(id)
);

CREATE TABLE booksToRead (
  id uuid PRIMARY KEY,
  isbn varchar(13) NOT NULL,
  title varchar(70) NOT NULL,
  description text,
  picture varchar(100),
  userId uuid,
  createdAt TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_user FOREIGN KEY(userId) REFERENCES users(id)
);

CREATE TABLE booksRecommended (
  id uuid PRIMARY KEY,
  isbn varchar(13) NOT NULL,
  title varchar(70) NOT NULL,
  description text,
  picture varchar(100),
  userId uuid,
  createdAt TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_user FOREIGN KEY(userId) REFERENCES users(id)
);
