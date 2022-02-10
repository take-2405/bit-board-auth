# bit-board-auth
## 実行方法(随時変更)
- go の開発環境が整っていることを前提とする

**事前準備**
- firebase プロジェクトを作成し、Authenticationの「method」で認証方法を設定し、秘密鍵を生成する
- 秘密鍵をenv fileか環境変数とする

**makeコマンドが使用できる場合**  
```cassandraql
make run
```

**makeコマンドが使用できない場合**  
```cassandraql
go run cmd/main.go
```

## 基本概要
#### リポジトリ概要
本リポジトリでは、firebaseを利用した認証を実現する

#### 開発フロー
- git-flow

#### 使用AFW
- chi

#### use package
- go-cache
- firebase.google.com/go
- github.com/pkg/errors

#### CICD
使用CI：github action
- lintチェック
- import packageの脆弱性チェック
- herokuへのデプロイ

#### 使用liter(静的解析ツール)
- golangci-lint  
  go lint は非推奨のため未使用

#### ディレクトリ構成
```
./
├── presentation
│   ├── controller
│   ├── request
│   ├── response
│   └── router
├── usecase
├── cmd
├── config
├── di
├── domain
└── infrastructure
```

#### 製作者
take-2405(Gitアカウント名)

### 参考情報
- chi 公式ドキュメント  
  https://github.com/go-chi/chi
