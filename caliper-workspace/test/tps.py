import os
import re


def change_tps(file):
    # file = "../benchmarks/myAssetBenchmark.yaml"
    f = open(file, "r")

    ff = open("%s.bak" % file, "w")
    for line in f:
        if "tps" in line:
            num = re.findall(r"[1-9]\d*$", line)
            a = int(num[0])
            a += 50
            line = "            tps: " + str(a) + "\n"
        ff.write(line)
    os.remove(file)
    os.rename("%s.bak" % file, file)


file1 = "../benchmarks/policy.yaml"
file2 = "../benchmarks/data.yaml"
file3 = "../benchmarks/record.yaml"

change_tps(file1)
change_tps(file2)
change_tps(file3)
