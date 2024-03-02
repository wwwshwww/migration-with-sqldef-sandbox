-- sqldef の差分出力機能の動作確認のためだけに存在するダミーのスキーマ定義

CREATE TABLE users (
    id bigint NOT NULL PRIMARY KEY,
    name text,
    age int
);
CREATE TABLE bigdata (
    data bigint
);