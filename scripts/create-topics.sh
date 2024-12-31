#!/bin/sh

# Wait for Kafka to be ready
echo "Waiting for Kafka to be ready..."
kafka-topics --bootstrap-server kafka:29092 --list > /dev/null 2>&1
while [ $? -ne 0 ]; do
  sleep 1
  kafka-topics --bootstrap-server kafka:29092 --list > /dev/null 2>&1
done

# Create topics
echo "Creating Kafka topics..."

# Add your topics here
kafka-topics --bootstrap-server kafka:29092 --create --if-not-exists --topic topic1 --partitions 1 --replication-factor 1
kafka-topics --bootstrap-server kafka:29092 --create --if-not-exists --topic topic2 --partitions 1 --replication-factor 1

echo "Topics created successfully!"
kafka-topics --bootstrap-server kafka:29092 --list
