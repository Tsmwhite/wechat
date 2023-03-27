package main

import (
	"wechat/config"
	"wechat/core/database/tools"
	"wechat/core/log"
)

func main() {
	err := config.SetupServer()
	if err != nil {
		log.PrintlnErr("setup error:", err)
	}
	//sql := "CREATE TABLE `$table` (\n  `id` int(11) NOT NULL AUTO_INCREMENT,\n  `uuid` char(32) NOT NULL DEFAULT '',\n  `sender` char(32) NOT NULL DEFAULT '',\n  `recipient` char(32) NOT NULL DEFAULT '' COMMENT '接受消息房间uuid',\n  `type` int(4) NOT NULL DEFAULT 0,\n  `second_type` int(4) NOT NULL DEFAULT 0,\n  `status` tinyint(1) NOT NULL DEFAULT 0,\n  `content` varchar(2500) NOT NULL DEFAULT '',\n  `parent` char(32) NOT NULL DEFAULT '',\n  `deletes` varchar(2500) NOT NULL DEFAULT '',\n  `reads` varchar(2500) NOT NULL DEFAULT '',\n  `send_time` int(10) NOT NULL DEFAULT 0,\n  `receive_time` int(10) NOT NULL DEFAULT 0,\n  `forward_time` int(10) NOT NULL DEFAULT 0,\n  `log_time` int(10) NOT NULL DEFAULT 0,\n  `update_time` int(10) NOT NULL DEFAULT 0,\n  `remark` varchar(50) NOT NULL DEFAULT '',\n  `is_del` tinyint(1) NOT NULL DEFAULT 0,\n  PRIMARY KEY (`id`,`uuid`) USING BTREE,\n  KEY `sender` (`sender`) USING BTREE,\n  KEY `recipient` (`recipient`) USING BTREE,\n  KEY `type` (`type`) USING BTREE,\n  KEY `second_type` (`second_type`) USING BTREE,\n  KEY `status` (`status`) USING BTREE,\n  KEY `parent` (`parent`) USING BTREE,\n  KEY `send_time` (`send_time`) USING BTREE,\n  KEY `is_del` (`is_del`)\n) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	//tools.ExecSql(sql, "messages", 10)

	sql := "ALTER TABLE `$table` \nADD INDEX `recipient_read_status`(`room`, `recipient`, `is_read`);"
	tools.ExecSql(sql, "receive_messages", 20)
}
