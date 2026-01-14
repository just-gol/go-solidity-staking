CREATE TABLE IF NOT EXISTS staking_event_staked (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
  log_index BIGINT UNSIGNED NOT NULL COMMENT '日志索引',
  block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
  contract VARCHAR(42) NOT NULL COMMENT '合约地址',
  user VARCHAR(42) NOT NULL COMMENT '用户地址',
  amount VARCHAR(78) NOT NULL COMMENT '质押数量(最小单位)',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tx_log (tx_hash, log_index),
  KEY idx_contract_block (contract, block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Staking: Staked 事件明细';

CREATE TABLE IF NOT EXISTS staking_event_withdrawn (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
  log_index BIGINT UNSIGNED NOT NULL COMMENT '日志索引',
  block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
  contract VARCHAR(42) NOT NULL COMMENT '合约地址',
  user VARCHAR(42) NOT NULL COMMENT '用户地址',
  amount VARCHAR(78) NOT NULL COMMENT '提现数量(最小单位)',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tx_log (tx_hash, log_index),
  KEY idx_contract_block (contract, block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Staking: Withdrawn 事件明细';

CREATE TABLE IF NOT EXISTS staking_event_rewards_claimed (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
  log_index BIGINT UNSIGNED NOT NULL COMMENT '日志索引',
  block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
  contract VARCHAR(42) NOT NULL COMMENT '合约地址',
  user VARCHAR(42) NOT NULL COMMENT '用户地址',
  amount VARCHAR(78) NOT NULL COMMENT '领取奖励(最小单位)',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tx_log (tx_hash, log_index),
  KEY idx_contract_block (contract, block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Staking: RewardsClaimed 事件明细';

CREATE TABLE IF NOT EXISTS staking_event_reward_rate_updated (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
  log_index BIGINT UNSIGNED NOT NULL COMMENT '日志索引',
  block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
  contract VARCHAR(42) NOT NULL COMMENT '合约地址',
  new_reward_rate VARCHAR(78) NOT NULL COMMENT '新奖励速率',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tx_log (tx_hash, log_index),
  KEY idx_contract_block (contract, block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Staking: RewardRateUpdated 事件明细';

CREATE TABLE IF NOT EXISTS erc20_event_transfer (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
  log_index BIGINT UNSIGNED NOT NULL COMMENT '日志索引',
  block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
  contract VARCHAR(42) NOT NULL COMMENT '合约地址',
  `from` VARCHAR(42) NOT NULL COMMENT '转出地址',
  `to` VARCHAR(42) NOT NULL COMMENT '转入地址',
  value VARCHAR(78) NOT NULL COMMENT '转账数量(最小单位)',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tx_log (tx_hash, log_index),
  KEY idx_contract_block (contract, block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ERC20: Transfer 事件明细';

CREATE TABLE IF NOT EXISTS erc20_event_approval (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
  log_index BIGINT UNSIGNED NOT NULL COMMENT '日志索引',
  block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
  contract VARCHAR(42) NOT NULL COMMENT '合约地址',
  owner VARCHAR(42) NOT NULL COMMENT '授权人',
  spender VARCHAR(42) NOT NULL COMMENT '被授权人',
  value VARCHAR(78) NOT NULL COMMENT '授权额度(最小单位)',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tx_log (tx_hash, log_index),
  KEY idx_contract_block (contract, block_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ERC20: Approval 事件明细';
