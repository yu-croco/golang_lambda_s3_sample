# Golang Lambda S3 Sample

Golang on AWS LambdaでS3へ画像をアップロードするサンプルコードです。ご自由にお使いください。


## セットアップ
- `bin/download.sh`を実行する

## deploy
- `bin/deploy.sh` を実行する

## リクエスト
- action: POST
- body
    - image: string(Base64形式の画像データ)
    - userId: int
