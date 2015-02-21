#goyish-typetalk
[Typetalk](https://typetalk.in/ "Typetalk") API の Go 言語ラッパーライブラリです。

[![Build Status](https://api.travis-ci.org/zaneli/goyish-typetalk.png?branch=master)](https://travis-ci.org/zaneli/goyish-typetalk)

##インストール
    go get github.com/zaneli/goyish-typetalk/typetalk

##認証
###アクセストークンの取得
    import (
        "fmt"
        "github.com/zaneli/goyish-typetalk/typetalk"
        "log"
    )

    client := typetalk.NewClient()
    auth, err := client.GetAccessToken(<YOUR_CLIENT_ID>, <YOUR_CLIENT_SECRET>, typetalk.My, typetalk.TopicRead, typetalk.TopicPost)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(auth.AccessToken)
    fmt.Println(auth.RefreshToken)

    // client を使用してAPIアクセス
    client.GetMyProfile()

(スコープはtypetalk.My, typetalk.TopicRead, typetalk.TopicPostから複数指定可)

###アクセストークンの更新
    client := typetalk.NewClient()
    auth, err := client.UpdateAccessToken(<YOUR_CLIENT_ID>, <YOUR_CLIENT_SECRET>, <RefreshToken>)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(auth.AccessToken)

    // client を使用してAPIアクセス
    client.GetMyProfile()

###事前に取得済みのアクセストークンを設定
    auth := typetalk.AuthedClient(<AccessToken>)

    // client を使用してAPIアクセス
    client.GetMyProfile()

##APIの実行
###プロフィールの取得
    res, err := client.GetMyProfile()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(res)

###投稿メッセージリストの取得
    res, err := client.GetTopicMessages(<TOPIC_ID>)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(res)
または

    res, err := client.GetTopicMessagesApi(<TOPIC_ID>).Count(<COUNT_NUM>).From(<POST_ID>).Forward().Call()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(res)

###メッセージの投稿
    res, err := client.PostMessage(<MESSAGE>, <TOPIC_ID>)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(res)
または

    res, err := client.PostMessageApi(<MESSAGE>, <TOPIC_ID>).ReplyTo(<POST_ID>).FileKeys(<FILE_KEY0>, <FILE_KEY1>, ...).Call()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(res)

その他、[API リファレンス](https://developer.nulab-inc.com/ja/docs/typetalk/ "Typetalk API リファレンス")を参照下さい。

