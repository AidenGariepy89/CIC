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

	// Init Questions Table
	result := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='question'")

	var name string
	err = result.Scan(&name)
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

	Db = db
	return nil
}

func setupGifts(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS gift (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        key INTEGER NOT NULL
    )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`BEGIN TRANSACTION;
        INSERT INTO gift (name, description, key) VALUES ("Administration", "", 65);
        INSERT INTO gift (name, description, key) VALUES ("Apostleship", "", 66);
        INSERT INTO gift (name, description, key) VALUES ("Crafting/Craftsmanship", "", 67);
        INSERT INTO gift (name, description, key) VALUES ("Creative Communication", "", 68);
        INSERT INTO gift (name, description, key) VALUES ("Discernment", "", 69);
        INSERT INTO gift (name, description, key) VALUES ("Encouragement", "", 70);
        INSERT INTO gift (name, description, key) VALUES ("Evangelism", "", 71);
        INSERT INTO gift (name, description, key) VALUES ("Faith", "", 72);
        INSERT INTO gift (name, description, key) VALUES ("Giving", "", 73);
        INSERT INTO gift (name, description, key) VALUES ("Helps", "", 74);
        INSERT INTO gift (name, description, key) VALUES ("Hospitality", "", 75);
        INSERT INTO gift (name, description, key) VALUES ("Intercession", "", 76);
        INSERT INTO gift (name, description, key) VALUES ("Knowledge", "", 77);
        INSERT INTO gift (name, description, key) VALUES ("Leadership", "", 78);
        INSERT INTO gift (name, description, key) VALUES ("Mercy", "", 79);
        INSERT INTO gift (name, description, key) VALUES ("Prophecy", "", 80);
        INSERT INTO gift (name, description, key) VALUES ("Shepherding", "", 81);
        INSERT INTO gift (name, description, key) VALUES ("Teaching", "", 82);
        INSERT INTO gift (name, description, key) VALUES ("Wisdom", "", 83);
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
        gift INTEGER NOT NULL
    )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`BEGIN TRANSACTION;
        INSERT INTO question (content, gift) VALUES ("I like to organize people, tasks, and events.", 65);
        INSERT INTO question (content, gift) VALUES ("I would like to start churches in places where they do not presently exist.", 66);
        INSERT INTO question (content, gift) VALUES ("I enjoy working creatively with wood, cloth, paints, metal, glass, or other materials.", 67);
        INSERT INTO question (content, gift) VALUES ("I enjoy challenging people’s perspective of God by using various forms of art.", 68);
        INSERT INTO question (content, gift) VALUES ("I can readily distinguish between spiritual truth and error, good and evil.", 69);
        INSERT INTO question (content, gift) VALUES ("I tend to see the potential in people.", 70);
        INSERT INTO question (content, gift) VALUES ("I communicate the gospel to others with clarity and effectiveness.", 71);
        INSERT INTO question (content, gift) VALUES ("I find it natural and easy to trust God to answer my prayers.", 72);
        INSERT INTO question (content, gift) VALUES ("I give liberally and joyfully to people in financial need or to projects requiring support.", 73);
        INSERT INTO question (content, gift) VALUES ("I enjoy working behind the scenes to support the work of others.", 74);
        INSERT INTO question (content, gift) VALUES ("I view my home as a place to minister to people in need.", 75);
        INSERT INTO question (content, gift) VALUES ("I take prayer requests from others and consistently pray for them.", 76);
        INSERT INTO question (content, gift) VALUES ("I am approached by people who want to know my perspective on a particular passage or Biblical truth.", 77);
        INSERT INTO question (content, gift) VALUES ("I am able to motivate others to accomplish a goal.", 78);
        INSERT INTO question (content, gift) VALUES ("I empathize with hurting people and desire to help in their healing process.", 79);
        INSERT INTO question (content, gift) VALUES ("I can speak in a way that results in conviction and change in the lives of others.", 80);
        INSERT INTO question (content, gift) VALUES ("I enjoy spending time nurturing and caring for others.", 81);
        INSERT INTO question (content, gift) VALUES ("I am able to communicate God’s Word effectively to adults, youth, or children.", 82);
        INSERT INTO question (content, gift) VALUES ("I am often sought out by others for advice about spiritual or personal matters.", 83);

        INSERT INTO question (content, gift) VALUES ("I am careful, thorough, and skilled at managing details.", 65);
        INSERT INTO question (content, gift) VALUES ("I am attracted to the idea of serving in another country or ethnic community.", 66);
        INSERT INTO question (content, gift) VALUES ("I am skilled in working with different kinds of tools.", 67);
        INSERT INTO question (content, gift) VALUES ("I enjoy developing and using my artistic skills (art, drama, music, photography, etc.).", 68);
        INSERT INTO question (content, gift) VALUES ("I frequently am able to judge a person’s character based upon first impressions.", 69);
        INSERT INTO question (content, gift) VALUES ("I enjoy reassuring and strengthening those who are discouraged.", 70);
        INSERT INTO question (content, gift) VALUES ("I consistently look for opportunities to build relationships with non-Christians.", 71);
        INSERT INTO question (content, gift) VALUES ("I have confidence in God’s continuing provision and help, even in difficult times.", 72);
        INSERT INTO question (content, gift) VALUES ("I give more than a tithe so that God’s work can be accomplished.", 73);
        INSERT INTO question (content, gift) VALUES ("I enjoy doing routine tasks that support the ministry.", 74);
        INSERT INTO question (content, gift) VALUES ("I enjoy meeting new people and helping them to feel welcomed.", 75);
        INSERT INTO question (content, gift) VALUES ("I enjoy praying for long periods of time and receive leadings as to what God wants me to pray for.", 76);
        INSERT INTO question (content, gift) VALUES ("With the help of appropriate study materials, I can find what God’s Word teaches on most topics.", 77);
        INSERT INTO question (content, gift) VALUES ("I am able to influence others to achieve a vision.", 78);
        INSERT INTO question (content, gift) VALUES ("I can patiently support those going through painful experiences as they try to stabilize their lives.", 79);
        INSERT INTO question (content, gift) VALUES ("People in trouble are encouraged when I talk with them.", 80);
        INSERT INTO question (content, gift) VALUES ("I have compassion for wandering believers and want to protect them.", 81);
        INSERT INTO question (content, gift) VALUES ("I can spend time in study knowing that presenting truth will make a difference in the lives of people–young or old.", 82);
        INSERT INTO question (content, gift) VALUES ("I can often find simple, practical solutions in the midst of conflict or confusion.", 83);

        INSERT INTO question (content, gift) VALUES ("I can clarify goals and develop strategies or plans to accomplish them.", 65);
        INSERT INTO question (content, gift) VALUES ("I am willing to take an active part in starting a new church.", 66);
        INSERT INTO question (content, gift) VALUES ("I enjoy making things for use in ministry.", 67);
        INSERT INTO question (content, gift) VALUES ("I help people understand themselves, their relationships, and God better through artistic expression.", 68);
        INSERT INTO question (content, gift) VALUES ("I can see through phoniness or deceit before it is evident to others.", 69);
        INSERT INTO question (content, gift) VALUES ("I give hope to others by directing them to the promises of God.", 70);
        INSERT INTO question (content, gift) VALUES ("I am effective at adapting the gospel message so that it connects with an individual’s felt need.", 71);
        INSERT INTO question (content, gift) VALUES ("I believe that God will help me to accomplish great things.", 72);
        INSERT INTO question (content, gift) VALUES ("I manage my money well in order to free more of it for giving.", 73);
        INSERT INTO question (content, gift) VALUES ("I willingly take on a variety of odd jobs around the church to meet the needs of others.", 74);
        INSERT INTO question (content, gift) VALUES ("I genuinely believe the Lord directs strangers to me who need to get connected to others.", 75);
        INSERT INTO question (content, gift) VALUES ("I am conscious of ministering to others as I pray.", 76);
        INSERT INTO question (content, gift) VALUES ("I am committed, and schedule blocks of time for reading and studying Scripture, to understand Biblical truth fully and accurately.", 77);
        INSERT INTO question (content, gift) VALUES ("I can adjust my leadership style to bring out the best in others.", 78);
        INSERT INTO question (content, gift) VALUES ("I enjoy helping people sometimes regarded as undeserving or beyond help.", 79);
        INSERT INTO question (content, gift) VALUES ("I expose cultural trends, teachings, or events which contradict Biblical principles.", 80);
        INSERT INTO question (content, gift) VALUES ("I like to provide guidance for the whole person–relationally, emotionally, spiritually, etc.", 81);
        INSERT INTO question (content, gift) VALUES ("I pay close attention to the words, phrases, and meaning of those who teach.", 82);
        INSERT INTO question (content, gift) VALUES ("I can easily select the most effective course of action from among several alternatives.", 83);

        INSERT INTO question (content, gift) VALUES ("I can identify and effectively use the resources needed to accomplish tasks.", 65);
        INSERT INTO question (content, gift) VALUES ("I can adapt well to different cultures and surroundings.", 66);
        INSERT INTO question (content, gift) VALUES ("I can visualize how something should be constructed before I build it.  ", 67);
        INSERT INTO question (content, gift) VALUES ("I like finding new and fresh ways of communicating God’s truth.", 68);
        INSERT INTO question (content, gift) VALUES ("I tend to see rightness or wrongness in situations.", 69);
        INSERT INTO question (content, gift) VALUES ("I reassure those who need to take courageous action in their faith, family, or life.", 70);
        INSERT INTO question (content, gift) VALUES ("I invite unbelievers to accept Christ as their Savior.", 71);
        INSERT INTO question (content, gift) VALUES ("I trust God in circumstances where success cannot be guaranteed by human effort alone.", 72);
        INSERT INTO question (content, gift) VALUES ("I am challenged to limit my lifestyle in order to give away a higher percentage of my income.", 73);
        INSERT INTO question (content, gift) VALUES ("I see spiritual significance in doing practical tasks.", 74);
        INSERT INTO question (content, gift) VALUES ("I like to create a place where people do not feel that they are alone.", 75);
        INSERT INTO question (content, gift) VALUES ("I pray with confidence because I know that God works in response to prayer.", 76);
        INSERT INTO question (content, gift) VALUES ("I am perfectly at ease answering people’s Bible questions.", 77);
        INSERT INTO question (content, gift) VALUES ("I set goals and manage people and resources effectively to accomplish them.", 78);
        INSERT INTO question (content, gift) VALUES ("I have great compassion for hurting people.", 79);
        INSERT INTO question (content, gift) VALUES ("People often tell me, “God used you.  You dealt exactly with my need.”", 80);
        INSERT INTO question (content, gift) VALUES ("I can faithfully provide long-term support and concern for others.", 81);
        INSERT INTO question (content, gift) VALUES ("I like to take a systematic approach to my study of the Bible.", 82);
        INSERT INTO question (content, gift) VALUES ("I can anticipate the likely consequences of an individual’s or a group’s action. ", 83);

        INSERT INTO question (content, gift) VALUES ("I like to help organizations or groups become more efficient.", 65);
        INSERT INTO question (content, gift) VALUES ("I can relate to others in culturally sensitive ways.", 66);
        INSERT INTO question (content, gift) VALUES ("I honor God with my handcrafted gifts.", 67);
        INSERT INTO question (content, gift) VALUES ("I apply various artistic expressions to communicate God’s truth.", 68);
        INSERT INTO question (content, gift) VALUES ("I receive affirmation from others concerning the reliability of my insights or perceptions. ", 69);
        INSERT INTO question (content, gift) VALUES ("I strengthen those who are wavering in their faith.", 70);
        INSERT INTO question (content, gift) VALUES ("I openly tell people that I am a Christian and want them to ask me about my faith.", 71);
        INSERT INTO question (content, gift) VALUES ("I am convinced of God’s daily presence and action in my life.", 72);
        INSERT INTO question (content, gift) VALUES ("I like knowing that my financial support makes a real difference in the lives and ministries of God’s people.", 73);
        INSERT INTO question (content, gift) VALUES ("I like to find small things that need to be done and often do them without being asked.", 74);
        INSERT INTO question (content, gift) VALUES ("I enjoy entertaining people and opening my home to others.", 75);
        INSERT INTO question (content, gift) VALUES ("When I hear about needy situations, I feel burdened to pray.", 76);
        INSERT INTO question (content, gift) VALUES ("Salvation by faith alone is a truth I clearly understand.", 77);
        INSERT INTO question (content, gift) VALUES ("I influence others to perform to the best of their capability.", 78);
        INSERT INTO question (content, gift) VALUES ("I can look beyond a person’s handicaps or problems to see a life that matters to God.", 79);
        INSERT INTO question (content, gift) VALUES ("I appreciate people who are honest and will speak the truth.", 80);
        INSERT INTO question (content, gift) VALUES ("I enjoy giving guidance and practical support to a small group of people.", 81);
        INSERT INTO question (content, gift) VALUES ("I can communicate Scripture in ways that motivate others to study and want to learn more.", 82);
        INSERT INTO question (content, gift) VALUES ("I give practical advice to help others through complicated situations.", 83);

        INSERT INTO question (content, gift) VALUES ("I enjoy learning about how organizations function.", 65);
        INSERT INTO question (content, gift) VALUES ("I enjoy pioneering new undertakings.", 66);
        INSERT INTO question (content, gift) VALUES ("I am good at and enjoy working with my hands. ", 67);
        INSERT INTO question (content, gift) VALUES ("I am creative and imaginative.", 68);
        INSERT INTO question (content, gift) VALUES ("I can identify preaching, teaching, or communication which is not true to the Bible. ", 69);
        INSERT INTO question (content, gift) VALUES ("I like motivating others to take steps for spiritual growth.", 70);
        INSERT INTO question (content, gift) VALUES ("I openly and confidently tell others what Christ has done for me.", 71);
        INSERT INTO question (content, gift) VALUES ("I am regularly challenging others to trust God.", 72);
        INSERT INTO question (content, gift) VALUES ("I give generously due to my commitment to stewardship.", 73);
        INSERT INTO question (content, gift) VALUES ("I feel comfortable being a helper, assisting others to do their job more effectively.  ", 74);
        INSERT INTO question (content, gift) VALUES ("I do whatever I can to make people feel that they belong.", 75);
        INSERT INTO question (content, gift) VALUES ("I am honored when someone asks me to pray for them.", 76);
        INSERT INTO question (content, gift) VALUES ("I discover important Biblical truths when reading or studying Scripture which benefit others in the body of Christ.", 77);
        INSERT INTO question (content, gift) VALUES ("I am able to cast a vision that others want to be a part of.", 78);
        INSERT INTO question (content, gift) VALUES ("I enjoy bringing hope and joy to people living in difficult circumstances.", 79);
        INSERT INTO question (content, gift) VALUES ("I will speak God’s truth, even in places there it is unpopular or difficult for others to accept.", 80);
        INSERT INTO question (content, gift) VALUES ("I can gently restore wandering believers to faith and fellowship.", 81);
        INSERT INTO question (content, gift) VALUES ("I can present information and skills to others at a level that makes it easy for them to grasp and apply to their lives.", 82);
        INSERT INTO question (content, gift) VALUES ("I can apply Scriptural truth that others regard as practical and helpful.", 83);

        INSERT INTO question (content, gift) VALUES ("I can visualize a coming event, anticipate potential problems, and develop backup plans.", 65);
        INSERT INTO question (content, gift) VALUES ("I am able to orchestrate or oversee several church ministries.", 66);
        INSERT INTO question (content, gift) VALUES ("I am able to design and construct things that help the church.", 67);
        INSERT INTO question (content, gift) VALUES ("I regularly need to get alone to reflect and develop my imagination.", 68);
        INSERT INTO question (content, gift) VALUES ("I can tell whether a person is being influenced by the Lord or Satan.", 69);
        INSERT INTO question (content, gift) VALUES ("I am often asked to help those in trouble resolve their problems.", 70);
        INSERT INTO question (content, gift) VALUES ("I seek opportunities to talk about spiritual matters with unbelievers.", 71);
        INSERT INTO question (content, gift) VALUES ("I can move forward in spite of opposition or lack of support when I sense God’s blessing on an undertaking.", 72);
        INSERT INTO question (content, gift) VALUES ("I believe I have been given an abundance of resources so that I may give more to the Lord’s work.", 73);
        INSERT INTO question (content, gift) VALUES ("I readily and happily use my natural or learned skills to help wherever needed.", 74);
        INSERT INTO question (content, gift) VALUES ("I can make people feel at ease even in unfamiliar surroundings.", 75);
        INSERT INTO question (content, gift) VALUES ("I often see specific results in direct response to my prayers.", 76);
        INSERT INTO question (content, gift) VALUES (".  I confidently share my knowledge and insights with others.", 77);
        INSERT INTO question (content, gift) VALUES ("I figure out where we need to go and help others to get there.", 78);
        INSERT INTO question (content, gift) VALUES ("I enjoy doing practical things for others who are in need.", 79);
        INSERT INTO question (content, gift) VALUES ("I feel compelled to expose sin wherever I see it and to challenge people to repentance.", 80);
        INSERT INTO question (content, gift) VALUES ("I enjoy patiently but firmly nurturing others in their development as believers.", 81);
        INSERT INTO question (content, gift) VALUES ("I enjoy explaining things to people so that they can grow spiritually and personally.", 82);
        INSERT INTO question (content, gift) VALUES ("I have insights into how to solve problems that others often do not see.", 83);

        COMMIT;
    `)

	if err != nil {
		return err
	}

	return nil
}
