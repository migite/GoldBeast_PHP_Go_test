# Go + PHP Integration Demo

フロントエンド（Go）とバックエンド（PHP）を連携させたデモアプリケーションです。

## プロジェクト構成

```
GoldBeast_PHP_Go_test/
├── go_frontend/          # Goフロントエンドサーバー
│   ├── main.go
│   └── go.mod
├── php_backend/          # PHPバックエンド API
│   └── api.php
└── README.md
```

## アーキテクチャ

```
クライアント（ブラウザ）
    ↓
Go（ポート 8000）- フロントエンド
    ↓
PHP（ポート 8080）- バックエンド API
    ↓
データベース
```

## インストール＆実行

### PHPバックエンド（API）

```bash
cd php_backend
php -S localhost:8080
```

### Goフロントエンド

```bash
cd go_frontend
go run main.go
```

ブラウザで `http://localhost:8000` にアクセス

## 機能

- ✅ ユーザー一覧表示
- ✅ ユーザー追加
- ✅ 商品一覧表示
- ✅ Go ↔ PHP 連携

## 技術スタック

- **フロントエンド**: Go（net/http）
- **バックエンド**: PHP
- **通信**: JSON形式のHTTP
- **CORSサポート**: 有効
