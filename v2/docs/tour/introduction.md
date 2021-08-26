# Tour/Introduction
Fyne 公式ドキュメントの [Tour/Introduction](https://developer.fyne.io/tour/introduction/) の一部を抜粋し、翻訳したものです。

## GUI Overview
Origin: https://developer.fyne.io/tour/introduction/guis.html

グラフィカルなアプリケーションは、ウェブベースやコマンドラインのアプリケーションに比べて作成が複雑になることがよくあります。Fyneは、Goの優れたデザインを利用して、美しいグラフィカルなアプリケーションを簡単かつ迅速に作成することで、この状況を変えます。

グラフィカルアプリケーションを動作させるには、ウィンドウを作成し、アプリケーションを実行するように指示する必要があります。そうすることで、ユーザーの入力に反応するイベント処理コードが開始され、コードの実行に合わせて画面が更新されます。

この例では、新しいアプリケーションを作成し、「Hello」というタイトルのウィンドウを1つ作成します。このウィンドウの中に、"Hello Fyne!"というテキストを含むラベルを1つ配置します。

ウィンドウの内容が設定されたら、ウィンドウを表示してアプリケーションを実行します（`Window.ShowAndRun()`は、`Window.Show()` && `App.Run()`のショートカットです）。`Run()`または`ShowAndRun()`を呼び出すと、アプリケーションが実行され、ウィンドウが閉じられた後に関数が返されます。

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello Fyne!"))
	myWindow.ShowAndRun()
}
```

## Organisation and Packages
Origin: https://developer.fyne.io/tour/introduction/organisation.html

Fyneプロジェクトは多くのパッケージに分かれており、それぞれが異なるタイプの機能を提供しています。それらは以下の通りです。

| package                           | desc                                                                                                                                                                            |
|-----------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `fyne.io/fyne/v2`                 | このインポートは、すべてのFyneコードに共通する基本的な定義を提供します。データタイプやインターフェースを含む、すべてのFyneコードに共通する基本的な定義を提供します。            |
| `fyne.io/fyne/v2/app`             | `app` パッケージは、新しいアプリケーションを開始するためのAPIを提供します。通常は`app.New()`または`app.NewWithID()`のみを必要とします。                                         |
| `fyne.io/fyne/v2/canvas`          | `canvas` パッケージは、Fyne内のすべての drawing API を提供します。Fyneのツールキットは、これらのプリミティブなグラフィカルタイプ上に構築されています。                          |
| `fyne.io/fyne/v2/container`       | `container` パッケージは、アプリケーションをレイアウトして整理するために使用される コンテナ を提供します。                                                                      |
| `fyne.io/fyne/v2/data/binding`    | `binding` パッケージは、データソースを Widget にバインドする方法を提供します。                                                                                                  |
| `fyne.io/fyne/v2/data/validation` | `validation` パッケージは、Widget 内のデータを検証するための ツール を提供します。                                                                                              |
| `fyne.io/fyne/v2/dialog`          | `dialog` パッケージは、 Confirm, Error, File-Open / File-Save などの ダイアログ を含みます。                                                                                    |
| `fyne.io/fyne/v2/layout`          | `layout` パッケージは、コンテナで使用するための様々なレイアウト実装を提供します。コンテナで使用するための様々なレイアウトの実装を提供します（後のチュートリアルで説明します）。 |
| `fyne.io/fyne/v2/storage`         | `storage` パッケージは、ストレージへのアクセスと管理機能を提供します。                                                                                                          |
| `fyne.io/fyne/v2/test`            | `test` パッケージ内のツールを使用して、アプリケーションをより簡単にテストすることができます。
| `fyne.io/fyne/v2/widget`          | ほとんどの グラフィカルアプリケーション は、複数のウィジェットを使用して作成されます。Fyneのすべてのウィジェットと動的な要素はこのパッケージに含まれています。                  |

## Packaging and Distribution
Origin: https://developer.fyne.io/tour/introduction/packaging.html

複数のOSに対応するためのパッケージ化は、複雑な作業となります。グラフィカルなアプリケーションには、アイコンやメタデータが関連付けられているだけでなく、各環境との統合に必要な特定のフォーマットが存在します。

`fyne` コマンドは、ツールキットがサポートするすべてのプラットフォームに配布するアプリケーションの準備をサポートします。`fyne package` を実行すると、自分のコンピュータにインストールできるアプリケーションが作成され、作成されたファイルをカレントディレクトリからコピーするだけで、他のコンピュータに配布できるようになります。

Windowsでは、アイコンが埋め込まれた `.exe` ファイルが作成されます。macOSの場合は `.app` bundle が、Linuxの場合は `.tar.xz` ファイルが作成され、通常の方法でインストールすることができます（または、解凍されたフォルダ内で make install を実行することでインストールできます）。

もちろん、標準の Go ツールチェーン を使ってアプリケーションを実行することもできます。

```shell
#!/bin/sh

go get fyne.io/fyne/v2/cmd/fyne

go build
fyne package -icon mylogo.png

# result is a platform specific package
# for the current operating system.
```

## Beginner to Expert
Origin: https://developer.fyne.io/tour/introduction/learning.html

Fyneは、簡単に使い始めることができ、複数のプラットフォームにまたがる大規模なアプリケーションをシンプルに構築できるように設計されています。また、カスタム要素の追加や、コントリビューションが可能なように設計されています。

このチュートリアルでは、徐々に複雑になっていくトピックを紹介していますが、熟練したプログラマーを超えるものではありません。各ステップでは、例題をIDEにコピーして、実際に動作させることができます。

このツアーを終える頃には、Fyneとそのツールの構成要素のすべてを知ることができるでしょう。あなたが何を作るのか、私たちは待ちきれません。

もし、あなたがお返しをしたいと思ったら、私たちは貢献、バグレポート、ツールキットを使っている人との会話を歓迎します。コントリビューションの詳細については、[コントリビューターページ](https://fyne.io/contribute.html) または [githubリポジトリ](https://github.com/fyne-io/fyne/) をご覧ください。

最初のチュートリアル [Basics](https://developer.fyne.io/tour/basics/) に進んでください。

- - -
以下余白
