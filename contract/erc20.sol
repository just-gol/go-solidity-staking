// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ierc20.sol";

contract ERC20Token is IERC20 {
    string public name;
    string public symbol;
    uint8 public immutable decimals = 18;

    uint256 private _totalSupply;
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;

    constructor(string memory name_, string memory symbol_, uint256 initialSupply) {
        name = name_;
        symbol = symbol_;
        _mint(msg.sender, initialSupply);
    }

    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) external view override returns (uint256) {
        return _balances[account];
    }

    function transfer(address to, uint256 value) external override returns (bool) {
        _transfer(msg.sender, to, value);
        return true;
    }

    function allowance(address owner, address spender) external view override returns (uint256) {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 value) external override returns (bool) {
        _approve(msg.sender, spender, value);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) external override returns (bool) {
        uint256 currentAllowance = _allowances[from][msg.sender];
        require(currentAllowance >= value, "ERC20: insufficient allowance");
        _approve(from, msg.sender, currentAllowance - value);
        _transfer(from, to, value);
        return true;
    }

    function _transfer(address from, address to, uint256 value) internal {
        require(to != address(0), "ERC20: transfer to zero address");
        uint256 fromBalance = _balances[from];
        require(fromBalance >= value, "ERC20: transfer amount exceeds balance");
        _balances[from] = fromBalance - value;
        _balances[to] += value;
        emit Transfer(from, to, value);
    }

    function _approve(address owner, address spender, uint256 value) internal {
        require(spender != address(0), "ERC20: approve to zero address");
        require(owner != address(0), "ERC20: approve from zero address");
        _allowances[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    function _mint(address to, uint256 value) internal {
        require(to != address(0), "ERC20: mint to zero address");
        _totalSupply += value;
        _balances[to] += value;
        emit Transfer(address(0), to, value);
    }
}
