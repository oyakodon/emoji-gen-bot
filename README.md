# Emoji Gen Bot

Discordで送られてきたテキストから絵文字を生成して返すボット。

## コマンド一覧

- Discord上で使うときは、補完が効くのでコマンドを覚えていなくても大丈夫。

`/emoji <サブコマンド> [オプション]`

|サブコマンド|機能|
|:-|:-|
|`gen`|絵文字を生成し、サーバーに登録する|
|`preview`|絵文字の生成のみ行う (絵文字の確認用)|

|オプション|説明|備考|
|:-|:-|:-|
|`name`|登録する絵文字の名前|(`gen`のみ必須)|
|`text`|絵文字のテキスト|必須|
|`color`|絵文字の文字色 (カラーコード, ex: FFFFFF)|デフォルト: 000000 (黒)|
|`align`|絵文字テキストの文字揃え|[center, left, right] デフォルト: left|
|`font`|絵文字の文字フォント|設定に依る, exampleと同じであれば [mp1pblack, notosansmono]|

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
