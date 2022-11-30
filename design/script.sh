#!/bin/bash
echo "Cluster testing script"
sleep 1

# Create a bunch of data
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Hello", "value": [119, 111, 114, 108, 100]}' # 1 world
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Data1", "value": [72, 111, 117, 115, 101]}' # 2 House
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Animal", "value": [67, 111, 119]}' # 3 Cow
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Data3", "value": [79, 98, 106, 101, 99, 116, 32, 111, 114, 105, 101, 110, 116, 101, 100, 32, 112, 114, 111, 103, 114, 97, 109, 109, 105, 110, 103]}' # 4 Object oriented programming
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Data4", "value": [83, 117, 110]}' # 5 Sun
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Data5", "value": [82, 111, 115, 101, 115]}' # 6 Roses
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Data6", "value": [83, 117, 114, 118, 105, 118, 111, 114, 32, 98, 105, 97, 115]}' # 7 Survivor bias
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Science", "value": [78, 101, 119, 116, 111, 110]}' # 8 Newton
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"University", "value": [85, 84, 77]}' # 9 UTM
curl http://localhost:8081/create -X POST -H "Content-Type: application/json" -d '{"key":"Data9", "value": [79, 115, 109, 111, 115, 105, 115]}' # 10 Osmosis
#sleep 3

# Update some data
curl http://localhost:8081/update -X PUT -H "Content-Type: application/json" -d '{"key":"Animal", "value": [72, 111, 114, 115, 101]}'      # 1 Horse
curl http://localhost:8081/update -X PUT -H "Content-Type: application/json" -d '{"key":"Science", "value": [68, 97, 114, 119, 105, 110]}' # 2 Darwin
curl http://localhost:8081/update -X PUT -H "Content-Type: application/json" -d '{"key":"University", "value": [85, 83, 77]}'              # 3 USM
#sleep 3

# Delete some data
curl http://localhost:8081/delete -X DELETE -H "Key: Data1" # 1
curl http://localhost:8081/delete -X DELETE -H "Key: Data3" # 2
curl http://localhost:8081/delete -X DELETE -H "Key: Data4" # 3
#sleep 3

# Kill a server
# echo "Kill Server"
# curl http://localhost:8083/kill
# sleep 1

# Get a bunch of data
printf "Fetch data:"
curl http://localhost:8081/read -H "Key: Hello" && printf "\n"      # 1
curl http://localhost:8081/read -H "Key: Animal" && printf "\n"     # 2
curl http://localhost:8081/read -H "Key: Science" && printf "\n"    # 3
curl http://localhost:8081/read -H "Key: University" && printf "\n" # 4
curl http://localhost:8081/read -H "Key: Data5" && printf "\n"      # 5
curl http://localhost:8081/read -H "Key: Data6" && printf "\n"      # 6
curl http://localhost:8081/read -H "Key: Data9" && printf "\n"      # 7