# Emoji Gen Bot

Discordで送られてきたテキストから絵文字を生成して返すボット。

## 動かすには

1. ボットを登録
    - <https://discord.com/developers/applications> でアプリケーション・ボットを作成
    - 下記権限を追加してボットをサーバーに招待
      - `Change Nickname`
      - `Manage Emojis and Stickers`
      - `Read Messages/View Channels`
      - `Send Messages`
      - (URL: `https://discord.com/api/oauth2/authorize?client_id=<client_id>&permissions=1140853760&scope=bot`)
2. `config/config.yml`を変更
    - `config_example.yml`を参考に
3. `docker-compose build`
    - (とても時間がかかる)
4. `docker-compose up -d`
    - air (Goのホットリロードツール)でアプリケーションが起動

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
