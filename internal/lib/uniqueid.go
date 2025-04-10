package lib

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	digits  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	decimal = int64(len(digits)) // N进制
)

var (
	snowflakeNode      [1 << 10]*snowflake.Node
	digitsReverseIndex = make(map[byte]int)
)

func init() {
	for j, digit := range []byte(digits) {
		digitsReverseIndex[digit] = j
	}
	for i := 0; i < 1<<10; i++ {
		node, err := snowflake.NewNode(int64(i))
		if err != nil {
			panic(err)
		}
		snowflakeNode[i] = node
	}
}

func GenUniqueID(nodeID ...int64) string {
	if len(nodeID) == 0 {
		nodeID = []int64{0}
	}
	idx := nodeID[rand.Intn(len(nodeID))]
	return encode(snowflakeNode[idx].Generate().Int64())
}

func DecodeTime(str string) (time.Time, error) {
	snowID, err := decode(str)
	if err != nil {
		return time.Time{}, err
	}
	return time.UnixMilli(snowflake.ID(snowID).Time()), nil
}

func encode(num int64) string {
	b := strings.Builder{}
	l := decimal
	for {
		if num < 1 {
			break
		}
		b.WriteByte(digits[num%l])
		num /= l
	}
	ret := []byte(b.String())
	reverse(ret)
	return string(ret)
}

func decode(str string) (int64, error) {
	bytes := []byte(str)
	reverse(bytes)
	var ret int64
	for i, byt := range bytes {
		digitIdx, ok := digitsReverseIndex[byt]
		if !ok {
			return 0, fmt.Errorf("invalid character found:%v", rune(byt))
		}
		ret += int64(digitIdx) * int64(math.Pow(float64(decimal), float64(i)))
	}
	return ret, nil
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
