# basics

## Usage
`go run` コマンドでアプリケーションを起動:

    go run main.go

## Box Layout
https://developer.fyne.io/tour/layout/

[Container and Layouts](https://developer.fyne.io/tour/basics/container.html) で議論した通り、
コンテナ内部の要素は layout を使って配置を指定することができます。このセクションでは Fyne に
同梱されているレイアウトとその使い方について説明します。

もっとも一般的なレイアウトは`layout.BoxLayout` で、これは _horizontal_ と _vertical_ の2つの変数をもっています。 
ボックスレイアウトは、すべての要素を1つの行または列に配置し、オプションでスペースを入れて配置を補助します。


