# basics

## Usage
`go run` コマンドでアプリケーションを起動:

    go run main.go

## アプリケーションとランループ について
https://developer.fyne.io/tour/basics

GUIアプリケーションが動作するためには、ユーザーの操作や描画イベントを処理するイベントループ（もしくは ランループ と呼ばれる）を実行する必要があります。Fyneでは、`App.Run()` または `Window.ShowAndRun()` 関数を使ってこれを開始します。これは、`main()` 関数から呼び出されなければなりません。なお、1つのアプリケーションは1つのランループしか持たないため、`Run()` を複数回呼び出してはいけません。

```go
func init() {
    // ...
}

func main() {
	App := app.New()
	Window := App.NewWindow("Hello")
	Window.SetContent(widget.NewLabel("Hello"))
 	Window.Show()
	App.Run() // 全てのウィンドウが閉じられるまでブロック

	fmt.Println("Bye ノシ")
}
```

デスクトップの実行環境では、 `App.Quit()` を呼び出すことでアプリを直接終了させることができます（モバイルアプリはこれをサポートしていないようです）。しかし、アプリケーションは、すべてのウィンドウが閉じられると終了するため、通常は明示的に `App.Quit()` を実行する必要はありません。前段のソースコード内部にコメントした通り、`Run()` の後に実行される関数は、アプリケーションが終了するまで呼び出されないため注意が必要です。

## ウィンドウの処理 について
https://developer.fyne.io/tour/basics/windows.html

ウィンドウは `App.NewWindow()` を使用して作成され、`Show()` 関数を使用して表示する必要があります。`fyne.Window` には、この両者を同時に実行する ヘルパー メソッド `ShowAndRun()` が用意されているため、通常はこちらを使用することになります。

```go
func main() {
	App := app.New()
	Window := App.NewWindow("Hello")
	Window.SetContent(widget.NewLabel("Hello"))

	myWindow.ShowAndRun() // Window.Show() & App.Run()

	fmt.Println("Bye ノシ")
}
```

2つ目のウィンドウを表示したい場合は、`Show()` 関数のみを呼び出す必要があります。例えば、以下の例では `showAnother()` 関数内部で 新規ウィンドウが生成され、 `Show()` 関数が実行されています。 

```go
func main() {
	App := app.New()
	Window := App.NewWindow("Hello")
	Window.SetContent(widget.NewLabel("Hello"))

	go showAnother(App)
	Window.ShowAndRun()
	fmt.Println("Bye ノシ")
}

func showAnother(a fyne.App) {
	defer func() { fmt.Println("another window has closed") }()
	time.Sleep(time.Second * 3)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()

	time.Sleep(time.Second * 5)
	win.Close()
}
```

<!-- TODO: main.go へリンクする -->
<!-- TODO: 動きがわかるように GIF をとってあげよう -->

デフォルトでは、ウィンドウ は内包されるウィジェットに対して `MinSize()` 関数をチェックすることで、コンテンツを表示するのにサイズが確保されます。 より大きなサイズを設定する場合には、明示的に `Window.Resize()` 関数を呼び出すことで、より大きなサイズを設定することができます。

```go
func main() {
	App := app.New()
	Window := App.NewWindow("Hello")
	Window.SetContent(widget.NewLabel("Hello"))

	Window.Resize(fyne.NewSize(200, 200)) // ウィンドウサイズを指定
	Window.ShowAndRun()

	fmt.Println("Bye ノシ")
}
```

デスクトップ環境には、要求されたサイズよりもウィンドウが小さくなるような制約がある場合があるので注意が必要です。[^1]

[^1]: どういう時に起こるのか、ちょっと不明 :thinking_face: 
<!-- TODO: どういう時に発生するのか、調べる -->

## キャンバスとキャンバスオブジェクト
https://developer.fyne.io/tour/basics/canvas.html

Fyneでは、Canvasはアプリケーションが描画される領域です。各ウィンドウには、`Window.Canvas()` でアクセスできるキャンバスがありますが、通常、キャンバスへのアクセスを回避する `Window` の関数があります。

Fyneで描画可能なものはすべて `type CanvasObject interface` です。この例では、新しいウィンドウを開き、ウィンドウのキャンバスの内容を設定することで、さまざまなタイプのプリミティブなグラフィック要素を表示しています。テキストや円の例のように、各タイプのオブジェクトをカスタマイズする方法はたくさんあります。

`Canvas.SetContent()` を使って表示内容を変更するだけでなく、表示される内容を変更することも可能です。たとえば、四角形の色を変更した場合、`canvas.Refresh(rect)` を使って、既存のコンポーネントの更新を要求できます。

