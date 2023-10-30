package gifts

import (
	"cic/site/pkg/db"
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
