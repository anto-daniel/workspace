#!/usr/bin/python3

# Python program to find instance id from IP address using AWS SDK
import boto3
import os
import sys

if len(sys.argv) != 2:
    print('Usage: python find_instance_id_from_ip.py <ip address>')
    sys.exit(1)

# Get the IP address from the command line
ip_address = sys.argv[1]
os.environ['AWS_REGION'] = 'us-east-1'

# Get the aws secret tokens from env
aws_access_key_id = os.environ['AWS_ACCESS_KEY_ID']
aws_secret_access_key = os.environ['AWS_SECRET_ACCESS_KEY']
aws_session_token = os.environ['AWS_SESSION_TOKEN']

# Create a session using your credentials
session = boto3.Session(
    region_name = "us-east-1",
    aws_access_key_id=aws_access_key_id,
    aws_secret_access_key=aws_secret_access_key,
    aws_session_token=aws_session_token
)

# Create a client using the session
client = session.client('ec2')

# Get the instance id from the IP address
response = client.describe_instances(
    Filters=[
        {
            'Name': 'private-ip-address',
            'Values': [
                ip_address
            ]
        },
    ]
)

# Print the instance id
print(response['Reservations'][0]['Instances'][0]['InstanceId'])
print(response['Reservations'][0]['Instances'][0]['InstanceType'])
