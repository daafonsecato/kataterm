CREATE TABLE questions (
  ID VARCHAR(255) NOT NULL UNIQUE,
  Content_Text VARCHAR(2047) NOT NULL,
  Hint VARCHAR(255) NOT NULL,
  Subtext VARCHAR(255) NOT NULL,
  Type_Question VARCHAR(255) NOT NULL,
  Staging_Message VARCHAR(255) NOT NULL,
  Options JSON,
  Before_Actions JSON,
  Answer VARCHAR(255),
  Test_spec_filename VARCHAR(255),
  Trials VARCHAR(255) NOT NULL
);
