package config

type Config struct {
	LoginStudentUsername	string	`toml:"login-student-username"`
	LoginStudentPassword	string	`toml:"login-student-password"`
	LoginTeacherUsername	string	`toml:"login-teacher-username"`
	LoginTeacherPassword	string	`toml:"login-teacher-password"`
}
