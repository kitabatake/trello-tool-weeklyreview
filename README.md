trelloでタスク管理をしているユーザー向けの、1週間の作業の振り返り用の情報（markdown）を生成するツールです。

# 前提とするtrelloの運用

- 終了したタスク（カード）は `done` リストへ移動する
- 毎朝、前日に `done` に入れたタスクを確認して、アーカイブする。

# 使い方

- 環境変数にtrelloの認証情報を入れる（ [こちら](https://trello.com/app-key)  から取得）
    - TRELLO_API_KEY: APIキー
    - TRELLO_TOKEN: Token
 
- コマンドを実行すると一週間の作業内容がmarkdownとして吐き出される
- エディタ等で開いて振り返り内容を書いたりする