package tools

import (
	"fmt"
	"strings"
	"wechat/core/database"
)

/****
`CREATE TABLE `$table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) NOT NULL DEFAULT '',
  `sender` char(32) NOT NULL DEFAULT '',
  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',
  `type` int(4) NOT NULL DEFAULT 0,
  `second_type` int(4) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `content` varchar(2500) NOT NULL DEFAULT '',
  `parent` char(32) NOT NULL DEFAULT '',
  `deletes` varchar(2500) NOT NULL DEFAULT '',
  `reads` varchar(2500) NOT NULL DEFAULT '',
  `send_time` int(10) NOT NULL DEFAULT 0,
  `receive_time` int(10) NOT NULL DEFAULT 0,
  `forward_time` int(10) NOT NULL DEFAULT 0,
  `log_time` int(10) NOT NULL DEFAULT 0,
  `update_time` int(10) NOT NULL DEFAULT 0,
  `remark` varchar(50) NOT NULL DEFAULT '',
  `is_del` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`,`uuid`) USING BTREE,
  KEY `sender` (`sender`) USING BTREE,
  KEY `recipient` (`recipient`) USING BTREE,
  KEY `type` (`type`) USING BTREE,
  KEY `second_type` (`second_type`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent` (`parent`) USING BTREE,
  KEY `send_time` (`send_time`) USING BTREE,
  KEY `is_del` (`is_del`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;`
*/

const TableIdentifier = "$table"

func ExecSql(sourceSql, table string, total int) {
	var tableCopy string
	for i := 0; i < total; i++ {
		if i < 10 {
			if i == 0 {
				tableCopy = table
			} else {
				tableCopy = fmt.Sprintf("%s_0%d", table, i)
			}
		} else {
			tableCopy = fmt.Sprintf("%s_%d", table, i)
		}
		sql := strings.Replace(sourceSql, TableIdentifier, tableCopy, 1)
		database.GetDB().Exec(sql)
	}
}
