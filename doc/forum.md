forum 对应表关系

### HYBBS
```
mysql> desc hy_forum;
+------------+-------------+------+-----+---------+-------+
| Field      | Type        | Null | Key | Default | Extra |
+------------+-------------+------+-----+---------+-------+
| id         | int(10)     | NO   | PRI | NULL    |       |
| fid        | int(10)     | NO   | MUL | -1      |       |
| fgid       | int(11)     | NO   |     | 1       |       |
| name       | varchar(12) | NO   |     | NULL    |       |
| name2      | varchar(18) | NO   |     | NULL    |       |
| threads    | int(10)     | NO   |     | 0       |       |
| posts      | int(10)     | NO   |     | 0       |       |
| forumg     | text        | NO   |     | NULL    |       |
| json       | text        | NO   |     | NULL    |       |
| html       | longtext    | NO   |     | NULL    |       |
| color      | varchar(30) | NO   |     | NULL    |       |
| background | varchar(30) | NO   |     | NULL    |       |
+------------+-------------+------+-----+---------+-------+
12 rows in set (0.00 sec)
```

### DISCUZ
```
mysql> desc cdb_forums;
+------------------+-----------------------------+------+-----+---------+----------------+
| Field            | Type                        | Null | Key | Default | Extra          |
+------------------+-----------------------------+------+-----+---------+----------------+
| fid              | smallint(6) unsigned        | NO   | PRI | NULL    | auto_increment |
| fup              | smallint(6) unsigned        | NO   | MUL | 0       |                |
| type             | enum('group','forum','sub') | NO   |     | forum   |                |
| name             | char(50)                    | NO   |     |         |                |
| status           | tinyint(1)                  | NO   | MUL | 0       |                |
| displayorder     | smallint(6)                 | NO   |     | 0       |                |
| styleid          | smallint(6) unsigned        | NO   |     | 0       |                |
| threads          | mediumint(8) unsigned       | NO   |     | 0       |                |
| posts            | mediumint(8) unsigned       | NO   |     | 0       |                |
| todayposts       | mediumint(8) unsigned       | NO   |     | 0       |                |
| lastpost         | char(110)                   | NO   |     |         |                |
| allowsmilies     | tinyint(1)                  | NO   |     | 0       |                |
| allowhtml        | tinyint(1)                  | NO   |     | 0       |                |
| allowbbcode      | tinyint(1)                  | NO   |     | 0       |                |
| allowimgcode     | tinyint(1)                  | NO   |     | 0       |                |
| allowmediacode   | tinyint(1)                  | NO   |     | 0       |                |
| allowanonymous   | tinyint(1)                  | NO   |     | 0       |                |
| allowshare       | tinyint(1)                  | NO   |     | 0       |                |
| allowpostspecial | smallint(6) unsigned        | NO   |     | 0       |                |
| allowspecialonly | tinyint(1) unsigned         | NO   |     | 0       |                |
| alloweditrules   | tinyint(1)                  | NO   |     | 0       |                |
| allowfeed        | tinyint(1)                  | NO   |     | 1       |                |
| recyclebin       | tinyint(1)                  | NO   |     | 0       |                |
| modnewposts      | tinyint(1)                  | NO   |     | 0       |                |
| jammer           | tinyint(1)                  | NO   |     | 0       |                |
| disablewatermark | tinyint(1)                  | NO   |     | 0       |                |
| inheritedmod     | tinyint(1)                  | NO   |     | 0       |                |
| autoclose        | smallint(6)                 | NO   |     | 0       |                |
| forumcolumns     | tinyint(3) unsigned         | NO   |     | 0       |                |
| threadcaches     | tinyint(1)                  | NO   |     | 0       |                |
| alloweditpost    | tinyint(1) unsigned         | NO   |     | 1       |                |
| simple           | tinyint(1) unsigned         | NO   |     | NULL    |                |
| allowtag         | tinyint(1)                  | NO   |     | 1       |                |
| modworks         | tinyint(1) unsigned         | NO   |     | NULL    |                |
| allowglobalstick | tinyint(1)                  | NO   |     | 1       |                |
+------------------+-----------------------------+------+-----+---------+----------------+
35 rows in set (0.00 sec)
```

### DISCUZ7.2 Tables
```
cdb_forums
cdb_forumfields
```

### 对应关系
```
+------------+-------------+---------+
| Hybbs      | Discuz      |   描述  |
+------------+-------------+---------+
| id         | fid         | 版块id   
| fid        | fup/-1      | 父版块id, Hybbs -1则为顶级id, Discuz 0则为顶级
| fgid       | fup         | 分组id   
| name       | name        | 版块名 
| threads    | threads     | 主题数
| posts      | posts       | 回复数
| html       |             | 描述  cdb_forumfields.description   
+------------+-------------+---------+
```

**注意:**   
discuz type 为group时，则其为hybbs的大分组, 不添加至此table   
当 type为forum, fup的值为大分组的值时,    
   则hybbs fid对应值应该为-1, fgid为fup的值   