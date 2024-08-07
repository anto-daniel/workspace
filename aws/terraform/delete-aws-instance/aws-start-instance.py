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

EC2 = boto3.client('ec2', region_name='us-east-1',
        aws_access_key_id=access_key_id,
        aws_secret_access_key=secret_access_key)



def get_instances():
    instances = ec2.instances.all()
    for instance in instances:
        return instance.id

def start_instance(instance_id):
    response = EC2.start_instances(InstanceIds=[instance_id], DryRun=False)
    return response


if __name__ == '__main__':
    instance_name = get_instances()
    print(instance_name)
    resp = start_instance(instance_name)
    print(resp)

