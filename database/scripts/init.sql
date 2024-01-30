CREATE TABLE questions (
  ID VARCHAR(255) NOT NULL UNIQUE,
  Content_Text VARCHAR(4095) NOT NULL,
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
CREATE TABLE machines (
    id SERIAL PRIMARY KEY,
    aws_instance_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    domain VARCHAR(255) NOT NULL
);
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    session_id UUID NOT NULL,
    machine_id INTEGER NOT NULL,
    FOREIGN KEY (machine_id) REFERENCES machines(id)
);
CREATE INDEX idx_session_id ON sessions(session_id);