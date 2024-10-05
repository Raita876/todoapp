# TODO Application

ポートフォリオとして公開しているタスク管理 API。主に以下のアーキテクチャおよびツールを使用。

- 設計
  - ドメイン駆動開発
  - Rest API

- 言語
  - Go

- ライブラリ
  - gin
  - gorm

- データベース
  - postgreSQL

- テスト
  - gotests
  - golangci
  - runn
  
- ドキュメント
  - swagger

- CI/CD
  - goreleaser
  
## API

- <http://localhost:8080/ping>
- <http://localhost:8080/api/v1/tasks>
- <http://localhost:8080/api/v1/tasks/b81240b0-7122-4d06-bdb2-8bcf512d6c63>
- <http://localhost:8080/swagger/index.html>

## Command Example

```bash
curl \
-X POST \
-H "Content-Type: application/json" \
-d '{"name":"sample　task", "description":"This is sample task", "status_id": 9}' \
http://localhost:8080/api/v1/tasks
```

```bash
curl \
-X PUT \
-H "Content-Type: application/json" \
-d '{"id": "fad796a1-e0ed-4ee5-9f88-9b7258d35ae9", "name":"updated task", "description":"This is updated task", "status_id": 5}' \
http://localhost:8080/api/v1/tasks
```

```bash
curl \
-X DELETE \
-H "Content-Type: application/json" \
-d '{"id": "07aaadbc-8967-406f-aebd-58b289377aef"}' \
http://localhost:8080/api/v1/tasks
```

## 参考

<https://github.com/sklinkert/go-ddd/tree/main>
