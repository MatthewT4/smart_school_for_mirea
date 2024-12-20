package storage

const (
	queryCreateUser = `
INSERT INTO users (id, email, password)
VALUES ($1, $2, $3)
RETURNING id, email, password
`

	queryGetUserByEmail = `
SELECT id, email, password
FROM users
WHERE email = $1
`

	queryGetUserByID = `
SELECT id, email, password
FROM users
WHERE id = $1
`

	queryGetTopic = `
SELECT * FROM topic WHERE id = $1
`
	queryGetCourseTopics = `
SELECT * FROM topic WHERE course_id = $1
`

	execAddTopicViewedRow = `
INSERT INTO viewed_topic (user_id, topic_id)
VALUES ($1, $2)
`

	queryGetCourse = `
SELECT id, title, description
FROM course
WHERE id = $1
`

	queryGetCourseElements = `
SELECT element_type, element_id, index
FROM course_element
WHERE course_id = $1
`
	queryUserCourses = `
SELECT course_id 
FROM user_invited_in_course
WHERE user_id = $1
`
	queryFindCourses = `
SELECT * 
FROM course 
WHERE
	($1::uuid[] IS NULL OR id = ANY($1))
AND
	($2::text IS NULL OR title ILIKE '%' || $2 || '%')
`
	execAddUserInCourse = `
INSERT INTO user_invited_in_course(user_id, course_id)
VALUES ($1, $2)
`
)
