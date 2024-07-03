#!/usr/bin/env python3

## Generate a script to get list of all aws instances

import boto3
import os

# Get aws access_key and secret_access_key from environment variables

# Get variables from environment variables
access_key_id = os.environ["AWS_ACCESS_KEY_ID"]
secret_access_key = os.environ["AWS_SECRET_ACCESS_KEY"]

# Specify aws region
ec2 = boto3.resource('ec2', region_name='us-east-1',
        aws_access_key_id=access_key_id,
        aws_secret_access_key=secret_access_key)

def get_instances():
    instances = ec2.instances.all()
    for instance in instances:
        print(instance.id, instance.state, instance.tags)


if __name__ == '__main__':
    get_instances()
