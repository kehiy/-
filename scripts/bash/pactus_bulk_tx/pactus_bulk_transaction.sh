#!/bin/bash

# Checking arguments
if [ $# -lt 4 ]; then
    echo "Usage: $0 <address> <min_amount> <max_amount> <file> <wallet_path>"
    exit 1
fi

# Get the addresses string and amount from the command-line arguments
address=$1
min=$2
max=$3
file=$4
wallet_path=$5

# Read all address
my_array=()
while IFS= read -r line; do
    my_array+=("$line")
done < $file

# Loop over all address
for addr in "${my_array[@]}"; do
    
    # Generating amount randomly
    random_number=$(( $(( RANDOM % ($max - $min + 1) )) + $min ))
    
    # Creating command string
    command="pactus-wallet tx transfer $address $addr $random_number --path=$wallet_path"

    echo "Going to send $random_number PAC to $addr"

    # Run the command and automatically input 'y' followed by enter
    echo -e "y\n" | $command

      # Check the exit status of the command
    if [ $? -ne 0 ]; then
        echo "Error transferring $amount to address: $address"
    else
        echo "Transferred $amount to address: $address"
    fi
done

echo "Bulk transfer completed."
