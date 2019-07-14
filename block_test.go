package block

import (
	"fmt"
	"testing"
	"time"

	"github.com/u6du/highwayhash"

	"github.com/u6du/block/xor"
)

func TestBlock(t *testing.T) {

	var msg [32]byte

	begin := uint64(time.Now().UnixNano())
	for count := uint64(1); count <= 10; count++ {
		salt := Begin0MoreThan(msg, 23)
		cost := (uint64(time.Now().UnixNano()) - begin) / uint64(time.Millisecond) / count
		fmt.Printf("MsgSaltLeadingZeroCount %d\n", MsgSaltLeadingZeroCount(msg, salt[:]))

		xor.Copy32(&salt, &msg)
		fmt.Printf("%dms %x\n", cost, highwayhash.Zero.Sum(append(msg[:], salt[:]...)))
		fmt.Printf("%x %x\n\n", msg, salt)
		Next(&msg)
	}
}
