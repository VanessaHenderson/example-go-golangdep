package sub

import "context"
import "strconv"
import "fmt"
import "crypto/rsa"
import "crypto/rand"
import "github.com/google/go-github/github"
import "github.com/square/go-jose"

func Foo() {
    client := github.NewClient(nil)

    // list public repositories for org "github"
    opt := &RepositoryListByOrgOptions{"forks", ListOptions{Page: 2}}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "o", opt)
    
    fmt.Println("Hello from sub! Num repos: " + strconv.Itoa(len(repos)))

    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    publicKey := &privateKey.PublicKey
    encrypter, err := jose.NewEncrypter(jose.RSA_OAEP, jose.A128GCM, publicKey)
    if err != nil {
       panic(err)
    }

    var plaintext = []byte("Lorem ipsum dolor sit amet")
    object, err := encrypter.Encrypt(plaintext)
    if err != nil {
        panic(err)
    }

    fmt.Println(object)
}
