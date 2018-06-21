#!/usr/bin/env pytho
# -*- coding:utf-8 -*-
'''
修改tushare源码 增加user_agent 解决403问题
目前美国hb,ok 美国禁止提供服务
数字货币封装
'''
import tushare as ts

def coins_tick(borker='ok',code='btc'):
	return ts.coins_tick(borker,code)

if __name__ == '__main__':
	ct=coins_tick('chbtc','btc')
	print(ct)
