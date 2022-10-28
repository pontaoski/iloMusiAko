package main

import "math/rand"
// import "fmt"   // for testing
// import "time"  //

// kiwen — lon musi la sitelen ‹.› li kama weka

const nPoka int = 12 // nanpa poka kiwen

var kiwenSuli = [5]string {
    "eioujklmnpst",
    "aiojklmnpstw",
    "alnpst......",
    "akmpsw......",
    "klmntw......",
}

var kiwenNamako = [5]string {
    "aouwljptknms",
    "aeiwljptknms",
    "aiouwlptknms",
    "aiowljptknms",
    "aeowljptknms",
}

func randomLetters(nNamako int) []string {
    // theoretical maximum number of letters
    nMuSewi := len(kiwenSuli) + nNamako * len(kiwenNamako)
    
    // empty array with capacity set at maximum theoretical
    mu := make([]string, 0, nMuSewi)

    /// o kama jo e mu tan kiwen suli
    for k := 0; k < len(kiwenSuli); k++ { // k = kiwen
        // rand.Seed(time.Now().UTC().UnixNano()) // will be done in commands.go, although here is probably better
        m := string(kiwenSuli[k][rand.Intn(nPoka)]) // roll one die
        if m != "." { // if not ‹.›: add;
            mu = append(mu, m)
        }             // if ‹.›: continue
    }
    
    /// o kama jo e mu tan kiwen namako
    for t := 0; t < nNamako; t++ { // t = tenpo
        for k := 0; k < len(kiwenNamako); k++ { // k = kiwen
            m := string(kiwenNamako[k][rand.Intn(nPoka)]) // roll one die
            if m != "." { // if not ‹.›: add;
                mu = append(mu, m)
            }             // if ‹.›: continue
        }
    }
    
    /// o nasa e mu
    nMu := len(mu) // the actual number of letters it has
    nasinNasa := rand.Perm(nMu) // random shuffle order
    
    muNasa := make([]string, nMu)
    
    // construct new array by using the random permutation
    for m := 0; m < nMu; m++ { // m = mu
        muNasa[m] = mu[nasinNasa[m]]
    }
    
	return muNasa
}

/* // for testing
func main() {
    for j := 0; j < 10; j++ {
        fmt.Println(randomLetters(0))
        // o ante e nanpa pi kulupu pi kiwen namako
    }
}
// */
