#!/bin/bash

mkdir -p $HOME/.aws

create_config_file() {
  cat >> $HOME/.aws/config<<EOM
[profile personal]
region=eu-north-1
output=json
EOM
}
create_cred_file() {
  cat >> $HOME/.aws/credentials<<EOM
[personal]
aws_access_key_id=$(op item get aws --fields ACCESS_KEY)
aws_secret_access_key=$(op item get aws --fields SECRET_ACCESS_KEY)
EOM
}
pushd $HOME/.aws
  create_config_file
  create_cred_file
popd
