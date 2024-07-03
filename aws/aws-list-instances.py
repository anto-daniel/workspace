#!/usr/bin/env python3

## Generate a script to get list of all aws instances

import boto3
import os


# Specify aws region
ec2 = boto3.resource('ec2', region_name='us-east-1')

def get_instances():
    instances = ec2.instances.all()
    for instance in instances:
        print(instance.id, instance.state, instance.tags)


if __name__ == '__main__':
    get_instances()
