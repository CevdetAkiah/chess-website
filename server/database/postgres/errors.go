package postgres

const UNIQUE_KEY_VIOLATION = "23505"

// errors to return
const USERNAME_DUPLICATE = "username exists"
const EMAIL_DUPLICATE = "email exists"

// pq constraints
const CONSTRAINT_USERNAME = "unique_name"
const CONSTRAINT_EMAIL = "users_email_key"
