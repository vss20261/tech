package validator

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	ID          uint
	Name        string
	Email       string
	PhoneNumber string
	Birthday    time.Time
}

type Member struct {
	ID          uint
	Name        string
	PhoneNumber string
	Birthday    time.Time
}

func TestCopyFields(t *testing.T) {
	user := User{
		ID:          1,
		Name:        "Kim",
		Email:       "example@example.com",
		PhoneNumber: "010-0000-0000",
		Birthday:    time.Now(),
	}
	memberA := Member{}
	memberB := Member{}

	err := CopyFields(user, &memberA, "Name", "Birthday", "PhoneNumber")
	if err != nil {
		t.Errorf("%s", err)
	}

	err = CopyAll(user, &memberB)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("memberA = %+v\n", memberA)
	fmt.Printf("memberB = %+v\n", memberB)
}
