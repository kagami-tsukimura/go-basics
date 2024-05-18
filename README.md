# GO lang

## コマンド一覧

- バージョン確認

```bash
go version
```

- モジュール作成

```bash
go mod init <モジュール名>
```

- build

```bash
go build -o <バイナリ名> <ファイル名>
```

- build ファイルの実行

```bash
./<バイナリ名>
```

- staticcheck の適用

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```

- settings.json の追記

```json: settings.json
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.formatOnSave": true
  },
  "gopls": { "ui.diagnostic.staticcheck": true }
```

## 言語仕様

- 小文字の変数や関数: パッケージ内でのみ使用可能(Private)
- 大文字の変数や関数: 外部パッケージでも使用可能(Public)
