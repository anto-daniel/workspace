#!/usr/bin/env python3

import boto3
import json
import os

# Get aws access_key and secret_access_key from environment variables

# Get variables from environment variables
access_key_id = os.environ["AWS_ACCESS_KEY_ID"]
secret_access_key = os.environ["AWS_SECRET_ACCESS_KEY"]

# Specify aws region
ec2 = boto3.resource('ec2', region_name='us-east-1',
    aws_access_key_id=access_key_id,
    aws_secret_access_key=secret_access_key)

client = boto3.client('secretsmanager', region_name='us-east-1',
    aws_access_key_id=access_key_id,
    aws_secret_access_key=secret_access_key)

secret_name = "test_secret"
secret = client.get_secret_value(SecretId=secret_name)
secret_value = secret['SecretString']
print(secret_value)