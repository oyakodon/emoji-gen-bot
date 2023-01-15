# Emoji Gen Bot

Discordで送られてきたテキストから絵文字を生成して返すボット。

## 使い方

1. ボットが反応するチャンネルでコマンドを入力する
    - `/emoji gen name:zoi text:ぞい！ color:f39800`
2. 生成された絵文字を確認し、OKなら「登録」を押す
   - <img src="https://github.com/oyakodon/emoji-gen-bot/raw/images/assets/readme_cmd_gen.png" width="128px" />
3. 登録完了！
    - <img src="https://github.com/oyakodon/emoji-gen-bot/raw/images/assets/readme_generated.png" width="128px" />
    - 登録された絵文字はすぐ使えます。
    - また、絵文字生成通知チャンネルに生成された旨が通知されます。

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
|`align`|絵文字テキストの文字揃え|[left, center, right] デフォルト: left|
|`font`|絵文字の文字フォント|設定に依る, exampleと同じであれば [mp1pblack, notosansmono]|

## 開発者向け

- Docker環境必須
  - Windows 10(Intel x86_64, Docker Desktop)とmacOS Ventura 13.1(M1 arm64, lima + Docker) で動作確認済み
- VSCode + devcontainer推奨 (CGOの警告が消えないので)

### 動かすには

1. ボットを登録
    - <https://discord.com/developers/applications> でアプリケーション・ボットを作成
    - 下記権限を追加してボットをサーバーに招待
      - `Change Nickname`
      - `Manage Emojis and Stickers`
      - `Read Messages/View Channels`
      - `Send Messages`
      - (URL: `https://discord.com/api/oauth2/authorize?client_id=<client_id>&permissions=1140853760&scope=bot`)
2. `config/config.yml`を変更
    - `config_example.yml`を参考にトークンなど埋める
3. `docker-compose build`
    - (とても時間がかかる)
4. `docker-compose up -d`
    - air (Goのホットリロードツール)でアプリケーションが起動

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
