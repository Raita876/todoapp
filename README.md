# TODO Application

ポートフォリオとして公開しているタスク管理 API。主に以下のアーキテクチャおよびツールを使用。

- 設計
  - ドメイン駆動開発
  - Rest API

- 言語
  - [Go](https://go.dev/)

- ライブラリ
  - [gin](https://gin-gonic.com/)
  - [gorm](https://gorm.io/)
  - [urfave/cli](https://github.com/urfave/cli)

- データベース
  - [postgreSQL](https://www.postgresql.org/)
  - [mysql](https://www.mysql.com/jp/)

- テスト
  - [gotests](https://github.com/cweill/gotests)
  - [golangci-lint](https://golangci-lint.run/)
  - [runn](https://github.com/k1LoW/runn)
  
- ドキュメント
  - [swagger](https://swagger.io/)

- CI/CD
  - [github-actions](https://docs.github.com/ja/actions)
  - [goreleaser](https://goreleaser.com/)
  - [docker](https://www.docker.com/)

## How to Use

```bash
make up run # ローカルデバッグ用の API サーバーと DB 起動
make test # 単体テスト
make golangci # 静的解析
make scenariotest # シナリオテスト
make down stop # サーバーと DB の停止
```

## API Document

<http://localhost:8080/swagger/index.html>

## 参考

<https://github.com/sklinkert/go-ddd/tree/main>
