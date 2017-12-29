package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type zalgoOpts struct {
	MinUp, MaxUp     int
	MinMid, MaxMid   int
	MinDown, MaxDown int
}

var zalgo_up = []rune{
	'\u030d', /*     ̍     */
	'\u030e', /*     ̎     */
	'\u0304', /*     ̄     */
	'\u0305', /*     ̅     */
	'\u033f', /*     ̿     */
	'\u0311', /*     ̑     */
	'\u0306', /*     ̆     */
	'\u0310', /*     ̐     */
	'\u0352', /*     ͒     */
	'\u0357', /*     ͗     */
	'\u0351', /*     ͑     */
	'\u0307', /*     ̇     */
	'\u0308', /*     ̈     */
	'\u030a', /*     ̊     */
	'\u0342', /*     ͂     */
	'\u0343', /*     ̓     */
	'\u0344', /*     ̈́     */
	'\u034a', /*     ͊     */
	'\u034b', /*     ͋     */
	'\u034c', /*     ͌     */
	'\u0303', /*     ̃     */
	'\u0302', /*     ̂     */
	'\u030c', /*     ̌     */
	'\u0350', /*     ͐     */
	'\u0300', /*     ̀     */
	'\u0301', /*     ́     */
	'\u030b', /*     ̋     */
	'\u030f', /*     ̏     */
	'\u0312', /*     ̒     */
	'\u0313', /*     ̓     */
	'\u0314', /*     ̔     */
	'\u033d', /*     ̽     */
	'\u0309', /*     ̉     */
	'\u0363', /*     ͣ     */
	'\u0364', /*     ͤ     */
	'\u0365', /*     ͥ     */
	'\u0366', /*     ͦ     */
	'\u0367', /*     ͧ     */
	'\u0368', /*     ͨ     */
	'\u0369', /*     ͩ     */
	'\u036a', /*     ͪ     */
	'\u036b', /*     ͫ     */
	'\u036c', /*     ͬ     */
	'\u036d', /*     ͭ     */
	'\u036e', /*     ͮ     */
	'\u036f', /*     ͯ     */
	'\u033e', /*     ̾     */
	'\u035b', /*     ͛     */
	'\u0346', /*     ͆     */
	'\u031a', /*     ̚     */
}

var zalgo_down = []rune{
	'\u0316', /*     ̖     */
	'\u0317', /*     ̗     */
	'\u0318', /*     ̘     */
	'\u0319', /*     ̙     */
	'\u031c', /*     ̜     */
	'\u031d', /*     ̝     */
	'\u031e', /*     ̞     */
	'\u031f', /*     ̟     */
	'\u0320', /*     ̠     */
	'\u0324', /*     ̤     */
	'\u0325', /*     ̥     */
	'\u0326', /*     ̦     */
	'\u0329', /*     ̩     */
	'\u032a', /*     ̪     */
	'\u032b', /*     ̫     */
	'\u032c', /*     ̬     */
	'\u032d', /*     ̭     */
	'\u032e', /*     ̮     */
	'\u032f', /*     ̯     */
	'\u0330', /*     ̰     */
	'\u0331', /*     ̱     */
	'\u0332', /*     ̲     */
	'\u0333', /*     ̳     */
	'\u0339', /*     ̹     */
	'\u033a', /*     ̺     */
	'\u033b', /*     ̻     */
	'\u033c', /*     ̼     */
	'\u0345', /*     ͅ     */
	'\u0347', /*     ͇     */
	'\u0348', /*     ͈     */
	'\u0349', /*     ͉     */
	'\u034d', /*     ͍     */
	'\u034e', /*     ͎     */
	'\u0353', /*     ͓     */
	'\u0354', /*     ͔     */
	'\u0355', /*     ͕     */
	'\u0356', /*     ͖     */
	'\u0359', /*     ͙     */
	'\u035a', /*     ͚     */
	'\u0323', /*     ̣     */
}

var zalgo_mid = []rune{
	'\u0315', /*     ̕     */
	'\u031b', /*     ̛     */
	'\u0340', /*     ̀     */
	'\u0341', /*     ́     */
	'\u0358', /*     ͘     */
	'\u0321', /*     ̡     */
	'\u0322', /*     ̢     */
	'\u0327', /*     ̧     */
	'\u0328', /*     ̨     */
	'\u0334', /*     ̴     */
	'\u0335', /*     ̵     */
	'\u0336', /*     ̶     */
	'\u034f', /*     ͏     */
	'\u035c', /*     ͜     */
	'\u035d', /*     ͝     */
	'\u035e', /*     ͞     */
	'\u035f', /*     ͟     */
	'\u0360', /*     ͠     */
	'\u0362', /*     ͢     */
	'\u0338', /*     ̸     */
	'\u0337', /*     ̷     */
	'\u0361', /*     ͡     */
	'\u0489', /*     ҉_    */
}

func randFromSlice(s []rune) rune {
	randIdx := rand.Intn(len(s))
	return s[randIdx]
}

func zalgoifyOne(r rune, z zalgoOpts) []rune {
	zalgoed := make([]rune, 0)
	zalgoed = append(zalgoed, r)
	for i := z.MinUp; i < z.MaxUp; i++ {
		zalgoed = append(zalgoed, randFromSlice(zalgo_up))
	}

	for i := z.MinMid; i < z.MaxMid; i++ {
		zalgoed = append(zalgoed, randFromSlice(zalgo_mid))
	}

	for i := z.MinDown; i < z.MaxDown; i++ {
		zalgoed = append(zalgoed, randFromSlice(zalgo_down))
	}
	return zalgoed
}

func zalgoify(s string, z zalgoOpts) string {
	runes := []rune(s)
	out := make([]rune, 0)
	for _, r := range runes {
		out = append(out, zalgoifyOne(r, z)...)
	}
	return string(out)
}

func newZalgoOpts(level int) zalgoOpts {
	z := zalgoOpts{}
	if level == 0 { // min
		z.MinUp = 0
		z.MaxUp = rand.Intn(8)
		z.MinMid = 0
		z.MaxMid = rand.Intn(2)
		z.MinDown = 0
		z.MaxDown = rand.Intn(8)
		return z
	}

	if level == 1 { // med
		z.MinUp = 1
		z.MaxUp = rand.Intn(16)/2 + 1
		z.MinMid = 1
		z.MaxMid = rand.Intn(6) / 2
		z.MinDown = 1
		z.MaxDown = rand.Intn(16)/2 + 1
		return z
	}

	// maximum zalgoooo
	z.MinUp = 3
	z.MaxUp = rand.Intn(64) + 3
	z.MinMid = 1
	z.MaxMid = rand.Intn(16) + 1
	z.MinDown = 3
	z.MaxDown = rand.Intn(64) + 3
	return z
}

func main() {
	rand.Seed(time.Now().Unix())
	var zalgoLevel int
	flag.IntVar(&zalgoLevel, "level", 0, "zalgo level (0-2)")
	flag.Parse()
	opts := newZalgoOpts(zalgoLevel)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(zalgoify(scanner.Text(), opts))
	}
}
