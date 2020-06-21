DROP TABLE IF EXISTS urls;

CREATE TABLE urls (
  id SERIAL primary key,
  short_path varchar(255) UNIQUE,
  long_url varchar(255)
);

-- INSERT INTO urls(short_path, long_url) values('/1234', 'https://donbackup.com/admin/');
