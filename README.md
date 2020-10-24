# Golang Lambda S3 Sample

Golang on AWS LambdaでS3へ画像をアップロードするサンプルコードです。ご自由にお使いください。

## 環境
- Golang: v1.15.2
- serverless framework 
    - Plugin: v4.1.1
    - SDK: v2.3.2
    - Components: v3.2.5

## セットアップ
- `bin/download.sh`を実行する

## deploy
- `bin/deploy.sh` を実行する

## リクエスト
- action: POST
- body
    - image: string(Base64形式の画像データ)
    - userId: int
