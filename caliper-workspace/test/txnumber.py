import os

file = "../benchmarks/myAssetBenchmark.yaml"
f = open(file, "r")

ff = open("%s.bak" % file, "w")

b = 10000
for line in f:
    if "txNumber" in line:
        line = "        txNumber: " + str(b) + "\n"
    ff.write(line)
os.remove(file)
os.rename("%s.bak" % file, file)