CREATE TABLE userinfo(
	id int unique not null auto_increment primary key,
	OAuth_service VARCHAR(255),
	OAuth_userinfo VARCHAR(255),
	email varchar(255) unique not null,
	password varchar(255),
	name varchar(255) DEFAULT '名無し',
	signup_date datetime not null DEFAULT CURRENT_TIMESTAMP,
	safe_search tinyint not null DEFAULT 0,
	NG_count int DEFAULT 0,
	dead tinyint not null DEFAULT 0
);

CREATE TABLE tmp_user(
	OAuth_service VARCHAR(255),
	OAuth_userinfo VARCHAR(255) unique not null,
	act varchar(255) unique not null,
	email varchar(255) unique not null,
	referer VARCHAR(255),
	send_time DATETIME not null
);

-- 最後の行で拡張している
CREATE TABLE page_status(
	id int unique not null auto_increment primary key,
	title varchar(255) not null,
	URL varchar(8190) not null,
	register_date datetime not null DEFAULT CURRENT_TIMESTAMP,
	last_update datetime DEFAULT '0001-01-01 01:01:01',
	admin_user_id int DEFAULT 0,
	genre text,
	media text,
	dead tinyint DEFAULT 0,
	tag1 varchar(30),
	tag2 varchar(30),
	tag3 varchar(30),
	tag4 varchar(30),
	tag5 varchar(30),
	tag6 varchar(30),
	tag7 varchar(30),
	tag8 varchar(30),
	tag9 varchar(30),
	tag10 varchar(30)
) ROW_FORMAT=DYNAMIC;

CREATE TABLE individual_eval(
	num int unique not null auto_increment,
	page_id int not null,
	evaluator_id int not null,
	posted datetime not null DEFAULT CURRENT_TIMESTAMP,
	browse_time datetime DEFAULT '0001-01-01 01:01:01',
	browse_purpose text not null,
	deliberate tinyint DEFAULT 0,
	description_eval text,
	recommend_good int DEFAULT 0,
	recommend_bad int DEFAULT 0,
	goodness_of_fit tinyint not null,
	because_goodness_of_fit text,
	device varchar(30),
	visibility tinyint not null,
	because_visibility text,
	num_typo tinyint not null,
	because_num_typo text,
	opt1 int DEFAULT 0,
	because_opt1 text,
	opt2 int DEFAULT 0,
	because_opt2 text,
	opt3 int DEFAULT 0,
	because_opt3 text,
	opt4 int DEFAULT 0,
	because_opt4 text,
	opt5 int DEFAULT 0,
	because_opt5 text,
	opt6 int DEFAULT 0,
	because_opt6 text,
	opt7 int DEFAULT 0,
	because_opt7 text,
	opt8 int DEFAULT 0,
	because_opt8 text,
	opt9 int DEFAULT 0,
	because_opt9 text,
	opt10 int DEFAULT 0,
	because_opt10 text,
	PRIMARY KEY(page_id, evaluator_id)
);

CREATE TABLE individual_eval_recom(
	eval_num int unique not null,
	user_id int unique not null,
	recommend varchar(30) not null,
	PRIMARY KEY(eval_num, user_id)
);

CREATE TABLE individual_eval_comment(
	num int unique not null auto_increment PRIMARY KEY,
	page_id int unique not null,
	commenter_id int not null,
	posted datetime not null DEFAULT CURRENT_TIMESTAMP,
	reply_eval_num int DEFAULT 0,
	reply_comment_num int DEFAULT 0,
	deliberate tinyint DEFAULT 0,
	comment text not null,
	recommend_good int DEFAULT 0,
	recommend_bad int DEFAULT 0
);

CREATE TABLE individual_eval_comment_recom(
	comment_num int not null,
	user_id int not null,
	recommend varchar(30) not null,
	PRIMARY KEY(comment_num, user_id)
);

CREATE TABLE typo(
	num int unique not null auto_increment,
	page_id int not null,
	evaluator_id int not null,
	individual_eval_num int unique DEFAULT 0,
	incorrect varchar(255) not null,
	correct varchar(255) not null,
	PRIMARY KEY(page_id, evaluator_id)
);

CREATE TABLE rating_item(
	num int unique not null auto_increment primary key,
	genre varchar(255) not null,
	media varchar(30) not null,
	opt1 text,
	opt2 text,
	opt3 text,
	opt4 text,
	opt5 text,
	opt6 text,
	opt7 text,
	opt8 text,
	opt9 text,
	opt10 text
);

CREATE TABLE page_status_item(
	num int unique not null auto_increment primary key,
	genre varchar(255),
	media varchar(30)
);

CREATE TABLE NG_word_Lv1(
	num int UNIQUE NOT NULL auto_increment PRIMARY KEY,
	user_id int NOT NULL,
	NG_word varchar(255) NOT NULL
);

CREATE TABLE NG_word_Lv2(
	num int UNIQUE NOT NULL auto_increment PRIMARY KEY,
	user_id int NOT NULL,
	NG_word varchar(255) NOT NULL
);

CREATE TABLE NG_word_Lv3(
	num int UNIQUE NOT NULL auto_increment PRIMARY KEY,
	user_id int NOT NULL,
	NG_word varchar(255) NOT NULL
);

CREATE TABLE NG_word_Lv4(
	num int UNIQUE NOT NULL auto_increment PRIMARY KEY,
	user_id int NOT NULL,
	NG_word varchar(255) NOT NULL
);

CREATE TABLE NG_word_Lv5(
	num int UNIQUE NOT NULL auto_increment PRIMARY KEY,
	user_id int NOT NULL,
	NG_word varchar(255) NOT NULL
);

INSERT INTO page_status_item (genre, media) VALUES(
	'ブログ', '文章'),
	('掲示板', '動画'),
	('5chまとめ', '画像'),
	('企業', '音楽'),
	('ニュース', 'ゲーム'),
	('学術', 'その他'),
	('通販', NULL),
	('漫画・アニメ', NULL),
	('ゲーム', NULL),
	('その他', NULL
);

