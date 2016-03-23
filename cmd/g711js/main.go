// g711js generates a Javascript G.711 decoder.
package main

import (
	"fmt"

	g711 "github.com/pd0mz/go-g711"
)

var (
	ad []int16
	µd []int16
)

func init() {
	ad = make([]int16, 256)
	µd = make([]int16, 256)

	var i uint16
	for i = 0; i < 0x100; i++ {
		ad[i] = g711.FromALaw(uint8(i))
		µd[i] = g711.FromMLaw(uint8(i))
	}
}

func dump(v []int16) {
	fmt.Print(`        `)
	for i, v := range ad {
		if v < 0 {
			fmt.Printf("%#05x", v)
		} else {
			fmt.Printf(" %#04x", v)
		}
		if i != 0xff {
			fmt.Print(", ")
			if (i+1)%8 == 0 {
				fmt.Print("\n        ")
			}
		}
	}
}

func main() {
	fmt.Println(`var G711 = function() {`)
	fmt.Println(`    this.aLaw = [`)
	dump(ad)
	fmt.Println("\n    ];")
	fmt.Println(`    this.uLaw = [`)
	dump(µd)
	fmt.Println("\n    ];")
	fmt.Println(`};

G711.prototype = {
    _decode: function(buffer, mapping) {
        var output = new Int16Array(buffer.length);
        for (var i = 0; i < buffer; i++) {
            output[i] = mapping[buffer[i]];
        }
        return output;
    },
	
    aLawDecode: function(buffer) {
        return this._decode(buffer, this.aLaw);
    },

    uLawDecode: function(buffer) {
        return this._decode(buffer, this.uLaw);
    }
};`)
}
