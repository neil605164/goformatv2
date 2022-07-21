package structer

import "time"

type UserList struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	MemberNo  string    `json:"member_no"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	Review    *Review   `json:"review"`
}

type Review struct {
	ID             uint64     `json:"id"`
	VerifyType     string     `json:"verify_type"`
	Status         string     `json:"status"`
	SubmissionedAt *time.Time `json:"submissionedAt"`
	ReviewedAt     *time.Time `json:"reviewedAt"`
	CreatedAt      time.Time  `json:"created_at" `
}
