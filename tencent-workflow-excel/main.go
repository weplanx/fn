package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	x := new(Inject)
	if err := Load(x); err != nil {
		log.Fatalln("environment configuration failed to load", err)
	}
	if err := Invoke(context.Background(), x); err != nil {
		log.Fatalln(err)
	}
}

func Invoke(ctx context.Context, x *Inject) (err error) {
	now := time.Now().String()
	key := fmt.Sprintf(`job_%s`, now)
	if _, err = x.Client.Object.Put(ctx, key, strings.NewReader(now), nil); err != nil {
		return
	}
	return
}
