# 这是一个示例 Python 脚本。

# 按 ⌃R 执行或将其替换为您的代码。
# 按 双击 ⇧ 在所有地方搜索类、文件、工具窗口、操作和设置。
import os
import re


def change_tx():
    file = "../../myfabric/configtx.yaml"
    f = open(file, "r")

    ff = open("%s.bak" % file, "w")
    for line in f:
        if "MaxMessageCount" in line:
            num = re.findall(r"[1-9]\d*$", line)
            a = int(num[0])
            a += 50
            line = "        MaxMessageCount: " + str(a) + "\n"
        ff.write(line)
    os.remove(file)
    os.rename("%s.bak" % file, file)


def change_tps(file):
    a = 50
    # b = 1000
    f = open(file, "r")

    ff = open("%s.bak" % file, "w")
    for line in f:
        if "tps" in line:
            line = "            tps: " + str(a) + "\n"
        # elif "txNumber" in line:
        #     line = "        txNumber: " + str(b) + "\n"
        ff.write(line)
    os.remove(file)
    os.rename("%s.bak" % file, file)


file1 = "../benchmarks/policy.yaml"
file2 = "../benchmarks/data.yaml"
file3 = "../benchmarks/record.yaml"

change_tps(file1)
change_tps(file2)
change_tps(file3)
change_tx()
