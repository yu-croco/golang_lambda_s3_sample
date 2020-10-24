#!/bin/bash
# 複数環境に対してデプロイ出来るようにベースを用意

MSG="引数にはデプロイ対象の環境(develop)を指定してください"

if [ $# -ne 1 ]; then
  echo $MSG 1>&2
  exit 1
fi

environment=$1

if [ $environment = "develop" ]; then
  echo "ビルド開始"
  sh ./bin/build.sh
  echo "ビルド終了"
  echo $environment"にデプロイします"
  sls deploy --stage $environment
else
  echo $MSG
fi
