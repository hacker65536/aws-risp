# AWS-RISP

AWS Reserved Instances and Savings Plans - AWSリザーブドインスタンス、セービングプランの情報を取得・表示するCLIツール

## 概要

AWS-RISPは、AWSのリザーブドインスタンス（RI）に関する情報を簡単に取得し、表示するためのコマンドラインツールです。主に以下の機能を提供します：

- リザーベーションのカバレッジ情報の取得と表示
- リザーベーションの利用率情報の取得と表示
- 複数のAWSサービス（EC2、RDS、ElastiCache、OpenSearch、Redshiftなど）に対応

## インストール

```bash
go install github.com/hacker65536/aws-risp@latest
```

または、ソースからビルドする場合：

```bash
git clone https://github.com/hacker65536/aws-risp.git
cd aws-risp
make build
```

makeを使用せずにビルドする場合：

```bash
git clone https://github.com/hacker65536/aws-risp.git
cd aws-risp
go build
```

## 使用方法

### リザーベーションカバレッジの確認

特定のサービスのリザーベーションカバレッジを確認するには：

```bash
aws-risp rsvCov ec2 rds
```

すべてのサポートされているサービスのカバレッジを確認するには：

```bash
aws-risp rsvCov ec2 rds elasticache opensearch memorydb redshift elasticsearch
```

### リザーベーション使用率の確認

```bash
aws-risp rsvUtil
```

### オプション

- `--start`: 期間の開始日（ISO 8601形式）
- `--end`: 期間の終了日（ISO 8601形式）
- `--sort`: ソート基準（例：`OnDemandCost`）

## サポートされているサービス

- Amazon EC2 - Compute (`ec2`)
- Amazon RDS (`rds`)
- Amazon ElastiCache (`elasticache`)
- Amazon OpenSearch Service (`opensearch`)
- Amazon MemoryDB (`memorydb`)
- Amazon Redshift (`redshift`)
- Amazon Elasticsearch Service (`elasticsearch`)

## ライセンス

[LICENSE](LICENSE) ファイルを参照してください。

## 開発

### Makefileターゲット

このプロジェクトでは、様々な開発・デプロイメントタスクを自動化するためのMakefileが用意されています：

```
make build      # プログラムをビルドする
make test       # テストを実行する
make cover      # カバレッジレポートを生成・表示する
make lint       # コードの静的解析を実行する
make fmt        # コードフォーマットを実行する
make clean      # ビルド成果物をクリーンアップする
make release    # リリースビルドを作成する
make doc        # ドキュメントを生成する
make version    # バージョン情報を表示する
make help       # Makefileのヘルプを表示する
```

### バージョン情報

バージョン情報は以下のコマンドで確認できます：

```bash
aws-risp --version
aws-risp version
```