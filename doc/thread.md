thread(主题) 对应表关系

### HYBBS
```
mysql> desc hy_thread;
+-----------+---------------------+------+-----+---------+----------------+
| Field     | Type                | Null | Key | Default | Extra          |
+-----------+---------------------+------+-----+---------+----------------+
| id        | int(10) unsigned    | NO   | PRI | NULL    | auto_increment |
| fid       | int(10)             | NO   | MUL | NULL    |                |
| uid       | int(10) unsigned    | NO   | MUL | NULL    |                |
| pid       | int(10) unsigned    | NO   |     | 0       |                |
| title     | char(128)           | NO   |     | NULL    |                |
| summary   | text                | NO   |     | NULL    |                |
| atime     | int(10) unsigned    | NO   | MUL | 0       |                |
| btime     | int(10) unsigned    | NO   | MUL | 0       |                |
| buid      | int(10)             | NO   |     | 0       |                |
| views     | int(10)             | NO   | MUL | 0       |                |
| posts     | int(10)             | NO   | MUL | 0       |                |
| goods     | int(10)             | NO   | MUL | 0       |                |
| nos       | int(10)             | NO   |     | 0       |                |
| img       | text                | NO   |     | NULL    |                |
| img_count | tinyint(3) unsigned | NO   | MUL | 0       |                |
| top       | tinyint(1)          | NO   | MUL | 0       |                |
| files     | tinyint(3) unsigned | NO   |     | 0       |                |
| hide      | tinyint(1)          | NO   |     | 0       |                |
| gold      | int(10) unsigned    | NO   |     | 0       |                |
| state     | tinyint(1)          | NO   |     | 0       |                |
+-----------+---------------------+------+-----+---------+----------------+
20 rows in set (0.00 sec)
```

### DISCUZ
```
mysql> desc cdb_threads;
+-----------------+-----------------------+------+-----+---------+----------------+
| Field           | Type                  | Null | Key | Default | Extra          |
+-----------------+-----------------------+------+-----+---------+----------------+
| tid             | mediumint(8) unsigned | NO   | PRI | NULL    | auto_increment |
| fid             | smallint(6) unsigned  | NO   | MUL | 0       |                |
| iconid          | smallint(6) unsigned  | NO   |     | 0       |                |
| typeid          | smallint(6) unsigned  | NO   |     | 0       |                |
| sortid          | smallint(6) unsigned  | NO   | MUL | 0       |                |
| readperm        | tinyint(3) unsigned   | NO   |     | 0       |                |
| price           | smallint(6)           | NO   |     | 0       |                |
| author          | char(15)              | NO   |     |         |                |
| authorid        | mediumint(8) unsigned | NO   | MUL | 0       |                |
| subject         | char(80)              | NO   |     |         |                |
| dateline        | int(10) unsigned      | NO   |     | 0       |                |
| lastpost        | int(10) unsigned      | NO   |     | 0       |                |
| lastposter      | char(15)              | NO   |     |         |                |
| views           | int(10) unsigned      | NO   |     | 0       |                |
| replies         | mediumint(8) unsigned | NO   |     | 0       |                |
| displayorder    | tinyint(1)            | NO   |     | 0       |                |
| highlight       | tinyint(1)            | NO   |     | 0       |                |
| digest          | tinyint(1)            | NO   | MUL | 0       |                |
| rate            | tinyint(1)            | NO   |     | 0       |                |
| special         | tinyint(1)            | NO   |     | 0       |                |
| attachment      | tinyint(1)            | NO   |     | 0       |                |
| moderated       | tinyint(1)            | NO   |     | 0       |                |
| closed          | mediumint(8) unsigned | NO   |     | 0       |                |
| itemid          | mediumint(8) unsigned | NO   |     | 0       |                |
| supe_pushstatus | tinyint(1)            | NO   |     | 0       |                |
| recommends      | smallint(6)           | NO   | MUL | NULL    |                |
| recommend_add   | smallint(6)           | NO   |     | NULL    |                |
| recommend_sub   | smallint(6)           | NO   |     | NULL    |                |
| heats           | int(10) unsigned      | NO   | MUL | 0       |                |
| status          | smallint(6) unsigned  | NO   |     | 0       |                |
+-----------------+-----------------------+------+-----+---------+----------------+
30 rows in set (0.00 sec)
```

### 对应关系
```
+-----------+---------------------+------+-----+---------+----------------+
| hybbs     | Discuz              |   描述
+-----------+---------------------+------+-----+---------+----------------+
| id        | tid                 |  主题id
| fid       | fid                 |  对应版块id
| uid       | authorid            |  用户uid
| pid       |                     | 对应主题详情在 dz 的 cdb_posts.pid
| title     | subject             | 标题
| summary   | 无                  | 概要
| atime     | dateline            | 创建时间
| btime     | lastpost            | 最后回复时间
| buid      | 无(只有lastposter)  | 最后回复人的uid
| views     | views               | 阅读数
| posts     | replies             | 回复数量
| goods     | 无                  | 点赞
| nos       | 无                  | 点踩
| img       | 无                  | 图片
| img_count | 无                  | 图片数量
| top       | 无                  | 置顶
| files     | attachment          | 附件数量
| hide      | 无                  | 主题是否隐藏
| gold      | 无                  | 主题售价(金币)
| state     | 无                  | 主题状态(1.锁定)
+-----------+---------------------+------+-----+---------+----------------+
```