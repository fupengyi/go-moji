package go_moji

import "fmt"

这个包提供了一个 Go 接口，用于在 Zenkaku（全角即全角）和 Hankaku（半角即半角）字符（主要用于日语）之间进行转换。该库在很大程度上受到 JavaScript 实现的 niwaringo/moji 的影响。
有关API 的详细信息，请参阅文档。

Installation
Use go get:
	$ go get github.com/ktnyt/go-moji

Requirements
这个包只在 Go >= 1.8 上测试过。使用较低版本时要小心。

Example
package main

import (
	"fmt"
	"github.com/ktnyt/moji"
)

func main() {
	s := "ＡＢＣ ABC　あがぱ　アガパ　ｱｶﾞﾊﾟ"
	
	// Convert Zenkaku Eisuu to Hankaku Eisuu			// Convert 全角 字母数字 to 半角 字母数字
	fmt.Println(moji.Convert(s, moji.ZE, moji.HE))
	
	// Convert Hankaku Eisuu to Zenkaku Eisuu			// Convert 半角 字母数字 to 全角 字母数字
	fmt.Println(moji.Convert(s, moji.HE, moji.ZE))
	
	// Convert HiraGana to KataKana						// 将平假名转换为片假名
	fmt.Println(moji.Convert(s, moji.HG, moji.KK))
	
	// Convert KataKana to HiraGana						// 将片假名转换为平假名
	fmt.Println(moji.Convert(s, moji.KK, moji.HG))
	
	// Convert Zenkaku Katakana to Hankaku Katakana		// 将全角 片假名 转换为 半角 平假名
	fmt.Println(moji.Convert(s, moji.ZK, moji.HK))
	
	// Convert Hankaku Katakana to Zenkaku Katakana		// 将半角 片假名 转换为 全角 平假名
	fmt.Println(moji.Convert(s, moji.HK, moji.ZK))
	
	// Convert Zenkaku Space to Hankaku Space
	fmt.Println(moji.Convert(s, moji.ZS, moji.HS))
	
	// Convert Hankaku Space to Zenkaku Space
	fmt.Println(moji.Convert(s, moji.HS, moji.ZS))
}

Variables
var HE = NewRangeDictionary(0x0021, 0x007e)			// HE 定义了 Hankaku Eisuu（即半角英文文本）词典
var HG = NewRangeDictionary(0x3041, 0x3096)			// HG 定义了 HiraGana (平假名) 词典
var HK = NewDictionary([]string{							// HK 定义半角片假名（即半角片假名）词典
	"ｶﾞ", "ｷﾞ", "ｸﾞ", "ｹﾞ", "ｺﾞ",
	"ｻﾞ", "ｼﾞ", "ｽﾞ", "ｾﾞ", "ｿﾞ",
	"ﾀﾞ", "ﾁﾞ", "ﾂﾞ", "ﾃﾞ", "ﾄﾞ",
	"ﾊﾞ", "ﾊﾟ", "ﾋﾞ", "ﾋﾟ", "ﾌﾞ", "ﾌﾟ", "ﾍﾞ", "ﾍﾟ", "ﾎﾞ", "ﾎﾟ",
	"ﾜﾞ", "ｦﾞ", "ｳﾞ",
	"｡", "｢", "｣", "､", "･", "ｰ", "ﾞ", "ﾟ",
	"ｱ", "ｲ", "ｳ", "ｴ", "ｵ",
	"ｶ", "ｷ", "ｸ", "ｹ", "ｺ",
	"ｻ", "ｼ", "ｽ", "ｾ", "ｿ", "ﾀ", "ﾁ", "ﾂ", "ﾃ", "ﾄ",
	"ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ",
	"ﾊ", "ﾋ", "ﾌ", "ﾍ", "ﾎ",
	"ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ",
	"ﾔ", "ﾕ", "ﾖ",
	"ﾗ", "ﾘ", "ﾙ", "ﾚ", "ﾛ",
	"ﾜ", "ｦ", "ﾝ",
	"ｧ", "ｨ", "ｩ", "ｪ", "ｫ",
	"ｬ", "ｭ", "ｮ", "ｯ",
})
var HS = NewDictionary([]string{" ", stringify(0x00a0)})	// HS 定义半角空间（即半角空间）字典
var KK = NewRangeDictionary(0x30a1, 0x30f6)					// KK 定义片假名词典
var ZE = NewRangeDictionary(0xff01, 0xff5e)					// ZE 定义了 Zenkaku Eisuu（即全角英文文本）词典
var ZK = NewDictionary([]string{							// ZK 定义了全角片假名（即全角片假名）词典
	"ガ", "ギ", "グ", "ゲ", "ゴ",
	"ザ", "ジ", "ズ", "ゼ", "ゾ",
	"ダ", "ヂ", "ヅ", "デ", "ド",
	"バ", "パ", "ビ", "ピ", "ブ", "プ", "ベ", "ペ", "ボ", "ポ",
	"ヷ", "ヺ", "ヴ",
	"。", "「", "」", "、", "・", "ー", "゛", "゜",
	"ア", "イ", "ウ", "エ", "オ",
	"カ", "キ", "ク", "ケ", "コ",
	"サ", "シ", "ス", "セ", "ソ",
	"タ", "チ", "ツ", "テ", "ト",
	"ナ", "ニ", "ヌ", "ネ", "ノ",
	"ハ", "ヒ", "フ", "ヘ", "ホ",
	"マ", "ミ", "ム", "メ", "モ",
	"ヤ", "ユ", "ヨ",
	"ラ", "リ", "ル", "レ", "ロ",
	"ワ", "ヲ", "ン",
	"ァ", "ィ", "ゥ", "ェ", "ォ",
	"ャ", "ュ", "ョ", "ッ",
})
var ZS = NewDictionary([]string{stringify(0x3000), stringify(0x3000)})		// ZS 定义了 Zenkaku Space（即全角空间）词典

Functions
func Convert(s string, from, to Dictionary) string							// 在两个字典之间转换字符串

Types
1.type Dictionary interface {												// 字典定义了一个字符串和索引之间映射的接口
	Encode([]MaybeIndex) string
	Decode(string) []MaybeIndex
}
func NewDictionary(d []string) Dictionary									// NewDictionary 从给定的字符串切片创建一个字典
func NewRangeDictionary(s, e rune) Dictionary								// NewRangeDictionary 从给定范围创建字典
2.type MaybeIndex struct {													// MaybeIndex 可能是一个索引或一个符文
	// contains filtered or unexported fields
}
