# jcba streamer

Streaming ogg stream from jcba to stdout.

jcbaのoggストリームを標準出力に流します。

## 利用までの方法

- 方法1: ビルド済みバイナリを使う
  - releasesからダウンロードしてよしなにしてください
- 方法2: ソースから
  - ```
    go install github.com/mikuta0407/jcba-streamer@latest
    ```

## オプション
- `-s`: 局名 (下記例はfmfukuro)
- `-d`: 流す秒 (下記例は3600秒)


## 例
```
./jcba-streamer -s fmfukuro -d 3600 | ffmpeg -i pipe: -c libvorbis -acodec copy rec.opus
```

内容: fmfukuroから3600秒感ストリームを標準出力に流し、ffmpegでpipeから受け取ってrec.opusに保存

## その他
標準出力なのでそのままリダイレクトで保存することも出来ますが、メタデータ内の再生時間やプレイヤーのシークバーがぶっ壊れるのでffmpegを通すことをおすすめします。

## 免責事項
- 作者はこのプログラムを使ったことによる被害等は一切責任を負いません。自己責任でご利用下さい。

## 参考にしたもの

[jcba.py](https://kemasoft.net/?scr/jcba)