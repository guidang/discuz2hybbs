post(帖子) 对应表关系

### HYBBS
```
mysql> desc hy_post;
+----------+------------+------+-----+---------+----------------+
| Field    | Type       | Null | Key | Default | Extra          |
+----------+------------+------+-----+---------+----------------+
| id       | int(10)    | NO   | PRI | NULL    | auto_increment |
| tid      | int(10)    | NO   | MUL | NULL    |                |
| fid      | int(10)    | NO   |     | NULL    |                |
| uid      | int(10)    | NO   | MUL | NULL    |                |
| isthread | tinyint(1) | NO   |     | 0       |                |
| content  | longtext   | NO   |     | NULL    |                |
| atime    | int(10)    | NO   | MUL | NULL    |                |
| goods    | int(10)    | YES  |     | 0       |                |
| nos      | int(10)    | NO   |     | 0       |                |
| posts    | int(10)    | NO   |     | 0       |                |
+----------+------------+------+-----+---------+----------------+
10 rows in set (0.00 sec)
```

### DISCUZ
```
mysql> desc cdb_posts;
+-------------+-----------------------+------+-----+---------+----------------+
| Field       | Type                  | Null | Key | Default | Extra          |
+-------------+-----------------------+------+-----+---------+----------------+
| pid         | int(10) unsigned      | NO   | PRI | NULL    | auto_increment |
| fid         | smallint(6) unsigned  | NO   | MUL | 0       |                |
| tid         | mediumint(8) unsigned | NO   | MUL | 0       |                |
| first       | tinyint(1)            | NO   |     | 0       |                |
| author      | varchar(15)           | NO   |     |         |                |
| authorid    | mediumint(8) unsigned | NO   | MUL | 0       |                |
| subject     | varchar(80)           | NO   |     |         |                |
| dateline    | int(10) unsigned      | NO   | MUL | 0       |                |
| message     | mediumtext            | NO   |     | NULL    |                |
| useip       | varchar(15)           | NO   |     |         |                |
| invisible   | tinyint(1)            | NO   | MUL | 0       |                |
| anonymous   | tinyint(1)            | NO   |     | 0       |                |
| usesig      | tinyint(1)            | NO   |     | 0       |                |
| htmlon      | tinyint(1)            | NO   |     | 0       |                |
| bbcodeoff   | tinyint(1)            | NO   |     | 0       |                |
| smileyoff   | tinyint(1)            | NO   |     | 0       |                |
| parseurloff | tinyint(1)            | NO   |     | 0       |                |
| attachment  | tinyint(1)            | NO   |     | 0       |                |
| rate        | smallint(6)           | NO   |     | 0       |                |
| ratetimes   | tinyint(3) unsigned   | NO   |     | 0       |                |
| status      | tinyint(1)            | NO   |     | 0       |                |
+-------------+-----------------------+------+-----+---------+----------------+
21 rows in set (0.00 sec)
```

### 对应关系
```
+----------+------------+------+-----+---------+----------------+
| hybbs    | discuz     | 备注
+----------+------------+------+-----+---------+----------------+
| id       | pid        | 帖子id
| tid      | tid        | 对应主题id
| fid      | fid        | 版块id
| uid      | authorid   | 发帖者id
| isthread | first      | 是否为主题(1)
| content  | message    | 内容(html)
| atime    | dateline   | 创建时间
| goods    | 无         | 点赞
| nos      | 无         | 点踩
| posts    |            | 未知
+----------+------------+------+-----+---------+----------------+
```