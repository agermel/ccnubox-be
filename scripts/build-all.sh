#!/usr/bin/env bash

# 这个脚本主要用于打包镜像

set -e  # 一旦有命令失败就退出

# 捕获 SIGINT 信号（Ctrl+C）并退出
trap 'echo "Script interrupted."; exit 1' SIGINT

# shellcheck disable=SC2034
ds=(
  "be-banner"
  "be-calendar"
  "be-ccnu"
  "be-class"
  "be-classlist"
  "be-counter"
  "be-department"
  "be-elecprice"
  "be-feed"
  "be-grade"
  "be-library"
  "be-infosum"
  "be-website"
  "bff"
)

imageRepo=$1

CRYPTO_KEY=$2

## 这里注意需要自己配置一个加密key，否则会使用默认的key，存在安全隐患
## 本地调试可忽略
if [[ -n "$CRYPTO_KEY" ]]; then
  CRYPTO_KEY="muxiStudio123456"
fi

for d in "${ds[@]}"; do
  echo -e "🔧🔧🔧 Building and pushing image for $d 🔧🔧🔧\n"

  # shellcheck disable=SC2046
  docker build -t "$d:v1" -f "./$d/Dockerfile" .

  if [[ -n "$imageRepo" ]]; then
    echo -e "📦 Tagging and pushing $d to $imageRepo ...\n"
    docker tag "$d:v1" "$imageRepo/$d:v1"
    docker push "$imageRepo/$d:v1"
  else
    echo -e "No imageRepo provided, skipping tag & push for $d  \n"
  fi

done


speciald="be-user"

echo -e "🔧🔧🔧 Building and pushing image for $speciald 🔧🔧🔧\n"

docker build -t "$speciald:v1" -f "./$speciald/Dockerfile" --build-arg KEY="$CRYPTO_KEY"  .


if [[ -n "$imageRepo" ]]; then
    echo -e "📦 Tagging and pushing $speciald to $imageRepo ... \n"
    docker tag "$speciald:v1" "$imageRepo/$speciald:v1"
    docker push "$imageRepo/$speciald:v1"
else
    echo -e "No imageRepo provided, skipping tag & push for $speciald   \n"
fi