import requests


url="http://baidu.com"
resp=requests.get(url)
print(resp.text)

print()
print("测试结束")

