package gifts

import (
	"cic/site/pkg/db"
	"fmt"
)

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
