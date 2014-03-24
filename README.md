#goyish-typetalk
[Typetalk](https://typetalk.in/ "Typetalk") API の Go 言語ラッパーライブラリです。

##認証
###アクセストークンの取得
    auth := typetalk.NewAuthClient(<YOUR_CLIENT_ID>, <YOUR_CLIENT_SECRET>, typetalk.My, typetalk.TopicRead, typetalk.TopicPost)
    err := auth.GetAccessToken()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(auth.AccessToken)
    fmt.Println(auth.RefreshToken)
    client := typetalk.NewClient(auth)
    // client を使用してAPIアクセス

(スコープはtypetalk.My, typetalk.TopicRead, typetalk.TopicPostから複数指定可)

###アクセストークンの更新
    auth := typetalk.NewAuthClient(<YOUR_CLIENT_ID>, <YOUR_CLIENT_SECRET>)
    auth.RefreshToken = <YOUR_REFRESH_TOKEN>
    err := auth.UpdateToken()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(auth.AccessToken)
    client := typetalk.NewClient(auth)
    // client を使用してAPIアクセス

###事前に取得済みのアクセストークンを設定
    auth := typetalk.NewAuthClient(<YOUR_CLIENT_ID>, <YOUR_CLIENT_SECRET>)
    auth.AccessToken = <YOUR_ACCESS_TOKEN>
    client := typetalk.NewClient(auth)
    // client を使用してAPIアクセス

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

その他、[API リファレンス](http://developers.typetalk.in/api_ja.html "Typetalk API リファレンス")を参照下さい。

##今後の対応
* レスポンスに含まれる `pendings`, `talks` に対応
* 単体テスト作成
