
for i in $(seq 50 50 200):
do

  for j in $(seq 50 50 1000):
  do
    cd ../myfabric
    bash start.sh
    cd ../caliper-workspace
    npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/networkConfig.yaml --caliper-benchconfig benchmarks/policy.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled > ./log/policy/${i}-${j}.log
    sleep 2
    npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/networkConfig.yaml --caliper-benchconfig benchmarks/data.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled > ./log/data/${i}-${j}.log
    sleep 2
    npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/networkConfig.yaml --caliper-benchconfig benchmarks/record.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled > ./log/record/${i}-${j}.log
    cd test
    python3 tps.py
    cd ..
    sleep 5
  done
  cd test
  python3 main.py
  cd ..
  sleep 5
done
