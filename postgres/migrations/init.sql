CREATE TABLE users
(
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE course
(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE topic
(
    id UUID PRIMARY KEY,
    course_id UUID NOT NULL REFERENCES course(id),
    title TEXT NOT NULL,
    body TEXT NOT NULL
);

CREATE TABLE viewed_topic
(
    user_id UUID NOT NULL REFERENCES users(id),
    topic_id UUID NOT NULL REFERENCES topic(id)
);

CREATE TYPE element_type AS ENUM('topic', 'test');
CREATE TABLE course_element
(
    course_id UUID NOT NULL,
    element_type element_type NOT NULL,
    element_id UUID NOT NULL,
    index UUID NOT NULL
);

CREATE TABLE user_invited_in_course
(
    user_id UUID NOT NULL REFERENCES users(id),
    course_id UUID NOT NULL REFERENCES course(id)
);