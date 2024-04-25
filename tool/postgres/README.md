# 注意
危険な変更はちゃんと弾かれる。例えば特定のカラムの型を VARCHAR から UUID に変更するようなパターン。
```
>> ALTER TABLE "public"."users" ALTER COLUMN "id" TYPE uuid;

2024/03/06 06:29:07 pq: column "id" cannot be cast automatically to type uuid
```

このケースでは以下のように `USING` 句で明示的にキャストする ALTER を手動で実行すればよい。既にレコードが存在する場合はデータの段階移行作業も必須。

```SQL
ALTER TABLE "public"."users" ALTER COLUMN "id" TYPE uuid USING (id::uuid);
```
