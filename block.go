package block

import (
	"math/bits"

	"github.com/u6du/highwayhash"

	"github.com/u6du/block/xor"
)

func Next(msg *[32]byte) bool {
	for i := range msg {
		t := msg[i]
		if t != 255 {
			msg[i] = t + 1
			for j := 0; j < i; j++ {
				msg[j] = 0
			}
			return true
		}
	}
	return false
}

func MsgSaltLeadingZeroCount(msg, salt [32]byte) int {
	xor.Copy32(&salt, &msg)
	return LeadingZeroCount(highwayhash.Zero.Sum(append(msg[:], salt[:]...)))
}

const Zero = uint8(0)

func LeadingZeroCount(msg [32]byte) (n int) {

	for i := range msg {
		if Zero == msg[i] {
			n += 8
		} else {
			n += bits.LeadingZeros8(uint8(0) ^ msg[i])
			break
		}
	}

	return
}

func Begin0MoreThan(msg [32]byte, atLest int) [32]byte {
	salt := [32]byte{}

	for {
		if LeadingZeroCount(highwayhash.Zero.Sum(append(msg[:], salt[:]...))) >= atLest {
			xor.Copy32(&salt, &msg)
			return salt
		}

		if !Next(&salt) {
			return salt
		}
	}
}
