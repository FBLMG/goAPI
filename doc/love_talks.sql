-- ----------------------------------------------------------------------------------------------------------
-- 情话板
-- ----------------------------------------------------------------------------------------------------------

CREATE TABLE `love_talks` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type_id` int(11) DEFAULT '0' COMMENT '分类Id',
  `motto` varchar(225) DEFAULT '' COMMENT '文案',
  `motto_card` varchar(552) DEFAULT '' COMMENT '文案封面',
  `card` varchar(552) DEFAULT '' COMMENT '卡片封面',
  `status` int(11) DEFAULT '1' COMMENT '状态 1：显示 2：隐藏',
  `sort` int(11) DEFAULT '0' COMMENT '序号',
  PRIMARY KEY (`id`),
  KEY `idx_type_id` (`type_id`),
  KEY `idx_motto` (`motto`),
  KEY `idx_status` (`status`),
  KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='情话板-素材列表';