package play

import (
	"fmt"
	"log"
	"strings"
)

// NoError はエラーがあればログを出力して終了します。
// en: NoError logs the error and exits if there is an error.
func NoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Debug は渡された値をまとめて詳細にログ出力します
//
// en: Debug logs the passed values in detail.
// If you pass arguments using play.KV(), it will map and output the names appropriately.
func Debug(values ...any) {
	builder := strings.Builder{}
	builder.WriteString("============= Debug =============\n")
	for index, value := range values {
		if kv, ok := value.(KeyValue); ok {
			builder.WriteString(fmt.Sprintf("    %s => %#v\n", kv.Key, kv.Value))
			continue
		}
		builder.WriteString(fmt.Sprintf("    %2d => %#v\n", index, value))
	}
	builder.WriteString("=====================================================\n")
	log.Print(builder.String())
}

// KeyValue はデバッグ用のキーと値を保持します。
// en: KeyValue holds a key and value for debugging.
type KeyValue struct {
	Key   string
	Value any
}

// KV は Debug に渡す KeyValue を生成します。reflectやruntimeを使っても関数呼び出し元での変数名等は取得できないので、この設計になっています。
//
// en: KV generates a KeyValue to pass to Debug. This design is used because it is not possible to get the variable name
// at the call site using the reflect or runtime packages.
func KV(key string, value any) KeyValue {
	return KeyValue{Key: key, Value: value}
}
