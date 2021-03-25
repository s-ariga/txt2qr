<!--
 Copyright 2021 Seiichi Ariga <seiichi.ariga@gmail.com>
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# QRコードを生成するコマンド

Go Lang。授業用のサンプルとして作成したコマンド。CLIのプログラム作成のヒント用

文字列からQRコードのPNGを作る

## オプション

| オプション名 | 意味 |　デフォルト |
 --- | --- | ---
| -i | 変換する文字列 | |
| -o | 出力する画像ファイル名(PNG) | qr_output.png |
| -s | 画像サイズ | 256 |
| -l | リカバリーレベル | 2 (Medium) |

## 依存ライブラリ

[github.com/skip2/go-qrcode](https://github.com/skip2/go-qrcode)