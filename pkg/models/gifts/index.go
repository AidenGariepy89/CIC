package gifts

import (
	"cic/site/pkg/db"
	"cic/site/pkg/models/user"
	"cmp"
	"fmt"
	"slices"
)

type Gift struct {
	Id          int
	Name        string
	Description string
}

type Question struct {
	Id      int
	Content string
	GiftId  int
}

type Answer struct {
	Id         int
	UserId     int
	QuestionId int `form:"questionId"`
	Answer     int `form:"answer"` // Range: [0, 3]
}

type Results struct {
	First        Gift
	Second       Gift
	Third        Gift
	FirstPoints  int
	SecondPoints int
	ThirdPoints  int
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
		var giftId int

		err = rows.Scan(&id, &content, &giftId)
		if err != nil {
			return nil, fmt.Errorf("Error scanning question: %w\n", err)
		}

		questions = append(questions, Question{id, content, giftId})
	}

	return &questions, nil
}

func GetQuestion(id int) (*Question, error) {
	row := db.Db.QueryRow(fmt.Sprintf("SELECT * FROM question WHERE id = %v", id))

	var content string
	var giftId int

	err := row.Scan(&id, &content, &giftId)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving row: %w", err)
	}

	return &Question{id, content, giftId}, nil
}

func GetQuestionsByGift(gift *Gift) (*[]Question, error) {
	rows, err := db.Db.Query(fmt.Sprintf("SELECT * FROM question WHERE giftId = %v", gift.Id))
	if err != nil {
		return nil, fmt.Errorf("Error fetching questions from db: %w\n", err)
	}

	var questions []Question
	for rows.Next() {
		var id int
		var content string
		var giftId int

		err = rows.Scan(&id, &content, &gift)
		if err != nil {
			return nil, fmt.Errorf("Error scanning question: %w\n", err)
		}

		questions = append(questions, Question{id, content, giftId})
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

		err = rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, fmt.Errorf("Error scanning question: %w\n", err)
		}

		gifts = append(gifts, Gift{id, name, description})
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
	rows, err := db.Db.Query(fmt.Sprintf(
		"select id from answer where userId = %v and questionId = %v",
		userId,
		questionId,
	))
	if err != nil {
		return fmt.Errorf("Error retrieving existing answers from db: %w\n", err)
	}

	count := 0
	for rows.Next() {
		count += 1
	}

	if count > 0 {
		_, err = db.Db.Exec(fmt.Sprintf(
			"update answer set answer = %v where userId = %v and questionId = %v",
			answer,
			userId,
			questionId,
		))

		return nil
	}

	_, err = db.Db.Exec(fmt.Sprintf(
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

func ProcessSpiritualGiftsResults(userId int) (*Results, error) {
	rows, err := db.Db.Query(
		"select answer, giftId from answer inner join question on answer.questionId = question.id where userId = ? order by answer desc",
		userId,
	)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving results from db: %w", err)
	}

	results := []struct {
		answer int
		giftId int
	}{}
	for rows.Next() {
		var answer int
		var giftId int

		err = rows.Scan(&answer, &giftId)
		if err != nil {
			return nil, fmt.Errorf("Error scanning data from db: %w", err)
		}

		results = append(results, struct {
			answer int
			giftId int
		}{answer, giftId})
	}

	buckets := make(map[int]int)

	for _, result := range results {
		buckets[result.giftId] += result.answer
	}

	buckets_slice := []struct {
		key int
		val int
	}{}
	for key, val := range buckets {
		buckets_slice = append(buckets_slice, struct {
			key int
			val int
		}{key, val})
	}

	slices.SortFunc(buckets_slice, func(a, b struct {
		key int
		val int
	}) int {
		return cmp.Compare(b.val, a.val)
	})

	firstGiftId := buckets_slice[0].key
	secondGiftId := buckets_slice[1].key
	thirdGiftId := buckets_slice[2].key
	firstLargestBucket := buckets_slice[0].val
	secondLargestBucket := buckets_slice[1].val
	thirdLargestBucket := buckets_slice[2].val

	// First Gift
	row := db.Db.QueryRow("select * from gift where id = ?", firstGiftId)

	var id int
	var name string
	var description string

	err = row.Scan(&id, &name, &description)
	if err != nil {
		return nil, fmt.Errorf("Error scanning gift from db: %w", err)
	}

	firstGift := Gift{Id: id, Name: name, Description: description}

	// Second Gift
	row = db.Db.QueryRow("select * from gift where id = ?", secondGiftId)

	err = row.Scan(&id, &name, &description)
	if err != nil {
		return nil, fmt.Errorf("Error scanning gift from db: %w", err)
	}

	secondGift := Gift{Id: id, Name: name, Description: description}

	// Third Gift
	row = db.Db.QueryRow("select * from gift where id = ?", thirdGiftId)

	err = row.Scan(&id, &name, &description)
	if err != nil {
		return nil, fmt.Errorf("Error scanning gift from db: %w", err)
	}

	thirdGift := Gift{Id: id, Name: name, Description: description}

	// return &Gift{Id: id, Name: name, Description: description}, nil
	return &Results{
		First:        firstGift,
		Second:       secondGift,
		Third:        thirdGift,
		FirstPoints:  firstLargestBucket,
		SecondPoints: secondLargestBucket,
		ThirdPoints:  thirdLargestBucket,
	}, nil
}
