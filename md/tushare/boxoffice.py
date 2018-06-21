#!/usr/bin/env pytho
# -*- coding:utf-8 -*-
'''
票房接口封装
'''

import tushare as ts

def get_realtime_boxoffice():
	return ts.realtime_boxoffice()

if __name__ == '__main__':
	df=get_realtime_boxoffice();
	print(df)