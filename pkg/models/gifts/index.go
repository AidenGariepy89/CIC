package gifts

import (
	"cic/site/pkg/db"
	"cic/site/pkg/models/user"
	"fmt"
)

type Gift struct {
	Id          int
	Name        string
	Description string
	Key         rune
}

type Question struct {
	Id      int
	Content string
	Gift    rune
}

type Answer struct {
	Id         int
	UserId     int
	QuestionId int
	Answer     int // Range: [0, 3]
}

func GetQuestions() (*[]Question, error) {
	rows, err := db.Db.Query("SELECT * FROM question")
	if err != nil {
		return nil, fmt.Errorf("Error fetching questions from db: %w\n", err)
	}

	var questions []Question
	for rows.Next() {
		var id int
		var content string
		var gift rune

		err = rows.Scan(&id, &content, &gift)
		if err != nil {
			return nil, fmt.Errorf("Error scanning question: %w\n", err)
		}

		questions = append(questions, Question{id, content, gift})
	}

	return &questions, nil
}

func GetQuestion(id int) (*Question, error) {
	row := db.Db.QueryRow(fmt.Sprintf("SELECT * FROM question WHERE id = %v", id))

	var content string
	var gift rune

	err := row.Scan(&id, &content, &gift)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving row: %w", err)
	}

	return &Question{id, content, gift}, nil
}

func GetQuestionsByGift(gift rune) (*[]Question, error) {
	rows, err := db.Db.Query(fmt.Sprintf("SELECT * FROM question WHERE gift = %v", gift))
	if err != nil {
		return nil, fmt.Errorf("Error fetching questions from db: %w\n", err)
	}

	var questions []Question
	for rows.Next() {
		var id int
		var content string
		var gift rune

		err = rows.Scan(&id, &content, &gift)
		if err != nil {
			return nil, fmt.Errorf("Error scanning question: %w\n", err)
		}

		questions = append(questions, Question{id, content, gift})
	}

	return &questions, nil
}

func GetGifts() (*[]Gift, error) {
	rows, err := db.Db.Query("SELECT * FROM gift;")
	if err != nil {
		return nil, fmt.Errorf("Error fetching questions from db: %w\n", err)
	}

	var gifts []Gift
	for rows.Next() {
		var id int
		var name string
		var description string
		var key rune

		err = rows.Scan(&id, &name, &description, &key)
		if err != nil {
			return nil, fmt.Errorf("Error scanning question: %w\n", err)
		}

		gifts = append(gifts, Gift{id, name, description, key})
	}

	return &gifts, nil
}

func GetUserAnswers(userId int) (*[]Answer, error) {
	u, err := user.GetUser(userId)
	if err != nil {
		return nil, err
	}

	rows, err := db.Db.Query(fmt.Sprintf("SELECT * FROM answer WHERE user_id = %v", u.Id))
	if err != nil {
		return nil, fmt.Errorf("Error retrieving answers from db: %w\n", err)
	}

	answers := []Answer{}
	for rows.Next() {
		var id int
		var userId int
		var questionId int
		var answer int

		err = rows.Scan(&id, &userId, &questionId, &answer)
		if err != nil {
			return nil, fmt.Errorf("Error scanning answer: %w\n", err)
		}

		answers = append(answers, Answer{id, userId, questionId, answer})
	}

	return &answers, nil
}

func SubmitAnswer(answer int, userId int, questionId int) error {
	_, err := db.Db.Exec(fmt.Sprintf(
		"insert into answer (answer, userId, questionId) values (%v, %v, %v)",
        answer,
        userId,
        questionId,
	))
	if err != nil {
		return fmt.Errorf("Error inserting answer: %w\n", err)
	}

	return nil
}
