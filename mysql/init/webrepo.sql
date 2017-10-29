CREATE TABLE userinfo(
	id int unique not null auto_increment primary key,
	OAuth_service VARCHAR(255),
	OAuth_userinfo VARCHAR(255),
	email varchar(255) unique not null,
	password varchar(255),
	name varchar(255) DEFAULT '名無し',
	signup_date datetime not null,
	safe_search tinyint not null DEFAULT 0,
	NG_count int DEFAULT 0
);

CREATE TABLE tmp_user(
	OAuth_service VARCHAR(255),
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
	regist_date datetime not null,
	last_update datetime,
	admin_user_id int,
	genre text,
	media text,
	dead tinyint,
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

CREATE TABLE individual_eval_template(
	num int unique not null auto_increment primary key,
	posted datetime not null,
	evaluator_id int unique not null,
	browse_time datetime not null,
	browse_purpose text not null,
	deliberate tinyint DEFAULT 0,
	description_eval text,
	recommend_good int,
	recommend_bad int,
	goodness_of_fit tinyint not null,
	because_goodness_of_fit text,
	visibility tinyint not null,
	because_visibility text,
	num_typo tinyint not null,
	because_num_typo text,
	opt1 int,
	because_opt1 text,
	opt2 int,
	because_opt2 text,
	opt3 int,
	because_opt3 text,
	opt4 int,
	because_opt4 text,
	opt5 int,
	because_opt5 text,
	opt6 int,
	because_opt6 text,
	opt7 int,
	because_opt7 text,
	opt8 int,
	because_opt8 text,
	opt9 int,
	because_opt9 text,
	opt10 int,
	because_opt10 text
);

CREATE TABLE individual_eval_comment_template(
	num int unique not null auto_increment primary key,
	posted datetime not null,
	commenter_id int not null,
	reply_eval_num int,
	reply_comment_num int,
	deliberate tinyint DEFAULT 0,
	comment text not null,
	recommend_good int,
	recommend_bad int
);

CREATE TABLE typo_template(num int unique not null auto_increment primary key,
	evaluator_id int not null,
	individual_eval_num int not null,
	incorrect varchar(255) not null,
	correct varchar(255) not null
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

CREATE TABLE all_NG_word(
	Lv1 varchar(255),
	Lv2 varchar(255),
	Lv3 varchar(255),
	Lv4 varchar(255),
	Lv5 varchar(255)
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
	('その他', NULL);
	
