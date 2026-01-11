# go-solidity-staking
## 测试网合约地址
- `0xf37894409dF321A66a2D9F167e5064b9B5C5da28`
- 本地链 staking + ERC20 示例，包含合约编译、部署、监听和 HTTP API。

## 环境
- Go
- solc (0.8.20+)
- 本地链 (Hardhat/Anvil/Ganache 任一)

## 合约
合约文件在 `contract/`：
- `staking.sol`
- `erc20.sol`
- `ierc20.sol`

### 编译 & 生成 Go 绑定
示例：
```bash
solc --abi --bin contract/staking.sol -o build
solc --abi --bin contract/erc20.sol -o build

abigen --bin=build/Staking.bin --abi=build/Staking.abi --pkg=staking --out=gen/staking/staking.go
abigen --bin=build/ERC20Token.bin --abi=build/ERC20Token.abi --pkg=erc20 --out=gen/erc20/erc20.go
```

## 配置
配置文件：`config/staking.ini`

```
[url]
rpc_url = http://localhost:8545
ws_url = ws://127.0.0.1:8545

[eth]
private_key = 0x...
contract_address = 0x...
staking_token = 0x...
reward_token = 0x...
start_block = 0
confirmations = 1
interval = 2
```

## 运行
```bash
go run .
```

## API
Base: `http://localhost:8080/api`

### Staking
- `POST /stake`
  - form: `privateKeyStr`, `contractAddress`, `amount`
- `POST /withdrawStakedTokens`
  - form: `privateKeyStr`, `contractAddress`, `amount`
- `GET /getReward`
  - query: `privateKeyStr`, `contractAddress`
- `POST /updateRewardRate`
  - form: `privateKeyStr`, `contractAddress`, `newRewardRate`

只读查询：
- `GET /earned?contractAddress=...&account=...`
- `GET /stakedBalance?contractAddress=...&account=...`
- `GET /rewardPerToken?contractAddress=...`
- `GET /rewardPerTokenStored?contractAddress=...`
- `GET /rewardRate?contractAddress=...`
- `GET /lastUpdateTime?contractAddress=...`
- `GET /userRewardPerTokenPaid?contractAddress=...&account=...`
- `GET /rewards?contractAddress=...&account=...`

### ERC20
- `POST /approve`
  - form: `privateKeyStr`, `contractAddress`, `spenderAddress`, `value`
- `POST /transfer`
  - form: `privateKeyStr`, `contractAddress`, `to`, `value`
- `GET /balanceOf`
  - query: `contractAddress`, `to`
- `GET /allowance`
  - query: `contractAddress`, `ownerAddress`, `spenderAddress`

## 已做优化
- listener 回放循环改为 ticker，避免只执行一次
- 确认区块回放逻辑修正：按 `confirmations` 回退最新区块
- staking 与 ERC20 事件回放补齐并统一抽象（减少重复代码）
- 事件入库结构统一（包含 signature 字段）
- ERC20 Approve/Transfer/Allowance 逻辑修正
- 新增 staking 只读查询接口（earned、rewardRate 等）
- 引入 logrus 结构化日志
- service 层错误包装，日志可追踪上下文
