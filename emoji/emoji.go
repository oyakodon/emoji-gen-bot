package emoji

// #cgo LDFLAGS: -L../third_party/libemoji/lib -lemoji -lstdc++ -lz -lm -lpthread -lfontconfig -lfreetype -lGL
/*
  #include "../third_party/libemoji/include/emoji.h"
*/
import "C"
import "fmt"

var (
	ErrGenerateEmoji = fmt.Errorf("error occured during generating emoji")
)

type EmojiTextAlign int

const (
	EmojiAlignLeft EmojiTextAlign = iota
	EmojiAlignCenter
	EmojiAlignRight
)

const (
	EmojiDefaultFontPath = "./assets/NotoSansMonoCJKjp-Bold.otf"

	EmojiSizeMDPI = 128

	EmojiColorTransparent = 0x00FFFFFF // 透明
)

// 絵文字を生成し、画像のバイナリデータを返す。
//
// fgColor, bgColor: ARGB
func GenerateEmoji(text string, fgColor, bgColor int, align EmojiTextAlign) ([]byte, error) {
	return generateEmojiPng(
		text,
		EmojiSizeMDPI,
		EmojiSizeMDPI,
		fgColor,
		bgColor,
		align,
		EmojiDefaultFontPath,
	)
}

func generateEmojiPng(text string, width, height, fgColor, bgColor int, align EmojiTextAlign, fontpath string) ([]byte, error) {
	params := C.EgGenerateParams{
		fText:            C.CString(text),
		fWidth:           C.uint(width),
		fHeight:          C.uint(height),
		fColor:           C.uint(fgColor),
		fBackgroundColor: C.uint(bgColor),
		fTextAlign:       C.EgAlign(align),
		fTextSizeFixed:   false, // 文字サイズを固定
		fDisableStretch:  false, //	自動で伸縮しない
		fTypefaceFile:    C.CString(fontpath),
		fFormat:          C.kPNG_Format,
		fQuality:         100,
	}

	var result C.EgGenerateResult
	defer C.emoji_free(&result)

	if C.emoji_generate(&params, &result) != C.EG_OK {
		return nil, ErrGenerateEmoji
	}

	b := C.GoBytes(result.fData, C.int(result.fSize))
	return b, nil
}
