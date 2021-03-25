/**
 * Copyright 2021 Seiichi Ariga <seiichi.ariga@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"errors"
	"flag"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

type Flags struct {
	originalText string
	outputFile   string
	level        int
	size         int
}

func main() {

	f, err := createAndParseFlags()
	if err != nil {
		panic("オプションがエラー")
	}

	fmt.Println(f.originalText)
	fmt.Println(f.outputFile)

	err = qrcode.WriteFile(f.originalText, qrcode.RecoveryLevel(f.level), f.size, f.outputFile)
	if err != nil {
		panic("画像生成でエラー")
	}

}

func createAndParseFlags() (Flags, error) {
	f := Flags{
		originalText: "",
		outputFile:   "",
	}

	// オプションを設定
	flag.StringVar(&f.originalText, "i", "", "QRコードに変換する文字列")
	flag.StringVar(&f.outputFile, "o", "qr_output.png", "QRコード出力ファイル名")
	flag.IntVar(&f.level, "l", 2, "QRコードのリカバリーレベル（1:低 ~ 4:高）")
	flag.IntVar(&f.size, "s", 256, "QRコード画像(正方形)のサイズ")

	if f.level > 4 || f.level < 1 {
		return f, errors.New("リカバリーレベルは 1~4")
	}

	// オプションを解析
	flag.Parse()

	return f, nil
}
