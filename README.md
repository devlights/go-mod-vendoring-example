# go-mod-vendoring-example
Go で go mod vendor を試してみるサンプルです。

## Vendoring

vendoring は、アプリケーションで利用するライブラリをプロジェクトのルートディレクトリに取り込み

バージョンを固定するやり方。業務によってはネットワークなどが利用できない環境でビルドを行ったりすることがある。

そのような場合に vendoring を利用して依存パッケージをプロジェクトのルートディレクトリに取り込んでおくことが出来る。

## やり方

予め ```go get``` を実行して必要な依存ライブラリを取得しておき

その後 ```go mod vendor``` を実行すると、アプリケーションが依存している部分のライブラリソースコードが vendor ディレクトリに出力される。

go で ビルド や 実行 を行う際に ランタイム は vendor ディレクトリがあれば、優先してそちらを参照するようになっている。

## vendoringしているライブラリの更新

```go get .; go mod tidy; go mod vendor``` とすると更新される。

## 試す

本リポジトリのサンプルは [go-task](https://taskfile.dev/) で実行できるようにしています。

```sh
$ task --list
task: Available tasks for this project:
* run-normal:           vendoring なし
* run-using-vendor:     vendoring あり
```

### Vendoring なしの場合（通常通りのやり方)

```sh
$ task run-normal
task: [run-normal] go clean
task: [run-normal] go clean -modcache -x
rm -rf /workspace/go/pkg/mod
task: [run-normal] go get .
go: downloading github.com/devlights/gomy v0.4.6
task: [run-normal] go build
task: [run-normal] ./go-mod-vendoring-example
hello world
task: [run-normal] bin/exa --tree $(go env GOPATH)/pkg
/workspace/go/pkg
├── mod
│  ├── cache
│  │  └── download
│  │     └── github.com
│  │        └── devlights
│  │           └── gomy
│  │              └── @v
│  │                 ├── list
│  │                 ├── v0.4.6.info
│  │                 ├── v0.4.6.lock
│  │                 ├── v0.4.6.mod
│  │                 ├── v0.4.6.zip
│  │                 └── v0.4.6.ziphash
│  └── github.com
│     └── devlights
│        └── gomy@v0.4.6
│           ├── async
│           │  ├── doc.go
│           │  ├── mergedwg.go
│           │  └── mergedwg_test.go
│           ├── bcd
│           │  ├── bcd.go
│           │  ├── bcd_test.go
│           │  └── doc.go
│           ├── chans
│           │  ├── bridge.go
│           │  ├── bridge_test.go
│           │  ├── buffer.go
│           │  ├── buffer_test.go
│           │  ├── chain.go
│           │  ├── chain_test.go
│           │  ├── chunk.go
│           │  ├── chunk_test.go
│           │  ├── concat.go
│           │  ├── concat_test.go
│           │  ├── convert.go
│           │  ├── convert_test.go
│           │  ├── doc.go
│           │  ├── enumerate.go
│           │  ├── enumerate_test.go
│           │  ├── example_test.go
│           │  ├── fanin.go
│           │  ├── fanin_test.go
│           │  ├── fanout.go
│           │  ├── fanout_test.go
│           │  ├── filter.go
│           │  ├── filter_test.go
│           │  ├── foreach.go
│           │  ├── foreach_test.go
│           │  ├── generator.go
│           │  ├── generator_test.go
│           │  ├── interval.go
│           │  ├── interval_test.go
│           │  ├── loop.go
│           │  ├── loop_test.go
│           │  ├── map.go
│           │  ├── map_test.go
│           │  ├── ordone.go
│           │  ├── ordone_test.go
│           │  ├── repeat.go
│           │  ├── repeat_test.go
│           │  ├── select.go
│           │  ├── select_test.go
│           │  ├── skip.go
│           │  ├── skip_test.go
│           │  ├── take.go
│           │  ├── take_test.go
│           │  ├── tee.go
│           │  ├── tee_test.go
│           │  ├── whenall.go
│           │  ├── whenall_test.go
│           │  ├── whenany.go
│           │  └── whenany_test.go
│           ├── cmd
│           │  ├── disphex
│           │  │  ├── main.go
│           │  │  └── README.md
│           │  ├── extract
│           │  │  ├── main.go
│           │  │  └── README.md
│           │  ├── splitbin
│           │  │  ├── main.go
│           │  │  └── README.md
│           │  └── splitrec
│           │     ├── main.go
│           │     └── README.md
│           ├── consts
│           │  ├── consts_test.go
│           │  ├── doc.go
│           │  └── returncode.go
│           ├── convert
│           │  ├── bin.go
│           │  ├── bin_test.go
│           │  ├── dec.go
│           │  ├── dec_test.go
│           │  ├── doc.go
│           │  ├── hex.go
│           │  └── hex_test.go
│           ├── ctxs
│           │  ├── doc.go
│           │  ├── donech.go
│           │  ├── donech_test.go
│           │  ├── export_test.go
│           │  ├── whenall.go
│           │  ├── whenall_test.go
│           │  ├── whenany.go
│           │  └── whenany_test.go
│           ├── deepcopy
│           │  ├── deepcopy.go
│           │  ├── deepcopy_test.go
│           │  └── doc.go
│           ├── enumerable
│           │  ├── doc.go
│           │  ├── range.go
│           │  └── range_test.go
│           ├── errs
│           │  ├── forget.go
│           │  ├── forget_test.go
│           │  ├── stderr.go
│           │  ├── stderr_test.go
│           │  ├── stdout.go
│           │  └── stdout_test.go
│           ├── exts
│           │  ├── doc.go
│           │  ├── num.go
│           │  └── num_test.go
│           ├── fileio
│           │  ├── copy.go
│           │  ├── copy_test.go
│           │  ├── dir.go
│           │  ├── dir_test.go
│           │  ├── doc.go
│           │  ├── file.go
│           │  ├── file_test.go
│           │  └── jp
│           │     ├── doc.go
│           │     └── encoding.go
│           ├── go.mod
│           ├── go.sum
│           ├── iter
│           │  ├── doc.go
│           │  ├── range.go
│           │  ├── range_test.go
│           │  ├── zip.go
│           │  └── zip_test.go
│           ├── latch
│           │  ├── countdownlatch.go
│           │  ├── countdownlatch_test.go
│           │  └── doc.go
│           ├── LICENSE
│           ├── logops
│           │  ├── default.go
│           │  ├── default_test.go
│           │  └── doc.go
│           ├── Makefile
│           ├── mem
│           │  ├── doc.go
│           │  ├── mem.go
│           │  └── mem_test.go
│           ├── misc
│           │  ├── doc.go
│           │  ├── misc.go
│           │  └── misc_test.go
│           ├── output
│           │  ├── doc.go
│           │  ├── output.go
│           │  └── output_test.go
│           ├── README.md
│           ├── signals
│           │  ├── doc.go
│           │  ├── ready.go
│           │  └── ready_test.go
│           ├── slices
│           │  ├── doc.go
│           │  ├── filter.go
│           │  └── filter_test.go
│           ├── strops
│           │  ├── doc.go
│           │  ├── newline.go
│           │  └── newline_test.go
│           └── times
│              ├── doc.go
│              ├── format.go
│              ├── format_test.go
│              ├── parse.go
│              └── parse_test.go
└── sumdb
   └── sum.golang.org
      └── latest
```

### Vendoring ありの場合にどのようになるか試す

```sh
$ task run-using-vendor
task: [run-using-vendor] go get .
task: [run-using-vendor] go mod vendor
task: [run-using-vendor] go clean
task: [run-using-vendor] go clean -modcache -x
rm -rf /workspace/go/pkg/mod
task: [run-using-vendor] go build
task: [run-using-vendor] ./go-mod-vendoring-example
hello world
task: [run-using-vendor] bin/exa --tree $(go env GOPATH)/pkg
/workspace/go/pkg
├── mod
└── sumdb
   └── sum.golang.org
      └── latest
task: [run-using-vendor] bin/exa --tree vendor/
vendor
├── github.com
│  └── devlights
│     └── gomy
│        ├── LICENSE
│        └── logops
│           ├── default.go
│           └── doc.go
└── modules.txt
```


## 関連リポジトリ

- [try-golang](https://github.com/devlights/try-golang)
  - Go のサンプルを置いているリポジトリです。
