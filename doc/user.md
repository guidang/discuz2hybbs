user(用户) 对应表关系

### HYBBS
```
mysql> desc hy_user;
+-----------+------------------+------+-----+---------+----------------+
| Field     | Type             | Null | Key | Default | Extra          |
+-----------+------------------+------+-----+---------+----------------+
| id        | int(10)          | NO   | PRI | NULL    | auto_increment |
| user      | varchar(18)      | NO   | MUL | NULL    |                |
| pass      | varchar(32)      | NO   |     | NULL    |                |
| email     | varchar(100)     | NO   | MUL | NULL    |                |
| salt      | varchar(8)       | NO   |     | NULL    |                |
| threads   | int(10) unsigned | NO   |     | NULL    |                |
| posts     | int(10) unsigned | NO   |     | NULL    |                |
| atime     | int(10) unsigned | NO   | MUL | NULL    |                |
| group     | smallint(2)      | NO   | MUL | 0       |                |
| gold      | int(10)          | NO   |     | 0       |                |
| credits   | int(10)          | NO   |     | 0       |                |
| mess      | int(10) unsigned | NO   |     | 0       |                |
| etime     | int(10) unsigned | NO   |     | 0       |                |
| ps        | varchar(40)      | YES  |     | NULL    |                |
| fans      | int(10) unsigned | NO   |     | 0       |                |
| follow    | int(10) unsigned | NO   |     | 0       |                |
| ctime     | int(10) unsigned | NO   |     | 0       |                |
| file_size | int(10) unsigned | NO   |     | 0       |                |
| chat_size | int(10) unsigned | NO   |     | 0       |                |
+-----------+------------------+------+-----+---------+----------------+
19 rows in set (0.05 sec)
```

### DISCUZ
```
mysql> desc cdb_members;
+---------------+-----------------------+------+-----+------------+----------------+
| Field         | Type                  | Null | Key | Default    | Extra          |
+---------------+-----------------------+------+-----+------------+----------------+
| uid           | mediumint(8) unsigned | NO   | PRI | NULL       | auto_increment |
| username      | char(15)              | NO   | UNI |            |                |
| password      | char(32)              | NO   |     |            |                |
| secques       | char(8)               | NO   |     |            |                |
| gender        | tinyint(1)            | NO   |     | 0          |                |
| adminid       | tinyint(1)            | NO   |     | 0          |                |
| groupid       | smallint(6) unsigned  | NO   | MUL | 0          |                |
| groupexpiry   | int(10) unsigned      | NO   |     | 0          |                |
| extgroupids   | char(20)              | NO   |     |            |                |
| regip         | char(15)              | NO   |     |            |                |
| regdate       | int(10) unsigned      | NO   |     | 0          |                |
| lastip        | char(15)              | NO   |     |            |                |
| lastvisit     | int(10) unsigned      | NO   |     | 0          |                |
| lastactivity  | int(10) unsigned      | NO   |     | 0          |                |
| lastpost      | int(10) unsigned      | NO   |     | 0          |                |
| posts         | mediumint(8) unsigned | NO   |     | 0          |                |
| threads       | mediumint(8) unsigned | NO   |     | 0          |                |
| digestposts   | smallint(6) unsigned  | NO   |     | 0          |                |
| oltime        | smallint(6) unsigned  | NO   |     | 0          |                |
| pageviews     | mediumint(8) unsigned | NO   |     | 0          |                |
| credits       | int(10)               | NO   |     | 0          |                |
| extcredits1   | int(10)               | NO   |     | 0          |                |
| extcredits2   | int(10)               | NO   |     | 0          |                |
| extcredits3   | int(10)               | NO   |     | 0          |                |
| extcredits4   | int(10)               | NO   |     | 0          |                |
| extcredits5   | int(10)               | NO   |     | 0          |                |
| extcredits6   | int(10)               | NO   |     | 0          |                |
| extcredits7   | int(10)               | NO   |     | 0          |                |
| extcredits8   | int(10)               | NO   |     | 0          |                |
| email         | char(40)              | NO   | MUL |            |                |
| uc_uid        | int(10)               | YES  | MUL | 0          |                |
| bday          | date                  | NO   |     | 0000-00-00 |                |
| sigstatus     | tinyint(1)            | NO   |     | 0          |                |
| tpp           | tinyint(3) unsigned   | NO   |     | 0          |                |
| ppp           | tinyint(3) unsigned   | NO   |     | 0          |                |
| styleid       | smallint(6) unsigned  | NO   |     | 0          |                |
| dateformat    | tinyint(1)            | NO   |     | 0          |                |
| timeformat    | tinyint(1)            | NO   |     | 0          |                |
| pmsound       | tinyint(1)            | NO   |     | 0          |                |
| showemail     | tinyint(1)            | NO   |     | 0          |                |
| newsletter    | tinyint(1)            | NO   |     | 0          |                |
| invisible     | tinyint(1)            | NO   |     | 0          |                |
| timeoffset    | char(4)               | NO   |     |            |                |
| prompt        | tinyint(1)            | NO   |     | 0          |                |
| accessmasks   | tinyint(1)            | NO   |     | 0          |                |
| editormode    | tinyint(1) unsigned   | NO   |     | 2          |                |
| customshow    | tinyint(1) unsigned   | NO   |     | 26         |                |
| xspacestatus  | tinyint(1)            | NO   |     | 0          |                |
| customaddfeed | tinyint(1)            | NO   |     | 0          |                |
| msnid         | char(16)              | NO   |     |            |                |
| conisbind     | tinyint(1) unsigned   | NO   | MUL | 0          |                |
| newbietaskid  | smallint(6) unsigned  | NO   |     | 0          |                |
+---------------+-----------------------+------+-----+------------+----------------+
52 rows in set (0.03 sec)
```

### 对应关系
```
+-----------+---------------------+------+-----+---------+----------------+
| hybbs     | Discuz              |   描述
+-----------+---------------------+------+-----+---------+----------------+
| id        | uid                 | 用户id
| user      | uername             | 用户名
| pass      | password            | 密码
| email     | email               | 邮箱
| salt      |                     | 盐值
| threads   | threads             | 主题数
| posts     | posts               | 帖子数
| atime     | regdate             | 注册时间
| group     | groupid             | 用户组id(由于dz和hy分组体系不同,建议不要使用)              
| gold      | extcredits1 ~ 8     | 金钱(由于dz和hy物品体系不同,建议不要使用)
| credits   | credits             | 积分
| mess      | 无                  |
| etime     | 无                  |
| ps        | 无                  |
| fans      | 无                  |
| follow    | 无                  |
| ctime     | lastvisit           | 最后访问时间
| file_size | 无                  | 
| chat_size | 无                  | 
+-----------+------------------+------+-----+---------+----------------+
19 rows in set (0.05 sec)
```

**注意:**   
group = 1 时, 为管理员, 2 时为普通用户.  
