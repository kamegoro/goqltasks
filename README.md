## goqltasks

### ディレクトリ構成

```
goqltasks/
├── backend/
│   ├── api/
│   │   ├── graphql/
│   │   │   ├── schema.graphql
│   │   │   └── resolver.go
│   │   └── main.go
│   ├── Dockerfile
│   └── go.mod
└── frontend/
    ├── components/
    │   ├── TodoList.tsx
    │   └── TodoForm.tsx
    ├── pages/
    │   ├── index.tsx
    │   └── _app.tsx
    ├── styles/
    │   └── globals.css
    ├── Dockerfile
    ├── next.config.js
    ├── package.json
    └── tsconfig.json
```

- backend: Golang のバックエンドアプリケーション
  - Dockerfile: Golang バックエンドの Docker イメージをビルドするための設定ファイル
  - go.mod: Golang のモジュール定義ファイル
  - api
    - graphql: GraphQL スキーマとリゾルバの定義
    - main.go: Golang のメインアプリケーションファイル
- frontend: Next.js のフロントエンドアプリケーション
  - components: React コンポーネント
  - pages: Next.js のページコンポーネント
  - styles: グローバルな CSS スタイル
  - Dockerfile: Next.js フロントエンドの Docker イメージをビルドするための設定ファイル
  - next.config.js: Next.js の設定ファイル
  - package.json: プロジェクトの依存関係とスクリプト
  - tsconfig.json: TypeScript の設定ファイル
