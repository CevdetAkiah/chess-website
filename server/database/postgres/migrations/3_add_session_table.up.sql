CREATE TABLE IF NOT EXISTS chessapp.sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references chessapp.users(id),
  created_at timestamp not null   
);
