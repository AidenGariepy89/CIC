package db

import (
	"database/sql"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var Db *sql.DB

func InitDb(url string) error {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return err
	}

	err = seedData(db)
	if err != nil {
		return err
	}

	Db = db
	return nil
}

func seedData(db *sql.DB) error {
	// Init Questions Table
	result := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='question'")

	var name string
	err := result.Scan(&name)
	if err != nil && name == "" {
		err = setupQuestions(db)
		if err != nil {
			return err
		}
	}

	// Init Gift Table
	result = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='gift'")

	err = result.Scan(&name)
	if err != nil && name == "" {
		err = setupGifts(db)
		if err != nil {
			return err
		}
	}

	// Init Gift Table
	result = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='user'")

	err = result.Scan(&name)
	if err != nil && name == "" {
		err = setupUsers(db)
		if err != nil {
			return err
		}
	}

	// Init Answer Table
	result = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='answer'")

	err = result.Scan(&name)
	if err != nil && name == "" {
		err = setupAnswers(db)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupAnswers(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS answer (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        userId INTEGER NOT NULL,
        questionId INTEGER NOT NULL,
        answer INTEGER NOT NULL,
        FOREIGN KEY(userId) REFERENCES user(id),
        FOREIGN KEY(questionId) REFERENCES question(id)
        )`)
	if err != nil {
		return err
	}

	return nil
}

func setupUsers(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL
        )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`BEGIN TRANSACTION;
        INSERT INTO user (username, password) VALUES ("admin", "admin");
        INSERT INTO user (username, password) VALUES ("demo", "demo");
        COMMIT;
        `)

	return nil
}

func setupGifts(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS gift (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL
        )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`BEGIN TRANSACTION;
        INSERT INTO gift (id, name, description) VALUES (1, "Administration", "");
        INSERT INTO gift (id, name, description) VALUES (2, "Apostleship", "");
        INSERT INTO gift (id, name, description) VALUES (3, "Crafting/Craftsmanship", "");
        INSERT INTO gift (id, name, description) VALUES (4, "Creative Communication", "");
        INSERT INTO gift (id, name, description) VALUES (5, "Discernment", "");
        INSERT INTO gift (id, name, description) VALUES (6, "Encouragement", "");
        INSERT INTO gift (id, name, description) VALUES (7, "Evangelism", "");
        INSERT INTO gift (id, name, description) VALUES (8, "Faith", "");
        INSERT INTO gift (id, name, description) VALUES (9, "Giving", "");
        INSERT INTO gift (id, name, description) VALUES (10, "Helps", "");
        INSERT INTO gift (id, name, description) VALUES (11, "Hospitality", "");
        INSERT INTO gift (id, name, description) VALUES (12, "Intercession", "");
        INSERT INTO gift (id, name, description) VALUES (13, "Knowledge", "");
        INSERT INTO gift (id, name, description) VALUES (14, "Leadership", "");
        INSERT INTO gift (id, name, description) VALUES (15, "Mercy", "");
        INSERT INTO gift (id, name, description) VALUES (16, "Prophecy", "");
        INSERT INTO gift (id, name, description) VALUES (17, "Shepherding", "");
        INSERT INTO gift (id, name, description) VALUES (18, "Teaching", "");
        INSERT INTO gift (id, name, description) VALUES (19, "Wisdom", "");
        COMMIT;
        `)
	if err != nil {
		return err
	}

	return nil
}

func setupQuestions(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS question (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        giftId INTEGER NOT NULL,
        FOREIGN KEY(giftId) REFERENCES gift(id)
        )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`BEGIN TRANSACTION;
        INSERT INTO question (content, giftId) VALUES ("I like to organize people, tasks, and events.", 1);
        INSERT INTO question (content, giftId) VALUES ("I would like to start churches in places where they do not presently exist.", 2);
        INSERT INTO question (content, giftId) VALUES ("I enjoy working creatively with wood, cloth, paints, metal, glass, or other materials.", 3);
        INSERT INTO question (content, giftId) VALUES ("I enjoy challenging people’s perspective of God by using various forms of art.", 4);
        INSERT INTO question (content, giftId) VALUES ("I can readily distinguish between spiritual truth and error, good and evil.", 5);
        INSERT INTO question (content, giftId) VALUES ("I tend to see the potential in people.", 6);
        INSERT INTO question (content, giftId) VALUES ("I communicate the gospel to others with clarity and effectiveness.", 7);
        INSERT INTO question (content, giftId) VALUES ("I find it natural and easy to trust God to answer my prayers.", 8);
        INSERT INTO question (content, giftId) VALUES ("I give liberally and joyfully to people in financial need or to projects requiring support.", 9);
        INSERT INTO question (content, giftId) VALUES ("I enjoy working behind the scenes to support the work of others.", 10);
        INSERT INTO question (content, giftId) VALUES ("I view my home as a place to minister to people in need.", 11);
        INSERT INTO question (content, giftId) VALUES ("I take prayer requests from others and consistently pray for them.", 12);
        INSERT INTO question (content, giftId) VALUES ("I am approached by people who want to know my perspective on a particular passage or Biblical truth.", 13);
        INSERT INTO question (content, giftId) VALUES ("I am able to motivate others to accomplish a goal.", 14);
        INSERT INTO question (content, giftId) VALUES ("I empathize with hurting people and desire to help in their healing process.", 15);
        INSERT INTO question (content, giftId) VALUES ("I can speak in a way that results in conviction and change in the lives of others.", 16);
        INSERT INTO question (content, giftId) VALUES ("I enjoy spending time nurturing and caring for others.", 17);
        INSERT INTO question (content, giftId) VALUES ("I am able to communicate God’s Word effectively to adults, youth, or children.", 18);
        INSERT INTO question (content, giftId) VALUES ("I am often sought out by others for advice about spiritual or personal matters.", 19);

        INSERT INTO question (content, giftId) VALUES ("I am careful, thorough, and skilled at managing details.", 1);
        INSERT INTO question (content, giftId) VALUES ("I am attracted to the idea of serving in another country or ethnic community.", 2);
        INSERT INTO question (content, giftId) VALUES ("I am skilled in working with different kinds of tools.", 3);
        INSERT INTO question (content, giftId) VALUES ("I enjoy developing and using my artistic skills (art, drama, music, photography, etc.).", 4);
        INSERT INTO question (content, giftId) VALUES ("I frequently am able to judge a person’s character based upon first impressions.", 5);
        INSERT INTO question (content, giftId) VALUES ("I enjoy reassuring and strengthening those who are discouraged.", 6);
        INSERT INTO question (content, giftId) VALUES ("I consistently look for opportunities to build relationships with non-Christians.", 7);
        INSERT INTO question (content, giftId) VALUES ("I have confidence in God’s continuing provision and help, even in difficult times.", 8);
        INSERT INTO question (content, giftId) VALUES ("I give more than a tithe so that God’s work can be accomplished.", 9);
        INSERT INTO question (content, giftId) VALUES ("I enjoy doing routine tasks that support the ministry.", 10);
        INSERT INTO question (content, giftId) VALUES ("I enjoy meeting new people and helping them to feel welcomed.", 11);
        INSERT INTO question (content, giftId) VALUES ("I enjoy praying for long periods of time and receive leadings as to what God wants me to pray for.", 12);
        INSERT INTO question (content, giftId) VALUES ("With the help of appropriate study materials, I can find what God’s Word teaches on most topics.", 13);
        INSERT INTO question (content, giftId) VALUES ("I am able to influence others to achieve a vision.", 14);
        INSERT INTO question (content, giftId) VALUES ("I can patiently support those going through painful experiences as they try to stabilize their lives.", 15);
        INSERT INTO question (content, giftId) VALUES ("People in trouble are encouraged when I talk with them.", 16);
        INSERT INTO question (content, giftId) VALUES ("I have compassion for wandering believers and want to protect them.", 17);
        INSERT INTO question (content, giftId) VALUES ("I can spend time in study knowing that presenting truth will make a difference in the lives of people–young or old.", 18);
        INSERT INTO question (content, giftId) VALUES ("I can often find simple, practical solutions in the midst of conflict or confusion.", 19);

        INSERT INTO question (content, giftId) VALUES ("I can clarify goals and develop strategies or plans to accomplish them.", 1);
        INSERT INTO question (content, giftId) VALUES ("I am willing to take an active part in starting a new church.", 2);
        INSERT INTO question (content, giftId) VALUES ("I enjoy making things for use in ministry.", 3);
        INSERT INTO question (content, giftId) VALUES ("I help people understand themselves, their relationships, and God better through artistic expression.", 4);
        INSERT INTO question (content, giftId) VALUES ("I can see through phoniness or deceit before it is evident to others.", 5);
        INSERT INTO question (content, giftId) VALUES ("I give hope to others by directing them to the promises of God.", 6);
        INSERT INTO question (content, giftId) VALUES ("I am effective at adapting the gospel message so that it connects with an individual’s felt need.", 7);
        INSERT INTO question (content, giftId) VALUES ("I believe that God will help me to accomplish great things.", 8);
        INSERT INTO question (content, giftId) VALUES ("I manage my money well in order to free more of it for giving.", 9);
        INSERT INTO question (content, giftId) VALUES ("I willingly take on a variety of odd jobs around the church to meet the needs of others.", 10);
        INSERT INTO question (content, giftId) VALUES ("I genuinely believe the Lord directs strangers to me who need to get connected to others.", 11);
        INSERT INTO question (content, giftId) VALUES ("I am conscious of ministering to others as I pray.", 12);
        INSERT INTO question (content, giftId) VALUES ("I am committed, and schedule blocks of time for reading and studying Scripture, to understand Biblical truth fully and accurately.", 13);
        INSERT INTO question (content, giftId) VALUES ("I can adjust my leadership style to bring out the best in others.", 14);
        INSERT INTO question (content, giftId) VALUES ("I enjoy helping people sometimes regarded as undeserving or beyond help.", 15);
        INSERT INTO question (content, giftId) VALUES ("I expose cultural trends, teachings, or events which contradict Biblical principles.", 16);
        INSERT INTO question (content, giftId) VALUES ("I like to provide guidance for the whole person–relationally, emotionally, spiritually, etc.", 17);
        INSERT INTO question (content, giftId) VALUES ("I pay close attention to the words, phrases, and meaning of those who teach.", 18);
        INSERT INTO question (content, giftId) VALUES ("I can easily select the most effective course of action from among several alternatives.", 19);

        INSERT INTO question (content, giftId) VALUES ("I can identify and effectively use the resources needed to accomplish tasks.", 1);
        INSERT INTO question (content, giftId) VALUES ("I can adapt well to different cultures and surroundings.", 2);
        INSERT INTO question (content, giftId) VALUES ("I can visualize how something should be constructed before I build it.  ", 3);
        INSERT INTO question (content, giftId) VALUES ("I like finding new and fresh ways of communicating God’s truth.", 4);
        INSERT INTO question (content, giftId) VALUES ("I tend to see rightness or wrongness in situations.", 5);
        INSERT INTO question (content, giftId) VALUES ("I reassure those who need to take courageous action in their faith, family, or life.", 6);
        INSERT INTO question (content, giftId) VALUES ("I invite unbelievers to accept Christ as their Savior.", 7);
        INSERT INTO question (content, giftId) VALUES ("I trust God in circumstances where success cannot be guaranteed by human effort alone.", 8);
        INSERT INTO question (content, giftId) VALUES ("I am challenged to limit my lifestyle in order to give away a higher percentage of my income.", 9);
        INSERT INTO question (content, giftId) VALUES ("I see spiritual significance in doing practical tasks.", 10);
        INSERT INTO question (content, giftId) VALUES ("I like to create a place where people do not feel that they are alone.", 11);
        INSERT INTO question (content, giftId) VALUES ("I pray with confidence because I know that God works in response to prayer.", 12);
        INSERT INTO question (content, giftId) VALUES ("I am perfectly at ease answering people’s Bible questions.", 13);
        INSERT INTO question (content, giftId) VALUES ("I set goals and manage people and resources effectively to accomplish them.", 14);
        INSERT INTO question (content, giftId) VALUES ("I have great compassion for hurting people.", 15);
        INSERT INTO question (content, giftId) VALUES ("People often tell me, “God used you.  You dealt exactly with my need.”", 16);
        INSERT INTO question (content, giftId) VALUES ("I can faithfully provide long-term support and concern for others.", 17);
        INSERT INTO question (content, giftId) VALUES ("I like to take a systematic approach to my study of the Bible.", 18);
        INSERT INTO question (content, giftId) VALUES ("I can anticipate the likely consequences of an individual’s or a group’s action. ", 19);

        INSERT INTO question (content, giftId) VALUES ("I like to help organizations or groups become more efficient.", 1);
        INSERT INTO question (content, giftId) VALUES ("I can relate to others in culturally sensitive ways.", 2);
        INSERT INTO question (content, giftId) VALUES ("I honor God with my handcrafted gifts.", 3);
        INSERT INTO question (content, giftId) VALUES ("I apply various artistic expressions to communicate God’s truth.", 4);
        INSERT INTO question (content, giftId) VALUES ("I receive affirmation from others concerning the reliability of my insights or perceptions. ", 5);
        INSERT INTO question (content, giftId) VALUES ("I strengthen those who are wavering in their faith.", 6);
        INSERT INTO question (content, giftId) VALUES ("I openly tell people that I am a Christian and want them to ask me about my faith.", 7);
        INSERT INTO question (content, giftId) VALUES ("I am convinced of God’s daily presence and action in my life.", 8);
        INSERT INTO question (content, giftId) VALUES ("I like knowing that my financial support makes a real difference in the lives and ministries of God’s people.", 9);
        INSERT INTO question (content, giftId) VALUES ("I like to find small things that need to be done and often do them without being asked.", 10);
        INSERT INTO question (content, giftId) VALUES ("I enjoy entertaining people and opening my home to others.", 11);
        INSERT INTO question (content, giftId) VALUES ("When I hear about needy situations, I feel burdened to pray.", 12);
        INSERT INTO question (content, giftId) VALUES ("Salvation by faith alone is a truth I clearly understand.", 13);
        INSERT INTO question (content, giftId) VALUES ("I influence others to perform to the best of their capability.", 14);
        INSERT INTO question (content, giftId) VALUES ("I can look beyond a person’s handicaps or problems to see a life that matters to God.", 15);
        INSERT INTO question (content, giftId) VALUES ("I appreciate people who are honest and will speak the truth.", 16);
        INSERT INTO question (content, giftId) VALUES ("I enjoy giving guidance and practical support to a small group of people.", 17);
        INSERT INTO question (content, giftId) VALUES ("I can communicate Scripture in ways that motivate others to study and want to learn more.", 18);
        INSERT INTO question (content, giftId) VALUES ("I give practical advice to help others through complicated situations.", 19);

        INSERT INTO question (content, giftId) VALUES ("I enjoy learning about how organizations function.", 1);
        INSERT INTO question (content, giftId) VALUES ("I enjoy pioneering new undertakings.", 2);
        INSERT INTO question (content, giftId) VALUES ("I am good at and enjoy working with my hands. ", 3);
        INSERT INTO question (content, giftId) VALUES ("I am creative and imaginative.", 4);
        INSERT INTO question (content, giftId) VALUES ("I can identify preaching, teaching, or communication which is not true to the Bible. ", 5);
        INSERT INTO question (content, giftId) VALUES ("I like motivating others to take steps for spiritual growth.", 6);
        INSERT INTO question (content, giftId) VALUES ("I openly and confidently tell others what Christ has done for me.", 7);
        INSERT INTO question (content, giftId) VALUES ("I am regularly challenging others to trust God.", 8);
        INSERT INTO question (content, giftId) VALUES ("I give generously due to my commitment to stewardship.", 9);
        INSERT INTO question (content, giftId) VALUES ("I feel comfortable being a helper, assisting others to do their job more effectively.  ", 10);
        INSERT INTO question (content, giftId) VALUES ("I do whatever I can to make people feel that they belong.", 11);
        INSERT INTO question (content, giftId) VALUES ("I am honored when someone asks me to pray for them.", 12);
        INSERT INTO question (content, giftId) VALUES ("I discover important Biblical truths when reading or studying Scripture which benefit others in the body of Christ.", 13);
        INSERT INTO question (content, giftId) VALUES ("I am able to cast a vision that others want to be a part of.", 14);
        INSERT INTO question (content, giftId) VALUES ("I enjoy bringing hope and joy to people living in difficult circumstances.", 15);
        INSERT INTO question (content, giftId) VALUES ("I will speak God’s truth, even in places there it is unpopular or difficult for others to accept.", 16);
        INSERT INTO question (content, giftId) VALUES ("I can gently restore wandering believers to faith and fellowship.", 17);
        INSERT INTO question (content, giftId) VALUES ("I can present information and skills to others at a level that makes it easy for them to grasp and apply to their lives.", 18);
        INSERT INTO question (content, giftId) VALUES ("I can apply Scriptural truth that others regard as practical and helpful.", 19);

        INSERT INTO question (content, giftId) VALUES ("I can visualize a coming event, anticipate potential problems, and develop backup plans.", 1);
        INSERT INTO question (content, giftId) VALUES ("I am able to orchestrate or oversee several church ministries.", 2);
        INSERT INTO question (content, giftId) VALUES ("I am able to design and construct things that help the church.", 3);
        INSERT INTO question (content, giftId) VALUES ("I regularly need to get alone to reflect and develop my imagination.", 4);
        INSERT INTO question (content, giftId) VALUES ("I can tell whether a person is being influenced by the Lord or Satan.", 5);
        INSERT INTO question (content, giftId) VALUES ("I am often asked to help those in trouble resolve their problems.", 6);
        INSERT INTO question (content, giftId) VALUES ("I seek opportunities to talk about spiritual matters with unbelievers.", 7);
        INSERT INTO question (content, giftId) VALUES ("I can move forward in spite of opposition or lack of support when I sense God’s blessing on an undertaking.", 8);
        INSERT INTO question (content, giftId) VALUES ("I believe I have been given an abundance of resources so that I may give more to the Lord’s work.", 9);
        INSERT INTO question (content, giftId) VALUES ("I readily and happily use my natural or learned skills to help wherever needed.", 10);
        INSERT INTO question (content, giftId) VALUES ("I can make people feel at ease even in unfamiliar surroundings.", 11);
        INSERT INTO question (content, giftId) VALUES ("I often see specific results in direct response to my prayers.", 12);
        INSERT INTO question (content, giftId) VALUES ("I confidently share my knowledge and insights with others.", 13);
        INSERT INTO question (content, giftId) VALUES ("I figure out where we need to go and help others to get there.", 14);
        INSERT INTO question (content, giftId) VALUES ("I enjoy doing practical things for others who are in need.", 15);
        INSERT INTO question (content, giftId) VALUES ("I feel compelled to expose sin wherever I see it and to challenge people to repentance.", 16);
        INSERT INTO question (content, giftId) VALUES ("I enjoy patiently but firmly nurturing others in their development as believers.", 17);
        INSERT INTO question (content, giftId) VALUES ("I enjoy explaining things to people so that they can grow spiritually and personally.", 18);
        INSERT INTO question (content, giftId) VALUES ("I have insights into how to solve problems that others often do not see.", 19);

        COMMIT;
        `)

	if err != nil {
		return err
	}

	return nil
}
