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


CREATE TABLE test
(
    id UUID PRIMARY KEY,
    course_id UUID NOT NULL REFERENCES course(id),

    title TEXT NOT NULL
);

CREATE TABLE test_element
(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    test_id UUID NOT NULL REFERENCES test(id),

    index INT NOT NULL,
    correct_answer TEXT NOT NULL
);

CREATE TABLE test_result
(
    id UUID PRIMARY KEY,
    test_id UUID NOT NULL REFERENCES test(id),
    user_id UUID NOT NULL REFERENCES users(id),

    count_correct_answers INTEGER NOT NULL,
    count_answers INTEGER NOT NULL
);

CREATE TABLE test_element_result
(
    id UUID PRIMARY KEY,
    test_result_id UUID NOT NULL REFERENCES test_result(id),
    element_id UUID NOT NULL REFERENCES test_element(id),

    user_answer TEXT NOT NULL,
    score INTEGER NOT NULL
);
