# Emoji Gen Bot

Discordで送られてきたテキストから絵文字を生成して返すボット。

## 動かすには

1. `config/config.yml`を変更
    - `config_example.yml`を参考に
2. `docker-compose build`
    - (とても時間がかかる)
3. `docker-compose up -d`

## 開発環境

- devcontainer推奨 (CGOの警告が消せないので)

### libemoji

<https://github.com/emoji-gen/libemoji>

- こちらの裏側で使っているっぽい。
  - <https://emoji-gen.ninja/>

- (Forkしてパッチ当てている)
  - <https://github.com/oyakodon/libemoji>

- その他依存ライブラリ
  - zlib1g-dev (-lz)
  - libfontconfig-dev (-lfontconfig, -lfreetype)
  - libgl1-mesa-dev (-lGL)
