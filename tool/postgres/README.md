# メモ
注意：sqldef の自動マイグレーション実行に失敗するケースが存在する。例えば、特定のカラムの型を VARCHAR から UUID に変更するパターンなど。
```
>> ALTER TABLE "public"."users" ALTER COLUMN "id" TYPE uuid;

2024/03/06 06:29:07 pq: column "id" cannot be cast automatically to type uuid
```

この場合、手作業で操作するしかない。

例えば、以下のように `USING` 句によって明示的にキャストした変更クエリを手動でDBへ流す。
```SQL
ALTER TABLE "public"."users" ALTER COLUMN "id" TYPE uuid USING (id::uuid);
```

ただし、事前に全データのカラムがキャスト後のフォーマット（ここでは UUID）に準拠していることを確認する必要がある。

本番データに対して実施する際は、例えば段階的な移行プロセスを打ち立てて、より慎重に実施したほうが良い。