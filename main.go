package main

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/math/uints"
	"github.com/reclaimprotocol/gnark-chacha20/chacha"

	"golang.org/x/crypto/chacha20"
)

// /go:embed prove/r1cs
var r1cs_embedded []byte

// /go:embed prove/pk
var pk_embedded []byte

type Witness struct {
	Key     []uints.U32
	Counter uints.U32 `gnark:",public"`
	Nonce   []uints.U32
	In      []uints.U32 `gnark:",public"`
	Out     []uints.U32 `gnark:",public"`
}

func (c *Witness) Define(api frontend.API) error {
	return nil
}

func main() {

	/*go func() {
		http.ListenAndServe("localhost:8088", nil)
	}()*/
	/*time.Sleep(time.Second * 10)
	for i := 0; i < 10; i++ {
		err := ProveG16()
		if err != nil {
			log.Fatal("groth16 error:", err)
		}
	}
	time.Sleep(time.Second * 1000)*/
	generateGroth16()
	// trySerialize()
}

// var r1css = groth16.NewCS(ecc.BN254)
// var pk groth16.ProvingKey

func generateGroth16() error {

	/*fmt.Println("about to read key & circuit")
	r1css = groth16.NewCS(ecc.BN254)
	_, err := r1css.ReadFrom(bytes.NewBuffer(r1cs_embedded))
	if err != nil {
		panic(err)
	}

	pk = groth16.NewProvingKey(ecc.BN254)
	_, err = pk.ReadFrom(bytes.NewBuffer(pk_embedded))
	if err != nil {
		panic(err)
	}
	fmt.Println("read key & circuit")

	bKey := make([]uint8, 32)
	rand.Read(bKey)
	bNonce := make([]uint8, 12)
	rand.Read(bNonce)
	counter := uints.NewU32(1)

	dataBytes := chacha.Blocks * 64
	bPt := make([]byte, dataBytes)
	rand.Read(bPt)
	bCt := make([]byte, dataBytes)

	cipher, err := chacha20.NewUnauthenticatedCipher(bKey, bNonce)
	if err != nil {
		return err
	}

	cipher.SetCounter(1)
	cipher.XORKeyStream(bCt, bPt)

	plaintext := chacha.BytesToUint32BE(bPt)
	ciphertext := chacha.BytesToUint32BE(bCt)

	fmt.Printf("%0X\n", bKey)
	fmt.Printf("%0X\n", bNonce)
	fmt.Printf("%0X\n", bPt)
	fmt.Printf("%0X\n", bCt)

	witness := chacha.Circuit{}
	copy(witness.Key[:], chacha.BytesToUint32LE(bKey))
	copy(witness.Nonce[:], chacha.BytesToUint32LE(bNonce))
	witness.Counter = counter
	copy(witness.In[:], plaintext)
	copy(witness.Out[:], ciphertext)

	wtns, err := frontend.NewWitness(&witness, ecc.BN254.ScalarField())

	proof, err := groth16.Prove(r1css, pk, wtns)

	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	_, err = proof.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	res := buf.Bytes()
	fmt.Printf("%0X\n", res)*/

	/*f, err := os.Open("f:\\r1cs")
	r1css := groth16.NewCS(ecc.BN254)
	r1css.ReadFrom(f)
	f.Close()

	f1, err := os.Open("f:\\pk")
	pk := groth16.NewProvingKey(ecc.BN254)
	pk.ReadFrom(f1)
	f1.Close()*/

	/*f2, err := os.Open("f:\\vk")
	vk := groth16.NewVerifyingKey(ecc.BN254)
	vk.ReadFrom(f2)
	f2.Close()*/

	/*p := profile.Start()
	_, _ = frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &chacha.Circuit{})
	p.Stop()

	fmt.Println(p.NbConstraints())
	fmt.Println(p.Top())*/

	bKey := make([]uint8, 32)
	rand.Read(bKey)
	bNonce := make([]uint8, 12)
	rand.Read(bNonce)

	cnt := uint32(1)

	counter := uints.NewU32(cnt)

	dataBytes := chacha.Blocks * 64
	bPt := make([]byte, dataBytes)
	rand.Read(bPt)
	bCt := make([]byte, dataBytes)

	cipher, err := chacha20.NewUnauthenticatedCipher(bKey, bNonce)
	if err != nil {
		return err
	}

	cipher.SetCounter(cnt)
	cipher.XORKeyStream(bCt, bPt)

	plaintext := chacha.BytesToUint32BE(bPt)
	ciphertext := chacha.BytesToUint32BE(bCt)
	key := chacha.BytesToUint32LE(bKey)
	nonce := chacha.BytesToUint32LE(bNonce)

	fmt.Printf("%0X\n", bKey)
	fmt.Printf("%0X\n", bNonce)
	// fmt.Printf("%0X\n", bPt)
	// fmt.Printf("%0X\n", bCt)

	witness := Witness{
		Counter: uints.NewU32(0),
		Key:     make([]uints.U32, len(key)),
		Nonce:   make([]uints.U32, len(nonce)),
		In:      make([]uints.U32, len(plaintext)),
		Out:     make([]uints.U32, len(ciphertext)),
	}

	curve := ecc.BN254.ScalarField()
	/*t := time.Now()
	r1css, err := frontend.Compile(curve, r1cs.NewBuilder, &witness, frontend.WithCompressThreshold(10), frontend.WithCapacity(500000))
	if err != nil {
		return err
	}
	fmt.Println("compile took ", time.Since(t))

	fmt.Printf("Blocks: %d, constraints: %d\n", chacha.Blocks, r1css.GetNbConstraints())

	f, err := os.OpenFile("f:\\r1cs", os.O_RDWR|os.O_CREATE, 0777)
	r1css.WriteTo(f)
	f.Close()

	pk1, vk1, err := groth16.Setup(r1css)
	if err != nil {
		return err
	}

	f2, err := os.OpenFile("f:\\pk", os.O_RDWR|os.O_CREATE, 0777)
	pk1.WriteTo(f2)
	f2.Close()

	f3, err := os.OpenFile("f:\\vk", os.O_RDWR|os.O_CREATE, 0777)
	vk1.WriteTo(f3)
	f3.Close()*/

	r1css := groth16.NewCS(ecc.BN254)
	f, err := os.OpenFile("f:\\r1cs", os.O_RDONLY, 0777)
	_, err = r1css.ReadFrom(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	witness = Witness{
		Key:     key,
		Nonce:   nonce,
		Counter: counter,
		In:      plaintext,
		Out:     ciphertext,
	}

	fmt.Println("witness")
	wtns, err := frontend.NewWitness(&witness, curve)
	if err != nil {
		panic(err)
	}
	fmt.Println("loading proving key")
	fpk, _ := os.OpenFile("f:\\pk", os.O_RDONLY, 0777)
	pk := groth16.NewProvingKey(ecc.BN254)
	pk.ReadFrom(fpk)
	fpk.Close()

	fmt.Println("Proving")
	proof, err := groth16.Prove(r1css, pk, wtns)
	buf := &bytes.Buffer{}
	_, err = proof.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	res := buf.Bytes()
	fmt.Printf("%0X\n", res)
	/*f3, err := os.OpenFile("f:\\proof", os.O_RDWR|os.O_CREATE, 0777)
	proof.WriteTo(f3)
	f3.Close()*/

	witness = Witness{
		Counter: counter,
		In:      plaintext,
		Out:     ciphertext,
	}
	wp, err := frontend.NewWitness(&witness, curve, frontend.PublicOnly())
	fvk, _ := os.OpenFile("f:\\vk", os.O_RDONLY, 0777)
	vk := groth16.NewVerifyingKey(ecc.BN254)
	vk.ReadFrom(fvk)
	fvk.Close()
	err = groth16.Verify(proof, vk, wp)
	fmt.Println("proof ok", err == nil)
	return nil
}

/*func ProveG16() error {
	key := make([]uint8, 32)
	rand.Read(key)
	nonce := make([]uint8, 12)
	rand.Read(nonce)
	cnt := uints.NewU32(1)

	dataBytes := chacha.Blocks * 64
	bPt := make([]byte, dataBytes)
	rand.Read(bPt)
	bCt := make([]byte, dataBytes)

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return err
	}

	cipher.SetCounter(1)
	cipher.XORKeyStream(bCt, bPt)

	plaintext := chacha.BytesToUint32BE(bPt)
	ciphertext := chacha.BytesToUint32BE(bCt)

	r1css := groth16.NewCS(ecc.BN254)
	_, err = r1css.ReadFrom(bytes.NewBuffer(r1cs_embedded))
	if err != nil {
		panic(err)
	}

	pk := groth16.NewProvingKey(ecc.BN254)
	_, err = pk.ReadFrom(bytes.NewBuffer(pk_embedded))
	if err != nil {
		panic(err)
	}

	witness := chacha.Circuit{}
	copy(witness.Key[:], chacha.BytesToUint32LE(key))
	copy(witness.Nonce[:], chacha.BytesToUint32LE(nonce))
	witness.Counter = cnt
	copy(witness.In[:], plaintext)
	copy(witness.Out[:], ciphertext)
	wtns, err := frontend.NewWitness(&witness, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}
	_, err = groth16.Prove(r1css, pk, wtns)
	return err
}*/

type Circuit1 struct {
	Val  uints.U32
	Val1 uints.U32
}

func (c *Circuit1) Define(api frontend.API) error {
	uapi, err := uints.New[uints.U32](api)
	if err != nil {
		return err
	}
	uapi.AssertEq(c.Val, c.Val1)
	return nil
}

func trySerialize() {
	witness := Circuit1{
		Val:  uints.NewU32(10),
		Val1: uints.NewU32(10),
	}

	curve := ecc.BN254.ScalarField()
	r1css, err := frontend.Compile(curve, r1cs.NewBuilder, &witness)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	r1css.WriteTo(buf)

	r1css = groth16.NewCS(ecc.BN254)
	r1css.ReadFrom(bytes.NewBuffer(buf.Bytes()))
	fmt.Println(r1css.GetNbConstraints())
}
