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

	result := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='question'")

	var name string
	err = result.Scan(&name)
	if err != nil && name == "" {
		err = setupQuestions(db)
		if err != nil {
			return err
		}
	}

	// err = setupQuestions(db)
	// if err != nil {
	// 	return err
	// }

	Db = db
	return nil
}

func setupQuestions(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS question (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        gift INTEGER NOT NULL
    )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`BEGIN TRANSACTION;
        insert into question (content, gift) values ("I like to organize people, tasks, and events.", 65);
        insert into question (content, gift) values ("I would like to start churches in places where they do not presently exist.", 66);
        insert into question (content, gift) values ("I enjoy working creatively with wood, cloth, paints, metal, glass, or other materials.", 67);
        insert into question (content, gift) values ("I enjoy challenging people’s perspective of God by using various forms of art.", 68);
        insert into question (content, gift) values ("I can readily distinguish between spiritual truth and error, good and evil.", 69);
        insert into question (content, gift) values ("I tend to see the potential in people.", 70);
        insert into question (content, gift) values ("I communicate the gospel to others with clarity and effectiveness.", 71);
        insert into question (content, gift) values ("I find it natural and easy to trust God to answer my prayers.", 72);
        insert into question (content, gift) values ("I give liberally and joyfully to people in financial need or to projects requiring support.", 73);
        insert into question (content, gift) values ("I enjoy working behind the scenes to support the work of others.", 74);
        insert into question (content, gift) values ("I view my home as a place to minister to people in need.", 75);
        insert into question (content, gift) values ("I take prayer requests from others and consistently pray for them.", 76);
        insert into question (content, gift) values ("I am approached by people who want to know my perspective on a particular passage or Biblical truth.", 77);
        insert into question (content, gift) values ("I am able to motivate others to accomplish a goal.", 78);
        insert into question (content, gift) values ("I empathize with hurting people and desire to help in their healing process.", 79);
        insert into question (content, gift) values ("I can speak in a way that results in conviction and change in the lives of others.", 80);
        insert into question (content, gift) values ("I enjoy spending time nurturing and caring for others.", 81);
        insert into question (content, gift) values ("I am able to communicate God’s Word effectively to adults, youth, or children.", 82);
        insert into question (content, gift) values ("I am often sought out by others for advice about spiritual or personal matters.", 83);

        insert into question (content, gift) values ("I am careful, thorough, and skilled at managing details.", 65);
        insert into question (content, gift) values ("I am attracted to the idea of serving in another country or ethnic community.", 66);
        insert into question (content, gift) values ("I am skilled in working with different kinds of tools.", 67);
        insert into question (content, gift) values ("I enjoy developing and using my artistic skills (art, drama, music, photography, etc.).", 68);
        insert into question (content, gift) values ("I frequently am able to judge a person’s character based upon first impressions.", 69);
        insert into question (content, gift) values ("I enjoy reassuring and strengthening those who are discouraged.", 70);
        insert into question (content, gift) values ("I consistently look for opportunities to build relationships with non-Christians.", 71);
        insert into question (content, gift) values ("I have confidence in God’s continuing provision and help, even in difficult times.", 72);
        insert into question (content, gift) values ("I give more than a tithe so that God’s work can be accomplished.", 73);
        insert into question (content, gift) values ("I enjoy doing routine tasks that support the ministry.", 74);
        insert into question (content, gift) values ("I enjoy meeting new people and helping them to feel welcomed.", 75);
        insert into question (content, gift) values ("I enjoy praying for long periods of time and receive leadings as to what God wants me to pray for.", 76);
        insert into question (content, gift) values ("With the help of appropriate study materials, I can find what God’s Word teaches on most topics.", 77);
        insert into question (content, gift) values ("I am able to influence others to achieve a vision.", 78);
        insert into question (content, gift) values ("I can patiently support those going through painful experiences as they try to stabilize their lives.", 79);
        insert into question (content, gift) values ("People in trouble are encouraged when I talk with them.", 80);
        insert into question (content, gift) values ("I have compassion for wandering believers and want to protect them.", 81);
        insert into question (content, gift) values ("I can spend time in study knowing that presenting truth will make a difference in the lives of people–young or old.", 82);
        insert into question (content, gift) values ("I can often find simple, practical solutions in the midst of conflict or confusion.", 83);

        insert into question (content, gift) values ("I can clarify goals and develop strategies or plans to accomplish them.", 65);
        insert into question (content, gift) values ("I am willing to take an active part in starting a new church.", 66);
        insert into question (content, gift) values ("I enjoy making things for use in ministry.", 67);
        insert into question (content, gift) values ("I help people understand themselves, their relationships, and God better through artistic expression.", 68);
        insert into question (content, gift) values ("I can see through phoniness or deceit before it is evident to others.", 69);
        insert into question (content, gift) values ("I give hope to others by directing them to the promises of God.", 70);
        insert into question (content, gift) values ("I am effective at adapting the gospel message so that it connects with an individual’s felt need.", 71);
        insert into question (content, gift) values ("I believe that God will help me to accomplish great things.", 72);
        insert into question (content, gift) values ("I manage my money well in order to free more of it for giving.", 73);
        insert into question (content, gift) values ("I willingly take on a variety of odd jobs around the church to meet the needs of others.", 74);
        insert into question (content, gift) values ("I genuinely believe the Lord directs strangers to me who need to get connected to others.", 75);
        insert into question (content, gift) values ("I am conscious of ministering to others as I pray.", 76);
        insert into question (content, gift) values ("I am committed, and schedule blocks of time for reading and studying Scripture, to understand Biblical truth fully and accurately.", 77);
        insert into question (content, gift) values ("I can adjust my leadership style to bring out the best in others.", 78);
        insert into question (content, gift) values ("I enjoy helping people sometimes regarded as undeserving or beyond help.", 79);
        insert into question (content, gift) values ("I expose cultural trends, teachings, or events which contradict Biblical principles.", 80);
        insert into question (content, gift) values ("I like to provide guidance for the whole person–relationally, emotionally, spiritually, etc.", 81);
        insert into question (content, gift) values ("I pay close attention to the words, phrases, and meaning of those who teach.", 82);
        insert into question (content, gift) values ("I can easily select the most effective course of action from among several alternatives.", 83);

        insert into question (content, gift) values ("I can identify and effectively use the resources needed to accomplish tasks.", 65);
        insert into question (content, gift) values ("I can adapt well to different cultures and surroundings.", 66);
        insert into question (content, gift) values ("I can visualize how something should be constructed before I build it.  ", 67);
        insert into question (content, gift) values ("I like finding new and fresh ways of communicating God’s truth.", 68);
        insert into question (content, gift) values ("I tend to see rightness or wrongness in situations.", 69);
        insert into question (content, gift) values ("I reassure those who need to take courageous action in their faith, family, or life.", 70);
        insert into question (content, gift) values ("I invite unbelievers to accept Christ as their Savior.", 71);
        insert into question (content, gift) values ("I trust God in circumstances where success cannot be guaranteed by human effort alone.", 72);
        insert into question (content, gift) values ("I am challenged to limit my lifestyle in order to give away a higher percentage of my income.", 73);
        insert into question (content, gift) values ("I see spiritual significance in doing practical tasks.", 74);
        insert into question (content, gift) values ("I like to create a place where people do not feel that they are alone.", 75);
        insert into question (content, gift) values ("I pray with confidence because I know that God works in response to prayer.", 76);
        insert into question (content, gift) values ("I am perfectly at ease answering people’s Bible questions.", 77);
        insert into question (content, gift) values ("I set goals and manage people and resources effectively to accomplish them.", 78);
        insert into question (content, gift) values ("I have great compassion for hurting people.", 79);
        insert into question (content, gift) values ("People often tell me, “God used you.  You dealt exactly with my need.”", 80);
        insert into question (content, gift) values ("I can faithfully provide long-term support and concern for others.", 81);
        insert into question (content, gift) values ("I like to take a systematic approach to my study of the Bible.", 82);
        insert into question (content, gift) values ("I can anticipate the likely consequences of an individual’s or a group’s action. ", 83);
        insert into question (content, gift) values ("I like to help organizations or groups become more efficient.", 84);
        insert into question (content, gift) values ("I can relate to others in culturally sensitive ways.", 85);
        insert into question (content, gift) values ("I honor God with my handcrafted gifts.", 65);
        insert into question (content, gift) values ("I apply various artistic expressions to communicate God’s truth.", 65);
        insert into question (content, gift) values ("I receive affirmation from others concerning the reliability of my insights or                        		perceptions. ", 65);
        insert into question (content, gift) values ("I strengthen those who are wavering in their faith.", 65);
        insert into question (content, gift) values ("I openly tell people that I am a Christian and want them to ask me about my faith.", 65);
        insert into question (content, gift) values ("I am convinced of God’s daily presence and action in my life.", 65);
        insert into question (content, gift) values ("I like knowing that my financial support makes a real difference in the lives and ministries of God’s people.", 65);
        insert into question (content, gift) values ("I like to find small things that need to be done and often do them without being asked.", 65);
        insert into question (content, gift) values ("I enjoy entertaining people and opening my home to others.", 65);

        commit;
    `)

        // insert into question (content, gift) values ("", 65);
        // insert into question (content, gift) values ("", 66);
        // insert into question (content, gift) values ("", 67);
        // insert into question (content, gift) values ("", 68);
        // insert into question (content, gift) values ("", 69);
        // insert into question (content, gift) values ("", 70);
        // insert into question (content, gift) values ("", 71);
        // insert into question (content, gift) values ("", 72);
        // insert into question (content, gift) values ("", 73);
        // insert into question (content, gift) values ("", 74);
        // insert into question (content, gift) values ("", 75);
        // insert into question (content, gift) values ("", 76);
        // insert into question (content, gift) values ("", 77);
        // insert into question (content, gift) values ("", 78);
        // insert into question (content, gift) values ("", 79);
        // insert into question (content, gift) values ("", 80);
        // insert into question (content, gift) values ("", 81);
        // insert into question (content, gift) values ("", 82);
        // insert into question (content, gift) values ("", 83);

	if err != nil {
		return err
	}

	return nil
}
