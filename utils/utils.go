package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gookit/cache"
	"github.com/kr/pretty"
	"goUrlShortener/repository/short_entity_repository"
	"gorm.io/gorm/schema"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

func getStringToRandomizeFrom() []string {
	var arrayOfString []string
	for char := 'a'; char <= 'z'; char++ {
		arrayOfString = append(arrayOfString, string(char), string(unicode.ToUpper(char)))
	}
	for integer := 0; integer <= 9; integer++ {
		arrayOfString = append(arrayOfString, strconv.Itoa(integer))
	}
	return arrayOfString
}
func RandomUrlShorten(stringLen int) string {
	var randomizedStringArray []string
	arrayOfString := getStringToRandomizeFrom()
	for loopIteration := 0; loopIteration < stringLen; loopIteration++ {
		randomIndex := rand.Intn(len(arrayOfString))
		randomizedStringArray = append(randomizedStringArray, arrayOfString[randomIndex])
	}
	return strings.Join(randomizedStringArray, "")
}
func RandomUrlShortenAndGenerateNewUntilNotUsed(stringLen int) string {
	shortenedUrl := RandomUrlShorten(stringLen)
	firstTime := time.Now().Format(time.StampMilli)
	if CheckIfShortenIsPersistentInRedis(shortenedUrl) {
		return RandomUrlShortenAndGenerateNewUntilNotUsed(stringLen)
	}
	secTime := time.Now().Format(time.StampMilli)
	pretty.Println("redis", firstTime, secTime)
	cache.Set(shortenedUrl, "", cache.FiveMinutes)
	if CheckIfShortenIsPersistentInDb(shortenedUrl) {
		cache.Del(shortenedUrl)
		return RandomUrlShortenAndGenerateNewUntilNotUsed(stringLen)
	}
	thirdTime := time.Now().Format(time.StampMilli)
	pretty.Println("db", secTime, thirdTime)
	return shortenedUrl
}
func CheckIfShortenIsPersistentInRedis(urlShortened string) bool {
	if cache.Has(urlShortened) {
		return true
	} else {
		return false
	}
}
func CheckIfShortenIsPersistentInDb(urlShortened string) bool {
	_, err := short_entity_repository.GetSingle(urlShortened)
	if err != nil {
		return false
	}
	return true
}
func GetEnvVarAsInt(name string) int {
	var varAsInt, _ = strconv.Atoi(os.Getenv(name))
	return varAsInt
}
func ParseDateTimeAsOnlyDate(dateTime time.Time) time.Time {
	dateTime = dateTime.UTC()
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, dateTime.Location())
}

type SQLDateTimeToGoTime struct {
}

// Scan implements serializer interface
func (SQLDateTimeToGoTime) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)
	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}
		str := BytesToString(bytes)
		timeParsed, _ := time.Parse("2006-01-02 15:04:05", str)
		timeParsed = timeParsed
	}
	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

// Value implements serializer interface
func (SQLDateTimeToGoTime) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return json.Marshal(fieldValue)
}
func BytesToString(bytes []byte) string {
	var s string
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	stringHeader.Data = sliceHeader.Data
	stringHeader.Len = sliceHeader.Len
	return s
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := t.Time.Format("2006-01-02 15:04:05")
	return []byte(stamp), nil
}
func (t JSONTime) UnmarshalJSON(b []byte) error {
	stringFromByte := BytesToString(b)
	fmt.Println(stringFromByte)
	fmt.Println("XDDDD")
	return nil
}
func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
