CREATE TABLE userinfo(id int unique not null auto_increment primary key, OAuth varchar(255), email varchar(255) unique, password varchar(255), name varchar(255), signup_date datetime, safe_search int not null, NG_count int);
CREATE TABLE page_status(id int unique not null auto_increment primary key, title varchar(255), URL varchar(8190), regist_date datetime, last_update datetime, admin int, media varchar(30), genre varchar(255), alive_flag int, tag1 varchar(30), tag2 varchar(30), tag3 varchar(30), tag4 varchar(30), tag5 varchar(30), tag6 varchar(30), tag7 varchar(30), tag8 varchar(30), tag9 varchar(30), tag10 varchar(30));
CREATE TABLE individual_eval_template(num int unique not null auto_increment primary key, posted datetime, evaluator_id int, browse_time datetime, browse_purpose text, deliberate int, describe_eval text, recommend_good int, recommend_bad int, goodness_of_fit int, num_typo int, opt1 int, opt2 int, opt3 int, opt4 int, opt5 int, opt6 int, opt7 int, opt8 int, opt9 int, opt10 int);
CREATE TABLE individual_eval_comment_template(num int unique not null auto_increment primary key, posted varchar(255), commenter_id int, reply_eval_num int, reply_comment_num int, deliberate int, comment text, recommend_good int, recommend_bad int);
CREATE TABLE typo_template(num int unique not null auto_increment primary key, individual_eval_num int, incorrect varchar(255), correct varchar(255));
CREATE TABLE rating_item(genre varchar(255), media varchar(30), opt1 int, opt2 int, opt3 int, opt4 int, opt5 int, opt6 int, opt7 int, opt8 int, opt9 int, opt10 int);
CREATE TABLE page_status_item(genre varchar(255), media varchar(30));
CREATE TABLE NG_word(Lv1 varchar(255), Lv2 varchar(255), Lv3 varchar(255), Lv4 varchar(255), Lv5 varchar(255));
